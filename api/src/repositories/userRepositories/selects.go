package userrepositories

import (
	"fmt"

	"github.com/caio1459/devbook/src/models"
)

func (u usersRepositorie) SelectUsers(valueFilter string) ([]models.User, error) {
	valueFilter = fmt.Sprintf("%%%s%%", valueFilter) //Cria o LIKE %Valor%
	rows, err := u.db.Query(
		"SELECT user_id, name, nick, email, image_url, register FROM users WHERE name LIKE ? OR nick LIKE ?",
		valueFilter, valueFilter,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		if err = rows.Scan(
			&user.ID, &user.Name, &user.Nick, &user.Email,
			&user.ImageUrl, &user.Register,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u usersRepositorie) SelectUser(id uint64) (models.User, error) {
	row, err := u.db.Query(
		"SELECT user_id, name, nick, email, image_url,register FROM users WHERE user_id = ?", id,
	)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	user := models.User{}
	if row.Next() {
		if err := row.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.ImageUrl, &user.Register); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

func (u usersRepositorie) SelectUserFromEmail(email string) (models.User, error) {
	row, err := u.db.Query("SELECT * FROM users WHERE email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	user := models.User{}
	if row.Next() {
		if err = row.Scan(
			&user.ID, &user.Name, &user.Nick, &user.Email,
			&user.Password, &user.ImageUrl, &user.Register); err != nil {
			return models.User{}, err
		}
	}
	return user, nil
}

// Seleciona os seguidores de um usuário
func (u usersRepositorie) SelectFollows(id uint64) ([]models.User, error) {
	rows, err := u.db.Query(
		`SELECT u.name, u.nick, u.email, u.image_url, u.register 
		FROM users u, followers f
		WHERE u.user_id = f.follower_id AND f.user_id = ?`, id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		if err = rows.Scan(&user.Name, &user.Nick, &user.Email, &user.ImageUrl, &user.Register); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// Busca quem está seguindo
func (u usersRepositorie) SelectFollowing(id uint64) ([]models.User, error) {
	rows, err := u.db.Query(
		`SELECT u.name, u.nick, u.email, u.image_url, u.register 
		FROM users u, followers f
		WHERE u.user_id  = f.user_id AND f.follower_id  = ?`, id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		if err = rows.Scan(&user.Name, &user.Nick, &user.Email, &user.ImageUrl, &user.Register); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (u usersRepositorie) SelectUserPassword(id uint64) (string, error) {
	row, err := u.db.Query("SELECT password FROM users WHERE user_id = ?", id)
	if err != nil {
		return "", err
	}
	defer row.Close()

	user := models.User{}
	if row.Next() {
		if err = row.Scan(&user.Password); err != nil {
			return "", err
		}
	}
	return user.Password, nil
}
