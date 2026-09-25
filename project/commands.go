package main

import (
	"fmt"
	"time"
)

func nextID(tasks []Task) int {
	maxID := 0
	for _, id := range tasks {
		if id.ID > maxID {
			maxID = id.ID
		}
	}

	return maxID + 1
}

func findIndex(tasks []Task, id int) int {
	for i := range tasks {
		if tasks[i].ID == id {
			return i
		}
	}

	return -1
}

func addTask(desc string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	now := time.Now()
	task := Task{
		ID:          nextID(tasks),
		Description: desc,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	tasks = append(tasks, task)

	if err := saveTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
	return nil
}

func updateTask(id int, desc string) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	i := findIndex(tasks, id)
	if i == -1 {
		return fmt.Errorf("task with id %d is not found", id)
	}

	tasks[i].Description = desc
	tasks[i].UpdatedAt = time.Now()

	if err := saveTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("Task %d updated successfully\n", id)
	return nil
}

func deleteTask(id int) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	i := findIndex(tasks, id)
	if i == -1 {
		return fmt.Errorf("task with id %d is not found", id)
	}

	tasks = append(tasks[:i], tasks[i+1:]...)

	if err := saveTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("Task %d deleted successfully\n", id)
	return nil
}

func markTask(id int, status Status) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	i := findIndex(tasks, id)
	if i == -1 {
		return fmt.Errorf("task with id %d is not found", id)
	}

	tasks[i].Status = status
	tasks[i].UpdatedAt = time.Now()

	if err := saveTasks(tasks); err != nil {
		return err
	}

	fmt.Printf("Task %d marked as %s\n", id, status)
	return nil
}

func listTasks(filter Status) error {
	tasks, err := loadTasks()
	if err != nil {
		return err
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks yet")
		return nil
	}

	printed := 0
	for _, t := range tasks {
		if filter != "" && t.Status != filter {
			continue
		}
		fmt.Printf("[%d] %s (%s)\n", t.ID, t.Description, t.Status)
		printed++
	}

	if printed == 0 {
		fmt.Println("No tasks with this status")
	}

	return nil
}
