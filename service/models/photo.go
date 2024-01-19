package database;

type Photo struct {
	pid  string	`json:"pid"`
	uid  string	`json:"uid"`
}

// Upload a Photo
func (u *User) UploadPhoto(p *Photo) error {
	// database query to Photos table
	
	return nil
}

// Delete a Photo
func (u *User) DeletePhoto(p *Photo) error {
	// database query to Photos table
	return nil
}

// Helper functions

func UpdatePhotoDatabase(photo *Photo) error {
	// Update database
	err := UpdatePhotoDatabase(photo)
	if err != nil {
		return err
	}
	return nil
}

func GetPhotoDatabase(photoId string) (*Photo, error) {
	// Get photo from database
	photo, err := GetPhotoDatabase(photoId)
	if err != nil {
		return nil, err
	}
	return photo, nil
}
