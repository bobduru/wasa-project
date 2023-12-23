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
		return nil, fmt.Errorf("error fetching user data: %v", err)
	}

	// Retrieve photos of the user
	photosQuery := `SELECT id, file_name, upload_time, likes, comments FROM images WHERE user_id = ?`
	rows, err := db.c.Query(photosQuery, userId)
	if err != nil {
		return nil, fmt.Errorf("error fetching photos: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var photo Image
		if err := rows.Scan(&photo.ID, &photo.FileName, &photo.UploadTime, &photo.Likes, &photo.Comments); err != nil {
			return nil, fmt.Errorf("error scanning photo: %v", err)
		}
		userProfile.Photos = append(userProfile.Photos, photo)
	}

	// Retrieve followers of the user
	followersQuery := `SELECT u.id, u.name FROM users u JOIN follows f ON u.id = f.follower_id WHERE f.following_id = ?`
	followersRows, err := db.c.Query(followersQuery, userId)
	if err != nil {
		return nil, fmt.Errorf("error fetching followers: %v", err)
	}
	defer followersRows.Close()

	for followersRows.Next() {
		var follower User
		if err := followersRows.Scan(&follower.ID, &follower.Name); err != nil {
			return nil, fmt.Errorf("error scanning follower: %v", err)
		}
		userProfile.Followers = append(userProfile.Followers, follower)
	}

	// Retrieve following of the user
	followingQuery := `SELECT u.id, u.name FROM users u JOIN follows f ON u.id = f.following_id WHERE f.follower_id = ?`
	followingRows, err := db.c.Query(followingQuery, userId)
	if err != nil {
		return nil, fmt.Errorf("error fetching following: %v", err)
	}
	defer followingRows.Close()

	for followingRows.Next() {
		var following User
		if err := followingRows.Scan(&following.ID, &following.Name); err != nil {
			return nil, fmt.Errorf("error scanning following: %v", err)
		}
		userProfile.Following = append(userProfile.Following, following)
	}

	return &userProfile, nil
}
