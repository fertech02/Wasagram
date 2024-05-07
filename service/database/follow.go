package database

// Follow an User
func (db *appdbimpl) Follow(followeeId int, followerId int) error {

	_, err := db.c.Exec("INSERT INTO Follow (followeeId, followerId) VALUES (?, ?)", followeeId, followerId)
	if err != nil {
		return err
	}

	return nil
}

// Unfollow an User
func (db *appdbimpl) Unfollow(followeeId int, followerId int) error {

	_, err := db.c.Exec("DELETE FROM Follow WHERE followeeId = ? AND followerId = ?", followeeId, followerId)
	if err != nil {
		return err
	}

	return nil
}

// Check if a User is Following another User
func (db *appdbimpl) CheckFollow(followeeId int, followerId int) (bool, error) {

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Follow WHERE followeeId = ? AND followerId = ?", followeeId, followerId).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Get Followers Count
func (db *appdbimpl) GetFollowersCount(uid int) (int, error) {

	var count int
	err := db.c.Query("SELECT COUNT(*) FROM Follow WHERE followeeId = ?", uid).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// Get Followees Count
func (db *appdbimpl) GetFolloweesCount(uid int) (int, error) {

	var count int
	err := db.c.Query("SELECT COUNT(*) FROM Follow WHERE followerId = ?", uid).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
