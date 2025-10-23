package envloader

import (
	"os"
	"path"
)

func GetEnvPath() (string, error) {
	directory, err := os.Getwd()
	if err != nil {
		return "", err
	}

	filepath := searchUp(directory, ".env")
	return filepath, nil
}

func searchUp(dir string, filename string) string {
	if dir == "/" || dir == "" || dir == "." {
		return ""
	}

	if _, err := os.Stat(path.Join(dir, filename)); err == nil {
		return path.Join(dir, filename)
	}

	return searchUp(path.Dir(dir), filename)
}
