package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	command := os.Args[1]
	tasks := LoadTasks()

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Description is required")
			return
		}
		desc := os.Args[2]
		tasks.AddTask(desc)
		SaveTasks(tasks)
	case "list":
		tasks.ListTask()
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("ID is required to mark as done")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID")
		}
		tasks.MarkDone(id)
		SaveTasks(tasks)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("ID is required to delete")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID")
		}
		tasks.Delete(id)
		SaveTasks(tasks)
	}

}

func printUsage() {
	fmt.Println("Usage:")
	fmt.Println("  add <task description>      - Add a new task")
	fmt.Println("  list                        - List all tasks")
	fmt.Println("  done <task id>              - Mark task as done")
	fmt.Println("  delete <task id>            - Delete task")
}
