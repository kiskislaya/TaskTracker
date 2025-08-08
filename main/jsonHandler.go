package main

import (
	"encoding/json"
	"os"
)

func encodeTask(tasks []Task, filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	t, err := json.Marshal(tasks)
	if err != nil {
		return err
	}
	_, err = file.Write(t)
	if err != nil {
		return err
	}

	return nil
}

func decodeTasks(tasks *[]Task, filename string) ([]Task, error) {
	file, err := os.OpenFile(filename,
		os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(tasks)
	if err != nil {
		return nil, err
	}
	return *tasks, nil
}
