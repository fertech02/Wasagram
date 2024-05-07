package database

// Like a Photo
func (db *appdbimpl) Like(pid string, uid string) error {

	_, err := db.c.Exec("INSERT INTO Likes (uid, pid) VALUES ($1, $2)", uid, pid)
	return err
}

// Unlike a Photo
func (db *appdbimpl) Unlike(pid string, uid string) error {

	_, err := db.c.Exec("DELETE FROM Likes WHERE uid = $1 AND pid = $2", uid, pid)
	return err
}

// Check if a User has Liked a Photo
func (db *appdbimpl) CheckLike(pid string, uid string) (bool, error) {

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Likes WHERE uid = $1 AND pid = $2", uid, pid).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Get Like Count for a Photo
func (db *appdbimpl) GetLikeCount(pid string) (int, error) {

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Likes WHERE pid = $1", pid).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}
