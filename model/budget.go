package model

import ()

// Creating a struct to keep track of my budget
type Budget struct {
	Month int `json:"month"`
	Amount float64 `json:"amount"`
}