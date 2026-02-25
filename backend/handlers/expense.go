package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

type ExpenseHandler struct {
	DB *sql.DB
}

type addExpenseRequest struct {
	Amount      int    `json:"amount"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Date        string `json:"date"`
}

func (h *ExpenseHandler) AddExpense(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req addExpenseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}

	if req.Amount <= 0 || req.Category == "" || req.Date == "" {
		http.Error(w, "missing required fields", http.StatusBadRequest)
		return
	}

	_, err := h.DB.Exec(
		`INSERT INTO expenses (id, amount, category, description, expense_date)
		 VALUES (?, ?, ?, ?, ?)`,
		uuid.New().String(),
		req.Amount,
		req.Category,
		req.Description,
		req.Date,
	)

	if err != nil {
		http.Error(w, "failed to insert expense", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
func (h *ExpenseHandler) ListExpenses(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	rows, err := h.DB.Query(`
		SELECT amount, category, description, expense_date
		FROM expenses
		ORDER BY expense_date DESC
	`)
	if err != nil {
		http.Error(w, "failed to fetch expenses", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	type Expense struct {
		Amount      int    `json:"amount"`
		Category    string `json:"category"`
		Description string `json:"description"`
		Date        string `json:"date"`
	}

	var expenses []Expense

	for rows.Next() {
		var e Expense
		if err := rows.Scan(&e.Amount, &e.Category, &e.Description, &e.Date); err != nil {
			http.Error(w, "failed to read expenses", http.StatusInternalServerError)
			return
		}
		expenses = append(expenses, e)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(expenses)
}
