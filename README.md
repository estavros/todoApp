# Todo CLI App

A simple command-line Todo application written in Go. It allows you to add, list, complete, delete, and export tasks. Tasks are saved to a text file so they persist between runs.

## Features

* Add new tasks
* View all existing tasks
* Mark tasks as completed
* Delete tasks by number
* **Export tasks to JSON (`tasks.json`)**
* **Export tasks to Toon format (`tasks.toon`)** *(for LLM-friendly task discussion)*
* Persistent storage using a `tasks.txt` file
* Automatic overdue detection (⚠ overdue tasks are highlighted)
* Filter tasks (overdue, due today, high priority, pending)
* Sort tasks by due date

## How It Works

The application stores tasks in memory and writes them into `tasks.txt` so they are available the next time you run the program. When the application starts, it loads tasks from the file (if it exists).

Each task now has:

* **Text** — the description of the task
* **Due Date** — optional, format `YYYY-MM-DD`
* **Priority** — optional, e.g., `low`, `medium`, `high`
* **Done status** — whether the task is completed

You can export tasks into different formats:

* **JSON** — a structured machine-readable format
* **Toon** — a lightweight, human- and LLM-friendly text format designed for future AI-based task discussion

> Note: Toon export currently exists as an internal function and is not exposed in the menu yet. It is intended for future integration with an LLM chat interface.

## ⏰ Overdue Detection

Tasks with a due date are automatically checked against the current date.

A task is considered **overdue** if:
- It has a due date
- It is not marked as completed
- The due date is before today

Overdue tasks are displayed in the task list with:

⚠ OVERDUE

Example:

3. [ ] Pay rent (Due: 2025-12-01) [Priority: high] ⚠ OVERDUE

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
   3. Mark Task as Completed
   4. Delete Task
   5. Exit
   6. Export Tasks to JSON
   7. Export Tasks to .toon file
   ```

3. Choose an option:

   * **Add Task**: Type your task and press Enter.
   * **List Tasks**: Shows all tasks with status, due date, priority, and overdue warnings.
   * **Mark as Completed**: Enter the number of the task to mark it done.
   * **Delete Task**: Enter the number of the task you want to delete.
   * **Export Tasks to JSON**: Saves all tasks to `tasks.json`.
   * **Export Tasks to .toon file: Saves tasks to `tasks.toon` with due dates and priorities.
   * **Exit**: Close the application.

## File Structure

```
.
├── main.go        # Application source code
├── tasks.txt      # Automatically generated task storage file
├── tasks.json     # Generated JSON export file
└── tasks.toon     # Generated Toon export file
```

## Example

```
Todo App
1. Add Task
2. List Tasks
3. Mark Task as Completed
4. Delete Task
5. Exit
6. Export Tasks to JSON
7. Export Tasks to .toon file
8. View filtered & sorted tasks
Choose an option: 1
Enter task: Buy groceries
Enter due date (YYYY-MM-DD) or leave empty: 2025-12-20
Enter priority (low,medium,high) or leave empty: high
Task added!

1. [ ] Buy groceries (Due: 2025-12-20) [Priority: high] ⚠ OVERDUE
```

Enjoy your improved Todo CLI app!
