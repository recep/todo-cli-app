package utils

import (
	"os"
)

func SaveDataToFile(data []byte, path string) error {
	// TODO refactor and fix end of line error
	if err := os.Remove(path) ; err != nil {
		return err
	}

	// Open json file
	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write new task to file
	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}

// Read data from the file
func ReadData(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, err
}