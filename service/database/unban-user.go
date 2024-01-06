package database

import (
	"fmt"
	"strconv"
)

func (db *appdbimpl) UnbanUser(bannerIdString string, bannedIdString string) error {
	// Convert string IDs to integers
	bannerID, err := strconv.Atoi(bannerIdString)
	if err != nil {
		return fmt.Errorf("invalid banner ID: %w", err)
	}

	bannedID, err := strconv.Atoi(bannedIdString)
	if err != nil {
		return fmt.Errorf("invalid banned ID: %w", err)
	}

	// Delete ban relationship
	deleteStmt := `
        DELETE FROM bans 
        WHERE banner_id = ? AND banned_id = ?
    `
	result, err := db.c.Exec(deleteStmt, bannerID, bannedID)
	if err != nil {
		return fmt.Errorf("error deleting ban relationship: %w", err)
	}

	// Check if a row was actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no ban relationship exists to be deleted")
	}

	return nil
}
