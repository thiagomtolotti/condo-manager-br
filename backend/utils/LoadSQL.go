package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func LoadSQL(filename string) (string, error) {
	path := filepath.Join("models/queries", filename)

	content, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("error while loading %v: %v", filename, err)
	}

	return string(content), nil
}
