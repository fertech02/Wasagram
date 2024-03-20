package database;

// Follow an User
func (u *User) FollowUser(followeeId int, followerId int) {

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.QueryRow("INSERT INTO Follow (followeeId, followerId) VALUES (?, ?)", followeeId, followerId)
	if err != nil {
		return nil, err
	}
	defer db.Close()
}

// Unfollow an User
func (u *User) UnfollowUser(followeeId int, followerId int) {

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.QueryRow("DELETE FROM Follow WHERE followeeId = ? AND followerId = ?", followeeId, followerId)
	if err != nil {
		return nil, err
	}
	defer db.Close()
}

// Get Followers List
func (u *User) GetFollowers(uid int){

	db,err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.QueryRow("SELECT followerId FROM Follow WHERE followeeId = ?", uid)
	if err != nil {
		return nil, err
	}
	defer db.Close()
}

// Get Followees List
func (u *User) GetFollowees(uid int){

	db,err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.QueryRow("SELECT followeeId FROM Follow WHERE followerId = ?", uid)
	if err != nil {
		return nil, err
	}
	defer db.Close()
}