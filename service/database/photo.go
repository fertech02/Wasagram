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
func (db *appdbimpl) postPhoto(uid string, pid string, url string) error {
	
	date := time.Now().Format("2006-01-02 15:04:05")
	query := "INSERT INTO Photos (uid, pid, url, date) VALUES (?, ?, ?, ?)"
	_, err = db.Exec(query, uid, pid, url, date)
	if err != nil {
		panic(err)
	}

	return nil
}

// Delete a Photo
func (db *appdbimpl) deletePhoto(uid string, pid string) error {

	query := "DELETE FROM Photos WHERE uid = $1 AND pid = $2"
	_, err = db.Exec(query, uid, pid)
	if err != nil {
		panic(err)
	}

	return nil
}

// Get all photos for a user	
func (db *appdbimpl) getPhotos(uid string) error {
	
	query := "SELECT * FROM Photos WHERE uid = $1"
	rows, err := db.Query(query, u.uid)
	if err != nil {
		panic(err)
	}
	
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

