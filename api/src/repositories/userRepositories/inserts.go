package userrepositories

import "github.com/caio1459/devbook/src/models"

func (u usersRepositorie) InsertUser(user models.User) (uint64, error) {
	statement, err := u.db.Prepare(
		"INSERT INTO users (name, nick, email, password, image_url) VALUES(?, ?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	results, err := statement.Exec(user.Name, user.Nick, user.Email, user.Password, user.ImageUrl)
	if err != nil {
		return 0, err
	}

	lastID, err := results.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastID), err
}

// Permite que um user siga outro
func (u usersRepositorie) InsertFollow(userID, followID uint64) error {
	statement, err := u.db.Prepare("INSERT IGNORE INTO followers (user_id, follower_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(userID, followID); err != nil {
		return err
	}
	return nil
}
