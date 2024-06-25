package database

import (
	"github.com/google/uuid"
)

// Create a new User
func (db *appdbimpl) CreateUser(username string) (*User, error) {

	_, err := db.c.Exec("INSERT INTO User (Uid, Username) VALUES (?, ?)", uuid.New().String(), username)
	if err != nil {
		return nil, err
	}
	return &User{Uid: uuid.New().String(), Username: username}, nil
}

// Update Username
func (db *appdbimpl) UpdateUsername(userid string, username string) error {

	_, err := db.c.Exec("UPDATE User SET Username=? WHERE Uid=?", username, userid)
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
	return strm, nil

}

// Get User Id
func (db *appdbimpl) GetUserId(username string) (string, error) {

	var uid string
	err := db.c.QueryRow("SELECT Uid FROM User WHERE Username=?", username).Scan(&uid)
	if err != nil {
		return "", err
	}
	return uid, nil
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
