package database

func (db *appdbimpl) GetAllUsers() ([]User, error) {
	// Slice to hold the users
	users := []User{}

	// SQL query to select all users
	rows, err := db.c.Query("SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate over the rows
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Check for errors from iterating over rows
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
