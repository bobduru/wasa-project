package database

func (db *appdbimpl) CheckUserId(userID string) bool {
	var exists bool

	// Prepare the SQL query
	query := "SELECT EXISTS(SELECT 1 FROM users WHERE id = ?)"

	// Execute the query
	err := db.c.QueryRow(query, userID).Scan(&exists)
	if err != nil {
		// fmt.Errorf("error checking if user exists: %w", err)
		return false
	}

	return exists
}
