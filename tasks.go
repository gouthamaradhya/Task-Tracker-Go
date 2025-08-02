package main

import (
	"errors"
	"fmt"
	"strings"
)

type Task struct {
	ID          int
	Description string
	Done        bool
}

type TaskList struct {
	Tasks []Task
}

func (t *TaskList) AddTask(desc string) error {
	if strings.TrimSpace(desc) == "" {
		return errors.New("description cannot be empty")
	}
	id := len(t.Tasks) + 1
	t.Tasks = append(t.Tasks, Task{ID: id, Description: desc, Done: false})
	fmt.Println("Task added succesfully")
	return nil
}

func (t *TaskList) ListTask() {
	if len(t.Tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}
	for _, tasks := range t.Tasks {
		status := "❌"
		if tasks.Done {
			status = "✅"
		}
		fmt.Printf("[%s] %d %s\n", status, tasks.ID, tasks.Description)
	}
}

func (t *TaskList) MarkDone(id int) {
	for i, task := range t.Tasks {
		if id == task.ID {
			t.Tasks[i].Done = true
			fmt.Println("Marked task as done")
			return
		}
	}
	fmt.Println("Task not found")

}

func (t *TaskList) Delete(id int) {
	for i, task := range t.Tasks {
		if id == task.ID {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			fmt.Println("Tasks deleted successfully")
		}
	}
	fmt.Println("Could not find the task")
}
