package database;
import (
	"database/sql"
	"errors"
	"github.com/google/uuid"
)

type User struct {
	uid	      string  		`json:"uid"`
	username  string		`json:"username"`
}

// Create a new User
func (db *appdbimpl) CreateUser(username string) (*User, error) {

	// Check if the username is already taken
	row := db.QueryRow("SELECT uid FROM User WHERE username = ?", username)
	var existingUsername string
	err = row.Scan(&existingUsername)
	if err != sql.ErrNoRows {
		return nil, errors.New("Username already taken")
	}

	// Generate a new UUID
	uid := uuid.New().String()
	
	stmt, err := db.Prepare("INSERT INTO User(uid, username) values(?, ?)")
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(uid, username)
	if err != nil {
		return nil, err
	}

	return &User{uid: uid, username: username}, nil
}

// Get User by ID
func (db *appdbimpl) GetUserProfile(uid string) (*User, error) {
	/*
		Dobbiamo implementare la ricerca di altri utenti.
		Per fare ciò, dobbiamo vedere se l'utente che cerca non appartiene
		alla lista degli utenti bannati di quello che viene cercato. 
	**/

	row := db.QueryRow("SELECT uid, username FROM User WHERE uid = ?", uid)
	var user User
	err = row.Scan(&user.uid, &user.username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update Username
func (db *appdbimpl) UpdateUsername(userid string, username string) error {

	// Check if the username is already taken
	row := db.QueryRow("SELECT uid FROM User WHERE username = ?", username)
	var existingUsername string
	err = row.Scan(&existingUsername)
	if err != sql.ErrNoRows {
		return errors.New("Username already taken")
	}

	stmt, err := db.Prepare("UPDATE User SET username = ? WHERE uid = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(username, userid)
	if err != nil {
		return err
	}

	return nil
}
