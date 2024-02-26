package commentrepositories

import "database/sql"

type commentRepositorie struct {
	db *sql.DB
}

func NewCommentRepositorie(db *sql.DB) *commentRepositorie {
	return &commentRepositorie{db}
}
