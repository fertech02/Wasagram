package database

// Comment a Photo
func (db *appdbimpl) Comment(c Comment) error {

	query := "INSERT INTO Comments (Cid, Uid, Pid, Message) VALUES ($1, $2, $3, $4)"
	_, err := db.c.Exec(query, c.Cid, c.Uid, c.Pid, c.Message)
	if err != nil {
		return err
	}

	return nil
}

// Uncomment a Photo
func (db *appdbimpl) Uncomment(cid string) error {

	query := "DELETE FROM Comments WHERE Cid = $1"
	_, err := db.c.Exec(query, cid)
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
		err = rows.Scan(&c.Cid, &c.Uid, &c.Pid, &c.Message)
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
