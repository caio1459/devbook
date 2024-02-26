package commentrepositories

import "github.com/caio1459/devbook/src/models"

func (c *commentRepositorie) SelectComments(id uint64) ([]models.Comment, error) {
	rows, err := c.db.Query(`
		SELECT c.pub_id, c.user_id, u.name, u.image_url, c.text, c.register
		FROM comments c, users u
		WHERE c.user_id = u.user_id AND c.pub_id = ?`, id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	comments := []models.Comment{}
	for rows.Next() {
		comment := models.Comment{}
		if rows.Scan(
			&comment.PubID, &comment.UserID, &comment.ActorNick,
			&comment.UserImage, &comment.Text, &comment.Register,
		); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, err
}
