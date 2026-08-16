package main

import (
	"fmt"
	"os"
	"strconv"
)

type Command string
type Status string

const (
	Add            Command = "add"
	Update         Command = "update"
	Delete         Command = "delete"
	MarkInProgress Command = "mark-in-progress"
	MarkDone       Command = "mark-done"
	List           Command = "list"
	Done           Status  = "done"
	Todo           Status  = "todo"
	InProgress     Status  = "in-progress"
)

const fileName = "tasks.json"

// Core task struct
type Task struct {
	Id     uint   `json:"id"`
	Status Status `json:"status"`
	Text   string `json:"text"`
}

// main is the entry point of the program.
func main() {

	tasks := []Task{}
	userArgs := os.Args[1:]
	cmd := userArgs[0]
	// fmt.Printf("User arguments: %v\n", userArgs)

	tasks, err := readFile()
	if err != nil {
		writeFile(&tasks)
	}

	switch Command(cmd) {
	case Add:
		text := userArgs[1]
		addTask(&tasks, text)
	case Update:
		id, err := strconv.Atoi(userArgs[1])
		if err != nil {
			fmt.Printf("Invalid parameter: %v\n", err)
			return
		}
		text := userArgs[2]
		err2 := updateTask(&tasks, id, text)
		if err2 != nil {
			fmt.Printf("Error: %v\n", err2)
			return
		}
	case Delete:
		id, err := strconv.Atoi(userArgs[1])
		if err != nil {
			fmt.Printf("Invalid parameter: %v\n", err)
			return
		}
		err2 := deleteTask(&tasks, id)
		if err2 != nil {
			fmt.Printf("Error: %v\n", err2)
			return
		}
	case MarkInProgress:
		id, err := strconv.Atoi(userArgs[1])
		if err != nil {
			fmt.Printf("Invalid parameter: %v\n", err)
			return
		}
		err2 := setStatus(&tasks, InProgress, id)
		if err2 != nil {
			fmt.Printf("Error: %v\n", err2)
			return
		}
	case MarkDone:
		id, err := strconv.Atoi(userArgs[1])
		if err != nil {
			fmt.Printf("Invalid parameter: %v\n", err)
			return
		}
		err2 := setStatus(&tasks, Done, id)
		if err2 != nil {
			fmt.Printf("Error: %v\n", err2)
			return
		}
	case List:
		var text Status
		if len(userArgs) == 2 {
			text = Status(userArgs[1])
		}
		listItems(&tasks, text)
	default:
		fmt.Print("Command not supported")
		return
	}

	// fmt.Println(tasks)

	writeFile(&tasks)
}
