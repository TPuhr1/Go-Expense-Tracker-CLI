package cmd

import "fmt"

// Delete an expense by ID
func DeleteExpense(id int) {
	data, err := loadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}
	// Find the expense using ID
	var found bool
	// Iterrates through Expenses
	for i, expense := range data.Expenses {
		// Checks ID's to find a match
		if expense.ID == id {
			// Makes the appropriate changes and reorders ID numbers
			data.Expenses = append(data.Expenses[:i], data.Expenses[i+1:]...)
			found = true
			break
		}
	}
	// Print message if ID is not found in Expenses
	if !found {
		fmt.Println("Expense ID not found.")
	}
	// Re-index the remaining expenses
	for i := range data.Expenses {
		data.Expenses[i].ID = i + 1
	}
	// Save the updated list of expenses
	err = saveExpenses(data)
	if err != nil {
		fmt.Println("Error saving expenses:", err)
		return
	}
	// Confirmation message for completed deletion
	fmt.Println("Expense deleted successfully.")
}