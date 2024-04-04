package database;

import (
	"database/sql"
)

type Like struct {
	uid string	`json:"uid"`
	pid string	`json:"pid"`
}

// Like a Photo
func (u *User) Like(pid string, uid string) error {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "INSERT INTO Likes (uid, pid) VALUES ($1, $2)"
	_, err = db.Exec(query, uid, pid)
	if err != nil {
		panic(err)
	}

	return nil
}

// Unlike a Photo
func (u *User) Unlike(pid string, uid string) error {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	query := "DELETE FROM Likes WHERE uid = $1 AND pid = $2"
	_, err = db.Exec(query, uid, pid)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return nil
}

// Get Like by ID
func GetLike(pid string) ([]*Like, error) {
	// database query to Likes table
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	query := "SELECT uid, pid FROM Likes WHERE pid = $1"
	rows, err := db.Query(query, pid)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer db.Close()
	
	likes := make([]*Like, 0)
	for rows.Next() {
		var l Like
		err := rows.Scan(&l.uid, &l.pid)
		if err != nil {
			panic(err)
			return nil, err
		}
		likes = append(likes, &l)
	}

	if err = rows.Err(); err != nil {
		panic(err)
		return nil, err
	}

	return likes, nil
}
