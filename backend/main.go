package main

import (
	"log"
	"net/http"

	"expense-tracker/db"
	"expense-tracker/handlers"
)

func main() {
	database := db.InitDB()

	expenseHandler := &handlers.ExpenseHandler{
		DB: database,
	}

	http.HandleFunc("/api/expenses", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			expenseHandler.AddExpense(w, r)

		case http.MethodGet:
			expenseHandler.ListExpenses(w, r)

		case http.MethodDelete:
			expenseHandler.DeleteExpense(w, r)

		default:
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Println("Backend running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
