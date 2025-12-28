package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
	"strings"
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
		fmt.Println("LIST command called")
	} else {
		fmt.Println("Unknown command called", command)
	}
	
}	