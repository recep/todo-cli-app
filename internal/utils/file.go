package utils

import (
	"log"
	"os"
	"strings"
)

func SaveDataToFile(data, path string) error {
	// This code appends string to the file "todos.txt". It creates the file if it does not already exist.
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write new task to file
	_, err = f.WriteString(data + ".")
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}

// TODO refactor this func
func UpdateFile(todos []string, path string) error {
	if err := os.Remove(path); err != nil {
		return err
	}

	f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	for _, t := range todos {
		if t == "" {
			continue
		}
		_, err := f.WriteString(t + ".")
		if err != nil {
			return err
		}
	}

	return nil
}

func ReadData(path string) ([]string, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	s := strings.Split(string(bytes), ".")

	return s, err
}
