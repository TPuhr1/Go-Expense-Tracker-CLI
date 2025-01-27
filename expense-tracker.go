package main

import (
	"flag"
	"fmt"
	"os"
	cmd "Go-Expense-Tracker-CLI/cmd"
)

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
		cmd.AddExpense(*desc, *amount, *categ)
	// Lists all expenses or just ones from a specific category
	case "list":
		listCmd.Parse(os.Args[2:])
		cmd.ListExpenses(*listCat)
	// Deletes an expense by ID
	case "delete":
		deleteCmd.Parse(os.Args[2:])
		// Checks for an ID
		if *delID == 0 {
			// If no ID is provided, informs the user
			fmt.Println("Please provide an ID to delete.")
			return
		}
		cmd.DeleteExpense(*delID)
	// Prints a summary of expenses
	case "summary":
		summaryCmd.Parse(os.Args[2:])
		cmd.Summary(*monthSummaryFlag)
	// Set or check a monthly budget
	case "budget":
		budgetCmd.Parse(os.Args[2:])
		if *budgetMonth <= 0 || *budgetAmount <= 0 {
			fmt.Println("Please provide a valid month (1-12) and amount for the budget.")
			return
		}
		// set the budget
		cmd.SetBudget(*budgetMonth, *budgetAmount)
	// Check if budget is exceeded for a month
	case "check-budget":
		budgetCmd.Parse(os.Args[2:])
		if *budgetMonth <= 0 {
			fmt.Println("Please provide a valid month (1-12) to check the budget.")
			return
		}
		// Check if the budget is exceeded
		cmd.CheckBudget(*budgetMonth)
	// Export expenses to a CSV file
	case "export":
		exportCmd.Parse(os.Args[2:])
		// Export expenses to CSV
		cmd.ExportCSV()
	// If there is a command that is not included
	default:
		// Prints message informing the user of unknown command
		fmt.Println("Unknown command:", os.Args[1])
	}
}