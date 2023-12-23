package database

import (
	"fmt"
	"strconv"
)

func (db *appdbimpl) UnfollowUser(loggedInUserId string, userIdToUnfollow string) error {
	// Convert string IDs to integers
	followerID, err := strconv.Atoi(loggedInUserId)
	if err != nil {
		return fmt.Errorf("invalid follower ID: %v", err)
	}

	followingID, err := strconv.Atoi(userIdToUnfollow)
	if err != nil {
		return fmt.Errorf("invalid following ID: %v", err)
	}

	// Delete follow relationship
	deleteStmt := `
        DELETE FROM follows 
        WHERE follower_id = ? AND following_id = ?
    `
	result, err := db.c.Exec(deleteStmt, followerID, followingID)
	if err != nil {
		return fmt.Errorf("error deleting follow relationship: %v", err)
	}

	// Check if a row was actually deleted
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("error checking rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no follow relationship exists to be deleted")
	}

	return nil
}