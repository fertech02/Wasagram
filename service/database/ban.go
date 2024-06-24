package database

// Ban an User
func (db *appdbimpl) Ban(bannerId string, bannedId string) error {

	_, err := db.c.Exec("INSERT INTO Ban (bannerId, bannedId) VALUES (?, ?)", bannerId, bannedId)
	if err != nil {
		return err
	}

	return nil
}

// Unban an User
func (db *appdbimpl) Unban(bannerId string, bannedId string) error {

	_, err := db.c.Exec("DELETE FROM Ban WHERE bannerId = ? AND bannedId = ?", bannerId, bannedId)
	if err != nil {
		return err
	}

	return nil
}

// Check if a User is Banned
func (db *appdbimpl) CheckBan(bannerId string, bannedId string) (bool, error) {

	var count int
	err := db.c.QueryRow("SELECT COUNT(*) FROM Ban WHERE bannerId = ? AND bannedId = ?", bannerId, bannedId).Scan(&count)
	if err != nil {
		return false, err
	}

	return count > 0, nil
}
