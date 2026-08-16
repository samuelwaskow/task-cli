package main

import (
	"encoding/json"
	"os"
)

// Write the task file to a local JSON file
func writeFile(tasks *[]Task) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	err = encoder.Encode(tasks)
	if err != nil {
		return err
	}
	return nil
}

// readFile reads the tasks from the JSON file and returns them as a slice of Task structs.
func readFile() ([]Task, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
