package models

type Image struct {
	ID       uint64 `json:"image_id,omitempty"`
	FileName string `json:"file_name,omitempty"`
	UserID   uint64 `json:"user_id,omitempty"`
}
