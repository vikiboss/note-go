package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const MaxFrameSize = 1 << 20

func WriteFrame(w io.Writer, payload []byte) error {
	if len(payload) > MaxFrameSize {
		return fmt.Errorf("frame size %d exceeds limit %d", len(payload), MaxFrameSize)
	}
	var header [4]byte
	binary.BigEndian.PutUint32(header[:], uint32(len(payload)))
	if err := writeAll(w, header[:]); err != nil {
		return fmt.Errorf("write header: %w", err)
	}
	if err := writeAll(w, payload); err != nil {
		return fmt.Errorf("write payload: %w", err)
	}
	return nil
}

func writeAll(w io.Writer, data []byte) error {
	for len(data) > 0 {
		n, err := w.Write(data)
		if err != nil {
			return err
		}
		if n == 0 {
			return io.ErrShortWrite
		}
		data = data[n:]
	}
	return nil
}

func ReadFrame(r io.Reader) ([]byte, error) {
	var header [4]byte
	if _, err := io.ReadFull(r, header[:]); err != nil {
		return nil, fmt.Errorf("read header: %w", err)
	}
	size := binary.BigEndian.Uint32(header[:])
	if size > MaxFrameSize {
		return nil, fmt.Errorf("frame size %d exceeds limit %d", size, MaxFrameSize)
	}
	payload := make([]byte, size)
	if _, err := io.ReadFull(r, payload); err != nil {
		return nil, fmt.Errorf("read payload: %w", err)
	}
	return payload, nil
}

func main() {
	server, client := net.Pipe()
	go func() {
		defer server.Close()
		_ = WriteFrame(server, []byte("hello over a byte stream"))
	}()
	defer client.Close()
	message, err := ReadFrame(client)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(message))
}
