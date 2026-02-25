package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "./db/expenses.db")
	if err != nil {
		log.Fatal("failed to open db:", err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS expenses (
		id TEXT PRIMARY KEY,
		amount INTEGER NOT NULL,
		category TEXT NOT NULL,
		description TEXT,
		expense_date DATE NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`

	if _, err := db.Exec(createTable); err != nil {
		log.Fatal("failed to create table:", err)
	}

	return db
}
