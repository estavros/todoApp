package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	Done     bool
	Text     string
	DueDate  string // YYYY-MM-DD
	Priority string // low, medium, high
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

			fmt.Print("Enter priority (low, medium, high) or leave empty: ")
			scanner.Scan()
			priority := scanner.Text()

			tasks = append(tasks, Task{
				Done: false, Text: taskText, DueDate: dueDate, Priority: priority,
			})
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

					if t.Priority != "" {
						fmt.Printf(" [Priority: %s]", t.Priority)
					}

					if isOverdue(t) {
						fmt.Print(" ⚠ OVERDUE")
					}

					fmt.Println()
				}
			}

		case "3":
			fmt.Print("Enter task number to mark completed: ")
			scanner.Scan()
			num, err := strconv.Atoi(scanner.Text())
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
			num, err := strconv.Atoi(scanner.Text())
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

func isOverdue(t Task) bool {
	if t.Done || t.DueDate == "" {
		return false
	}

	due, err := time.Parse("2006-01-02", t.DueDate)
	if err != nil {
		return false
	}

	today := time.Now().Truncate(24 * time.Hour)

	return due.Before(today)
}

// Load tasks
func loadTasks() {
	file, err := os.Open(tasksFile)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "|", 4)
		if len(parts) != 4 {
			continue
		}

		tasks = append(tasks, Task{
			Done:     parts[0] == "1",
			Text:     parts[1],
			DueDate:  parts[2],
			Priority: parts[3],
		})
	}
}

// Save tasks
func saveTasks() {
	file, err := os.Create(tasksFile)
	if err != nil {
		fmt.Println("Error saving:", err)
		return
	}
	defer file.Close()

	for _, t := range tasks {
		done := "0"
		if t.Done {
			done = "1"
		}
		file.WriteString(done + "|" + t.Text + "|" + t.DueDate + "|" + t.Priority + "\n")
	}
}

// Export JSON
func exportToJSON() {
	file, err := os.Create("tasks.json")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")
	enc.Encode(tasks)

	fmt.Println("Tasks exported to tasks.json!")
}

// Export Toon
func exportToToon() {
	file, err := os.Create("tasks.toon")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	defer w.Flush()

	w.WriteString("TASKS:\n\n")

	for i, t := range tasks {
		status := "pending"
		if t.Done {
			status = "done"
		}

		w.WriteString(fmt.Sprintf(
			"- id: %d\n  status: %s\n  text: %s\n  due: %s\n  priority: %s\n\n",
			i+1, status, t.Text, t.DueDate, t.Priority,
		))
	}

	fmt.Println("Tasks exported to tasks.toon!")
}
