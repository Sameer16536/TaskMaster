# Feature Checklist

## 1. Project Setup
- [✅] Initialize Go module (`go mod init`)
- [✅] Create project structure

## 2. Data Structure (The "Model")
- [✅] Define a `Task` struct with fields:
	- ID (int)
	- Name (string)
	- Complete (bool)
- [✅] Define a slice to hold multiple tasks

## 3. The "Add" Command
- [✅] Read command line arguments (`os.Args`)
- [✅] If command is `add`, parse the task name
- [✅] Create a new `Task` struct
- [✅] **Hard Part:**
	- Read existing JSON file
	- Decode to slice
	- Append new task
	- Encode back to JSON
	- Write to file

## 4. The "List" Command
- [✅] Read `tasks.json`
- [✅] Loop through the slice
- [✅] Print tasks in a formatted way (e.g., `[ ] 1. Buy Milk`, `[x] 2. Learn Go`)

## 5. The "Complete" Command
- [ ] Accept a Task ID (e.g., `./taskmaster complete 1`)
- [ ] Find the task in the slice
- [ ] Toggle `Complete` to true
- [ ] Save back to file

## 6. The "Delete" Command
- [ ] Accept a Task ID
- [ ] Remove the element from the slice (learn how to delete from slices in Go)
- [ ] Save back to file