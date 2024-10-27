package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
)

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Summary expenses",
	Long:  `Summary of all expenses`,

	Run: func(cmd *cobra.Command, args []string) {
		summaryAllExpenses()
	},
}
var monthPrint = map[int]string{
	1:  "January",
	2:  "February",
	3:  "March",
	4:  "April",
	5:  "May",
	6:  "June",
	7:  "July",
	8:  "August",
	9:  "September",
	10: "October",
	11: "November",
	12: "December",
}
var month int

func init() {
	rootCmd.AddCommand(summaryCmd)

	summaryCmd.Flags().IntVarP(&month, "month", "m", 0, "Months for which you need to calculate the total")
}

func summaryAllExpenses() {

	var data []Expense
	var summary int

	f, err := os.ReadFile("expenses.json")
	if err != nil {
		log.Fatalf("Не удалось прочитать файл: %v", err)
	}
	err = json.Unmarshal(f, &data)
	if err != nil {
		log.Fatalf("Ошибка при парсинге JSON: %v", err)
	}

	if month != 0 {
		for _, expense := range data {
			if expense.UpdatedAt[5:7] == strconv.Itoa(month) {
				summary += expense.Amount
			}
		}
		fmt.Printf("Total expenses in month %s: $%d", monthPrint[month], summary)
		os.Exit(0)
	}

	for _, expense := range data {
		summary += expense.Amount
	}
	fmt.Printf("Total expenses: $%d", summary)
	os.Exit(0)
}
