package models

import "time"

type Expense struct {
	ExpenseID       int    		`json:"expenses_id"`
	CategoryID      int    		`json:"category_id"`
	PaymentMethodID int    		`json:"payment_id"`
	Amount          int    		`json:"amount"`
	Description     string 		`json:"description"`
	ExpenseDate     time.Time 	`json:"expense_date"`
}