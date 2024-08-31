package database

import (
	"database/sql"
)

func InsertUser(db *sql.DB, username, email string) error {
    _, err := db.Exec("INSERT INTO users (username, email) VALUES (?, ?)", username, email)
    return err
}

// Add other user-related database functions here, such as:
// GetUser, UpdateUser, DeleteUser, etc.