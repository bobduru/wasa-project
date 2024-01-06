package database

import (
	"fmt"
	"strconv"
)

func (db *appdbimpl) BanUser(bannerIdString string, bannedIdString string) error {
	// Convert string IDs to integers
	bannerID, err := strconv.Atoi(bannerIdString)
	if err != nil {
		return fmt.Errorf("invalid banner ID: %w", err)
	}

	bannedID, err := strconv.Atoi(bannedIdString)
	if err != nil {
		return fmt.Errorf("invalid banned ID: %w", err)
	}

	// Check if the ban relationship already exists
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM bans WHERE banner_id = ? AND banned_id = ?)`
	err = db.c.QueryRow(checkQuery, bannerID, bannedID).Scan(&exists)
	if err != nil {
		return fmt.Errorf("error checking existing ban relationship: %w", err)
	}
	if exists {
		return fmt.Errorf("ban relationship already exists")
	}

	// Insert ban relationship
	insertStmt := `
        INSERT INTO bans (banner_id, banned_id) 
        VALUES (?, ?)
    `
	_, err = db.c.Exec(insertStmt, bannerID, bannedID)
	if err != nil {
		return fmt.Errorf("error inserting ban relationship: %w", err)
	}

	return nil
}
