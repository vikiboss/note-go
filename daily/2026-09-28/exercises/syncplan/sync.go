package syncplan

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

type Entry struct {
	Path string
	Hash string
}

type Action struct {
	Path string
	Kind string
}

func Plan(source, target []Entry) []Action {
	targetHash := make(map[string]string, len(target))
	for _, entry := range target {
		targetHash[entry.Path] = entry.Hash
	}
	var actions []Action
	for _, entry := range source {
		if hash, ok := targetHash[entry.Path]; !ok {
			actions = append(actions, Action{Path: entry.Path, Kind: "copy"})
		} else if hash != entry.Hash {
			actions = append(actions, Action{Path: entry.Path, Kind: "update"})
		}
	}
	sort.Slice(actions, func(i, j int) bool { return actions[i].Path < actions[j].Path })
	return actions
}

func Apply(ctx context.Context, sourceRoot, targetRoot string, actions []Action) error {
	for _, action := range actions {
		if err := ctx.Err(); err != nil {
			return err
		}
		if err := copyAtomic(ctx, sourceRoot, targetRoot, action.Path); err != nil {
			return fmt.Errorf("%s %q: %w", action.Kind, action.Path, err)
		}
	}
	return nil
}

func copyAtomic(ctx context.Context, sourceRoot, targetRoot, relative string) error {
	clean := filepath.Clean(filepath.FromSlash(relative))
	if clean == "." || filepath.IsAbs(clean) || clean == ".." || len(clean) >= 3 && clean[:3] == ".."+string(filepath.Separator) {
		return fmt.Errorf("unsafe relative path %q", relative)
	}
	sourcePath := filepath.Join(sourceRoot, clean)
	targetPath := filepath.Join(targetRoot, clean)
	source, err := os.Open(sourcePath)
	if err != nil {
		return err
	}
	defer source.Close()
	if err := os.MkdirAll(filepath.Dir(targetPath), 0o755); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(targetPath), ".sync-*")
	if err != nil {
		return err
	}
	tmpName := tmp.Name()
	defer os.Remove(tmpName)
	if _, err := copyWithContext(ctx, tmp, source); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Sync(); err != nil {
		tmp.Close()
		return err
	}
	if err := tmp.Close(); err != nil {
		return err
	}
	return os.Rename(tmpName, targetPath)
}

func copyWithContext(ctx context.Context, dst io.Writer, src io.Reader) (int64, error) {
	buffer := make([]byte, 32<<10)
	var written int64
	for {
		if err := ctx.Err(); err != nil {
			return written, err
		}
		n, readErr := src.Read(buffer)
		if n > 0 {
			wn, writeErr := dst.Write(buffer[:n])
			written += int64(wn)
			if writeErr != nil {
				return written, writeErr
			}
			if wn != n {
				return written, io.ErrShortWrite
			}
		}
		if readErr != nil {
			if readErr == io.EOF {
				return written, nil
			}
			return written, readErr
		}
		if n == 0 {
			return written, io.ErrNoProgress
		}
	}
}
