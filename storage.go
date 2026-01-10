package main

import (
	"encoding/json"
	"os"
)

func saveTasks(tasks []Task) error {
	//Marshal the tasks to JSON and save to file
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile("tasks.json", data, 0644)
}

func loadTasks(tasks []Task) error {
	//Check if file exists
	if _, err := os.Stat(`tasks.json`); os.IsNotExist(err) {
		return nil
	}

	data, err := os.ReadFile(`tasks.json`)
	if err != nil {
		return err
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return err
	}

	return nil

}

func deleteTask() {}

func updateTask() {}
