package database

// Follow an User
func (db *appdbimpl) Follow(followeeId string, followerId string) error {

	_, err := db.c.Exec("INSERT INTO Follow (FolloweeId, FollowerId) VALUES (?, ?)", followeeId, followerId)
	if err != nil {
		return err
	}

	return nil
}

// Unfollow an User
func (db *appdbimpl) Unfollow(followeeId string, followerId string) error {

	_, err := db.c.Exec("DELETE FROM Follow WHERE FolloweeId = ? AND FollowerId = ?", followeeId, followerId)
	if err != nil {
		return err
	}

	return nil
}

// Check if a User is Following another User
func (db *appdbimpl) CheckFollow(followeeId string, followerId string) (bool, error) {

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Follow WHERE FolloweeId = ? AND FollowerId = ?", followeeId, followerId).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Get Followers Count
func (db *appdbimpl) GetFollowersCount(uid string) (int, error) {

	var count int
	rows, err := db.c.Query("SELECT COUNT(*) FROM Follow WHERE FolloweeId = ?", uid)
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

	if err = rows.Err(); err != nil {
		return 0, err
	}

	return count, nil
}

// Get Followees Count
func (db *appdbimpl) GetFolloweesCount(uid string) (int, error) {

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Follow WHERE FollowerId = ?", uid).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// GetFollowers
func (db *appdbimpl) GetFollowers(uid string) ([]string, error) {

	var followers []string
	rows, err := db.c.Query("SELECT FollowerId FROM Follow WHERE FolloweeId = ?", uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var follower string
		err = rows.Scan(&follower)
		if err != nil {
			return nil, err
		}
		followers = append(followers, follower)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return followers, nil
}

// GetFollowees
func (db *appdbimpl) GetFollowees(uid string) ([]string, error) {

	var followees []string
	rows, err := db.c.Query("SELECT FolloweeId FROM Follow WHERE FollowerId = ?", uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var followee string
		err = rows.Scan(&followee)
		if err != nil {
			return nil, err
		}
		followees = append(followees, followee)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	
	return followees, nil
}
