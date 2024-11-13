package database

import (
	"time"

	"github.com/gofrs/uuid"
)

// Post a Photo
func (db *appdbimpl) PostPhoto(uid string) (string, error) {

	currentTime := time.Now().UTC()
	uuid, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	pid := uuid.String()
	_, err = db.c.Exec(
		"INSERT INTO Photos(Pid, Uid, Date) VALUES (?, ?, ?)",
		pid, uid, currentTime.Format("2006-01-02 15:04:05"))
	if err != nil {
		return "", err
	}
	return pid, nil
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

// Get a Photo
func (db *appdbimpl) GetPhoto(pid string) (*Photo, error) {

	query := "SELECT * FROM Photos WHERE Pid = $1"
	row := db.c.QueryRow(query, pid)
	if row == nil {
		return nil, nil
	}

	var p Photo
	err := row.Scan(&p.Pid, &p.Uid, &p.Date)
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
		err = rows.Scan(&p.Pid, &p.Uid, &p.Date)
		if err != nil {
			return nil, err
		}
		photos = append(photos, &p)
	}

	if err = rows.Err(); err != nil {
		return nil, err
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

	if err = rows.Err(); err != nil {
		return 0, err
	}

	return count, nil
}
