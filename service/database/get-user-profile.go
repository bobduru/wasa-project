package database

import (
	"fmt"
)

func (db *appdbimpl) GetUserProfile(userId int64) (*UserProfile, error) {
	// Retrieve user information
	var userProfile UserProfile
	userQuery := `SELECT id, name FROM users WHERE id = ?`
	err := db.c.QueryRow(userQuery, userId).Scan(&userProfile.UserId, &userProfile.Name)
	if err != nil {
		return nil, fmt.Errorf("error fetching user data: %w", err)
	}

	// Retrieve photos of the user
	photosQuery := `SELECT id, user_id, file_name, upload_time FROM images WHERE user_id = ? ORDER BY upload_time DESC`
	rows, err := db.c.Query(photosQuery, userId)
	if err != nil {
		return nil, fmt.Errorf("error fetching photos: %w", err)
	}
	defer rows.Close()
	userProfile.Photos = []Image{}
	for rows.Next() {
		var photo Image
		if err := rows.Scan(&photo.ID, &photo.UserID, &photo.FileName, &photo.UploadTime); err != nil {
			return nil, fmt.Errorf("error scanning photo: %w", err)
		}

		photo.Likes, err = db.GetLikesForPhoto(photo.ID)
		if err != nil {
			return nil, fmt.Errorf("error fetching likes: %w", err)
		}

		photo.Comments, err = db.GetCommentsForPhoto(photo.ID)
		if err != nil {
			return nil, fmt.Errorf("error fetching comments: %w", err)
		}

		photo.UserName, err = db.GetNameFromUserId(photo.UserID)
		if err != nil {
			return nil, fmt.Errorf("error fetching user name: %w", err)
		}

		userProfile.Photos = append(userProfile.Photos, photo)
	}

	userProfile.Followers, err = db.GetFollowersForUser(userId)
	if err != nil {
		return nil, fmt.Errorf("error fetching followers: %w", err)
	}

	// Retrieve following of the user
	followingQuery := `SELECT u.id, u.name FROM users u JOIN follows f ON u.id = f.following_id WHERE f.follower_id = ?`
	followingRows, err := db.c.Query(followingQuery, userId)
	if err != nil {
		return nil, fmt.Errorf("error fetching following: %w", err)
	}
	defer followingRows.Close()
	userProfile.Following = []User{}
	for followingRows.Next() {
		var following User
		if err := followingRows.Scan(&following.ID, &following.Name); err != nil {
			return nil, fmt.Errorf("error scanning following: %w", err)
		}
		userProfile.Following = append(userProfile.Following, following)
	}

	return &userProfile, nil
}
