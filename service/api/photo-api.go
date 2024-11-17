package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	filesystem "github.com/fertech02/Wasa-repository/service/filesystem"
	"github.com/julienschmidt/httprouter"
)

type FromFile struct {
	File   multipart.File
	Header multipart.FileHeader
	Mime   string
}

const uploadDirectory = "/tmp/filesystem/"

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

	fileName := filepath.Join(uploadDirectory, fmt.Sprintf("%s%s", pid, ".jpg"))

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

	pid := ps.ByName("pid")

	uid, err := rt.db.GetPhotoAuthor(pid)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting photo author")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	ans := CheckIdAuthorized(r, uid)
	if ans != 0 {
		ctx.Logger.Error("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// delete the photo in the filesystem
	err = filesystem.RemovePhoto(pid)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error removing photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// delete the photo in the db
	err = rt.db.DeletePhoto(pid)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error deleting photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

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

	pid := ps.ByName("pid")

	// Check Authorization
	if !CheckValidAuth(r) {
		ctx.Logger.Error("Invalid Authorization")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	myId := GetIdFromBearer(r)

	hisId, err := rt.db.GetPhotoAuthor(pid)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting photo author")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	isBan, err := rt.db.CheckBan(hisId, myId)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error checking ban")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if isBan {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// get the photo
	path := "/tmp/filesystem/" + pid + ".jpg"
	photoFile, err := os.Open(path)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error opening photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "image/png")
		buf := bytes.NewBuffer(nil)
		_, err := io.Copy(buf, photoFile)
		if err != nil {
			ctx.Logger.WithError(err).Error("Error copying photo")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			_, err := w.Write(buf.Bytes())
			if err != nil {
				ctx.Logger.WithError(err).Error("Error writing photo")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
		}
	}
}
