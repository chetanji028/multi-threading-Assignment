package repository

import (
	"database/sql"

	"github.com/google/uuid"
)

type FileRepository interface {
	SaveFileParts(fileID uuid.UUID, parts [][]byte) error
	GetFileParts(fileID uuid.UUID) ([][]byte, error)
}

type fileRepository struct {
	db *sql.DB
}

func NewFileRepository(db *sql.DB) FileRepository {
	return &fileRepository{db: db}
}

func (r *fileRepository) SaveFileParts(fileID uuid.UUID, parts [][]byte) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare("INSERT INTO files (file_id, part_number, data) VALUES ($1, $2, $3)")
	if err != nil {
		tx.Rollback()
		return err
	}
	defer stmt.Close()

	for i, part := range parts {
		_, err := stmt.Exec(fileID, i, part)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *fileRepository) GetFileParts(fileID uuid.UUID) ([][]byte, error) {
	rows, err := r.db.Query("SELECT part_number, data FROM files WHERE file_id = $1 ORDER BY part_number", fileID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var parts [][]byte
	for rows.Next() {
		var partNumber int
		var data []byte
		if err := rows.Scan(&partNumber, &data); err != nil {
			return nil, err
		}
		parts = append(parts, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return parts, nil
}
