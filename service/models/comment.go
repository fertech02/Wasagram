package database;

type Comment struct {
	uid string	`json:"uid"`
	pid string	`json:"pid"`
	message string	`json:"message"`
}

// Comment a Photo
func (u *User) Comment(p *Post, message string) error {
	// database query to Comments table
	return nil
}

// Uncomment a Photo
func Uncomment(p *Post) error {
	// database query to Comments table
	return nil
}

