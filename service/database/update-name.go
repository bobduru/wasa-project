package database

import (
	"errors"
)

func (db *appdbimpl) UpdateName(id string, name string) (string, error) {
	result, err := db.c.Exec("UPDATE example_table SET name=? WHERE id=?", name, id)
	if err != nil {
		return "", err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return "", err
	}

	if rowsAffected == 0 {
		// No rows were updated, return an error or handle the case accordingly
		return "", errors.New("no rows were updated")
	}

	return name, nil
}
