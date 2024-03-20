package database;

// Create a new User
func CreateUser(username string) (*User, error) {
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Generate a new UUID
	uid := uuid.New().String()
	
	stmt, err := db.Prepare("INSERT INTO User(userid, username) values(?, ?)")
	if err != nil {
		return nil, err
	}

	res, err := stmt.Exec(username)
	if err != nil {
		return nil, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return nil, err
	}

	user: = &User{
		uid:   id,
		username: username,
	}

	return user, nil
}

// Get User by ID
func GetUserProfile(id string) (*User, error) {
	db, err := sql.Open("sqlite3","./foo.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	row := db.QueryRow("SELECT uid, username FROM User WHERE uid = ?", uid)
	var user User
	err:= row.Scan(&user.uid, &user.username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// Update Username
func UpdateUsernameDB(u *User) error {
	db, err := sql.Open("sqlite3","./foo.db")
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("UPDATE User SET username = ? WHERE uid = ?")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.username, u.uid)
	if err != nil {
		return err
	}

	return nil
}
