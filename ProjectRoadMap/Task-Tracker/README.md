Task CLI Application - README
Overview
The Task CLI Application is a command-line interface (CLI) tool for managing personal tasks. It allows users to add, update, delete, and mark tasks with different statuses such as "in-progress" or "done." Additionally, users can list all tasks or filter them based on their status. Idea for training taken from here: https://roadmap.sh/projects/task-tracker

Features
Add new tasks
Update existing tasks
Delete tasks
Mark tasks as "in-progress" or "done"
List all tasks
List tasks by status (todo, in-progress, done)
Installation
Clone the repository and navigate to the project directory. Then, ensure the CLI application is executable:

```
git clone <repository-url>
cd <project-directory>
chmod +x task-cli
```

Usage
Here are the available commands for managing tasks:

1. Add a new task
To add a new task, use the add command followed by the task description:
```
task-cli add "Buy groceries"
```
2. Update an existing task
To update a task, use the update command followed by the task ID and the updated description:
```
task-cli update 1 "Buy groceries and cook dinner"
```
3. Delete a task
To delete a task, use the delete command followed by the task ID:
```
task-cli delete 1
```
4. Mark a task as "in-progress"
To mark a task as "in-progress," use the mark-in-progress command followed by the task ID:
```
task-cli mark-in-progress 1
```
6. List all tasks
To view all tasks, use the list command:
```
task-cli list
```
7. List tasks by status
To view tasks filtered by their status, use the list command followed by the status keyword (e.g., done, todo, or in-progress):
```
task-cli list done
task-cli list todo
task-cli list in-progress
```
License
This project is licensed under the MIT License. See the LICENSE file for more details.
