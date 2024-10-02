package database

import "database/sql"

// Create a new User
func (db *appdbimpl) CreateUser(username string) (dbUser User, err error) {

	query := `INSERT INTO Users (Uid, Username) VALUES (?,?);`
	_, err = db.c.Exec(query, username, username)
	if err != nil {
		return
	}
	dbUser.Username = username
	dbUser.Uid = username
	return dbUser, nil
}

// Update Username
func (db *appdbimpl) UpdateUsername(uid string, username string) error {

	_, err := db.c.Exec("UPDATE User SET Username=? WHERE Uid=?", username, uid)
	if err != nil {
		return err
	}
	return nil
}

// Get User Profile
func (db *appdbimpl) GetUserProfile(uid string) (*Profile, error) {

	var p Profile
	err := db.c.QueryRow("SELECT Uid, Username, (SELECT COUNT(*) FROM Follow WHERE FolloweeId=?) AS Followers, (SELECT COUNT(*) FROM Follow WHERE FollowerId=?) AS Followees, (SELECT COUNT(*) FROM Photo WHERE Uid=?) FROM User AS Photos WHERE Uid=?", uid, uid, uid, uid).Scan(&p.Uid, &p.Username, &p.Followers, &p.Followees, &p.PhotosCount)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// Get My Stream
func (db *appdbimpl) GetMyStream(uid string) ([]*Photo, error) {

	var strm []*Photo
	rows, err := db.c.Query("SELECT Pid, Uid, File, Date FROM Photo WHERE Uid IN (SELECT FollowerId From Follow WHERE FolloweeId=? AND FollowerId NOT IN (SELECT Uid FROM Ban where BannedID=?)) ORDER BY Date DESC LIMIT 20", uid, uid)
	if err != nil {
		return strm, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Photo
		err = rows.Scan(&p.Pid, &p.Uid, &p.File, &p.Date)
		if err != nil {
			return strm, err
		}
		strm = append(strm, &p)
	}

	if err = rows.Err(); err != nil {
		return strm, err
	}

	return strm, nil
}

// Search User

func (db *appdbimpl) SearchUser(usernameToSearch string) (usersList []User, err error) {

	query := "SELECT * FROM Users WHERE Username LIKE ?"

	rows, err := db.c.Query(query, usernameToSearch+"%")
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err = rows.Scan(&user.Uid, &user.Username)
		if err != nil {
			return
		}
		usersList = append(usersList, user)
	}

	err = rows.Err()
	return
}

// Get User Id
func (db *appdbimpl) GetUserId(username string) (user User,present bool, err error) {

	query := `SELECT * FROM Users WHERE Uid = ?;`
	err = db.c.QueryRow(query, username).Scan(&user.Uid, &user.Username)
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

// Get Username
func (db *appdbimpl) GetUsername(uid string) (string, error) {

	var username string
	err := db.c.QueryRow("SELECT Username FROM User WHERE Uid=?", uid).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}
