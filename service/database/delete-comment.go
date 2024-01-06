package database

import (
	"fmt"
	"strconv"
)

func (db *appdbimpl) DeleteComment(userId string, commentId string) error {
	// Convert string IDs to integers
	userIDInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid user ID: %w", err)
	}

	commentIDInt, err := strconv.ParseInt(commentId, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid comment ID: %w", err)
	}

	// SQL statement to delete the comment record, ensuring it belongs to the user
	deleteStmt := `DELETE FROM comments WHERE id = ? AND user_id = ?`

	// Executing the SQL statement to delete
	result, err := db.c.Exec(deleteStmt, commentIDInt, userIDInt)
	if err != nil {
		return fmt.Errorf("error deleting comment: %w", err)
	}

	// Check if a comment was actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no comment found with the given ID for this user")
	}

	return nil
}
