package userrepositories

import (
	"database/sql"
)

// Struct com os métodos de crud
type usersRepositorie struct {
	db *sql.DB
}

// Recebe um banco de dados que é aberto no controller
func NewRepositorieUser(db *sql.DB) *usersRepositorie {
	return &usersRepositorie{db}
}