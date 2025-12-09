# Todo CLI App

A simple command-line Todo application written in Go. It allows you to add, list, and delete tasks, while automatically saving them to a text file so they persist between runs.

## Features

* Add new tasks
* View all existing tasks
* Delete tasks by number
* Persistent storage using a `tasks.txt` file

## How It Works

The application stores tasks in memory and also writes them into `tasks.txt` so they are available the next time you run the program. When the application starts, it loads tasks from the file (if it exists).

## Usage

1. Run the application using:

   ```bash
   go run main.go
   ```

2. You will see a menu:

   ```
   Todo App
   1. Add Task
   2. List Tasks
   3. Delete Task
   4. Exit
   ```

3. Choose an option:

   * **Add Task**: Type your task and press Enter.
   * **List Tasks**: Shows all tasks with numbering.
   * **Delete Task**: Enter the number of the task you want to delete.
   * **Exit**: Close the application.

## File Structure

```
.
├── main.go       # Application source code
└── tasks.txt     # Automatically generated task storage file
```

## Example

```
Todo App
1. Add Task
2. List Tasks
3. Delete Task
4. Exit
Choose an option: 1
Enter task: Buy groceries
Task added!
```

Enjoy your minimal Todo CLI app!
