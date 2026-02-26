package main

import (
	"expensesapi/database"
	"expensesapi/services"
	"log"
	"net/http"
)

func main() {
	database.GetConnection()
	http.HandleFunc("/expenses/", func(w http.ResponseWriter, r* http.Request){
		log.Println("HIT /expenses", r.Method)
		switch r.Method{
		case http.MethodGet:
			services.GetExpenses(w,r)
		case http.MethodPost:
			services.CreateExpense(w,r)
		default:
			http.Error(w, "method not allowed", 405)

		}
		
	})
	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}