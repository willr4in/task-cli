package main

import (
	"encoding/json"
	"errors"
	"os"
)

const tasksFile = "tasks.json"

func loadTasks() ([]Task, error) {
	data, err := os.ReadFile(tasksFile)
	if errors.Is(err, os.ErrNotExist) {
		return []Task{}, nil
	}
	if err != nil {
		return nil, err
	}

	var tasks []Task
	if err := json.Unmarshal(data, &tasks); err != nil {
		return nil, err
	}

	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(tasksFile, data, 0644)
}
