package database

// Ban an User
func (db *appdbimpl) Ban(bannerId string, bannedId string) error {

	_, err := db.c.Exec("INSERT INTO Bans (BannerId, BannedId) VALUES (?, ?)", bannerId, bannedId)
	if err != nil {
		return err
	}

	return nil
}

// Unban an User
func (db *appdbimpl) Unban(bannerId string, bannedId string) error {

	_, err := db.c.Exec("DELETE FROM Bans WHERE BannerId = ? AND BannedId = ?", bannerId, bannedId)
	if err != nil {
		return err
	}

	return nil
}

// Check if a User is Banned
func (db *appdbimpl) CheckBan(bannerId string, bannedId string) (bool, error) {

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Bans WHERE BannerId = ? AND BannedId = ?", bannerId, bannedId).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

// Get Banned Users
func (db *appdbimpl) GetBannedUsers(bannerId string) ([]string, error) {

	var bannedUsers []string
	rows, err := db.c.Query("SELECT BannedId FROM Bans WHERE BannerId = ?", bannerId)
	if err != nil {
		return bannedUsers, err
	}
	defer rows.Close()

	for rows.Next() {
		var bannedId string
		err = rows.Scan(&bannedId)
		if err != nil {
			return bannedUsers, err
		}
		bannedUsers = append(bannedUsers, bannedId)
	}

	if err = rows.Err(); err != nil {
		return bannedUsers, err
	}

	return bannedUsers, nil
}
