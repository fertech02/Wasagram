package database

// Like a Photo
func (db *appdbimpl) Like(pid string, uid string) error {

	_, err := db.c.Exec("INSERT INTO Likes (Uid, Pid) VALUES ($1, $2)", uid, pid)
	return err
}

// Unlike a Photo
func (db *appdbimpl) Unlike(pid string, uid string) error {

	_, err := db.c.Exec("DELETE FROM Likes WHERE Uid = ? AND Pid = ?", uid, pid)
	return err
}

// Check if a User has Liked a Photo
func (db *appdbimpl) CheckLike(pid string, uid string) (bool, error) {

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Likes WHERE Uid = ? AND Pid = ?", uid, pid).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Get Like Count for a Photo
func (db *appdbimpl) GetLikeCount(pid string) (int, error) {

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Likes WHERE pid = ?", pid).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

// Get Likes
func (db *appdbimpl) GetLikes(pid string) ([]Like, error) {

	rows, err := db.c.Query("SELECT Uid FROM Likes WHERE Pid = ?", pid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	likes := []Like{}
	for rows.Next() {
		var like Like
		err := rows.Scan(&like.Uid)
		if err != nil {
			return nil, err
		}
		likes = append(likes, like)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return likes, nil
}
