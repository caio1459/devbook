package commentrepositories

import "github.com/caio1459/devbook/src/models"

func (c *commentRepositorie) InsertComment(comment models.Comment) (uint64, error) {
	statemet, err := c.db.Prepare("INSERT INTO comments (user_id, pub_id, text) VALUES (?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statemet.Close()

	results, err := statemet.Exec(comment.UserID, comment.PubID, comment.Text)
	if err != nil {
		return 0, err
	}

	lastID, err := results.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastID), err
}
