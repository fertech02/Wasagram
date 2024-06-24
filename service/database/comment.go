package database

// Comment a Photo
func (db *appdbimpl) Comment(c *Comment) error {

	query := "INSERT INTO Comments (uid, pid, message) VALUES ($1, $2, $3)"
	_, err := db.c.Exec(query, c.uid, c.pid, c.message)
	if err != nil {
		return err
	}

	return nil
}

// Uncomment a Photo
func (db *appdbimpl) Uncomment(pid string, uid string) error {

	query := "DELETE FROM Comments WHERE uid = $1 AND pid = $2"
	_, err := db.c.Exec(query, uid, pid)
	if err != nil {
		return err
	}

	return nil

}

// Get Comments For a Photo
func (db *appdbimpl) GetComments(pid string) ([]Comment, error) {

	query := "SELECT * FROM Comments WHERE pid = $1"
	rows, err := db.c.Query(query, pid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	var comments []Comment
	for rows.Next() {
		var c Comment
		err = rows.Scan(&c.uid, &c.pid, &c.message)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	return comments, nil
}
