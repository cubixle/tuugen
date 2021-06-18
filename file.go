package main

import (
	"os"
	"path/filepath"
)

func createFile(path string) (*os.File, error) {
	fp := filepath.Dir(path)
	if err := os.MkdirAll(fp, 0777); err != nil {
		return nil, err
	}

	f, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return f, nil
}
