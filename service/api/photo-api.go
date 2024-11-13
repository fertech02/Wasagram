package api

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

// Photo path
const directory = "/tmp/filesystem/"

type FromFile struct {
	File   multipart.File
	Header multipart.FileHeader
	Mime   string
}

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	if !CheckValidAuth(r) {
		ctx.Logger.Error("Invalid Authorization")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Parse the multipart form
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error parsing form")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the photo from the form
	file, _, err := r.FormFile("file")
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting file")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()

	publisher := GetIdFromBearer(r)

	// Add the photo in the db
	pid, err := rt.db.PostPhoto(publisher)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error posting photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Save the photo
	err = rt.savePhoto(file, pid)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error saving photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (rt *_router) savePhoto(file multipart.File, pid string) error {

	// filename such that id.ext
	fileName := filepath.Join(directory, fmt.Sprintf("%s%s", pid, ".jpg"))

	// create file
	out, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	// copy file content to new file
	_, err = io.Copy(out, file)
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
