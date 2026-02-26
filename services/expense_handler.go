package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"expensesapi/database"
	"log"
	"net/http"
	"time"
)

type ExpenseView struct {
	ExpenseID         int       `json:"expense_id"`
	CategoryName      string    `json:"category_name"`
	PaymentMethodName string    `json:"payment_name"`
	Amount            int       `json:"amount"`
	Description       string    `json:"description"`
	ExpenseDate       time.Time `json:"expense_date"`
}

type ExpenseInput struct {
	CategoryName      string `json:"category_name"`
	PaymentMethodName string `json:"payment_name"`
	Amount            int    `json:"amount"`
	Description       string `json:"description"`
	ExpenseDate       string `json:"expense_date"`
}

func GetExpenses(w http.ResponseWriter, r *http.Request) {
	log.Println("test")
	ctx := context.Background()
	script := `
		SELECT 
			e.expenses_id, 
			c.category_name,
			pm.payment_name,
			e.amount,
			e.description,
			e.expense_date
		FROM 
			expenses e JOIN categories c 
			ON e.category_id = c.category_id
			JOIN payment_methods pm
			ON e.payment_id = pm.payment_id
	`

	rows, err := database.DB.QueryContext(ctx, script)
	if err != nil {
		http.Error(w, err.Error(), 500)
	}
	defer rows.Close()
	var expenses []ExpenseView
	for rows.Next() {
		var e ExpenseView
		rows.Scan(&e.ExpenseID, &e.CategoryName, &e.PaymentMethodName, &e.Amount, &e.Description, &e.ExpenseDate)
		expenses = append(expenses, e)
	}

	json.NewEncoder(w).Encode(expenses)
	log.Println("GET SUCCESS")
}

func CreateExpense(w http.ResponseWriter, r *http.Request) {
	var temp ExpenseInput
	reqErr := json.NewDecoder(r.Body).Decode(&temp)
	if reqErr != nil {
		http.Error(w, "invalid body request", 400)
		return
	}

	expense_date, err := time.Parse("2006-01-02", temp.ExpenseDate)

	if err != nil {
		http.Error(w, "invalid date format", http.StatusBadRequest)
	}
	var category_id int
	var payment_id int
	ctx := context.Background()
	script1 := `
			SELECT 
				category_id
			FROM 
				categories
			WHERE category_name = ?
		`
	err1 := database.DB.QueryRowContext(ctx, script1, temp.CategoryName).Scan(&category_id)

	if err1 == sql.ErrNoRows {
		http.Error(w, "category not found", http.StatusBadRequest)
		return
	}
	log.Println(category_id)

	script2 := `
			SELECT 
				payment_id
			FROM 
				payment_methods
			WHERE payment_name = ?
		`
	err2 := database.DB.QueryRowContext(ctx, script2, temp.PaymentMethodName).Scan(&payment_id)

	if err2 == sql.ErrNoRows {
		http.Error(w, "payment method not found", http.StatusBadRequest)
		return
	}
	log.Println(payment_id)

	scriptInput := "INSERT INTO expenses (category_id, payment_id, amount, description, expense_date) VALUES(?,?,?,?,?)"

	_, err3 := database.DB.ExecContext(ctx, scriptInput, category_id, payment_id, temp.Amount, temp.Description, expense_date)

	if err3 != nil {
		http.Error(w, err3.Error(), 500)
		return
	}
	w.WriteHeader(http.StatusCreated)
	log.Println("POST SUCCESS")

}
