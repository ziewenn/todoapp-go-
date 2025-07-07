# Go To-Do App

This is a simple and efficient command-line to-do application built with Go. It allows you to manage your tasks directly from the terminal.

## Features

- **Add new tasks:** Quickly add new to-do items to your list.
- **List all tasks:** View all your current to-do items with their status.
- **Edit existing tasks:** Modify the title of a task.
- **Mark tasks as completed:** Toggle the status of a task between pending and completed.
- **Delete tasks:** Remove tasks from your list.
- **Persistent storage:** Your to-do list is saved to a `todos.json` file, so your data persists between sessions.

## Getting Started

### Prerequisites

Make sure you have Go installed on your system. You can download it from [golang.org](https://golang.org/dl/).

### Installation

1. Clone the repository to your local machine:
   ```sh
   git clone https://github.com/ziewenn/todoapp-go-.git
   ```
2. Navigate to the project directory:
   ```sh
   cd todoapp-go-
   ```
3. Build the application:
   ```sh
   go build
   ```
   This will create an executable file named `gotodo` (or `gotodo.exe` on Windows).

## Dependencies

This project uses the following external dependency:

- **github.com/aquasecurity/table:** For formatting the to-do list into a clean, readable table in the terminal.

You can install it by running:

```sh
go get github.com/aquasecurity/table
```

## Usage

You can use the following flags to interact with the to-do list application.

### Add a new task

Use the `-add` flag to add a new task to your list.

```sh
./gotodo -add "Buy groceries"
```

### List all tasks

Use the `-list` flag to display all your tasks.

```sh
./gotodo -list
```

The output will show the index, status, and title of each task.

### Edit a task

Use the `-edit` flag to modify an existing task. The format is `<index>:<new title>`.

```sh
./gotodo -edit "0:Buy fresh milk"
```

### Toggle task status

Use the `-toggle` flag to mark a task as completed or pending.

```sh
./gotodo -toggle 0
```

### Delete a task

Use the `-del` flag to remove a task from the list by its index.

```sh
./gotodo -del 0
```
