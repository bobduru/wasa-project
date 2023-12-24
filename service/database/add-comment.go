package database

import (
	"fmt"
	"strconv"
)

func (db *appdbimpl) AddComment(userId string, photoId string, commentText string) (*Comment, error) {
	// Convert string IDs to integers
	userIDInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %v", err)
	}

	photoIDInt, err := strconv.ParseInt(photoId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid photo ID: %v", err)
	}

	insertStmt := `
        INSERT INTO comments (user_id, image_id, text) 
        VALUES (?, ?, ?)
    `

	// Executing the SQL statement to insert
	res, err := db.c.Exec(insertStmt, userIDInt, photoIDInt, commentText)
	if err != nil {
		return nil, fmt.Errorf("error inserting new comment: %v", err)
	}

	// Get the last inserted ID
	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert ID: %w", err)
	}

	// Fetch the newly inserted comment
	var newComment Comment
	selectStmt := `SELECT id, user_id, image_id, text, create_time FROM comments WHERE id = ?`
	err = db.c.QueryRow(selectStmt, lastId).Scan(&newComment.ID, &newComment.UserID, &newComment.ImageID, &newComment.Text, &newComment.CreateTime)
	if err != nil {
		return nil, fmt.Errorf("error fetching inserted comment: %w", err)
	}

	return &newComment, nil
}