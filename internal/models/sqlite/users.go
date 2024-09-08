package sqlite

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

type UserModel struct {
	DB *sql.DB
}

func (m *UserModel) Insert(name, email, password string) error {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	stmt := `INSERT INTO users (name, email, password)
	VALUES (?, ?, ?)`

	_, err = m.DB.Exec(stmt, name, email, passwordHash)
	if err != nil {
		return err
	}

	return nil
}