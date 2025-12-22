package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func menu() {
	fmt.Print("task-cli > ")
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}

	input := scanner.Text()
	parts := strings.Fields(input)

	if len(parts) == 0 {
		return
	}

	command := parts[0]
	args := parts[1:]

	switch command {
	case "add":
		if len(args) < 1 {
			fmt.Println("Usage: add \"task description\"")
			return
		}
		description := strings.Join(args, " ")
		addTask(description)

	case "list":
		var tasks []Task
		if len(args) > 0 {
			status := TaskStatus(args[0])
			tasks = getTasksByStatus(status)
		} else {
			tasks = getAllTask()
		}

		if len(tasks) == 0 {
			fmt.Println("No tasks found.")
			return
		}

		printTasksHeader()
		for _, task := range tasks {
			printTaskRow(task)
		}

	case "update":
		if len(args) < 2 {
			fmt.Println("Usage: update [ID] \"new description\"")
			return
		}
		taskID := args[0]
		description := strings.Join(args[1:], " ")
		updateTaskDescription(taskID, description)

	case "delete":
		if len(args) < 1 {
			fmt.Println("Usage: delete [ID]")
			return
		}
		deleteTask(args[0])

	case "mark-in-progress":
		if len(args) < 1 {
			fmt.Println("Usage: mark-in-progress [ID]")
			return
		}
		updateTaskStatus(args[0], StatusInProgress)

	case "mark-done":
		if len(args) < 1 {
			fmt.Println("Usage: mark-done [ID]")
			return
		}
		updateTaskStatus(args[0], StatusDone)

	case "exit", "quit":
		os.Exit(0)

	default:
		fmt.Println("Unknown command. Available: add, list, update, delete, mark-in-progress, mark-done, exit")
	}
}
