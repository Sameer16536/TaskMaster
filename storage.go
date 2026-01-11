package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func saveTasks(tasks []Task) error {
	//Marshal the tasks to JSON and save to file
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

func loadTasks() ([]Task, error) {
	//Check if file exists
	if _, err := os.Stat(`tasks.json`); os.IsNotExist(err) {
		return []Task{}, nil
	}

	//Read the file
	data, err := os.ReadFile(`tasks.json`)
	if err != nil {
		return nil, err
	}

	//Unmarshal the JSON data
	tasks := []Task{}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil

}

func addTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Error: Missing task description")
		return
	}

	// Recombine arguments into a single string
	description := strings.Join(args, " ")

	// Load existing tasks from file
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	// Create the new task
	newTask := Task{
		ID:          len(tasks) + 1,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
	}

	// Append to slice
	tasks = append(tasks, newTask)

	// Save back to file
	err = saveTasks(tasks)
	if err != nil {
		fmt.Printf("Error saving task: %v\n", err)
		return
	}

	fmt.Printf("✅ Task added: %s (ID: %d)\n", description, newTask.ID)
}
func deleteTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Missing task ID")
		return
	}
	//Convert ID from string to int
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}
	//Load tasks
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}
	//Find and delete task
	found := false
	for i, task := range tasks {
		if task.ID == id {
			//Remove task from slice
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		fmt.Printf("Task with ID %d not found\n", id)
		return
	}
	//Save tasks
	err = saveTasks(tasks)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		return
	}
	fmt.Printf("✅ Task with ID %d deleted\n", id)

}

func completeTask(args []string) {
	if len(args) == 0 {
		fmt.Println("Missing task ID")
		return
	}

	//Convert ID from string to int
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Invalid task ID")
		return
	}

	//Load tasks
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	//Find and update task
	found := false
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			tasks[i].Description = fmt.Sprintf("%s (Done)", task.Description)
			found = true
			break
		}
	}

	if !found {
		fmt.Printf("Task with ID %d not found\n", id)
		return
	}

	//Save tasks
	err = saveTasks(tasks)
	if err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
		return
	}

}

func listTasks() {
	tasks, err := loadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}

	//Print header
	fmt.Println("ID  | Status | Task")
	fmt.Println("----|--------|----------------")

	//Loop and print
	for _, task := range tasks {
		status := " "
		if task.Completed {
			status = "✅"
		}
		// %-3d = Integer, padded to 3 spaces, left aligned
		// %s   = String
		fmt.Printf("%-3d |   [%s]  | %s\n", task.ID, status, task.Description)
	}
}
