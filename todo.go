package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tasks []string

func main() {
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
			}
		case "4":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
