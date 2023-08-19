package main

import (
	"fmt"
	"os"
	"os/exec"
)

// Task represents a task with a description and a status.
type Task struct {
	Description string
	Done        bool
}

// Function to add a new task to the task list.
func addTask(tasks []Task, description string) []Task {
	newTask := Task{Description: description, Done: false}
	return append(tasks, newTask)
}

// Function to mark a task as done.
func completeTask(tasks []Task, index int) []Task {
	if index >= 0 && index < len(tasks) {
		tasks[index].Done = true
	}
	return tasks
}

// Function to display the list of tasks.
func displayTasks(tasks []Task) {
	clearScreen()
	fmt.Println("Tasks:")
	for i, task := range tasks {
		status := " "
		if task.Done {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, task.Description)
	}
	fmt.Println()
}

// Function to clear the terminal screen.
func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	tasks := []Task{}

	for {
		clearScreen()
		fmt.Println("Task Management")

		displayTasks(tasks)

		fmt.Println("1. Add Task")
		fmt.Println("2. Mark Task as Done")
		fmt.Println("3. Exit")
		fmt.Print("Select an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter task description: ")
			var description string
			fmt.Scanln(&description)
			tasks = addTask(tasks, description)
		case 2:
			fmt.Print("Enter task number to mark as done: ")
			var taskNumber int
			fmt.Scanln(&taskNumber)
			tasks = completeTask(tasks, taskNumber-1)
		case 3:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please select again.")
		}
	}
}
