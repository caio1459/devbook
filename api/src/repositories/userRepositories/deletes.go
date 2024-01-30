package userrepositories

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

// Permite que um user deixe de seguir outro
func (u usersRepositorie) DeleteFollow(userID, followID uint64) error {
	statement, err := u.db.Prepare("DELETE FROM followers WHERE user_id = ? AND follower_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followID); err != nil {
		return err
	}
	return nil
}