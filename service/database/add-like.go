package database

import (
	"fmt"
	"strconv"
)

func (db *appdbimpl) AddLike(userId string, photoId string) error {
	// Convert string IDs to integers
	userIDInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid user ID: %v", err)
	}

	photoIDInt, err := strconv.ParseInt(photoId, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid photo ID: %v", err)
	}

	// SQL statement to insert a new like record
	insertStmt := `
        INSERT INTO likes (user_id, image_id) 
        VALUES (?, ?)
    `

	// Executing the SQL statement to insert
	_, err = db.c.Exec(insertStmt, userIDInt, photoIDInt)
	if err != nil {
		return fmt.Errorf("error inserting new like: %v", err)
	}

	return nil
}