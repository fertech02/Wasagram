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
)

type User struct {
	uid      string `json:"uid"`
	username string `json:"username"`
}

/*
	type Profile struct {
		uid         string `json:"uid"`
		username    string `json:"username"`
		followers   int    `json:"followersCount"`
		followees   int    `json:"followeesCount"`
		photosCount int    `json:"photosCount"`
	}

*
*/
type Photo struct {
	pid  string `json:"pid"`
	uid  string `json:"uid"`
	file []byte `json:"file"`
	date string `json:"date"`
}

type Follow struct {
	followeeId string
	followerId string
}

type Comment struct {
	uid     string `json:"uid"`
	pid     string `json:"pid"`
	message string `json:"message"`
}

type Like struct {
	uid string `json:"uid"`
	pid string `json:"pid"`
}

type Ban struct {
	bannerId string
	bannedId string
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	Ping() error

	// User
	CreateUser(username string) (*User, error)
	GetUserId(username string) (string, error)
	GetUsername(uid string) (string, error)
	UpdateUsername(userid string, username string) (*User, error)
	GetMyStream(uid string) ([]*Photo, error)
	// GetUserProfile(uid string) (*Profile, error)

	// Photo
	PostPhoto(p *Photo) (*Photo, error)
	DeletePhoto(pid string) error
	GetPhotos(uid string) ([]*Photo, error)
	GetPhotoCount(uid string) (int, error)

	// Like
	Like(pid string, uid string) error
	Unlike(pid string, uid string) error
	CheckLike(pid string, uid string) (bool, error)
	GetLikeCount(pid string) (int, error)

	// Ban
	Ban(bannerId string, bannedId string) error
	Unban(bannerId string, bannedId string) error
	CheckBan(bannerId string, bannedId string) (bool, error)

	// Follow
	Follow(followeeId string, followerId string) error
	Unfollow(followeeId string, followerId string) error
	CheckFollow(followeeId string, followerId string) (bool, error)
	GetFollowersCount(uid string) (int, error)
	GetFolloweesCount(uid string) (int, error)

	// Comment
	Comment(c *Comment) error
	Uncomment(pid string, uid string) error
	GetComments(photoId string) ([]Comment, error)
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {

	// Check if db is nil
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}

	// Enable foreign keys
	_, err := db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		return nil, fmt.Errorf("enabling foreign keys: %w", err)
	}

	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string

	tableName = "Photos"
	err := db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name = ?;`, tableName).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Photos (

				pid TEXT PRIMARY KEY,
				uid TEXT NOT NULL,
				file BLOB NOT NULL,
				date TEXT NOT NULL

				FOREIGN KEY (uid) REFERENCES Users(uid) ON DELETE CASCADE
					
		); `
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	tableName = "Users"
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name = ?;`, tableName).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Users (

				uid TEXT PRIMARY KEY,
				username TEXT NOT NULL

		); `
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	tableName = "Comments"
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name = ?;`, tableName).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Comments (

				pid TEXT NOT NULL,
				uid TEXT NOT NULL,
				message TEXT NOT NULL

				PRIMARY KEY (pid, uid)
				FOREIGN KEY (pid) REFERENCES Photos(pid) ON DELETE CASCADE
				FOREIGN KEY (uid) REFERENCES Users(uid) ON DELETE CASCADE

		); `
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	tableName = "Likes"
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name = ?;`, tableName).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Likes (

				pid TEXT NOT NULL,
				uid TEXT NOT NULL

				PRIMARY KEY (pid, uid)
				FOREIGN KEY (pid) REFERENCES Photos(pid) ON DELETE CASCADE
				FOREIGN KEY (uid) REFERENCES Users(uid) ON DELETE CASCADE

		); `
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	tableName = "Follows"
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name = ?;`, tableName).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Follows (

				followeeId TEXT NOT NULL,
				followerId TEXT NOT NULL

				PRIMARY KEY (followeeId, followerId)
				FOREIGN KEY (followeeId) REFERENCES Users(uid) ON DELETE CASCADE
				FOREIGN KEY (followerId) REFERENCES Users(uid) ON DELETE CASCADE

		); `
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	tableName = "Bans"
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name = ?;`, tableName).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Bans (

				bannerId TEXT NOT NULL,
				bannedId TEXT NOT NULL

				PRIMARY KEY (bannerId, bannedId)
				FOREIGN KEY (bannerId) REFERENCES Users(uid) ON DELETE CASCADE
				FOREIGN KEY (bannedId) REFERENCES Users(uid) ON DELETE CASCADE
			
		); `
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure: %w", err)
		}
	}

	return &appdbimpl{
		c: db,
	}, nil
}

// Ping checks if the database is reachable
func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
