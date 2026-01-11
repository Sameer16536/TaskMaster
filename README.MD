# TaskMaster

TaskMaster is a simple command-line task management tool written in Go. It allows you to add, list, complete, and delete tasks, storing them in a local JSON file.

## Project Setup

1. **Clone the repository** :
   ```sh
   git clone <repo-url>
   cd TaskMaster
   ```

2. **Build the project:**
   ```sh
   go build -o taskmaster
   ```

3. **Run commands:**
   - Add a task:
     ```sh
     ./taskmaster add "Your task description"
     ```
   - List all tasks:
     ```sh
     ./taskmaster list
     ```
   - Mark a task as complete:
     ```sh
     ./taskmaster complete <task_id>
     ```
   - Delete a task:
     ```sh
     ./taskmaster delete <task_id>
     ```

## Explanation

TaskMaster is designed for quick and easy task management from the terminal. Tasks are stored in a `tasks.json` file in the project directory. Each task has an ID, description, completion status, and creation timestamp.

### Features
- **Add Tasks:** Add new tasks with a description.
- **List Tasks:** View all tasks with their status and IDs.
- **Complete Tasks:** Mark tasks as done by their ID.
- **Delete Tasks:** Remove tasks by their ID.

### File Overview
- `main.go`: Entry point, handles command-line arguments and routes commands.
- `storage.go`: Handles reading/writing tasks to `tasks.json` and implements task operations.
- `tasks.go`: Defines the `Task` struct.
- `tasks.json`: Stores your tasks (auto-created).
- `go.mod`: Go module definition.

---

Enjoy managing your tasks with TaskMaster!