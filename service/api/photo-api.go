package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	"github.com/fertech02/Wasa-repository/service/database"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

// Photo path
const directory = "/tmp/photos"

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the file path from the request
	file, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if len(file) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Create a new photo
	uid := ps.ByName("uid")
	photo := database.Photo{
		Pid:  uuid.New().String(),
		Uid:  uid,
		File: file,
		Date: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Post the photo
	_, err = rt.db.PostPhoto(&photo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Save Photo
	err = rt.savePhoto(file, photo.Pid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (rt *_router) savePhoto(file []byte, pid string) error {

	// Create a new file
	f, err := os.Create(directory + "/" + pid + ".jpg")
	if err != nil {
		return err
	}
	defer f.Close()

	// Write the file
	_, err = f.Write(file)
	if err != nil {
		return err
	}

	return nil
}

func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the photo ID
	pid := ps.ByName("pid")
	if pid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the photo
	photo, err := rt.db.GetPhoto(pid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check if the photo exists
	if photo == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Check if the user is the owner of the photo
	uid := ps.ByName("uid")
	if photo.Uid != uid {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// Delete the photo
	err = rt.db.DeletePhoto(pid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rt *_router) getPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the user ID
	userID := ps.ByName("uid")
	if userID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the user's photos
	photos, err := rt.db.GetPhotos(userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the photos
	err = json.NewEncoder(w).Encode(photos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (rt *_router) getPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Get the photo ID
	pid := ps.ByName("pid")
	if pid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the photo
	photo, err := rt.db.GetPhoto(pid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check if the photo exists
	if photo == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Return the photo
	err = json.NewEncoder(w).Encode(photo)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
