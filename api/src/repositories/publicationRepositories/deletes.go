package publicationrepositories

func (p publicationRepositorie) DeletePublication(id uint64) error {
	statement, err := p.db.Prepare("DELETE FROM publications WHERE pub_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		return err
	}
	return nil
}
