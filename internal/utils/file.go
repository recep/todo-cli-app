package utils

import (
	"log"
	"os"
)

func SaveDataToFile(data string) error {
	// This code appends string to the file "todos.txt". It creates the file if it does not already exist.
	f, err := os.OpenFile("./storage/todos.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write new task to file
	_, err = f.WriteString(data + "\n")
	if err != nil {
		log.Fatalln(err)
	}

	return nil
}
