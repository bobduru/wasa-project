package database

import (
	"fmt"
	"strconv"
)

func (db *appdbimpl) GetUserStatus(loggedInUser string, userToFind string) (isFollowing bool, isBanned bool, err error) {
	loggedInUserID, err := strconv.Atoi(loggedInUser)
	if err != nil {
		return false, false, fmt.Errorf("invalid logged in user ID: %w", err)
	}

	userToFindID, err := strconv.Atoi(userToFind)
	if err != nil {
		return false, false, fmt.Errorf("invalid user to find ID: %w", err)
	}

	// Check follow status
	followQuery := `SELECT EXISTS(SELECT 1 FROM follows WHERE follower_id = ? AND following_id = ?)`
	err = db.c.QueryRow(followQuery, loggedInUserID, userToFindID).Scan(&isFollowing)
	if err != nil {
		return false, false, fmt.Errorf("error checking follow status: %w", err)
	}

	// Check ban status
	banQuery := `SELECT EXISTS(SELECT 1 FROM bans WHERE banner_id = ? AND banned_id = ?)`
	err = db.c.QueryRow(banQuery, loggedInUserID, userToFindID).Scan(&isBanned)
	if err != nil {
		return false, false, fmt.Errorf("error checking ban status: %w", err)
	}

	return isFollowing, isBanned, nil
}
