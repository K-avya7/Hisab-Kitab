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
		if r.Method == http.MethodPost {
			expenseHandler.AddExpense(w, r)
			return
		}
		if r.Method == http.MethodGet {
			expenseHandler.ListExpenses(w, r)
			return
		}
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	})

	log.Println("Backend running on http://localhost:8080")
	log.Println("testing the new github workflow for ai agent review")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
