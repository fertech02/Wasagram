package database;

import (
	"database/sql"
)

type Comment struct {
	uid string		`json:"uid"`
	pid string		`json:"pid"`
	message string	`json:"message"`
}

// Comment a Photo
func (u *User) Comment(c *Comment) error {
    db, err := sql.Open("sqlite3", "./foo.db")
    if err != nil {
        panic(err)
    }
    query := "INSERT INTO Comments (uid, pid, message) VALUES ($1, $2, $3)"
    _, err = db.Exec(query, c.uid, c.pid, c.message)
    if err != nil {
        panic(err)
    }
    defer db.Close()

    return nil
}

// Uncomment a Photo
func (u *User) Uncomment(p *Photo) error {
	
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query := "DELETE FROM Comments WHERE uid = $1 AND pid = $2"
	_, err = db.Exec(query, u.uid, p.pid)
	if err != nil {
		panic(err)
	}
	
	return nil
}

// Get Comments For a Photo
func (p *Photo	) GetComments() ([]Comment, error) {
	
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