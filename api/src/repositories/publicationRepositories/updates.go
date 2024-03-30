package publicationrepositories

import "github.com/caio1459/devbook/src/models"

func (p publicationRepositorie) UpdatePublication(id uint64, publication models.Publications) error {
	statemet, err := p.db.Prepare("UPDATE publications SET title = ?, content = ? WHERE pub_id = ?")
	if err != nil {
		return err
	}
	defer statemet.Close()

	if _, err = statemet.Exec(publication.Title, publication.Content, id); err != nil {
		return err
	}
	return nil
}

func (p publicationRepositorie) UpdateLikes(id uint64) (uint64, error) {
	statement, err := p.db.Prepare("UPDATE publications SET likes = likes + 1 WHERE pub_id = ?")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		return 0, err
	}
	// Consulta para obter o número atualizado de likes
	var updatedLikes uint64
	row := p.db.QueryRow("SELECT likes FROM publications WHERE pub_id = ?", id)
	if err := row.Scan(&updatedLikes); err != nil {
		return 0, err
	}
	return updatedLikes, nil
}

func (p publicationRepositorie) UpdateDeslikes(id uint64) (uint64, error) {
	statemet, err := p.db.Prepare(`
		UPDATE publications SET likes = 
			CASE 
				WHEN likes > 0 THEN likes - 1 
				ELSE 0 
			END
		WHERE pub_id = ?`,
	)
	if err != nil {
		return 0, err
	}
	defer statemet.Close()

	if _, err = statemet.Exec(id); err != nil {
		return 0, err
	}
	// Consulta para obter o número atualizado de likes
	var updatedLikes uint64
	row := p.db.QueryRow("SELECT likes FROM publications WHERE pub_id = ?", id)
	if err := row.Scan(&updatedLikes); err != nil {
		return 0, err
	}
	return updatedLikes, nil
}
