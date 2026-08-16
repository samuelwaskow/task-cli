# Task Tracker CLI

A small command-line task tracker written in Go. It stores tasks in a local `tasks.json` file and lets you add, update, delete, and filter tasks by status.

## Features

- Add new tasks
- Update task text
- Delete tasks
- Mark tasks as `todo`, `in-progress`, or `done`
- List tasks, optionally filtered by status
- Persists tasks to a local JSON file

## Requirements

- Go 1.20 or newer

## Build

From the project root:

```bash
go build -o task-cli ./cmd/task-cli
```

This creates a binary named `task-cli` in the project root.

## Run

```bash
./task-cli <command> [arguments]
```

## Commands

### Add a task

```bash
./task-cli add "Buy groceries"
```

Creates a new task with default status `todo`.

### Update a task

```bash
./task-cli update 1 "Buy groceries and fruit"
```

Replaces the text for task `1`.

### Delete a task

```bash
./task-cli delete 1
```

Deletes task `1`.

### Mark task as in progress

```bash
./task-cli mark-in-progress 1
```

### Mark task as done

```bash
./task-cli mark-done 1
```

### List tasks

Show all tasks:

```bash
./task-cli list
```

Show only tasks with a specific status:

```bash
./task-cli list done
./task-cli list todo
./task-cli list in-progress
```

## Supported statuses

- `todo`
- `in-progress`
- `done`

## Storage

Tasks are saved to a file named `tasks.json` in the directory where the command is run.

Example:

```json
[
  {
    "id": 1,
    "status": "todo",
    "text": "Buy groceries"
  }
]
```

## Example workflow

```bash
go build -o task-cli ./cmd/task-cli
./task-cli add "Finish project proposal"
./task-cli add "Book dentist appointment"
./task-cli list
./task-cli mark-in-progress 1
./task-cli mark-done 1
./task-cli update 2 "Book dentist appointment for Friday"
./task-cli delete 2
./task-cli list
```

## Notes

- If the task file does not exist yet, the CLI will create one automatically.
- The command expects valid IDs and will print an error if the ID is invalid.
