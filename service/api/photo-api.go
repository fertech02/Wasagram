package api

import (
	"net/http"
	"errors"
	"time"
	"io"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	"github.com/fertech02/Wasa-repository/service/database"
	"github.com/julienschmidt/httprouter"
	"github.com/google/uuid"
)

// Photo path
const directory = "/tmp/photos"

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		ctx.Error(w, http.StatusUnauthorized, "missing Authorization header")
		return
	}

	// Validate the Authorization header
	uid, err := rt.db.ValidateToken(authHeader)
	if err != nil {
		ctx.Error(w, http.StatusUnauthorized, "invalid token")
		return
	}

	// Get the file from the request
	file, _, err := io.ReadAll(r.Body)
	if err != nil {
		ctx.Error(w, http.StatusBadRequest, "invalid request")
		return
	}

	// Check if the file is empty
	if len(file) == 0 {
		ctx.Error(w, http.StatusBadRequest, "empty file")
		return
	}

	// Create a new photo
	photo := database.Photo{
		pid:  uuid.New().String(),
		uid:  uid,
		file: file,
		date: time.Now().Format(time.RFC3339),
	}

	// Post the photo
	_, err = rt.db.PostPhoto(photo)
	if err != nil {
		ctx.Error(w, http.StatusInternalServerError, "error posting photo")
		return
	}

	// Save Photo
	err = rt.savePhoto(file, photo.pid)
	if err != nil {
		ctx.Error(w, http.StatusInternalServerError, "error saving photo")
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
		ctx.Error(w, http.StatusUnauthorized, "missing Authorization header")
		return
	}

	// Validate the Authorization header
	uid, err := rt.db.ValidateToken(authHeader)
	if err != nil {
		ctx.Error(w, http.StatusUnauthorized, "invalid token")
		return
	}

	// Get the photo ID
	pid := ps.ByName("pid")
	if pid == "" {
		ctx.Error(w, http.StatusBadRequest, "photo ID is required")
		return
	}

	// Get the photo
	photo, err := rt.db.GetPhoto(pid)
	if err != nil {
		ctx.Error(w, http.StatusInternalServerError, "error getting photo")
		return
	}

	// Check if the photo exists
	if photo == nil {
		ctx.Error(w, http.StatusNotFound, "photo not found")
		return
	}

	// Check if the user is the owner of the photo
	if photo.uid != uid {
		ctx.Error(w, http.StatusForbidden, "forbidden")
		return
	}

	// Delete the photo
	err = rt.db.DeletePhoto(pid)
	if err != nil {
		ctx.Error(w, http.StatusInternalServerError, "error deleting photo")
		return
	}

	w.WriteHeader(http.StatusOK)
}

fun (rt *_router) getPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the Authorization header	
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		ctx.Error(w, http.StatusUnauthorized, "missing Authorization header")
		return
	}

	// Validate the Authorization header
	uid, err := rt.db.ValidateToken(authHeader)
	if err != nil {
		ctx.Error(w, http.StatusUnauthorized, "invalid token")
		return
	}

	// Get the user ID
	uid := ps.ByName("uid")
	if uid == "" {
		ctx.Error(w, http.StatusBadRequest, "user ID is required")
		return
	}

	// Get the user's photos
	photos, err := rt.db.GetPhotos(uid)
	if err != nil {
		ctx.Error(w, http.StatusInternalServerError, "error getting photos")
		return
	}

	// get photo from filesystem
	for _, photo := range photos {
		


}
