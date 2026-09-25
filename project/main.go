package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		os.Exit(1)
	}
	cmd := os.Args[1]
	args := os.Args[2:]

	var err error

	switch cmd {
	case "add":
		if len(args) < 1 {
			err = fmt.Errorf("usage: task-cli add <description>")
			break
		}

		err = addTask(args[0])

	case "update":
		if len(args) < 2 {
			err = fmt.Errorf("usage: task-cli update <id> <description>")
			break
		}

		id, perr := strconv.Atoi(args[0])
		if perr != nil {
			err = fmt.Errorf("invalid id")
			break
		}

		err = updateTask(id, args[1])

	case "delete":
		if len(args) < 1 {
			err = fmt.Errorf("usage: task-cli delete <id>")
			break
		}

		id, perr := strconv.Atoi(args[0])
		if perr != nil {
			err = fmt.Errorf("invalid id")
			break
		}

		err = deleteTask(id)

	case "mark-in-progress":
		if len(args) < 1 {
			err = fmt.Errorf("usage: task-cli mark-in-progress <id>")
			break
		}

		id, perr := strconv.Atoi(args[0])
		if perr != nil {
			err = fmt.Errorf("invalid id")
			break
		}

		err = markTask(id, StatusInProgress)

	case "mark-done":
		if len(args) < 1 {
			err = fmt.Errorf("usage: task-cli mark-done <id>")
			break
		}

		id, perr := strconv.Atoi(args[0])
		if perr != nil {
			err = fmt.Errorf("invalid id")
			break
		}

		err = markTask(id, StatusDone)

	case "list":
		if len(args) == 0 {
			err = listTasks("")
			break
		}

		s := Status(args[0])
		if s != StatusDone && s != StatusInProgress && s != StatusTodo {
			err = fmt.Errorf("invalid status")
			break
		}

		err = listTasks(s)

	default:
		fmt.Println("Unknown command:", cmd)
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
