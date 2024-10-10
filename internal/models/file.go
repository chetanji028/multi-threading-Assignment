package models

import "time"

type FilePart struct {
	ID         int       `json:"id"`
	FileID     string    `json:"file_id"`
	PartNumber int       `json:"part_number"`
	Data       []byte    `json:"data"`
	CreatedAt  time.Time `json:"created_at"`
}
