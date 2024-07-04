package database

// Comment a Photo
func (db *appdbimpl) Comment(c *Comment) error {

	query := "INSERT INTO Comments (Uid, Pid, Message) VALUES ($1, $2, $3)"
	_, err := db.c.Exec(query, c.Uid, c.Pid, c.Message)
	if err != nil {
		return err
	}

	return nil
}

// Uncomment a Photo
func (db *appdbimpl) Uncomment(pid string, uid string) error {

	query := "DELETE FROM Comments WHERE Uid = $1 AND Pid = $2"
	_, err := db.c.Exec(query, uid, pid)
	if err != nil {
		return err
	}

	return nil

}

// Get Comments For a Photo
func (db *appdbimpl) GetComments(pid string) ([]Comment, error) {

	query := "SELECT * FROM Comments WHERE Pid = $1"
	rows, err := db.c.Query(query, pid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []Comment
	for rows.Next() {
		var c Comment
		err = rows.Scan(&c.Uid, &c.Pid, &c.Message)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return comments, nil
}
