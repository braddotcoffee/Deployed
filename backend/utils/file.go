package utils

import (
	"log"
	"os"
)

// WriteExistingFile opens an existing file
// and overwrites its contents
func WriteExistingFile(filePath string, contents string) error {
	info, err := os.Stat(filePath)
	if err != nil {
		log.Fatalf("Failed to stat file %s: %s", filePath, err.Error())
		return err
	}

	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, info.Mode())
	if err != nil {
		log.Fatalf("Failed to open file %s: %s", filePath, err.Error())
		return err
	}

	defer f.Close()
	f.WriteString(contents)
	return nil
}
