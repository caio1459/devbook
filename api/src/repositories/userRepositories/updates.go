package userrepositories

import "github.com/caio1459/devbook/src/models"

func (u usersRepositorie) UpdateUser(id uint64, user models.User) error {
	statement, err := u.db.Prepare(
		"UPDATE users SET name = ?, nick = ?, email = ?, image_url = ? WHERE user_id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, user.ImageUrl, id); err != nil {
		return err
	}
	return nil
}

func (u usersRepositorie) UpdatePassword(id uint64, password string) error {
	statemet, err := u.db.Prepare("UPDATE users SET password = ? WHERE user_id = ?")
	if err != nil {
		return err
	}
	defer statemet.Close()
	
	if _, err = statemet.Exec(password, id); err != nil {
		return err
	}
	return nil
}
