package database;

import (
	"database/sql"
)

type Ban struct {
	bannerId string
	bannedId string
}

// Ban an User
func (u *User) AddBan(bannerId, bannedId string) error{

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO Ban (bannerId, bannedId) VALUES (?, ?)", bannerId, bannedId)
	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}

// Unban an User
func (u *User) UnbanUser(bannerId string, bannedId string) error {

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM Ban WHERE bannerId = ? AND bannedId = ?", bannerId, bannedId)
	if err != nil {
		return err
	}
	defer db.Close()

	return nil
}

// Get Ban List
func (u *User) GetBanList(bannerId string) ([]string, error) {
	
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT bannedId FROM Ban WHERE bannerId = ?", bannerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bannedUsers []string
	for rows.Next() {
		var bannedId string
		if err := rows.Scan(&bannedId); err != nil {
			return nil, err
		}
		bannedUsers = append(bannedUsers, bannedId)
	}

	return bannedUsers, nil
}
