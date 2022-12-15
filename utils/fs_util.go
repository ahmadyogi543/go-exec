package utils

import (
	"os"
	"path/filepath"
)

func WriteFile(dirname string, filename string, content string) (string, error) {
	_, err := os.Stat(dirname)

	if err != nil {
		os.MkdirAll(dirname, os.ModePerm)
	}

	codePath := filepath.Join(dirname, filename)
	err = os.WriteFile(codePath, []byte(content), 0644)
	return codePath, err
}

func RemoveFile(path string) error {
	err := os.Remove(path)
	return err
}
