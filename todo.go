package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID       int
	Done     bool
	Text     string
	DueDate  string
	Priority string
}

var tasks []Task
var nextID = 1
const tasksFile = "tasks.txt"

const (
	Red    = "\033[31m"
	Yellow = "\033[33m"
	Green  = "\033[32m"
	Reset  = "\033[0m"
)

func main() {
	loadTasks()
	showStartupDashboard()

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("\nTodo App")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Task as Completed (by ID)")
		fmt.Println("4. Delete Task (by ID)")
		fmt.Println("5. Exit")
		fmt.Println("6. Export Tasks to JSON")
		fmt.Println("7. Export Tasks to .toon file")
		fmt.Println("8. View filtered & sorted tasks")
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
				ID:       nextID,
				Done:     false,
				Text:     taskText,
				DueDate:  dueDate,
				Priority: priority,
			})
			nextID++
			saveTasks()
			fmt.Println("Task added!")

		case "2":
			printTasks(tasks)

		case "3":
			fmt.Print("Enter task ID to mark completed: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			if markDoneByID(id) {
				saveTasks()
				fmt.Println("Task marked as completed!")
			} else {
				fmt.Println("Task not found.")
			}

		case "4":
			fmt.Print("Enter task ID to delete: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			if deleteByID(id) {
				saveTasks()
				fmt.Println("Task deleted!")
			} else {
				fmt.Println("Task not found.")
			}

		case "5":
			fmt.Println("Goodbye!")
			return

		case "6":
			exportToJSON()

		case "7":
			exportToToon()

		case "8":
			filterMenu(scanner)

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func printTasks(list []Task) {
	if len(list) == 0 {
		fmt.Println("No tasks.")
		return
	}

	for _, t := range list {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}

		fmt.Printf("#%d %s %s", t.ID, status, t.Text)

		if t.DueDate != "" {
			fmt.Printf(" (Due: %s)", t.DueDate)
		}
		if t.Priority != "" {
			fmt.Printf(" [Priority: %s]", t.Priority)
		}
		if isOverdue(t) {
			fmt.Print(Red + " ⚠ OVERDUE" + Reset)
		}
		fmt.Println()
	}
}

func markDoneByID(id int) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			return true
		}
	}
	return false
}

func deleteByID(id int) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return true
		}
	}
	return false
}

// ---------- FILE HANDLING ----------

func loadTasks() {
	file, err := os.Open(tasksFile)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), "|", 5)

		var t Task

		if len(parts) == 4 {
			// legacy format
			t.ID = nextID
			t.Done = parts[0] == "1"
			t.Text = parts[1]
			t.DueDate = parts[2]
			t.Priority = parts[3]
		} else {
			t.ID, _ = strconv.Atoi(parts[0])
			t.Done = parts[1] == "1"
			t.Text = parts[2]
			t.DueDate = parts[3]
			t.Priority = parts[4]
		}

		if t.ID >= nextID {
			nextID = t.ID + 1
		}

		tasks = append(tasks, t)
	}
}

func saveTasks() {
	file, _ := os.Create(tasksFile)
	defer file.Close()

	for _, t := range tasks {
		done := "0"
		if t.Done {
			done = "1"
		}
		file.WriteString(fmt.Sprintf("%d|%s|%s|%s|%s\n", t.ID, done, t.Text, t.DueDate, t.Priority))
	}
}

func exportToJSON() {
	file, _ := os.Create("tasks.json")
	defer file.Close()
	json.NewEncoder(file).Encode(tasks)
	fmt.Println("Tasks exported to tasks.json!")
}

func exportToToon() {
	file, _ := os.Create("tasks.toon")
	defer file.Close()

	for _, t := range tasks {
		status := "pending"
		if t.Done {
			status = "done"
		}
		fmt.Fprintf(file, "- id: %d\n  status: %s\n  text: %s\n  due: %s\n  priority: %s\n\n",
			t.ID, status, t.Text, t.DueDate, t.Priority)
	}
	fmt.Println("Tasks exported to tasks.toon!")
}
