package filesync

import (
	"context"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
)

type Action struct {
	Path string
	Kind string
}

func Sync(ctx context.Context, sourceRoot, targetRoot string, dryRun bool) ([]Action, error) {
	source, err := scan(sourceRoot)
	if err != nil {
		return nil, fmt.Errorf("scan source: %w", err)
	}
	target, err := scan(targetRoot)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, fmt.Errorf("scan target: %w", err)
		}
		target = map[string]string{}
	}
	var actions []Action
	for path, hash := range source {
		if old, ok := target[path]; !ok {
			actions = append(actions, Action{Path: path, Kind: "copy"})
		} else if old != hash {
			actions = append(actions, Action{Path: path, Kind: "update"})
		}
	}
	sort.Slice(actions, func(i, j int) bool { return actions[i].Path < actions[j].Path })
	if dryRun {
		return actions, nil
	}
	for _, action := range actions {
		if err := ctx.Err(); err != nil {
			return actions, err
		}
		if err := copyAtomic(ctx, sourceRoot, targetRoot, action.Path); err != nil {
			return actions, fmt.Errorf("%s %q: %w", action.Kind, action.Path, err)
		}
	}
	return actions, nil
}

func scan(root string) (map[string]string, error) {
	hashes := make(map[string]string)
	err := filepath.WalkDir(root, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}
		info, err := d.Info()
		if err != nil || !info.Mode().IsRegular() {
			return err
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		hash := sha256.New()
		_, copyErr := io.Copy(hash, file)
		closeErr := file.Close()
		if copyErr != nil {
			return copyErr
		}
		if closeErr != nil {
			return closeErr
		}
		relative, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}
		hashes[filepath.ToSlash(relative)] = fmt.Sprintf("%x", hash.Sum(nil))
		return nil
	})
	return hashes, err
}

func copyAtomic(ctx context.Context, sourceRoot, targetRoot, relative string) error {
	clean := filepath.Clean(filepath.FromSlash(relative))
	if clean == "." || filepath.IsAbs(clean) || clean == ".." || len(clean) >= 3 && clean[:3] == ".."+string(filepath.Separator) {
		return fmt.Errorf("unsafe relative path")
	}
	source, err := os.Open(filepath.Join(sourceRoot, clean))
	if err != nil {
		return err
	}
	defer source.Close()
	target := filepath.Join(targetRoot, clean)
	if err := os.MkdirAll(filepath.Dir(target), 0o755); err != nil {
		return err
	}
	tmp, err := os.CreateTemp(filepath.Dir(target), ".filesync-*")
	if err != nil {
		return err
	}
	name := tmp.Name()
	defer os.Remove(name)
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
	return os.Rename(name, target)
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
