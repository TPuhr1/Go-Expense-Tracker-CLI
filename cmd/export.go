package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Exports your expenses to a CSV file
func ExportCSV() {
	// Create CSV file
	file, err := os.Create("expenses.csv")
	if err != nil {
		fmt.Println("Error creating CSV file:", err)
		return
	}
	defer file.Close()
	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// Load data
	data, err := loadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}
	// Create a header for the CSV file
	writer.Write([]string{"ID", "Date", "Description", "Amount", "Category"})
	// Copy all of expenses into the variable record
	for _, expense := range data.Expenses {
		record := []string{
			strconv.Itoa(expense.ID),
			expense.Date,
			expense.Description,
			fmt.Sprintf("%.2f", expense.Amount),
			expense.Category,
		}
		// Write all of record to expenses.csv
		writer.Write(record)
	}
	// Confirmation message for completion of writing
	fmt.Println("Expenses exported to expenses.csv")
}