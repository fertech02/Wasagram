package models

type User struct {
	UserId    string
	username  string
	followers []string
	followees []string
	banned    []string
}

// Get User by ID
func GetUser(id string) (*User, error) {
	// database query to Users table
	return &User{}, nil
}

// Set Username
func (u *User) SetUsername(username string) {
	// Set new username
	u.username = username

	// Update database
	err := UpdateUserDataBase(u)
	if err != nil {
		return err
	}
	return nil
}

// Follow an User
func (follower *User) Follow(followee *User) error {
	// add followee to follower's followees
	followee.followers = append(followee.followers, follower.UserId)
	// add follower to followee's followers
	follower.followees = append(follower.followees, followee.UserId)

	// Update database
	err := UpdateFollowerDatabase(follower)
	if err != nil {
		return err
	}

	err = UpdateFolloweeDatabase(followee)
	if err != nil {
		return err
	}
	return nil
}

// Unfollow an User
func (follower *User) Unfollow(followee *User) error {
	// remove followee from follower's followees
	for i, id := range follower.followees {
		if id == followee.UserId {
			follower.followees = append(follower.followees[:i], follower.followees[i+1:]...)
			break
		}
	}
	// remove follower from followee's followers
	for i, id := range followee.followers {
		if id == follower.UserId {
			followee.followers = append(followee.followers[:i], followee.followers[i+1:]...)
			break
		}
	}

	// Update database
	err := UpdateFollowerDatabase(follower)
	if err != nil {
		return err
	}

	err = UpdateFolloweeDatabase(followee)
	if err != nil {
		return err
	}
	return nil
}

// Ban an User
func (u *User) Ban(user *User) error {
	// Add banned user to user's banned list
	u.banned = append(u.banned, user.UserId)

	// Update database
	err := AddBannedUser(u)
	if err != nil {
		return err
	}
	return nil
}

// Unban an User
func (u *User) Unban(user *User) error {
	// Remove banned user from user's banned list
	for i, id := range u.banned {
		if id == user.UserId {
			u.banned = append(u.banned[:i], u.banned[i+1:]...)
			break
		}
	}

	// Update database
	err := AddBannedUser(u)
	if err != nil {
		return err
	}
	return nil
}

// Helper functions

func UpdateFollowerDatabase(follower *User) error {
	return nil
}

func UpdateFolloweeDatabase(followee *User) error {
	return nil
}

func UpdateUserDataBase(user *User) error {
	// Connect to database
	// Convert User struct to Database model
	userModel := convertToModel(user)

	// Update usernae in database
	stmt := "UPDATE Users SET username = $1 WHERE userId = $2"
	_, err := db.Exec(stmt, userModel.username, userModel.userId)
	if err != nil {
		return err
	}
	return nil
}

func convertToModel(user *User) User {
	return &User{
		UserId:   user.UserId,
		username: user.username,
	}
}

func AddBannedUser(user *User) error {
	return nil
}
