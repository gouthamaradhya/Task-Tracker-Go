package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func LoadTasks() TaskList {
	var taskList TaskList
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		return taskList
	}
	err = json.Unmarshal(file, &taskList)
	if err != nil {
		fmt.Println("Error parsing the json")
	}
	return taskList

}

func SaveTasks(taskList TaskList) {
	data, err := json.MarshalIndent(taskList, "", "  ")
	if err != nil {
		fmt.Println("Error saving tasks")
		return
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing tasks")
	}
}
