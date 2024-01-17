package database

import (
	"fmt"
)

func (db *appdbimpl) GetFollowersForUser(userId int64) ([]User, error) {
	followers := []User{}
	// Retrieve followers of the user
	followersQuery := `SELECT u.id, u.name FROM users u JOIN follows f ON u.id = f.follower_id WHERE f.following_id = ?`
	followersRows, err := db.c.Query(followersQuery, userId)
	if err != nil {
		return nil, fmt.Errorf("error fetching followers: %w", err)
	}
	defer followersRows.Close()

	for followersRows.Next() {
		var follower User
		if err := followersRows.Scan(&follower.ID, &follower.Name); err != nil {
			return nil, fmt.Errorf("error scanning follower: %w", err)
		}
		followers = append(followers, follower)
	}

	// Check for errors from iterating over rows
	if err = followersRows.Err(); err != nil {
		return nil, err
	}
	return followers, nil
}
