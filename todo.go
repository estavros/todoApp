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
	showReminders()

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
		fmt.Println("9. Edit Task (by ID)")
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

		case "9":
			fmt.Print("Enter task ID to edit: ")
			scanner.Scan()
			id, _ := strconv.Atoi(scanner.Text())
			if editTaskByID(id, scanner) {
				saveTasks()
				fmt.Println("Task updated!")
			} else {
				fmt.Println("Task not found.")
			}

		default:
			fmt.Println("Invalid choice.")
		}
	}
}

/* ---------------- EDIT TASK ---------------- */

func editTaskByID(id int, scanner *bufio.Scanner) bool {
	for i := range tasks {
		if tasks[i].ID == id {
			fmt.Printf("Current text: %s\nNew text (leave empty to keep): ", tasks[i].Text)
			scanner.Scan()
			if input := scanner.Text(); input != "" {
				tasks[i].Text = input
			}

			fmt.Printf("Current due date: %s\nNew due date (YYYY-MM-DD, empty to keep): ", tasks[i].DueDate)
			scanner.Scan()
			if input := scanner.Text(); input != "" {
				tasks[i].DueDate = input
			}

			fmt.Printf("Current priority: %s\nNew priority (low/medium/high, empty to keep): ", tasks[i].Priority)
			scanner.Scan()
			if input := scanner.Text(); input != "" {
				tasks[i].Priority = input
			}

			return true
		}
	}
	return false
}

/* ---------------- FILTER SYSTEM ---------------- */

func filterMenu(scanner *bufio.Scanner) {
	filtered := make([]Task, 0)
	fmt.Println("\nFilter:")
	fmt.Println("1. All")
	fmt.Println("2. Pending")
	fmt.Println("3. Completed")
	fmt.Println("4. Overdue")
	fmt.Println("5. Due Today")
	fmt.Println("6. By Priority")
	fmt.Print("Choose filter: ")
	scanner.Scan()
	f := scanner.Text()

	today := time.Now().Format("2006-01-02")

	for _, t := range tasks {
		switch f {
		case "1":
			filtered = append(filtered, t)
		case "2":
			if !t.Done {
				filtered = append(filtered, t)
			}
		case "3":
			if t.Done {
				filtered = append(filtered, t)
			}
		case "4":
			if isOverdue(t) {
				filtered = append(filtered, t)
			}
		case "5":
			if t.DueDate == today {
				filtered = append(filtered, t)
			}
		case "6":
			fmt.Print("Enter priority (low, medium, high): ")
			scanner.Scan()
			p := scanner.Text()
			if strings.EqualFold(t.Priority, p) {
				filtered = append(filtered, t)
			}
		}
	}

	fmt.Println("\nSort by:")
	fmt.Println("1. ID")
	fmt.Println("2. Due date")
	fmt.Println("3. Priority")
	fmt.Print("Choose sort: ")
	scanner.Scan()
	s := scanner.Text()

	switch s {
	case "2":
		sort.Slice(filtered, func(i, j int) bool {
			return filtered[i].DueDate < filtered[j].DueDate
		})
	case "3":
		sort.Slice(filtered, func(i, j int) bool {
			return priorityRank(filtered[i]) > priorityRank(filtered[j])
		})
	}

	printTasks(filtered)
}

func priorityRank(t Task) int {
	switch strings.ToLower(t.Priority) {
	case "high":
		return 3
	case "medium":
		return 2
	case "low":
		return 1
	}
	return 0
}

func isOverdue(t Task) bool {
	if t.Done || t.DueDate == "" {
		return false
	}
	d, err := time.Parse("2006-01-02", t.DueDate)
	if err != nil {
		return false
	}
	return d.Before(time.Now())
}

/* ---------------- DISPLAY ---------------- */

func printTasks(list []Task) {
	if len(list) == 0 {
		fmt.Println("No tasks.")
		return
	}

	for _, t := range list {
		status := "[ ]"
		color := Yellow
		if t.Done {
			status = "[x]"
			color = Green
		}

		fmt.Printf("%s#%d %s %s", color, t.ID, status, t.Text)

		if t.DueDate != "" {
			fmt.Printf(" (Due: %s)", t.DueDate)
		}
		if t.Priority != "" {
			fmt.Printf(" [Priority: %s]", t.Priority)
		}
		if isOverdue(t) {
			fmt.Print(Red + " ⚠ OVERDUE")
		}
		fmt.Println(Reset)
	}
}

/* ---------------- CORE ---------------- */

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

/* ---------------- FILE HANDLING ---------------- */

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

		t.ID, _ = strconv.Atoi(parts[0])
		t.Done = parts[1] == "1"
		t.Text = parts[2]
		t.DueDate = parts[3]
		t.Priority = parts[4]

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
		file.WriteString(fmt.Sprintf("%d|%s|%s|%s|%s\n",
			t.ID, done, t.Text, t.DueDate, t.Priority))
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
		fmt.Fprintf(file,
			"- id: %d\n  status: %s\n  text: %s\n  due: %s\n  priority: %s\n\n",
			t.ID, status, t.Text, t.DueDate, t.Priority)
	}
	fmt.Println("Tasks exported to tasks.toon!")
}

/* ---------------- STARTUP DASHBOARD ---------------- */

func showStartupDashboard() {
	fmt.Println(Green + "Loaded", len(tasks), "tasks." + Reset)
}

func showReminders() {
	today := time.Now().Format("2006-01-02")
	hasReminders := false

	for _, t := range tasks {
		if !t.Done && t.DueDate != "" {
			if t.DueDate == today {
				if !hasReminders {
					fmt.Println(Yellow + "\nTasks Due Today:" + Reset)
					hasReminders = true
				}
				fmt.Printf("  #%d %s\n", t.ID, t.Text)
			} else if isOverdue(t) {
				if !hasReminders {
					fmt.Println(Red + "\nOverdue Tasks:" + Reset)
					hasReminders = true
				}
				fmt.Printf("  #%d %s (Due: %s)\n", t.ID, t.Text, t.DueDate)
			}
		}
	}
	if !hasReminders {
		fmt.Println(Green + "No overdue or due today tasks." + Reset)
	}
}
