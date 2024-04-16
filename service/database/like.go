package database;

import (
	"database/sql"
)

type Like struct {
	uid string	`json:"uid"`
	pid string	`json:"pid"`
}

// Like a Photo
func (db *appdbimpl) Like(pid string, uid string) error {

	query := "INSERT INTO Likes (uid, pid) VALUES ($1, $2)"
	_, err = db.Exec(query, uid, pid)
	if err != nil {
		panic(err)
	}

	return nil
}

// Unlike a Photo
func (db *appdbimpl) Unlike(pid string, uid string) error {

	query := "DELETE FROM Likes WHERE uid = $1 AND pid = $2"
	_, err = db.Exec(query, uid, pid)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	return nil
}

// Get Like by ID
func (db *appdbimpl) GetLike(pid string) ([]*Like, error) {

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
