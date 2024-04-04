package database;

import (
	"database/sql"
	"time"
)

type Photo struct {
	pid  string	`json:"pid"`
	uid  string	`json:"uid"`
	url  string	`json:"url"`
	date string	`json:"date"`
}

// Post a Photo
func (u *User) postPhoto(uid string, pid string, url string) error {
	// database query to Photos table
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	date := time.Now().Format("2006-01-02 15:04:05")
	query := "INSERT INTO Photos (uid, pid, url, date) VALUES (?, ?, ?, ?)"
	_, err = db.Exec(query, uid, pid, url, date)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return nil
}

// Delete a Photo
func (u *User) deletePhoto(uid string, pid string) error {
	// database query to Photos table
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	query := "DELETE FROM Photos WHERE uid = $1 AND pid = $2"
	_, err = db.Exec(query, uid, pid)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return nil
}

// Get all photos for a user	
func (u *User) getPhotos(uid string) error {
	// database query to Photos table
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	query := "SELECT * FROM Photos WHERE uid = $1"
	rows, err := db.Query(query, u.uid)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	
	// iterate over rows
	for rows.Next() {
		var pid string
		var url string
		err = rows.Scan(&pid, &url)
		if err != nil {
			panic(err)
		}
	}

	if err = rows.Err(); err != nil {
		return err
	}

	return nil
}

