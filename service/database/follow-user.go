package database

import (
	"fmt"
	"strconv"
)

func (db *appdbimpl) FollowUser(loggedInUserId string, userIdToFollow string) ([]User, error) {
	// Convert string IDs to integers
	followerID, err := strconv.Atoi(loggedInUserId)
	if err != nil {
		return nil, fmt.Errorf("invalid follower ID: %w", err)
	}

	followingID, err := strconv.Atoi(userIdToFollow)
	if err != nil {
		return nil, fmt.Errorf("invalid following ID: %w", err)
	}

	// Check if the follow relationship already exists
	var exists bool
	checkQuery := `SELECT EXISTS(SELECT 1 FROM follows WHERE follower_id = ? AND following_id = ?)`
	err = db.c.QueryRow(checkQuery, followerID, followingID).Scan(&exists)
	if err != nil {
		return nil, fmt.Errorf("error checking existing follow relationship: %w", err)
	}
	if exists {
		return nil, fmt.Errorf("follow relationship already exists")
	}

	// Insert follow relationship
	insertStmt := `
        INSERT INTO follows (follower_id, following_id) 
        VALUES (?, ?)
    `
	_, err = db.c.Exec(insertStmt, followerID, followingID)
	if err != nil {
		return nil, fmt.Errorf("error inserting follow relationship: %w", err)
	}

	followers, err := db.GetFollowersForUser(int64(followingID))
	if err != nil {
		return nil, fmt.Errorf("error fetching followers: %w", err)
	}
	return followers, nil
}
