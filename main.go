package main

import (
	"fmt"
	"os"
)

type Task struct {
	ID        int
	Name      string
	Completed bool
}

var tasks []Task

func main() {
	loadTasks()
	if len(os.Args) < 2 {
		fmt.Println("Usage: todo [add|list|complete|delete]")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task name")
			return
		}
		addTask(os.Args[2])
	case "list":
		listTasks()
	case "complete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID")
			return
		}
		completeTask(os.Args[2])
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task ID")
			return
		}
		deleteTask(os.Args[2])
	default:
		fmt.Println("Invalid command")
	}
	saveTasks()
}

func addTask(name string) {
	task := Task{ID: len(tasks) + 1, Name: name, Completed: false}
	tasks = append(tasks, task)
	fmt.Printf("Added task: %s\n", name)
}

func listTasks() {
	for _, task := range tasks {
		status := "Incomplete"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("%d. %s [%s]\n", task.ID, task.Name, status)
	}
}

func completeTask(id string) {
	// Implement task completion logic
}

func deleteTask(id string) {
	// Implement task deletion logic
}

func loadTasks() {
	// Load tasks from a file
}

func saveTasks() {
	// Save tasks to a file
}
