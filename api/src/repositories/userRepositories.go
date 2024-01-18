package repositories

import (
	"database/sql"

	"github.com/caio1459/devbook/src/models"
)

// Struct com os métodos de crud
type usersRepositorie struct {
	db *sql.DB
}

// Recebe um banco de dados que é aberto no controller
func NewRepositorieUser(db *sql.DB) *usersRepositorie {
	return &usersRepositorie{db}
}

func (u usersRepositorie) PostUser(user models.User) (uint64, error) {
	statement, err := u.db.Prepare(
		"INSERT INTO users (name, nick, email, password) VALUES(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	results, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastID, err := results.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastID), err
}
