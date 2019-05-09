package cache

import (
	"io"
	"os"
	"path/filepath"

	"github.com/knqyf263/fanal/utils"

	"golang.org/x/xerrors"
)

var (
	cacheDir = utils.CacheDir()
)

func Get(key string) io.Reader {
	filePath := filepath.Join(cacheDir, key)
	f, err := os.Open(filePath)
	if err != nil {
		return nil
	}
	return f
}

func Set(key string, file io.Reader) (io.Reader, error) {
	filePath := filepath.Join(cacheDir, key)
	if err := os.MkdirAll(cacheDir, os.ModePerm); err != nil {
		return nil, xerrors.Errorf("failed to mkdir all: %w", err)
	}
	cacheFile, err := os.Create(filePath)
	if err != nil {
		return file, xerrors.Errorf("failed to create cache file: %w", err)
	}

	tee := io.TeeReader(file, cacheFile)
	return tee, nil
}

func Clear() error {
	if err := os.RemoveAll(utils.CacheDir()); err != nil {
		return xerrors.New("failed to remove cache")
	}
	return nil
}