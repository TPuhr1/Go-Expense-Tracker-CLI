package model

import ()

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