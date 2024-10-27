# Expense Tracker

Sample solution for the [expense-tracker](https://roadmap.sh/projects/expense-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

## How to run

Clone the repository and run the following command:

```bash
git clone https://github.com/SolitaryGreed/expense-tracker.git
```

Run the following command to build and run the project:

```bash
go build -o expense-tracker
./expense-tracker --help # To see the list of available commands

# To add an expense
./expense-tracker add --description "Buy groceries" --amount 100

# List all expenses
./expense-tracker list

# Delete an expense
./expense-tracker delete --id 1

# Summarize expenses
./expense-tracker summary
./expense-tracker summary --month 8
```
