package models

type Image struct {
	ID       uint64 `json:"image_id,omitempty"`
	FileName string `json:"file_name,omitempty"`
	Size     int8   `json:"size,omitempty"`
}
