# Task Tracker CLI

A simple command-line interface (CLI) application written in Go to manage your tasks (to-do list). Tasks are stored locally in a JSON file and can be added, listed, updated, marked as done or in progress, and deleted.

This project was built as a learning exercise to practice:

* Go basics and project structure
* Working with files and JSON
* Building CLI applications
* Handling user input and edge cases
* Writing unit tests

---

## Features

* Add new tasks
* List all tasks
* Filter tasks by status (`todo`, `in-progress`, `done`)
* Update task descriptions
* Mark tasks as done or in progress
* Delete tasks by ID
* Persistent storage using a local `tasks.json` file

---

## Task Structure

Each task has the following fields:

* `id` – unique task identifier
* `description` – short task description
* `status` – `todo`, `in-progress`, or `done`
* `createdAt` – date and time when the task was created
* `updatedAt` – date and time of the last update

---

## Requirements

* Go 1.20 or later
* No external libraries (standard library only)

---

## Installation

Clone the repository:

```
git clone <your-repository-url>
cd task-tracker
```

Build the application:

```
go build -o task-cli
```

Or run directly:

```
go run main.go <command>
```

---

## Usage

```
task-cli <command> [arguments]
```

### Commands

#### Add a task

```
task-cli add "Buy groceries"
```

#### List all tasks

```
task-cli list
```

#### List tasks by status

```
task-cli list todo
task-cli list in-progress
task-cli list done
```

#### Update a task description

```
task-cli update 1 "Buy milk and bread"
```

#### Mark a task as done

```
task-cli mark-done 1
```

#### Mark a task as in progress

```
task-cli mark-in-progress 1
```

#### Delete a task

```
task-cli delete 1
```

#### Show help

```
task-cli help
```

---

## Data Storage

Tasks are stored in a local file named `tasks.json` in the project directory.

* The file is created automatically if it does not exist
* All changes are persisted immediately

---

## Testing

Unit tests are included for core helper functions.

Run tests with:

```
go test
```

---

## Project Structure

```
.
├── main.go        # Entry point and command routing
├── tasks.json     # Task storage (created automatically)
├── main_test.go   # Unit tests
└── README.md
```

---

## Notes

* Task IDs are never reused after deletion
* The application uses only Go's standard library
* Designed as a learning project and foundation for more advanced CLI tools

---

## License

This project is provided for educational purposes and is free to use and modify.

---

Happy coding! 🚀
