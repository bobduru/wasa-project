package database

import (
	"fmt"
	"strconv"
)

func (db *appdbimpl) GetStream(userId string) ([]Image, error) {
	// Convert string IDs to integers
	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid user ID: %w", err)
	}

	var images []Image

	// SQL query to fetch images from followed users that have not blocked the user, in reverse chronological order
	streamQuery := `
		SELECT i.id, i.user_id, i.file_name, i.upload_time 
		FROM images i
		INNER JOIN follows f ON i.user_id = f.following_id
		LEFT JOIN bans b ON b.banner_id = i.user_id AND b.banned_id = ?
		WHERE f.follower_id = ? AND b.banner_id IS NULL
		ORDER BY i.upload_time DESC
	`
	rows, err := db.c.Query(streamQuery, userIdInt, userIdInt)
	if err != nil {
		return nil, fmt.Errorf("error fetching image stream: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var img Image
		if err := rows.Scan(&img.ID, &img.UserID, &img.FileName, &img.UploadTime); err != nil {
			return nil, fmt.Errorf("error scanning image: %w", err)
		}

		// Assuming GetLikesForPhoto and GetCommentsForPhoto functions exist and work correctly
		img.Likes, err = db.GetLikesForPhoto(img.ID)
		if err != nil {
			return nil, fmt.Errorf("error fetching likes for photo: %w", err)
		}

		img.Comments, err = db.GetCommentsForPhoto(img.ID)
		if err != nil {
			return nil, fmt.Errorf("error fetching comments for photo: %w", err)
		}

		// Assuming GetNameFromUserId function exists and works correctly
		img.UserName, err = db.GetNameFromUserId(img.UserID)
		if err != nil {
			return nil, fmt.Errorf("error fetching user name: %w", err)
		}

		images = append(images, img)
	}

	return images, nil
}
