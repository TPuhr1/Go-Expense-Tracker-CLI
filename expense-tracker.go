package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"
)

// Expense structure
type Expense struct {
	ID 			int		`json:"id"`
	Date 		string	`json:"date"`
	Description string	`json:"description"`
	Amount 		float64	`json:"amount"`
	Category 	string  `json:"category"`
}
// Creating a struct to hold the expenses list
type ExpenseData struct {
	Expenses []Expense `json:"expenses"`
}
// Creating a struct to keep track of my budget
type Budget struct {
	Month int `json:"month"`
	Amount float64 `json:"amount"`
}
// Creating variables for my files
const expenseFile = "expenses.json"
const budgetFile = "budget.json"

// Load expenses from the file
func loadExpenses() (ExpenseData, error) {
	var data ExpenseData
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
func loadBudget() (map[int]Budget, error) {
	var data map[int]Budget
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
func saveExpenses(data ExpenseData) error {
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
func saveBudget(data map[int]Budget) error {
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
// Add a new expense
func addExpense(description string, amount float64, category string){
	data, err := loadExpenses()
	if err != nil {
		fmt.Println("Error loading expenses:", err)
		return
	}
	// Get the next expense ID
	expenseID := len(data.Expenses) + 1
	// Create a new expense
	expense := Expense {
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
	defer checkBudget(int(time.Month(expenseDate.Month())))
	// Confirmation message for completing the addition
	fmt.Printf("Expense added successfully (ID: %d)\n", expenseID)
}
// Set the monthly budget
func setBudget(month int, amount float64) {
	budgetData, err := loadBudget()
	if err != nil {
		fmt.Println("Error loading budget:", err)
		return
	}
	// Initialize the map if it is nil
	if budgetData == nil {
		budgetData = make(map[int]Budget)
	}
	// Set the new budget
	budgetData[month] = Budget{
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
// Checks monthly budget and warns user if exceeded
func checkBudget(month int) {
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
// List all expenses
func listExpenses(category string) {
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
// Delete an expense by ID
func deleteExpense(id int) {
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
// Display summary of expenses
func summary(month int) {
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
// Exports your expenses to a CSV file
func exportCSV() {
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

func main() {
	// Define flags for command-line arguments
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	desc := addCmd.String("description", "", "Description of the expense")
	amount := addCmd.Float64("amount", 0, "Amount of the expense")
	categ := addCmd.String("category", "", "Category of the expense")

	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listCat := listCmd.String("category", "", "Lists expenses of specified category")

	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	delID := deleteCmd.Int("id", 0, "ID of the expense to delete")

	summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
	monthSummaryFlag := summaryCmd.Int("month", 0, "Month for the summary (1-12)")

	budgetCmd := flag.NewFlagSet("budget", flag.ExitOnError)
	budgetMonth := budgetCmd.Int("month", 0, "Month for the budget (1-12)")
	budgetAmount := budgetCmd.Float64("amount", 0, "Amount for the budget")

	exportCmd := flag.NewFlagSet("export", flag.ExitOnError)

	// Parse the command-line arguments
	if len(os.Args) < 2 {
		// Prints message if there are missing arguments
		fmt.Println("Usage: expense-tracker <command> [arguments]")
		return
	}
	// Continues if there are a correct number of arguments
	switch os.Args[1] {
	// Adds new expenses to your Expenses file
	case "add":
		addCmd.Parse(os.Args[2:])
		// Checks to make sure that all of the flags are present
		if *desc == "" || *amount <=0 || *categ == "" {
			// If they are not, then prints this message to inform the user
			fmt.Println("Please provide a valid description, amount for the expense, and the category of the expense.")
			return
		}
		// Adds the expense if everything is present
		addExpense(*desc, *amount, *categ)
	// Lists all expenses or just ones from a specific category
	case "list":
		listCmd.Parse(os.Args[2:])
		listExpenses(*listCat)
	// Deletes an expense by ID
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		// Checks for an ID
		if *delID == 0 {
			// If no ID is provided, informs the user
			fmt.Println("Please provide an ID to delete.")
			return
		}
		deleteExpense(*delID)
	// Prints a summary of expenses
	case "summary":
		summaryCmd.Parse(os.Args[2:])
		summary(*monthSummaryFlag)
	// Set or check a monthly budget
	case "budget":
		budgetCmd.Parse(os.Args[2:])
		if *budgetMonth <= 0 || *budgetAmount <= 0 {
			fmt.Println("Please provide a valid month (1-12) and amount for the budget.")
			return
		}
		// set the budget
		setBudget(*budgetMonth, *budgetAmount)
	// Check if budget is exceeded for a month
	case "check-budget":
		budgetCmd.Parse(os.Args[2:])
		if *budgetMonth <= 0 {
			fmt.Println("Please provide a valid month (1-12) to check the budget.")
			return
		}
		// Check if the budget is exceeded
		checkBudget(*budgetMonth)
	// Export expenses to a CSV file
	case "export":
		exportCmd.Parse(os.Args[2:])
		// Export expenses to CSV
		exportCSV()
	// If there is a command that is not included
	default:
		// Prints message informing the user of unknown command
		fmt.Println("Unknown command:", os.Args[1])
	}
}