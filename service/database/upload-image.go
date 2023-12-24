package database

import (
	"fmt"

	"github.com/gofrs/uuid"
)

func (db *appdbimpl) UploadImage(userId string) (*Image, error) {
	// SQL statement to insert a new photo record
	insertStmt := `
        INSERT INTO images (user_id, file_name) 
        VALUES (?, ?)
    `
	newUUID := uuid.Must(uuid.NewV4())
	fileName := newUUID.String() + ".jpg"

	// Executing the SQL statement to insert
	res, err := db.c.Exec(insertStmt, userId, fileName)
	if err != nil {
		return nil, fmt.Errorf("error inserting new photo: %w", err)
	}

	// Get the last inserted ID
	lastId, err := res.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("error getting last insert ID: %w", err)
	}

	selectStmt := `SELECT id, user_id, file_name, upload_time, likes, comments FROM images WHERE id = ?`

	// Image object to hold the data
	var image Image

	// Executing the SQL statement to select
	err = db.c.QueryRow(selectStmt, lastId).Scan(&image.ID, &image.UserID, &image.FileName, &image.UploadTime, &image.Likes, &image.Comments)
	if err != nil {
		return nil, fmt.Errorf("error fetching inserted image: %w", err)
	}

	return &image, nil
}
