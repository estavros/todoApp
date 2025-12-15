package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	Done    bool
	Text    string
	DueDate string // New field for due date
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
		fmt.Println("6. Export Tasks to JSON")
		fmt.Println("7. Export Tasks to .toon file")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter task: ")
			scanner.Scan()
			taskText := scanner.Text()

			fmt.Print("Enter due date (YYYY-MM-DD) or leave empty: ")
			scanner.Scan()
			dueDate := scanner.Text()

			tasks = append(tasks, Task{Done: false, Text: taskText, DueDate: dueDate})
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
					fmt.Printf("%d. %s %s", i+1, status, t.Text)
					if t.DueDate != "" {
						fmt.Printf(" (Due: %s)", t.DueDate)
					}
					fmt.Println()
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

		case "6":
			exportToJSON()

		case "7":
			exportToToon()

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
		parts := strings.SplitN(line, "|", 3) // now expecting 3 parts
		if len(parts) != 3 {
			continue
		}
		done := parts[0] == "1"
		tasks = append(tasks, Task{Done: done, Text: parts[1], DueDate: parts[2]})
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
		file.WriteString(doneFlag + "|" + t.Text + "|" + t.DueDate + "\n")
	}
}

// Export tasks to JSON
func exportToJSON() {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error creating JSON file:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // pretty print

	if err := encoder.Encode(tasks); err != nil {
		fmt.Println("Error writing JSON:", err)
		return
	}

	fmt.Println("Tasks exported to tasks.json!")
}

// Export tasks to .toon file
func exportToToon() {
	file, err := os.Create("tasks.toon")
	if err != nil {
		fmt.Println("Error creating Toon file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	writer.WriteString("TASKS:\n\n")

	for i, t := range tasks {
		status := "pending"
		if t.Done {
			status = "done"
		}

		writer.WriteString(fmt.Sprintf(
			"- id: %d\n  status: %s\n  text: %s\n  due: %s\n\n",
			i+1,
			status,
			t.Text,
			t.DueDate,
		))
	}

	fmt.Println("Tasks exported to tasks.toon!")
}
