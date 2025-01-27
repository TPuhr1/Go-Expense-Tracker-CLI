package cmd

import (
	"fmt"
	model "Go-Expense-Tracker-CLI/model"
)

// Set the monthly budget
func SetBudget(month int, amount float64) {
	budgetData, err := loadBudget()
	if err != nil {
		fmt.Println("Error loading budget:", err)
		return
	}
	// Initialize the map if it is nil
	if budgetData == nil {
		budgetData = make(map[int]model.Budget)
	}
	// Set the new budget
	budgetData[month] = model.Budget{
		Month: month,
		Amount: amount,
	}
	// Save the updated budget
	err = saveBudget(budgetData)
	if err != nil {
		fmt.Println("Error saving budget:", err)
		return
	}
	// Confirmation message for setting the budget
	fmt.Printf("Budget for month %d set to $%.2f\n", month, amount)
}