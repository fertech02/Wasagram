package database;

import (
	"database/sql"
)

type Follow struct {
	followeeId string
	followerId string
}

// Follow an User
func (u *User) FollowUser(followeeId int, followerId int) error {

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Follow (followeeId, followerId) VALUES (?, ?)", followeeId, followerId)
	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}

// Unfollow an User
func (u *User) UnfollowUser(followeeId int, followerId int) error {

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Follow WHERE followeeId = ? AND followerId = ?", followeeId, followerId)
	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}

// Get Followers List
func (u *User) GetFollowers(uid int) error{

	db,err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("SELECT followerId FROM Follow WHERE followeeId = ?", uid)
	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}

// Get Followees List
func (u *User) GetFollowees(uid int) ([]string, error) {

	db,err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

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