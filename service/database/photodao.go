package database;

// Post a Photo
func (u *User) PostPhoto(p *Photo) error {
	// database query to Photos table
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	query := "INSERT INTO Photos (uid, pid) VALUES ($1, $2)"
	_, err = db.QueryRow(query, u.uid, p.pid)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Posted")
}

// Delete a Photo
func (u *User) DeletePhoto(p *Photo) error {
	// database query to Photos table
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		panic(err)
	}
	query := "DELETE FROM Photos WHERE uid = $1 AND pid = $2"
	_, err = db.QueryRow(query, u.uid, p.pid)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	fmt.Println("Deleted")
}
