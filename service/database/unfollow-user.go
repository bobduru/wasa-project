package database

import (
	"fmt"
	"strconv"
)

func (db *appdbimpl) UnfollowUser(loggedInUserId string, userIdToUnfollow string) ([]User, error) {
	// Convert string IDs to integers
	followerID, err := strconv.Atoi(loggedInUserId)
	if err != nil {
		return nil, fmt.Errorf("invalid follower ID: %w", err)
	}

	followingID, err := strconv.Atoi(userIdToUnfollow)
	if err != nil {
		return nil, fmt.Errorf("invalid following ID: %w", err)
	}

	// Delete follow relationship
	deleteStmt := `
        DELETE FROM follows 
        WHERE follower_id = ? AND following_id = ?
    `
	result, err := db.c.Exec(deleteStmt, followerID, followingID)
	if err != nil {
		return nil, fmt.Errorf("error deleting follow relationship: %w", err)
	}

	// Check if a row was actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("error checking rows affected: %w", err)
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no follow relationship exists to be deleted")
	}

	followers, err := db.GetFollowersForUser(int64(followingID))
	if err != nil {
		return nil, fmt.Errorf("error fetching followers: %w", err)
	}
	return followers, nil
}
