package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type TaskStatus string

const (
	StatusTodo       TaskStatus = "todo"
	StatusInProgress TaskStatus = "in-progress"
	StatusDone       TaskStatus = "done"
)

type Task struct {
	ID          string     `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func readFile() []Task {
	if _, err := os.Stat("tasks.json"); os.IsNotExist(err) {
		return []Task{}
	}

	t, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Error reading task.json:", err)
		return []Task{}
	}

	if len(t) == 0 {
		return []Task{}
	}

	tasks := []Task{}
	err = json.Unmarshal(t, &tasks)
	if err != nil {
		fmt.Println("Error unmarshalling tasks:", err)
		return []Task{}
	}

	return tasks
}

func saveToFile(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling tasks:", err)
		return err
	}

	return os.WriteFile("tasks.json", data, 0644)
}

func getAllTask() []Task {
	tasks := readFile()
	return tasks
}

func getTasksByStatus(status TaskStatus) []Task {
	tasks := readFile()

	var result []Task
	for _, task := range tasks {
		if task.Status == status {
			result = append(result, task)
		}
	}

	return result
}

func NewTask(id string, description string) Task {
	now := time.Now()
	return Task{
		ID:          id,
		Description: description,
		Status:      StatusTodo,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func addTask(taskDescription string) {
	tasks := readFile()
	newID := fmt.Sprintf("%d", len(tasks)+1)
	task := NewTask(newID, taskDescription)
	tasks = append(tasks, task)

	if err := saveToFile(tasks); err != nil {
		fmt.Println("Error saving to file:", err)
		return
	}

	fmt.Println("Task added successfully (ID:", task.ID, ")")
}

func updateTaskDescription(taskID string, newDescription string) {
	tasks := readFile()
	found := false
	for i := range tasks {
		if tasks[i].ID == taskID {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Task ID not found:", taskID)
		return
	}

	err := saveToFile(tasks)
	if err != nil {
		fmt.Println("Error saving task description:", err)
	}
}

func printTasksHeader() {
	fmt.Printf("%-20s | %-30s | %-12s | %-20s\n", "ID", "Description", "Status", "Updated At")
	fmt.Println("---------------------+--------------------------------+--------------+---------------------")
}

func printTaskRow(t Task) {
	fmt.Printf("%-20s | %-30s | %-12s | %-20s\n",
		t.ID, t.Description, t.Status, t.UpdatedAt.Format("2006-01-02 15:04:05"))
}

func updateTaskStatus(taskID string, newStatus TaskStatus) {
	tasks := readFile()
	found := false
	for i := range tasks {
		if tasks[i].ID == taskID {
			tasks[i].Status = newStatus
			tasks[i].UpdatedAt = time.Now()
			found = true
			break
		}
	}

	if !found {
		fmt.Println("Task ID not found:", taskID)
		return
	}

	err := saveToFile(tasks)
	if err != nil {
		fmt.Println("Error saving task status:", err)
	}
}

func deleteTask(taskID string) {
	tasks := readFile()
	var newTasks []Task
	found := false
	for _, task := range tasks {
		if task.ID != taskID {
			newTasks = append(newTasks, task)
		} else {
			found = true
		}
	}

	if !found {
		fmt.Println("Task ID not found:", taskID)
		return
	}

	err := saveToFile(newTasks)
	if err != nil {
		fmt.Println("Error deleting task:", err)
	}
}
