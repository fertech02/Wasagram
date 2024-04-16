package database;

import (
	"database/sql"
)

type Follow struct {
	followeeId string
	followerId string
}

// Follow an User
func (db *appdbimpl) FollowUser(followeeId int, followerId int) error {

	_, err = db.Exec("INSERT INTO Follow (followeeId, followerId) VALUES (?, ?)", followeeId, followerId)
	if err != nil {
		return err
	}

	return nil
}

// Unfollow an User
func (db *appdbimpl) UnfollowUser(followeeId int, followerId int) error {

	_, err = db.Exec("DELETE FROM Follow WHERE followeeId = ? AND followerId = ?", followeeId, followerId)
	if err != nil {
		return err
	}

	return nil
}

// Get Followers List
func (db *appdbimpl) GetFollowers(uid int) error{

	_, err = db.Exec("SELECT followerId FROM Follow WHERE followeeId = ?", uid)
	if err != nil {
		return err
	}

	return nil
}

// Get Followees List
func (db *appdbimpl) GetFollowees(uid int) ([]string, error) {

	rows, err := db.Query("SELECT followeeId FROM Follow WHERE followerId = ?", uid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followees []string
	for rows.Next() {
        var followeeId string
        if err := rows.Scan(&followeeId); err != nil {
            return nil, err
        }
        followees = append(followees, followeeId)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return followees, nil

}