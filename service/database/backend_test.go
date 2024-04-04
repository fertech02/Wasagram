package database;
import (
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestPostPhoto(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Create a new User
	u := &User{}

	// Create a new photo
	uid := "1234"
	pid := "5678"
	url := "/photos/1234.jpg"

	// Add expectations
	mock.ExpectExec("INSERT INTO Photos").WithArgs(uid, pid, url).WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the function
	err = u.postPhoto(uid, pid, url)
	if err != nil {
		t.Fatalf("error was not expected while posting a photo: %s", err)
	}

	// Assert that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDeletePhoto(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Create a new User
	u := &User{}

	// Create a new photo
	uid := "1234"
	pid := "5678"

	// Add expectations
	mock.ExpectExec("DELETE FROM Photos").WithArgs(uid, pid).WillReturnResult(sqlmock.NewResult(1, 1))

	// Call the function
	err = u.deletePhoto(uid, pid)
	if err != nil {
		t.Fatalf("error was not expected while deleting a photo: %s", err)
	}

	// Assert that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestGetPhotos(t *testing.T) {
	// Create a new mock database
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	// Create a new User
	u := &User{}

	// Create a new photo
	uid := "1234"
	pid := "5678"
	url := "/photos/1234.jpg"

	// Add expectations
	rows := sqlmock.NewRows([]string{"pid", "url"}).AddRow(pid, url)
	mock.ExpectQuery("SELECT * FROM Photos").WithArgs(uid).WillReturnRows(rows)

	// Call the function
	err = u.getPhotos(uid)
	if err != nil {
		t.Fatalf("error was not expected while getting photos: %s", err)
	}

	// Assert that all expectations were met
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}




