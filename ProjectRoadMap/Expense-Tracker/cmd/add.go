package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strconv"
	"time"
)

// Create new var
var amount int
var description string

type Expense struct {
	Id          string `json:"id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new expense",
	Long:  `Add a new expense with a specified amount and category`,

	Run: func(cmd *cobra.Command, args []string) {
		// Логика добавления расходов
		addExpense()
	},
}

// init function is used flag to command
func init() {
	rootCmd.AddCommand(addCmd)

	// Adding flags for the add command
	addCmd.Flags().StringVarP(&description, "description", "d", "", "description of the expense")
	addCmd.Flags().IntVarP(&amount, "amount", "a", 0, "Expense amount")
}

// addExpense create a new expense
func addExpense() {

	// Logic for adding expenses
	_, err := os.OpenFile("expenses.json", os.O_WRONLY|os.O_RDWR, 0777)
	if err != nil {
		if os.IsNotExist(err) {
			CreateNewFileAndAddExpense()
		}
	}
	CreateNewAddExpense()

}

func CreateNewFileAndAddExpense() {

	now := time.Now().Format("2006-01-02 15:04:05")

	var writerData = []Expense{
		{
			Id:          "1",
			Description: description,
			Amount:      amount,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}

	f, err := os.Create("expenses.json")
	if err != nil {
		log.Fatalf("Ошибка при создании файла: %v", err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatalf("Ошибка при закрытии файла: %v", err)
		}
	}(f)

	// Convert writerMap to JSON string with correct field order
	jsonData, err := json.MarshalIndent(writerData, "", "  ")
	if err != nil {
		log.Fatalf("Ошибка при преобразовании в JSON: %v", err)
	}

	_, err = f.WriteString(string(jsonData))
	if err != nil {
		log.Fatalf("Ошибка при записи в файл: %v", err)
	}

	fmt.Print("Expense added successfully (ID: 1)")
	os.Exit(0)

}

func CreateNewAddExpense() {

	// Create a new expense
	var data []Expense

	// Read the existing data from the JSON file
	f, err := os.ReadFile("expenses.json")
	if err != nil {
		log.Fatalf("Не удалось прочитать файл: %v", err)
	}
	err = json.Unmarshal(f, &data)
	if err != nil {
		log.Fatalf("Ошибка при парсинге JSON: %v", err)
	}

	// Generate new ID
	newID := strconv.Itoa(len(data) + 1)
	now := time.Now().Format("2006-01-02 15:04:05")
	// New expense
	var writerData = []Expense{
		{
			Id:          newID,
			Description: description,
			Amount:      amount,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}
	data = append(data, writerData...)

	// Encoding the updated data in JSON
	updatedData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Ошибка при кодировании JSON: %v", err)
	}

	err = os.WriteFile("expenses.json", updatedData, 0644)
	if err != nil {
		log.Fatalf("Ошибка при записи в файл: %v", err)
	}

	fmt.Printf("Expense added successfully (ID: %v)", newID)
	os.Exit(0)
}

// Execute adds all child commands to the root command and sets flags appropriately.
