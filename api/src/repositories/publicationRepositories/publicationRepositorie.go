package publicationrepositories

import "database/sql"

type publicationRepositorie struct {
	db *sql.DB
}

func NewRepositoriePublication(db *sql.DB) *publicationRepositorie {
	return &publicationRepositorie{db}
}
