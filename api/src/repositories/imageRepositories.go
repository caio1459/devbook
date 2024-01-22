package repositories

import (
	"database/sql"

	"github.com/caio1459/devbook/src/models"
)

type imagesRepositorie struct {
	db *sql.DB
}

func NewRepositorieImage(db *sql.DB) *imagesRepositorie {
	return &imagesRepositorie{db}
}

func (i imagesRepositorie) PostImage(image models.Image, id uint64) (uint64, error) {
	statement, err := i.db.Prepare("INSERT INTO images (filename, user_id) VALUES(?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	_, err = statement.Exec(image.FileName, id)
	if err != nil {
		return 0, err
	}
	return id, err
}
