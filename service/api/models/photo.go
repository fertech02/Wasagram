package models

type Photo struct {
	PhotoId  string
	owner    string
	likes    []Like
	comments []Comment
}

// Like a Photo
func (liker *User) Like(photo *Photo) error {
	// add liker to photo's likes
	photo.likes = append(photo.likes, Like{userId: liker.UserId, photoId: photo.PhotoId})

	// Update database
	err := UpdatePhotoDatabase(photo)
	if err != nil {
		return err
	}
	return nil
}

// Unlike a Photo
func (unliker *User) Unlike(photo *Photo) error {
	// remove liker from photo's likes
	for i, l := range photo.likes {
		if l.userId == unliker.UserId {
			photo.likes = append(photo.likes[:i], photo.likes[i+1:]...)
			break
		}
	}

	// Update database
	err := UpdatePhotoDatabase(photo)
	if err != nil {
		return err
	}
	return nil
}

// Comment on a Photo
func (commenter *User) Comment(photo *Photo, comment string) error {
	// add comment to photo's comments
	photo.comments = append(photo.comments, Comment{owner: commenter.UserId, text: comment})

	// Update database
	err := UpdatePhotoDatabase(photo)
	if err != nil {
		return err
	}
	return nil
}

// Uncomment on a Photo
func (uncommenter *User) Uncomment(photo *Photo, comment string) error {
	// remove comment from photo's comments
	for i, c := range photo.comments {
		if c.text == comment {
			photo.comments = append(photo.comments[:i], photo.comments[i+1:]...)
			break
		}
	}

	// Update database
	err := UpdatePhotoDatabase(photo)
	if err != nil {
		return err
	}
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
