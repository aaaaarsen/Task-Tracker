package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
	"strings"
	"strconv"
)

const (
	StatusTodo       = "todo"
	StatusDone       = "done"
	StatusInProgress = "in-progress"
)


type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func loadTasks() []Task{
	var tasks []Task

	fileBytes, err := os.ReadFile("tasks.json")

	if err != nil{
		if os.IsNotExist(err){
			return []Task{}
		} else {
			log.Fatalf("Can't read file: %v", err)
		}
	}

	err = json.Unmarshal(fileBytes, &tasks)
	if err != nil {
		log.Fatalf("failed to unmarshal json data: %v", err)
	}

	return tasks

}

func saveTasks(tasks []Task){
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil{
		log.Fatalf("failed to marshall text: %v", err)
	}

	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		log.Fatalf("failed to write file: %v", err)
	}
}

func printTask(t Task) {
    fmt.Printf(
        "ID: %d\nDescription: %s\nStatus: %s\nCreated at: %s\nUpdated at: %s\n-------\n",
        t.ID,
        t.Description,
        t.Status,
        t.CreatedAt.Format("2006-01-02 15:04"),
        t.UpdatedAt.Format("2006-01-02 15:04"),
    )
}

func findTaskByID(tasks []Task, id int)(*Task, bool){
	for i := range tasks{
		if tasks[i].ID == id {
			return &tasks[i], true
		}
	}

	return nil, false
}

func findTaskIndexByID(tasks []Task, id int)(int, bool){
	for i := range tasks{
		if tasks[i].ID == id {
			return i, true
		}
	}

	return -1, false
}

func handleAdd(cmdArgs []string) {
		if len(cmdArgs) == 0{
			fmt.Println("Error: description is required")
			os.Exit(1)
		}

		description := strings.Join(cmdArgs, " ")

		tasks := loadTasks()
		maxID := 0

		for _, t := range tasks {
			if t.ID > maxID {
				maxID = t.ID
			}
		}

		newTask := Task{
			ID: maxID+1,
			Description: description,
			Status: "todo",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		tasks = append(tasks, newTask)

		saveTasks(tasks)

		fmt.Printf("Task added succesfully! ID: %d\n", newTask.ID)
}

func handleList(cmdArgs []string){
		tasks:= loadTasks()
		found := false

		if len(tasks) == 0{
			fmt.Println("No tasks found")
			os.Exit(1)
		}

		if len(cmdArgs) == 0 {
			for _, t := range tasks {
				printTask(t)
			}

			return
		}

		if len(cmdArgs) == 1 {
			status := cmdArgs[0]

			validStatusses := map[string]bool{
				StatusTodo: true,
				StatusDone: true,
				StatusInProgress: true,
			}

			if !validStatusses[status]{
				fmt.Println("Invalid status. Use: todo | done | in-progress")
				return
			}

			found = false

			for _,t := range tasks{
				if t.Status  == status{
					printTask(t)
					found = true
				}
				
			}

			if !found {
				fmt.Println("No tasks with this status")
			}
			return
		}
}

func handleDelete(cmdArgs []string){
		if len(cmdArgs) < 1{
			fmt.Println("Error: ID is required")
			os.Exit(1)
		} 

		val, err := strconv.Atoi(cmdArgs[0])
		if err != nil {
			fmt.Println("Error: Input is not a valid integer")
			os.Exit(1)
		}
		
		tasks := loadTasks()
		var updatedTasks []Task
		taskIndex, found := findTaskIndexByID(tasks, val)

		if !found {
			fmt.Printf("Task with ID %d not found\n", val)
			os.Exit(1)
		}

		updatedTasks = append(tasks[:taskIndex], tasks[taskIndex+1:]...)

		saveTasks(updatedTasks)
		fmt.Printf("Task with ID %d deleted successfully\n", val)
}

func handleMarkDone(cmdArgs []string){
		if len(cmdArgs) < 1{
			fmt.Println("Error: ID is required")
			os.Exit(1)
		}

		val, err := strconv.Atoi(cmdArgs[0])
		if err != nil {
			fmt.Println("Error: Input is not a valid integer")
			os.Exit(1)
		}

		tasks := loadTasks()
		task, found := findTaskByID(tasks, val)

		if !found {
			fmt.Printf("Task with ID %d not found\n", val)
			os.Exit(1)
		}

		task.Status = StatusDone
		task.UpdatedAt = time.Now()

		saveTasks(tasks)
		fmt.Printf("Status of the task with ID %d changed successfully\n", val)
}

func handleMarkInProgress(cmdArgs []string){
	if len(cmdArgs) < 1{
			fmt.Println("Error: ID is required")
			os.Exit(1)
		}

		val, err := strconv.Atoi(cmdArgs[0])
		if err != nil {
			fmt.Println("Error: Input is not a valid integer")
			os.Exit(1)
		}

		tasks := loadTasks()
		task, found := findTaskByID(tasks, val)

		if !found {
			fmt.Printf("Task with ID %d not found\n", val)
			os.Exit(1)
		}

		task.Status = StatusInProgress
		task.UpdatedAt = time.Now()

		saveTasks(tasks)
		fmt.Printf("Status of the task with ID %d changed successfully\n", val)
}

func handleUpdate(cmdArgs []string){
		if len(cmdArgs) < 2 {
			fmt.Println("description is required")
			os.Exit(1)
		}

		val, err := strconv.Atoi(cmdArgs[0])
		if err != nil {
			fmt.Println("Error: Input is not a valid integer")
			os.Exit(1)
		}
		description := strings.Join(cmdArgs[1:], " ")

		tasks := loadTasks()
		task, found := findTaskByID(tasks, val)

		if !found {
			fmt.Printf("Task with ID %d not found\n", val)
			os.Exit(1)
		}

		task.Description = description
		task.UpdatedAt = time.Now()

		saveTasks(tasks)
		fmt.Printf("Description of the task with ID %d changed successfully\n", val)
}

func handleHelp(cmdArgs []string) {
    helpText := `
Task Tracker CLI

Usage:
  task-cli <command> [arguments]

Commands:
  add <description>          Add a new task
  list [status]              List all tasks or filter by status
  delete <id>                Delete a task by ID
  mark-done <id>             Mark task as done
  mark-in-progress <id>      Mark task as in progress
  update <id> <description>  Update task description
  help                       Show this help message

Examples:
  task-cli add "Buy groceries"
  task-cli list
  task-cli list done
  task-cli mark-done 1
  task-cli update 1 "Buy milk and bread"
  task-cli delete 1
`
    fmt.Print(helpText)
}

func main() {
	args := os.Args[1:]
	argsLen := len(args)
	
	if(argsLen == 0){
		handleHelp([]string{})
		os.Exit(1)
	}

	command := args[0]
	switch command {
	case "add":
		handleAdd(args[1:])
	case "list":
		handleList(args[1:])
	case "delete":
		handleDelete(args[1:])
	case "mark-done":
		handleMarkDone(args[1:])
	case "mark-in-progress":
		handleMarkInProgress(args[1:])
	case "update":
		handleUpdate(args[1:])
	case "help", "-h", "--help":
		handleHelp(args[1:])
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Run 'task-cli help' for usage information")
		os.Exit(1)
	}
	
}	