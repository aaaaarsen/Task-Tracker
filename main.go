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

func main() {
	args := os.Args[1:]
	argsLen := len(args)
	
	if(argsLen == 0){
		fmt.Println("Please provide a command")
		os.Exit(1)
	}

	command := args[0]
	if command == "add" {
		if len(args) < 2{
			fmt.Println("Error: description is required")
			os.Exit(1)
		}

		description := strings.Join(args[1:], " ")

		tasks := loadTasks()

		newTask := Task{
			ID: len(tasks)+1,
			Description: description,
			Status: "todo",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		tasks = append(tasks, newTask)

		saveTasks(tasks)

		fmt.Printf("Task added succesfully! ID: %d\n", newTask.ID)
		
	} else if command == "list" {
		tasks:= loadTasks()

		if len(tasks) == 0{
			fmt.Println("No tasks found")
			os.Exit(1)
		}

		for _, t := range tasks {
			fmt.Printf(
				"ID: %d, description: %s\nStatus: %s\nCreated at: %v\nUpdated at: %v\n-------\n",
				t.ID, t.Description, t.Status, t.CreatedAt.Format("2006-01-02 15:04"), t.UpdatedAt.Format("2006-01-02 15:04"))
		}

	} else if command == "delete"{
		if len(args) < 2{
			fmt.Println("Error: ID is required")
			os.Exit(1)
		} 

		val, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: Input is not a valid integer")
			os.Exit(1)
		}
		
		tasks := loadTasks()
		var updatedTasks []Task
		found := false

		for _, v := range tasks {
			if v.ID != val{
				updatedTasks = append(updatedTasks, v)
			} else {
				found = true
			}
		}

		if !found {
			fmt.Printf("Task with ID %d not found\n", val)
			os.Exit(1)
		}

		saveTasks(updatedTasks)
		fmt.Printf("Task with ID %d deleted successfully\n", val)

	} else if command == "mark-done"{
		if len(args) < 2{
			fmt.Println("Error: ID is required")
			os.Exit(1)
		}

		val, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: Input is not a valid integer")
			os.Exit(1)
		}

		tasks := loadTasks()
		found := false

		for i, _ := range tasks{
			if val == tasks[i].ID{
				tasks[i].Status = "done"
				tasks[i].UpdatedAt = time.Now()
			} else {
				found = true
			}
		}

		if !found {
			fmt.Printf("Task with ID %d not found\n", val)
			os.Exit(1)
		}

		saveTasks(tasks)
		fmt.Printf("Status of the task with ID %d changed successfully\n", val)

	} else if command == "mark-in-progress"{
		if len(args) < 2{
			fmt.Println("Error: ID is required")
			os.Exit(1)
		}

		val, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: Input is not a valid integer")
			os.Exit(1)
		}

		tasks := loadTasks()
		found := false

		for i, _ := range tasks{
			if val == tasks[i].ID{
				tasks[i].Status = "in progress"
				tasks[i].UpdatedAt = time.Now()
			} else {
				found = true
			}
		}

		if !found {
			fmt.Printf("Task with ID %d not found\n", val)
			os.Exit(1)
		}

		saveTasks(tasks)
		fmt.Printf("Status of the task with ID %d changed successfully\n", val)

	} else if command == "update"{
		if len(args) < 2 {
			fmt.Println("description is required")
			os.Exit(1)
		}

		val, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Error: Input is not a valid integer")
			os.Exit(1)
		}
		description := strings.Join(args[2:], " ")

		tasks := loadTasks()

		for i,_ := range tasks {
			if tasks[i].ID == val{
				tasks[i].Description = description
			}
		}

		saveTasks(tasks)
		fmt.Printf("Description of the task with ID %d changed successfully\n", val)

	} else {
		fmt.Println("Unknown command called", command)
	}
	
}	