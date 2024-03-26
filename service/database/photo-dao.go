package database;

import (
	"database/sql"
	"github.com/google/uuid"
)

// Post a Photo
func (u *User) postPhoto(p *Photo) error {
	// database query to Photos table
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	query := "INSERT INTO Photos (uid, pid, url) VALUES (?, ?, ?)"
	_, err = db.QueryRow(query, u.uid, pid, p.url)
	if err != nil {
		panic(err)
	}
	defer db.Close()

}

// Delete a Photo
func (u *User) deletePhoto(uid string, pid string) error {
	// database query to Photos table
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	query := "DELETE FROM Photos WHERE uid = $1 AND pid = $2"
	_, err = db.QueryRow(query, uid, pid)
	if err != nil {
		panic(err)
	}
	defer db.Close()
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
}

