package database

import (
	"fmt"
	"strconv"
)

// SetName sets the name in the database and returns the identifier
func (db *appdbimpl) SetName(name string) (string, error) {

	result, err := db.c.Exec("INSERT INTO users (name) VALUES (?)", name)
	if err != nil {
		return "", err
	}

	// Retrieve the ID of the last inserted record
	id, err := result.LastInsertId()
	if err != nil {
		return "", err
	}

	// Convert the ID to a string
	idString := strconv.FormatInt(id, 10)
	fmt.Printf("ID: %s\n", idString)
	return idString, nil
}
