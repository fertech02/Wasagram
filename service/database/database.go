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
	Uid      string `json:"uid"`
	Username string `json:"username"`
}

type Photo struct {
	Pid  string `json:"pid"`
	Uid  string `json:"uid"`
	Date string `json:"date"`
}

type Follow struct {
	FolloweeId string `json:"followeeId"`
	FollowerId string `json:"followerId"`
}

type Comment struct {
	Uid     string `json:"uid"`
	Pid     string `json:"pid"`
	Message string `json:"message"`
}

type Like struct {
	Uid string `json:"uid"`
	Pid string `json:"pid"`
}

type Ban struct {
	BannerId string `json:"bannerId"`
	BannedId string `json:"bannedId"`
}

type Profile struct {
	PhotoList     []Photo `json:"photoList"`
	Username      string  `json:"username"`
	FollowCount   int     `json:"followCount"`
	FollowedCount int     `json:"followedCount"`
	PhotoCount    int     `json:"photoCount"`
	IsFollowed    bool    `json:"isFollowed"`
	IsBanned      bool    `json:"isBanned"`
}

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	Ping() error

	// User
	CreateUser(username string) (User, error)
	GetUserId(username string) (string, error)
	GetUsername(uid string) (string, error)
	UpdateUsername(userid string, username string) error
	GetMyStream(uid string) ([]*Photo, error)
	SearchUser(username string) ([]User, error)

	// Profile
	GetProfilePhotos(uid string) ([]Photo, error)

	// Photo
	PostPhoto(uid string) (string, error)
	DeletePhoto(pid string) error
	GetPhotos(uid string) ([]*Photo, error)
	GetPhotoAuthor(pid string) (string, error)
	GetPhotoCount(uid string) (int, error)

	// Like
	Like(pid string, uid string) error
	Unlike(pid string, uid string) error
	CheckLike(pid string, uid string) (bool, error)
	GetLikeCount(pid string) (int, error)
	GetLikes(pid string) ([]Like, error)

	// Ban
	Ban(bannerId string, bannedId string) error
	Unban(bannerId string, bannedId string) error
	CheckBan(bannerId string, bannedId string) (bool, error)
	GetBannedUsers(bannerId string) ([]string, error)

	// Follow
	Follow(followeeId string, followerId string) error
	Unfollow(followeeId string, followerId string) error
	CheckFollow(followeeId string, followerId string) (bool, error)
	GetFollowersCount(uid string) (int, error)
	GetFolloweesCount(uid string) (int, error)
	GetFollowers(uid string) ([]string, error)
	GetFollowees(uid string) ([]string, error)

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
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name = ?;`, tableName).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE Photos (

				Pid TEXT PRIMARY KEY,
				Uid TEXT NOT NULL,
				Date TEXT NOT NULL,

				FOREIGN KEY (Uid) REFERENCES Users(Uid) ON DELETE CASCADE
					
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

				Uid TEXT PRIMARY KEY,
				Username TEXT NOT NULL

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

				Pid TEXT NOT NULL,
				Uid TEXT NOT NULL,
				Message TEXT NOT NULL,

				PRIMARY KEY (Pid, Uid),
				FOREIGN KEY (Pid) REFERENCES Photos(Pid) ON DELETE CASCADE,
				FOREIGN KEY (Uid) REFERENCES Users(Uid) ON DELETE CASCADE

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

				Pid TEXT NOT NULL,
				Uid TEXT NOT NULL,

				PRIMARY KEY (Pid, Uid),
				FOREIGN KEY (Pid) REFERENCES Photos(Pid) ON DELETE CASCADE,
				FOREIGN KEY (Uid) REFERENCES Users(Uid) ON DELETE CASCADE

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

				FolloweeId TEXT NOT NULL,
				FollowerId TEXT NOT NULL,

				PRIMARY KEY (FolloweeId, FollowerId),
				FOREIGN KEY (FolloweeId) REFERENCES Users(Uid) ON DELETE CASCADE,
				FOREIGN KEY (FollowerId) REFERENCES Users(Uid) ON DELETE CASCADE

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

				BannerId TEXT NOT NULL,
				BannedId TEXT NOT NULL,

				PRIMARY KEY (BannerId, BannedId),
				FOREIGN KEY (BannerId) REFERENCES Users(Uid) ON DELETE CASCADE,
				FOREIGN KEY (BannedId) REFERENCES Users(Uid) ON DELETE CASCADE
			
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
