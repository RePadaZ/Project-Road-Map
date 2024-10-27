from command import run_command

# List of commands
commandList = {
    "task-cli add": True,
    "task-cli update": True,
    "task-cli delete": True,
    "task-cli mark-in-progress": True,
    "task-cli mark-done": True,
    "task-cli list": True,
    "exit": True
}

# Command processor
def processor_command(command):
    try:
        commandeering = command.split()
        params = commandeering[2:]
        command = " ".join(commandeering[:2])

        if commandList[command.lower()]:
            return run_command(command, params)
    except KeyError:
        return False

# Star app
def main():
    try:
        while True:
            command = input("Enter the command: ")
            result = processor_command(command)
            if not result:
                print("Unknown team")
            elif result == "exit":
                print("Application completed")
                break
            elif not result and command == "task-cli list":
                print("You don't have such tasks")
            else:
                print(result)
    except KeyboardInterrupt:
        print("Application completed")

if __name__ == "__main__":
    main()
