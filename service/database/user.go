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

// Get Profile Photos
func (db *appdbimpl) GetProfilePhotos(uid string) ([]*Photo, error) {

	var photos []*Photo
	rows, err := db.c.Query("SELECT Pid, Uid, File, Date FROM Photo WHERE Uid=? ORDER BY Date DESC", uid)
	if err != nil {
		return photos, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Photo
		err = rows.Scan(&p.Pid, &p.Uid, &p.File, &p.Date)
		if err != nil {
			return photos, err
		}
		photos = append(photos, &p)
	}
	return photos, nil
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
func (db *appdbimpl) GetUserId(username string) (user string, err error) {

	var Id string
	err = db.c.QueryRow("SELECT Uid FROM Users WHERE Username = ?", username).Scan(&Id)
	if err == sql.ErrNoRows {
		return "", nil
	}
	return Id, err
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
