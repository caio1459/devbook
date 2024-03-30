package publicationrepositories

import "github.com/caio1459/devbook/src/models"

func (p publicationRepositorie) SelectPublication(id uint64) (models.Publications, error) {
	row, err := p.db.Query(`
		SELECT p.pub_id, p.title, p.content, p.image_url, p.likes, u.nick, u.user_id, u.image_url, p.register
		FROM publications p, users u 
		WHERE p.user_id = u.user_id AND p.pub_id = ?`, id,
	)
	if err != nil {
		return models.Publications{}, err
	}
	defer row.Close()

	publication := models.Publications{}
	if row.Next() {
		if err = row.Scan(
			&publication.PubID, &publication.Title, &publication.Content,
			&publication.ImageUrl, &publication.Likes, &publication.ActorNick,
			&publication.UserID, &publication.UserImage, &publication.Register,
		); err != nil {
			return models.Publications{}, err
		}
	}
	return publication, nil
}

//Busca as publicações do user logado e do seus seguidores
func (p publicationRepositorie) SelectPublications(id uint64) ([]models.Publications, error) {
	rows, err := p.db.Query(`
		SELECT DISTINCT p.pub_id, p.title, p.content, p.image_url, p.likes, p.user_id, u.nick, u.image_url, p.register
		FROM publications p
		JOIN users u ON p.user_id = u.user_id 
		JOIN followers f ON u.user_id = f.user_id 
		WHERE u.user_id = ? OR f.follower_id = ?
		ORDER BY p.register DESC`, id, id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	publications := []models.Publications{}
	for rows.Next() {
		publication := models.Publications{}
		if rows.Scan(
			&publication.PubID, &publication.Title, &publication.Content,
			&publication.ImageUrl, &publication.Likes, &publication.UserID,
			&publication.ActorNick, &publication.UserImage, &publication.Register,
		); err != nil {
			return nil, err
		}
		publications = append(publications, publication)
	}
	return publications, err
}

func (p publicationRepositorie) SelectPublicationsFromUser(id uint64) ([]models.Publications, error) {
	rows, err := p.db.Query(`
		SELECT p.pub_id, p.title, p.content, p.image_url, p.likes, p.user_id, u.nick, u.image_url, p.register 
		FROM publications p, users u
		WHERE p.user_id = u.user_id AND u.user_id = ?`, id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	publications := []models.Publications{}
	for rows.Next() {
		publication := models.Publications{}
		if err = rows.Scan(
			&publication.PubID, &publication.Title, &publication.Content,
			&publication.ImageUrl, &publication.Likes, &publication.UserID,
			&publication.ActorNick, &publication.UserImage, &publication.Register,
		); err != nil {
			return nil, err
		}
		publications = append(publications, publication)
	}
	return publications, nil
}