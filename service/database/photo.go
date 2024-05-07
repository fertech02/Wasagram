package database

// Post a Photo
func (db *appdbimpl) PostPhoto(p Photo) (*Photo, error) {

	query := "INSERT INTO Photos (pid, uid, file, date) VALUES (?, ?, ?, ?)"
	_, err = db.c.Exec(query, p.pid, p.uid, p.file, p.date)
	if err != nil {
		return p, err
	}
	return p, nil
}

// Delete a Photo
func (db *appdbimpl) DeletePhoto(pid string) error {

	_, err := db.c.Exec("DELETE FROM Like WHERE pid=?", pid)
	if err != nil {
		return err
	}

	_, er := db.c.Exec("DELETE FROM Comments WHERE pid=?", pid)
	if err != nil {
		return er
	}

	query := "DELETE FROM Photos WHERE pid = $1"
	_, err = db.c.Exec(query, pid)
	if err != nil {
		return err
	}

	return nil
}

// Get User's Photos
func (db *appdbimpl) GetPhotos(uid string) ([]*Photo, error) {

	query := "SELECT * FROM Photos WHERE uid = $1"
	rows, err := db.c.Query(query, uid)
	if err != nil {
		panic(err)
	}

	photos := []*Photo{}
	for rows.Next() {
		var p Photo
		err = rows.Scan(&p.pid, &p.uid, &p.file, &p.date)
		if err != nil {
			panic(err)
		}
		photos = append(photos, &p)
	}

	return photos, nil
}

// Get Photo Count
func (db *appdbimpl) GetPhotoCount(uid string) (int, error) {

	query := "SELECT COUNT(*) FROM Photos WHERE uid = $1"
	rows, err := db.c.Query(query, uid)
	if err != nil {
		panic(err)
	}

	var count int
	for rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			panic(err)
		}
	}

	return count, nil
}
