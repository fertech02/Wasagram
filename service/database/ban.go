package database;

import (
	"database/sql"
)

type Ban struct {
	bannerId string
	bannedId string
}

// Ban an User
func (db *appdbimpl) AddBan(bannerId, bannedId string) error{

	_, err = db.Exec("INSERT INTO Ban (bannerId, bannedId) VALUES (?, ?)", bannerId, bannedId)
	if err != nil {
		return err
	}

	return nil
}

// Unban an User
func (db *appdbimpl) UnbanUser(bannerId string, bannedId string) error {

	_, err = db.Exec("DELETE FROM Ban WHERE bannerId = ? AND bannedId = ?", bannerId, bannedId)
	if err != nil {
		return err
	}

	return nil
}

// Get Ban List
func (db *appdbimpl) GetBanList(bannerId string) ([]string, error) {
	
	rows, err := db.Query("SELECT bannedId FROM Ban WHERE bannerId = ?", bannerId)
	if err != nil {
		return nil, err
	}

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
