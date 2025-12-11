package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Done bool
	Text string
}

var tasks []Task
const tasksFile = "tasks.txt"

func main() {
	loadTasks()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTodo App")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task: ")
			scanner.Scan()
			task := scanner.Text()
			tasks = append(tasks, Task{Done: false, Text: task})
			fmt.Println("Task added!")
			saveTasks()

		case "2":
			fmt.Println("\nTasks:")
			if len(tasks) == 0 {
				fmt.Println("No tasks yet.")
			} else {
				for i, t := range tasks {
					status := "[ ]"
					if t.Done {
						status = "[x]"
					}
					fmt.Printf("%d. %s %s\n", i+1, status, t.Text)
				}
			}

		case "3":
			fmt.Print("Enter task number to mark completed: ")
			scanner.Scan()
			numStr := scanner.Text()
			num, err := strconv.Atoi(numStr)
			if err != nil || num < 1 || num > len(tasks) {
				fmt.Println("Invalid task number.")
			} else {
				tasks[num-1].Done = true
				fmt.Println("Task marked as completed!")
				saveTasks()
			}

		case "4":
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

		case "5":
			fmt.Println("Goodbye!")
			return

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

// Load tasks from file
func loadTasks() {
	file, err := os.Open(tasksFile)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "|", 2)
		if len(parts) != 2 {
			continue
		}
		done := parts[0] == "1"
		tasks = append(tasks, Task{Done: done, Text: parts[1]})
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
		doneFlag := "0"
		if t.Done {
			doneFlag = "1"
		}
		file.WriteString(doneFlag + "|" + t.Text + "\n")
	}
}
