package main

import (
	"testing"
)

func TestFindTaskByID(t *testing.T) {
	tasks := []Task{
		{ID: 1, Description: "Task 1", Status: StatusTodo},
		{ID: 2, Description: "Task 2", Status: StatusDone},
		{ID: 3, Description: "Task 3", Status: StatusInProgress},
	}

	task, found := findTaskByID(tasks,2)
	if !found {
		t.Error("Expected to find task with ID: 2")
	}

	if task.ID != 2 {
		t.Errorf("Expected ID: 2, got %d", task.ID)
	}

	if task.Description != "Task 2"{
		t.Errorf("Expected 'Task 2', got %s", task.Description)
	}


	_,found = findTaskByID(tasks,999)
	if found {
		t.Error("Expected not to find task with ID: 999")
	}
}

func TestFindTaskIndexByID(t *testing.T) {
	tasks := []Task {
		{ID: 10, Description: "First"},
		{ID: 20, Description: "Second"},
		{ID: 30, Description: "Third"},
	}
	//test 1
	index, found := findTaskIndexByID(tasks, 10)
	if !found {
		t.Error("Expected to find task index by ID: 10")
	}

	if index != 0 {
		t.Errorf("Expected task index by ID: %d", index)
	}
	//test 2
	index, found = findTaskIndexByID(tasks, 20)
	if !found {
		t.Error("Expected to find task with ID 20")
	}
	if index != 1 {
		t.Errorf("Expected index 1 for ID 20, got %d", index)
	}
	//test 3
	index, found = findTaskIndexByID(tasks, 30)
	if !found {
		t.Error("Expected to find task with ID 30")
	}
	if index != 2 {
		t.Errorf("Expected index 2 for ID 30, got %d", index)
	}
	//test 4
	index, found = findTaskIndexByID(tasks, 999)
	if found {
		t.Errorf("Expected to find task with ID: 999")
	}
	if index != -1{
		t.Errorf("Expected index -1 for non-existent task, got %d", index)
	}

}