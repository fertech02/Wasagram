package database;

import (
	"database/sql"
)


// Like a Photo
func (u *User) Like(pid string, uid string) error {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	query := "INSERT INTO Likes (uid, pid) VALUES ($1, $2)"
	_, err = db.QueryRow(query, uid, pid)
	if err != nil {
		panic(err)
	}
}

// Unlike a Photo
func (u *User) Unlike(pid string, uid string) error {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	query := "DELETE FROM Likes WHERE uid = $1 AND pid = $2"
	_, err = db.QueryRow(query, uid, pid)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

// Get Like by ID
func GetLike(pid string) (*Like, error) {
	// database query to Likes table
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	query := "SELECT uid, pid FROM Likes WHERE pid = $1"
	rows, err := db.Query(query, pid)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}
