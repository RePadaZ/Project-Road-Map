package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all expenses",
	Long:  `List all all all expenses`,

	Run: func(cmd *cobra.Command, args []string) {
		ListAllExpenses()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func ListAllExpenses() {

	var data []Expense

	f, err := os.ReadFile("expenses.json")
	if err != nil {
		log.Fatalf("Не удалось прочитать файл: %v", err)
	}
	err = json.Unmarshal(f, &data)
	if err != nil {
		log.Fatalf("Ошибка при парсинге JSON: %v", err)
	}

	fmt.Println("ID       Date          Description   Amount")
	for _, expense := range data {
		fmt.Printf("%s  %s     %s      $%d\n", expense.Id, expense.UpdatedAt, expense.Description, expense.Amount)
	}
	os.Exit(0)
}
