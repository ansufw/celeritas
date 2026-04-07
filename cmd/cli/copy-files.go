package main

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/ansufw/celeritas/cmd/cli/templates"
)

var templateFS = templates.Templates

func copyFilefromTemplate(templatePath string, destinationPath string) error {

	if fileExists(destinationPath) {
		return errors.New(destinationPath + " file already exists")
	}

	data, err := templateFS.ReadFile(templatePath)
	if err != nil {
		exitGracefully(err)
	}

	err = copyDataToFile(data, destinationPath)
	if err != nil {
		exitGracefully(err)
	}

	return nil
}

func copyDataToFile(data []byte, to string) error {
	// Create directory if it doesn't exist
	dir := filepath.Dir(to)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}
	return os.WriteFile(to, data, 0644)
}

func fileExists(fileToCheck string) bool {
	if _, err := os.Stat(fileToCheck); os.IsNotExist(err) {
		return false
	}
	return true
}
