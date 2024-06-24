package database

// Follow an User
func (db *appdbimpl) Follow(followeeId string, followerId string) error {

	_, err := db.c.Exec("INSERT INTO Follow (followeeId, followerId) VALUES (?, ?)", followeeId, followerId)
	if err != nil {
		return err
	}

	return nil
}

// Unfollow an User
func (db *appdbimpl) Unfollow(followeeId string, followerId string) error {

	_, err := db.c.Exec("DELETE FROM Follow WHERE followeeId = ? AND followerId = ?", followeeId, followerId)
	if err != nil {
		return err
	}

	return nil
}

// Check if a User is Following another User
func (db *appdbimpl) CheckFollow(followeeId string, followerId string) (bool, error) {

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Follow WHERE followeeId = ? AND followerId = ?", followeeId, followerId).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Get Followers Count
func (db *appdbimpl) GetFollowersCount(uid string) (int, error) {

	var count int
	rows, err := db.c.Query("SELECT COUNT(*) FROM Follow WHERE followeeId = ?", uid)
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, err
		}
	}

	return count, nil
}

// Get Followees Count
func (db *appdbimpl) GetFolloweesCount(uid string) (int, error) {

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Follow WHERE followerId = ?", uid).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
