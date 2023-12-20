package database

import (
	"database/sql"
	"errors"
)

func (db *appdbimpl) FindName(nameToFind string) (string, error) {
	var id string
	err := db.c.QueryRow("SELECT id FROM users WHERE name=?", nameToFind).Scan(&id)

	if errors.Is(err, sql.ErrNoRows) {
		// Name not found, return an empty string
		return "", nil
	}

	return id, err
}
