package database;

type Like struct {
	uid  string	`json:"uid"`
	pid string	`json:"pid"`
}

// Get Like by ID
func GetLike(id string) (*Like, error) {
	// database query to Likes table
	return &Like{}, nil
}

// Like a Photo
func (u *User) Like(p *Post) error {
}


// Unlike a Photo
func (u *User) Unlike(p *Post) error {
}
