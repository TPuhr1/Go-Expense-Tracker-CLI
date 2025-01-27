package cmd

import (
	"fmt"
	"time"
	model "Go-Expense-Tracker-CLI/model"
)

// Add a new expense
func AddExpense(description string, amount float64, category string){
	data, err := loadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}
	// Get the next expense ID
	expenseID := len(data.Expenses) + 1
	// Create a new expense
	expense := model.Expense {
		ID: 			expenseID,
		Description:	description,
		Amount:			amount,
		Date:			time.Now().Format("2006-01-02"),
		Category: 		category,
	}
	// Adds new expense to Expenses
	data.Expenses = append(data.Expenses, expense)
	// Saves new addition
	err = saveExpenses(data)
	if err != nil {
		fmt.Println("Error saving expense:", err)
		return
	}
	// Parsing date to get the month
	expenseDate, err := time.Parse("2006-01-02", expense.Date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return
	}
	// Checks the expense against current month's budget
	defer CheckBudget(int(time.Month(expenseDate.Month())))
	// Confirmation message for completing the addition
	fmt.Printf("Expense added successfully (ID: %d)\n", expenseID)
	return
}