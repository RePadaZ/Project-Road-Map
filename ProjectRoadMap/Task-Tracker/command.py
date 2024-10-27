import json
import os
from datetime import datetime

# Command run processor
def run_command(command, params):

    if command == "task-cli add":
        return add_command(params)
    elif command == "task-cli update":
        return update_command(params)
    elif command == "task-cli delete":
        return delete_command(params)
    elif command == "task-cli mark-in-progress":
        return mark_in_progress_command(params)
    elif command == "task-cli mark-done":
        return mark_done_command(params)
    elif command == "task-cli list":
        return list_command(params)
    elif command == "exit":
        return "exit"
    else:
        return False

# Add command
def add_command(params):

    current_time = datetime.now()
    formatted_data_time  = current_time.strftime('%H:%M %d %B %Y')

    if not os.path.exists("dbTask.json"):
        writedoc = [
            {

            "id": 1,
            "description" : params[0],
            "status" : "todo",
            "createdAt" : formatted_data_time,
            "updatedAt" : formatted_data_time

            }
        ]
        with open("dbTask.json", 'w') as file:
            json.dump(writedoc, file)
    else:
        with open("dbTask.json", 'r') as file:
            data = json.load(file)
            data.append({
                "id": data[-1]["id"] + 1,
                "description" : params[0],
                "status" : "todo",
                "createdAt" : formatted_data_time,
                "updatedAt" : formatted_data_time
            })
        with open("dbTask.json", 'w') as file:
            json.dump(data, file)
    return f'Task added successfully (ID:{data[-1]["id"]})'

# Update command
def update_command(params):
    if not os.path.exists("dbTask.json"):
        return False
    with open("dbTask.json", 'r') as file:
        data = json.load(file)
        for task in data:
            if task["id"] == int(params[0]):
                task["description"] = params[1]
                task["updatedAt"] = datetime.now().strftime('%H:%M %d %B %Y')
                with open("dbTask.json", 'w') as file:
                    json.dump(data, file)
                return f'Task updated successfully (ID:{task["id"]})'
    return False

# Delete command
def delete_command(params):
    if not os.path.exists("dbTask.json"):
        return False
    with open("dbTask.json", 'r') as file:
        data = json.load(file)
        for task in data:
            if task["id"] == int(params[0]):
                data.remove(task)
                with open("dbTask.json", 'w') as file:
                    json.dump(data, file)
                return f'Task deleted successfully (ID:{task["id"]})'
    return False

# Mark in progress command
def mark_in_progress_command(params):
    if not os.path.exists("dbTask.json"):
        return False
    with open("dbTask.json", 'r') as file:
        data = json.load(file)
        for task in data:
            if task["id"] == int(params[0]):
                task["status"] = "in progress"
                task["updatedAt"] = datetime.now().strftime('%H:%M %d %B %Y')
                with open("dbTask.json", 'w') as file:
                    json.dump(data, file)
                return f'Task marked in progress successfully (ID:{task["id"]})'

# Mark done command
def mark_done_command(params):
    if not os.path.exists("dbTask.json"):
        return False
    with open("dbTask.json", 'r') as file:
        data = json.load(file)
        for task in data:
            if task["id"] == int(params[0]):
                task["status"] = "done"
                task["updatedAt"] = datetime.now().strftime('%H:%M %d %B %Y')
                with open("dbTask.json", 'w') as file:
                    json.dump(data, file)
                return f'Task marked done successfully (ID:{task["id"]})'

# List command
def list_command(params):
    if not os.path.exists("dbTask.json"):
        return False
    if not params:
        with open("dbTask.json", 'r') as file:
            data = json.load(file)
            return data
    if params[0] == "in progress":
        with open("dbTask.json", 'r') as file:
            data = json.load(file)
            return [task for task in data if task["status"] == "in progress"]
    if params[0] == "done":
        with open("dbTask.json", 'r') as file:
            data = json.load(file)
            return [task for task in data if task["status"] == "done"]
    if params[0] == "todo":
        with open("dbTask.json", 'r') as file:
            data = json.load(file)
            return [task for task in data if task["status"] == "todo"]
