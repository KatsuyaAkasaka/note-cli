package utils

import (
	"os"
	"strings"
)

func AbsolutePath(path string) string {
	homeDir, _ := os.UserHomeDir()

	return strings.ReplaceAll(path, "$HOME", homeDir)
}
