package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete expenses",
	Long:  `Delete expenses by id`,

	Run: func(cmd *cobra.Command, args []string) {
		deleteExpenses()
	},
}
var id int

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().IntVarP(&id, "id", "i", 0, "id of the expense for deletion")
}

func deleteExpenses() {

	var data []Expense

	f, err := os.ReadFile("expenses.json")
	if err != nil {
		log.Fatalf("Не удалось прочитать файл: %v", err)
	}
	err = json.Unmarshal(f, &data)
	if err != nil {
		log.Fatalf("Ошибка при парсинге JSON: %v", err)
	}

	flag := false
	// Find the index of the expense to delete
	for i, expense := range data {
		if expense.Id == strconv.Itoa(id) {
			// Delete the expense from the slice
			data = append(data[:i], data[i+1:]...)
			flag = true
			break
		}
	}

	// Encoding the updated data in JSON
	updatedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Ошибка при кодировании JSON: %v", err)
	}

	err = os.WriteFile("expenses.json", updatedData, 0644)
	if err != nil {
		log.Fatalf("Ошибка при записи в файл: %v", err)
	}

	if !flag {
		fmt.Println("Expense not found")
		os.Exit(0)
	}
	fmt.Printf("Expense added successfully (ID: %v)", id)
	os.Exit(0)
}
