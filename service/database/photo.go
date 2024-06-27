package database

// Post a Photo
func (db *appdbimpl) PostPhoto(p *Photo) (*Photo, error) {

	query := "INSERT INTO Photos (Pid, Uid, File, Date) VALUES (?, ?, ?, ?)"
	_, err := db.c.Exec(query, p.Pid, p.Uid, p.File, p.Date)
	if err != nil {
		return p, err
	}
	return p, nil
}

// Delete a Photo
func (db *appdbimpl) DeletePhoto(pid string) error {

	_, err := db.c.Exec("DELETE FROM Like WHERE Pid=?", pid)
	if err != nil {
		return err
	}

	_, er := db.c.Exec("DELETE FROM Comments WHERE Pid=?", pid)
	if err != nil {
		return er
	}

	query := "DELETE FROM Photos WHERE Pid = $1"
	_, err = db.c.Exec(query, pid)
	if err != nil {
		return err
	}

	return nil
}

// Get a Photo
func (db *appdbimpl) GetPhoto(pid string) (*Photo, error) {

	query := "SELECT * FROM Photos WHERE Pid = $1"
	row := db.c.QueryRow(query, pid)
	if row == nil {
		return nil, nil
	}

	var p Photo
	err := row.Scan(&p.Pid, &p.Uid, &p.File, &p.Date)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// Get User's Photos
func (db *appdbimpl) GetPhotos(uid string) ([]*Photo, error) {

	query := "SELECT * FROM Photos WHERE Uid = $1"
	rows, err := db.c.Query(query, uid)
	if err != nil {
		return nil, err
	}

	photos := []*Photo{}
	for rows.Next() {
		var p Photo
		err = rows.Scan(&p.Pid, &p.Uid, &p.File, &p.Date)
		if err != nil {
			return nil, err
		}
		photos = append(photos, &p)
	}

	return photos, nil
}

// Get Photo Count
func (db *appdbimpl) GetPhotoCount(uid string) (int, error) {

	query := "SELECT COUNT(*) FROM Photos WHERE Uid = $1"
	rows, err := db.c.Query(query, uid)
	if err != nil {
		return 0, err
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}
