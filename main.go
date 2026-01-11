package main

import (
	"fmt"
	"os"
)

func main() {
	// SAFETY CHECK: Ensure the user typed a command
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		// Pass everything after "add" to the function
		addTask(os.Args[2:])
	case "list":
		listTasks()
	case "complete":
		completeTask(os.Args[2:])
	case "delete":
		deleteTask(os.Args[2:])
	default:
		printUsage()
	}
}

func printUsage() {
	fmt.Println("Welcome to TaskMaster CLI")
	fmt.Println("Usage:")
	fmt.Println("  add [task name]   -> Add a new task")
	fmt.Println("  list              -> List all tasks")
	fmt.Println("  complete [id]     -> Mark a task as done")
	fmt.Println("  delete [id]       -> Remove a task")
}
