package cmd

import "fmt"

// List all expenses
func ListExpenses(category string) {
	data, err := loadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}
	// Prints message if there are no expenses
	if len(data.Expenses) == 0 {
		fmt.Println("No expenses recorded.")
		return
	}
	// Checks for a category flag
	if category != "" {
		// If there is it will address the category
		fmt.Printf("Here are all of your %s expenses:\n", category)
		fmt.Printf("%-4s %-12s %-15s %-11s %-15s\n", "ID", "Date", "Description", "Amount", "Category")
		// Iterrating through Expenses
		for _, expense := range data.Expenses {
			// Searching for category matches
			if expense.Category == category {
				// Prints matches found in Expenses
				fmt.Printf("%-5d %-12s %-15s $%-10.2f %-15s\n", expense.ID, expense.Date, expense.Description, expense.Amount, expense.Category)
			}
		}
	// If there is no category flag
	} else {
		fmt.Printf("%-4s %-12s %-15s %-11s %-15s\n", "ID", "Date", "Description", "Amount", "Category")
		// Iterrates through Expenses
		for _, expense := range data.Expenses {
			// Prints all expenses
			fmt.Printf("%-5d %-12s %-15s $%-10.2f %-15s\n", expense.ID, expense.Date, expense.Description, expense.Amount, expense.Category)
		}
	}
}