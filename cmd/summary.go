package cmd

import (
	"fmt"
	"time"
)

// Display summary of expenses
func Summary(month int) {
	data, err := loadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}
	// Creating a variable to hold the total 
	var total float64
	// Iterrating through Expenses
	for _, expense := range data.Expenses {
		// Checks for a month flag
		if month > 0 {
			// Parses Expense Dates for comarison
			expenseDate, err := time.Parse("2006-01-02", expense.Date)
			if err != nil {
				fmt.Println("Error parsing date:", err)
				return
			}
			// Checks month with parsed Expense Date 
			if expenseDate.Month() == time.Month(month) {
				// Adds to the total once matched
				total += expense.Amount
			}
		// Adds all if there is no month flag
		} else {
			total += expense.Amount
		}
	}
	// Special format displaying month if the flag is present
	if month > 0 {
		fmt.Printf("Total expenses for month %d: $%.2f\n", month, total)
	// Prints out the total for all if no month flag is present
	} else {
		fmt.Printf("Total expenses: $%.2f\n", total)
	}
}