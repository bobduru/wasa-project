package database

import (
	"fmt"
)

func (db *appdbimpl) GetCommentsForPhoto(photoId int64) ([]Comment, error) {
	comments := []Comment{}

	// SQL statement to select comments for a specific photo
	query := `SELECT user_id, image_id, text, create_time FROM comments WHERE image_id = ?`
	rows, err := db.c.Query(query, photoId)
	if err != nil {
		return nil, fmt.Errorf("error querying comments: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var comment Comment
		if err := rows.Scan(&comment.UserID, &comment.ImageID, &comment.Text, &comment.CreateTime); err != nil {
			return nil, fmt.Errorf("error scanning comment: %w", err)
		}

		comment.UserName, err = db.GetNameFromUserId(comment.UserID)
		if err != nil {
			return nil, fmt.Errorf("error fetching user name: %w", err)
		}

		comments = append(comments, comment)
	}

	return comments, nil
}
