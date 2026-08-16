package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

// Adds a new task to the list
func addTask(tasks *[]Task, text string) {

	var id uint = 1
	if len(*tasks) > 0 {
		id = (*tasks)[len(*tasks)-1].Id + 1
	}

	t := Task{
		Id:     id,
		Status: Todo,
		Text:   text,
	}

	*tasks = append(*tasks, t)
}

// Updates a task's test given the provided ID
func updateTask(tasks *[]Task, id int, text string) error {

	index, err := search(tasks, id)
	if err != nil {
		return fmt.Errorf("Invalid ID: %v", id)
	} else {
		t := &(*tasks)[index]
		t.Text = text
		return nil
	}
}

// Deletes a task
func deleteTask(tasks *[]Task, id int) error {

	index, err := search(tasks, id)
	if err != nil {
		return fmt.Errorf("Invalid ID: %v", id)
	} else {
		*tasks = append((*tasks)[:index], (*tasks)[index+1:]...)
		return nil
	}
}

// Set status for a given task
func setStatus(tasks *[]Task, status Status, id int) error {

	index, err := search(tasks, id)
	if err != nil {
		return fmt.Errorf("Invalid ID: %v", id)
	} else {
		t := &(*tasks)[index]
		t.Status = status
		return nil
	}
}

func listItems(tasks *[]Task, filter Status) {

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tSTATUS\tTEXT")

	for i := 0; i < len(*tasks); i++ {
		t := &(*tasks)[i]
		if filter == "" || t.Status == filter {
			fmt.Fprintf(w, "%v\t%v\t%v\n", t.Id, t.Status, t.Text)
		}
	}
	w.Flush()
}

// Retrieve an index given an Id
func search(tasks *[]Task, id int) (int, error) {
	searchIndex := -1
	for i := 0; i < len(*tasks); i++ {
		t := &(*tasks)[i]
		if t.Id == uint(id) {
			searchIndex = i
			break
		}
	}

	if searchIndex == -1 {
		return 0, fmt.Errorf("Invalid ID: %v", id)
	} else {
		return searchIndex, nil
	}
}
