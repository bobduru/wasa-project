package database

import (
	"database/sql"
	"fmt"
)

func (db *appdbimpl) DeleteImage(photoId string) (string, error) {
	var fileName string
	getFileNameStmt := `SELECT file_name FROM images WHERE id = ?`
	err := db.c.QueryRow(getFileNameStmt, photoId).Scan(&fileName)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no image found for photo ID %s: %w", photoId, err)
		}
		return "", fmt.Errorf("error retrieving image filename: %w", err)
	}

	deleteStmt := `DELETE FROM images WHERE id = ?`
	_, err = db.c.Exec(deleteStmt, photoId)
	if err != nil {
		return "", fmt.Errorf("error deleting image: %w", err)
	}

	return fileName, nil
}
