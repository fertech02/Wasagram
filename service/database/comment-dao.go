package database;

// Comment a Photo
func (u *User) Comment(p *Post, message string) error {
	
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	query := "INSERT INTO Comments (uid, pid, message) VALUES ($1, $2, $3)"
	_, err = db.QueryRow(query, u.uid, p.pid, message)
	if err != nil {
		panic(err)
	}
	defer db.Close()
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
