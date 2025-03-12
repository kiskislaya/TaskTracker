package main

import (
	"time"
)

func addTask(description string) (int64, error) {
	// Create a new task
	task := Task{
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now().Format(time.RFC3339),
		UpdatedAt:   time.Now().Format(time.RFC3339),
	}

	return id, nil
}

func deleteTask(id int64) error {
	// Delete the task
	return nil
}

func updateTask(id int64, description string) error {
	// Update the task
	return nil
}

func markTaskAsDone(id int64) error {
	// Mark the task as done
	return nil
}

func markTaskAsInProgress(id int64) error {
	// Mark the task as in progress
	return nil
}

func getAllTasks() ([]Task, error) {
	// Get all tasks
	return nil, nil
}

func getTasksAsDone() ([]Task, error) {
	// Get all tasks as done
	return nil, nil
}

func getTasksAsInProgress() ([]Task, error) {
	// Get all tasks as in progress
	return nil, nil
}

func getTasksAsToDo() ([]Task, error) {
	// Get all tasks as to do
	return nil, nil
}
