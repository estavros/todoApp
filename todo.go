package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tasks []string
const tasksFile = "tasks.txt"

func main() {
	loadTasks()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTodo App")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task: ")
			scanner.Scan()
			task := scanner.Text()
			tasks = append(tasks, task)
			fmt.Println("Task added!")
			saveTasks()
		case "2":
			fmt.Println("\nTasks:")
			if len(tasks) == 0 {
				fmt.Println("No tasks yet.")
			} else {
				for i, t := range tasks {
					fmt.Printf("%d. %s\n", i+1, t)
				}
			}
		case "3":
			fmt.Print("Enter task number to delete: ")
			scanner.Scan()
			numStr := scanner.Text()
			num, err := strconv.Atoi(numStr)
			if err != nil || num < 1 || num > len(tasks) {
				fmt.Println("Invalid task number.")
			} else {
				tasks = append(tasks[:num-1], tasks[num:]...)
				fmt.Println("Task deleted!")
				saveTasks()
			}
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

// Load tasks from file at startup
func loadTasks() {
	file, err := os.Open(tasksFile)
	if err != nil {
		return // file doesn't exist yet
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())
	}
}

// Save tasks to file
func saveTasks() {
	file, err := os.Create(tasksFile)
	if err != nil {
		fmt.Println("Error saving tasks:", err)
		return
	}
	defer file.Close()

	for _, t := range tasks {
		file.WriteString(t + "\n")
	}
}
