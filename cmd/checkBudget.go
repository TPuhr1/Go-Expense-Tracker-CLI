package cmd

import (
	"fmt"
	"time"
)

// Checks monthly budget and warns user if exceeded
func CheckBudget(month int) {
	// load expenses
	data, err := loadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}
	// load budget
	budgetData, err := loadBudget()
	if err != nil {
		fmt.Println("Error loading budget:", err)
		return
	}
	// Check if a budget is set for the month
	budget, exists := budgetData[month]
	if !exists {
		fmt.Printf("No budget set for month %d\n", month)
		return
	}
	// Creating a variable to hold total expenses
	var total float64
	for _, expense := range data.Expenses {
		// Check for expenses in the right month
		expenseDate, err := time.Parse("2006-01-02", expense.Date)
		if err != nil {
			fmt.Println("Error parsing date:", err)
			return
		}
		// Once the expense belongs to the month, it is added to the total
		if expenseDate.Month() == time.Month(month) {
			total += expense.Amount
		}
	}
	// Comparing the total to the budget
	if total > budget.Amount {
		// If total exceeds budget
		fmt.Printf("Warning: You have exceeded your budget for month %d!\n Total: $%.2f, Budget: $%.2f\n", month, total, budget.Amount)
		// If total is still within the budget
		} else {
		fmt.Printf("Total expenses for month %d: $%.2f (Budget: $%.2f)\n", month, total, budget.Amount)
	}
}