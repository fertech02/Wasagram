package database;

import (
	"database/sql"
)

// Ban an User
func (u *User) AddBan(bannerId, bannedId string) {

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.QueryRow("INSERT INTO Ban (bannerId, bannedId) VALUES (?, ?)", bannerId, bannedId)
	if err != nil {
		return nil, err
	}
	defer db.Close()
}

// Unban an User
func (u *User) UnbanUser(bannerId string, bannedId string) {

	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.QueryRow("DELETE FROM Ban WHERE bannerId = ? AND bannedId = ?", bannerId, bannedId)
	if err != nil {
		return nil, err
	}
	defer db.Close()
}

// Get Ban List
func (u *User) GetBanList(bannerId string) {
	
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.QueryRow("SELECT * FROM Ban WHERE bannerId = ?", bannerId)
	if err != nil {
		return nil, err
	}
	defer db.Close()
}
