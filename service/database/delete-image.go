package database

import (
	"database/sql"
	"fmt"
	"errors"
)

func (db *appdbimpl) DeleteImage(userId string, photoId string) (string, error) {
	var fileName string

	// Retrieve the filename of the image to be deleted, ensuring it belongs to the user
	getFileNameStmt := `SELECT file_name FROM images WHERE id = ? AND user_id = ?`
	err := db.c.QueryRow(getFileNameStmt, photoId, userId).Scan(&fileName)
	if err != nil {
		
		if errors.Is(err,sql.ErrNoRows) {
			return "", fmt.Errorf("no image found for photo ID %s and user ID %s: %w", photoId, userId, err)
		}
		return "", fmt.Errorf("error retrieving image filename: %w", err)
	}

	// Delete the image record from the database, ensuring it belongs to the user
	deleteStmt := `DELETE FROM images WHERE id = ? AND user_id = ?`
	result, err := db.c.Exec(deleteStmt, photoId, userId)
	if err != nil {
		return "", fmt.Errorf("error deleting image: %w", err)
	}

	// Check if a row was actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return "", fmt.Errorf("no image found with the given ID for this user")
	}

	return fileName, nil
}
