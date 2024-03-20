package database;

// Ban an User
func (u *User) BanUser() {
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
func (u *User) UnbanUser() {
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
func (u *User) GetBanList() {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query("SELECT bannerId, bannedId FROM Ban WHERE bannerId = ? AND bannedId = ?", bannerId, bannedId)
	if err != nil {
		return nil, err
	}
	defer db.Close()
}
