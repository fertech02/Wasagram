package database

import (
	"github.com/google/uuid"
)

// Create a new User
func (db *appdbimpl) CreateUser(username string) (*User, error) {

	_, err := db.c.Exec("INSERT INTO User (uid, username) VALUES (?, ?)", uuid.New().String(), username)
	if err != nil {
		return nil, err
	}
	return &User{uid: uuid.New().String(), username: username}, nil
}

// Update Username
func (db *appdbimpl) UpdateUsername(userid string, username string) error {

	_, err := db.c.Exec("UPDATE User SET username=? WHERE uid=?", username, userid)
	if err != nil {
		return err
	}
	return nil
}

// Get User Profile
/*
func (db *appdbimpl) GetUserProfile(username string) (*Profile, error) {

	var p Profile
	p.uid = GetUserId(username)
	p.username = username
	p.followers = GetFollowersCount(p.uid)
	p.followees = GetFolloweesCount(p.uid)
	p.photosCount = GetPhotoCount(p.uid)
	return &p, nil
}
**/

// Get My Stream
func (db *appdbimpl) GetMyStream(uid string) ([]*Photo, error) {

	var strm []*Photo
	rows, err := db.c.Query("SELECT pid, uid, file, date FROM Photo WHERE uid IN (SELECT followerId From Follow WHERE followeeId=? AND followerId NOT IN (SELECT uid FROM Ban where bannedID=?)) ORDER BY date DESC LIMIT 20", uid, uid)
	if err != nil {
		return strm, err
	}
	defer rows.Close()

	for rows.Next() {
		var p Photo
		err = rows.Scan(&p.pid, &p.uid, &p.file, &p.date)
		if err != nil {
			return strm, err
		}
		strm = append(strm, &p)
	}
	return strm, nil

}

// Get User Id
func (db *appdbimpl) GetUser(username string) (string, error) {

	var uid string
	err := db.c.QueryRow("SELECT uid FROM User WHERE username=?", username).Scan(&uid)
	if err != nil {
		return "", err
	}
	return uid, nil
}

// Get Username
func (db *appdbimpl) GetUsername(uid string) (string, error) {

	var username string
	err := db.c.QueryRow("SELECT username FROM User WHERE uid=?", uid).Scan(&username)
	if err != nil {
		return "", err
	}
	return username, nil
}
