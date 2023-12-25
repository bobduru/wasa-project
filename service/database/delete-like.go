package database

import (
	"fmt"
	"strconv"
)

func (db *appdbimpl) DeleteLike(userId string, photoId string) error {
	// Convert string IDs to integers
	userIDInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid user ID: %v", err)
	}

	photoIDInt, err := strconv.ParseInt(photoId, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid photo ID: %v", err)
	}

	// SQL statement to delete the like record, ensuring it belongs to the user
	deleteStmt := `DELETE FROM likes WHERE user_id = ? AND image_id = ?`

	// Executing the SQL statement to delete
	result, err := db.c.Exec(deleteStmt, userIDInt, photoIDInt)
	if err != nil {
		return fmt.Errorf("error deleting like: %w", err)
	}

	// Check if a like was actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no like found with the given photo ID and user ID")
	}

	return nil
}
