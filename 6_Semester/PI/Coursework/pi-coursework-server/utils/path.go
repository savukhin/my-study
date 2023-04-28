package utils

import (
	"os"
	"path/filepath"
)

func GetExecutablePath() string {
	ex, _ := os.Executable()

	return filepath.Dir(ex)
}
