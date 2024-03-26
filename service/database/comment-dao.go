package database;

import (
	"database/sql"
)


// Comment a Photo
func (u *User) Comment(c *Comment) error {
    db, err := sql.Open("sqlite3", "./foo.db")
    if err != nil {
        panic(err)
    }
    query := "INSERT INTO Comments (uid, pid, message) VALUES ($1, $2, $3)"
    _, err = db.QueryRow(query, c.Uid, c.Pid, c.Message)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    return nil
}

// Uncomment a Photo
func (u *User) Uncomment(p *Post) error {
	
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	query := "DELETE FROM Comments WHERE uid = $1 AND pid = $2"
	_, err = db.QueryRow(query, u.uid, p.pid)
	if err != nil {
		panic(err)
	}
	defer db.Close()
}

// Get Comments For a Photo
func (p *Post) GetComments() ([]Comment, error) {
	
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	query := "SELECT * FROM Comments WHERE pid = $1"
	rows, err := db.Query(query, p.pid)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	var comments []Comment
	for rows.Next() {
		var c Comment
		err := rows.Scan(&c.uid, &c.pid, &c.message)
		if err != nil {
			panic(err)
		}
		comments = append(comments, c)
	}
	return comments, nil
}