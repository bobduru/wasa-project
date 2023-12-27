package database

func (db *appdbimpl) GetNameFromUserId(userId int64) (string, error) {
	var name string

	query := "SELECT name FROM users WHERE id = ?"
	err := db.c.QueryRow(query, userId).Scan(&name)
	if err != nil {
		return "", err
	}

	return name, nil
}
