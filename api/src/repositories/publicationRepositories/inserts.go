package publicationrepositories

import "github.com/caio1459/devbook/src/models"

func (p publicationRepositorie) InsertPublication(publication models.Publications) (uint64, error) {
	statemet, err := p.db.Prepare("INSERT INTO publications (user_id, title, content, image_url) VALUES (?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statemet.Close()

	results, err := statemet.Exec(publication.UserID, publication.Title, publication.Content, publication.ImageUrl)
	if err != nil {
		return 0, err
	}

	lastID, err := results.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(lastID), err
}
