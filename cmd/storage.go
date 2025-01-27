package cmd

import (
	"encoding/json"
	"os"
	model "Go-Expense-Tracker-CLI/model"
)
// Creating variables for my files
const expenseFile = "expenses.json"
const budgetFile = "budget.json"
// Load expenses from the file
func loadExpenses() (model.ExpenseData, error) {
	var data model.ExpenseData
	// Check if the file exists
	_, err := os.Stat(expenseFile)
	if os.IsNotExist(err) {
		// Return an empty ExpenseData struct if the file doesn't exist
		return data, nil
	}
	// Opens the file if it already exists
	file, err := os.Open(expenseFile)
	if err != nil {
		return data, err
	}
	defer file.Close()
	// Decoder
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
// Load budget from a file
func loadBudget() (map[int]model.Budget, error) {
	var data map[int]model.Budget
	// Check if the file exists
	_, err := os.Stat(budgetFile)
	if os.IsNotExist(err){
		// Return an empty map if the file doesn't exist
		return data, nil
	}
	// Open the file if it already exists
	file, err := os.Open(budgetFile)
	if err != nil {
		return data, err
	}
	defer file.Close()
	// Decoder
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&data)
	if err != nil {
		return data, err
	}
	return data, nil
}
// Save expenses to the file
func saveExpenses(data model.ExpenseData) error {
	file, err := os.Create("expenses.json")
	if err != nil {
		return err
	}
	defer file.Close()
	// Encoder
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " 	")
	return encoder.Encode(data)
}
// Save budget to the file
func saveBudget(data map[int]model.Budget) error {
	file, err := os.Create(budgetFile)
	if err != nil {
		return err
	}
	defer file.Close()
	// Encoder
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "		")
	return encoder.Encode(data)
}