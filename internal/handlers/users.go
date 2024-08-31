package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/MuhammadISayyed/tov.git/internal/database"
)

func CreateUser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var user struct {
            Username string `json:"username"`
            Email    string `json:"email"`
        }

        if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        if err := database.InsertUser(db, user.Username, user.Email); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
    }
}

// Add other user-related handlers here