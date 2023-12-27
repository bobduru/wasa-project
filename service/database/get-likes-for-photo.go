package database

import (
	"fmt"
)

func (db *appdbimpl) GetLikesForPhoto(photoId int64) ([]Like, error) {
	likes := []Like{}

	// SQL statement to select likes for a specific photo
	query := `SELECT user_id, image_id FROM likes WHERE image_id = ?`
	rows, err := db.c.Query(query, photoId)
	if err != nil {
		return nil, fmt.Errorf("error querying likes: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var like Like
		if err := rows.Scan(&like.UserID, &like.ImageID); err != nil {
			return nil, fmt.Errorf("error scanning like: %w", err)
		}
		likes = append(likes, like)
	}

	return likes, nil
}
