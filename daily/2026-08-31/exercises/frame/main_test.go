package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"testing"
)

type slowReader struct{ r io.Reader }

func (s slowReader) Read(p []byte) (int, error) {
	if len(p) > 1 {
		p = p[:1]
	}
	return s.r.Read(p)
}

func TestFrameRoundTripWithFragmentedReads(t *testing.T) {
	var buf bytes.Buffer
	if err := WriteFrame(&buf, []byte("hello")); err != nil {
		t.Fatal(err)
	}
	got, err := ReadFrame(slowReader{r: &buf})
	if err != nil || string(got) != "hello" {
		t.Fatalf("ReadFrame() = %q, %v", got, err)
	}
}

func TestReadFrameRejectsOversizedFrame(t *testing.T) {
	var header [4]byte
	binary.BigEndian.PutUint32(header[:], MaxFrameSize+1)
	if _, err := ReadFrame(bytes.NewReader(header[:])); err == nil {
		t.Fatal("expected oversized frame error")
	}
}

func TestReadFrameRejectsTruncatedPayload(t *testing.T) {
	var buf bytes.Buffer
	var header [4]byte
	binary.BigEndian.PutUint32(header[:], 3)
	buf.Write(header[:])
	buf.WriteByte('x')
	if _, err := ReadFrame(&buf); err == nil {
		t.Fatal("expected truncated payload error")
	}
}
