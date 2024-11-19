package database

import "database/sql"

// Post a Photo
func (db *appdbimpl) PostPhoto(p Photo) (Photo, error) {

	_, err := db.c.Exec(`INSERT INTO Photos (Pid, Uid, Date) VALUES (?, ?, ?)`, p.Pid, p.Uid, p.Date)
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

	_, err = db.c.Exec("DELETE FROM Comments WHERE Pid=?", pid)
	if err != nil {
		return err
	}

	query := "DELETE FROM Photos WHERE Pid = $1"
	_, err = db.c.Exec(query, pid)
	if err != nil {
		return err
	}

	return nil
}

func (db *appdbimpl) GetPhoto(pid string) (dbPhoto Photo, present bool, err error) {

	query := "SELECT Pid, Uid, Date FROM Photos WHERE Pid = ?;"

	row := db.c.QueryRow(query, pid)
	err = row.Scan(&dbPhoto.Pid, &dbPhoto.Uid, &dbPhoto.Date)
	if err != nil && err != sql.ErrNoRows {
		return
	} else if err == sql.ErrNoRows {
		err = nil
		return
	} else {
		err = nil
		present = true
		return
	}
}

// Get a Photo
func (db *appdbimpl) GetPhotoAuthor(pid string) (string, error) {

	query := "SELECT Uid FROM Photos WHERE Pid = $1"
	rows, err := db.c.Query(query, pid)
	if err != nil {
		return "", err
	}

	var uid string
	for rows.Next() {
		err = rows.Scan(&uid)
		if err != nil {
			return "", err
		}
	}

	if err = rows.Err(); err != nil {
		return "", err
	}

	return uid, nil
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

	if err = rows.Err(); err != nil {
		return 0, err
	}

	return count, nil
}
