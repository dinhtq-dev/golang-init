package fileutil

import (
	"fmt"
	"io/ioutil"
	"os"
)

// ReadFile reads the content of a file and returns it as a string
func ReadFile(filePath string) (string, error) {
	data, err := ioutil.ReadFile(filePath)
	
	if err != nil {
		return "", fmt.Errorf("failed to read file: %v", err)
	}

	return string(data), nil
}

// WriteFile writes content to a file
func WriteFile(filePath, content string) error {
	err := ioutil.WriteFile(filePath, []byte(content), 0644)

	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}

	return nil
}

// FileExists checks if a file exists
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
