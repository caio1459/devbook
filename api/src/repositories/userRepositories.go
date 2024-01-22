package repositories

import (
	"database/sql"
	"fmt"

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

func (u usersRepositorie) SelectUsers(valueFilter string) ([]models.User, error) {
	valueFilter = fmt.Sprintf("%%%s%%", valueFilter) //Cria o LIKE %Valor%
	rows, err := u.db.Query(
		"SELECT user_id, name, nick, email, register FROM users WHERE name LIKE ? OR nick LIKE ?",
		valueFilter, valueFilter,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		if err = rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Register); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (u usersRepositorie) SelectUser(id uint64) (models.User, error) {
	row, err := u.db.Query("SELECT user_id, name, nick, email, register FROM users WHERE user_id = ?", id)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	user := models.User{}
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.Register); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (u usersRepositorie) UpdateUser(id uint64, user models.User) error {
	statement, err := u.db.Prepare("UPDATE users SET name = ?, nick = ?, email = ? WHERE user_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, id); err != nil {
		return err
	}

	return nil
}

func (u usersRepositorie) DeleteUser(id uint64) error {
	statemet, err := u.db.Prepare("DELETE FROM users WHERE user_id = ?")
	if err != nil {
		return err
	}
	defer statemet.Close()

	_, err = statemet.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
