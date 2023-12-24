/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
	"time"
)

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	SetName(name string) (string, error)
	FindName(nameToFind string) (string, error)
	UpdateName(id string, name string) (string, error)
	UploadImage(userID string) (*Image, error)
	DeleteImage(photoId string) (string, error)
	CheckUserId(userID string) bool
	FollowUser(loggedInUserId string, userIdToFollow string) error
	UnfollowUser(loggedInUserId string, userIdToUnfollow string) error
	BanUser(bannerIdString string, bannedIdString string) error
	UnbanUser(bannerIdString string, bannedIdString string) error
	GetUserProfile(userId int64) (*UserProfile, error)
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

type UserProfile struct {
	UserId    int64
	Name      string
	Photos    []Image
	Followers []User
	Following []User
}

type User struct {
	ID   int64
	Name string
}

type Image struct {
	ID         int64
	UserID     int64
	FileName   string
	UploadTime time.Time // or string if you prefer
	Likes      int
	Comments   int
}

type Comment struct {
	ID         int64
	UserID     int64
	ImageID    int64
	Text       string
	CreateTime time.Time // or string if you prefer
}

type Follow struct {
	FollowerID  int64
	FollowingID int64
}

type Ban struct {
	BannerID int64
	BannedID int64
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Create User table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT UNIQUE NOT NULL
		)
	`)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	// Create image table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS images (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			file_name TEXT NOT NULL,
			upload_time DATETIME DEFAULT CURRENT_TIMESTAMP,
			likes INTEGER DEFAULT 0,
			comments INTEGER DEFAULT 0,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		)
	`)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	// Create Comment table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS comments (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER,
			image_id INTEGER,
			text TEXT NOT NULL,
			create_time DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY (image_id) REFERENCES images(id) ON DELETE CASCADE
		)
	`)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	// Create Follow table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS follows (
			follower_id INTEGER,
			following_id INTEGER,
			PRIMARY KEY (follower_id, following_id),
			FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY (following_id) REFERENCES users(id) ON DELETE CASCADE
		)
	`)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	// Create Ban table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS bans (
			banner_id INTEGER,
			banned_id INTEGER,
			PRIMARY KEY (banner_id, banned_id),
			FOREIGN KEY (banner_id) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY (banned_id) REFERENCES users(id) ON DELETE CASCADE
		)
	`)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}

	fmt.Println("Tables created successfully")

	return &appdbimpl{
		c: db,
	}, nil
}

// // New returns a new instance of AppDatabase based on the SQLite connection `db`.
// // `db` is required - an error will be returned if `db` is `nil`.
// func New(db *sql.DB) (AppDatabase, error) {
// 	if db == nil {
// 		return nil, errors.New("database is required when building a AppDatabase")
// 	}

// 	// Check if table exists. If not, the database is empty, and we need to create the structure
// 	var tableName string
// 	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='example_table';`).Scan(&tableName)
// 	fmt.Printf("Table name: %s\n", tableName)
// 	fmt.Printf("Error: %s\n", err.Error())
// 	if errors.Is(err, sql.ErrNoRows) {
// 		sqlStmt := `CREATE TABLE example_table (id INTEGER NOT NULL PRIMARY KEY, name TEXT);`
// 		_, err = db.Exec(sqlStmt)
// 		if err != nil {
// 			return nil, fmt.Errorf("error creating database structure: %w", err)
// 		}
// 	}

// 	return &appdbimpl{
// 		c: db,
// 	}, nil
// }

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
