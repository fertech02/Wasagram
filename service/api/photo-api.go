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
	"time"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	"github.com/fertech02/Wasa-repository/service/database"
	"github.com/fertech02/Wasa-repository/service/filesystem"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

type FromFile struct {
	File   []byte `json:"file"`
	Header string `json:"header"`
	Mime   string `json:"mime"`
}

const uploadDirectory = "/tmp/filesystem/"

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// check auth
	if !CheckValidAuth(r) {
		ctx.Logger.Error("Auth header invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Parse the form data
	err := r.ParseMultipartForm(10 << 20) // 10 MB
	if err != nil {
		ctx.Logger.WithError(err).Error("Error parsing form")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// get the photo from the form
	file, handler, err := r.FormFile("file")
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	additionalData := r.FormValue("additionalData")
	var photoData database.Caption
	err = json.Unmarshal([]byte(additionalData), &photoData)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error unmarshalling additional data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	uid := GetIdFromBearer(r)
	pid := uuid.New().String()

	var photo database.Photo
	photo.Uid = uid
	photo.Pid = pid
	photo.Date = time.Now().String()

	_, err = rt.db.PostPhoto(photo)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error posting photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// save the photo in the file system
	err = saveUploadedFile(file, handler, photo.Pid)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error saving photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

// saveUploadedFile saves the uploaded file to the filesystem
func saveUploadedFile(file multipart.File, handler *multipart.FileHeader, photoID string) error {

	filename := filepath.Join(uploadDirectory, fmt.Sprintf("%s%s", photoID, ".jpg"))
	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

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

	// delete the photo in the file system
	err = filesystem.RemovePhoto(pid)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error deleting photo")
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

	// get the photo from the file system
	path := filepath.Join(uploadDirectory, fmt.Sprintf("%s%s", pid, ".jpg"))
	photoFile, err := os.Open(path)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error opening photo file")
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		w.Header().Set("Content-Type", "image/png")
		buf := bytes.NewBuffer(nil)
		_, err := io.Copy(buf, photoFile)
		if err != nil {
			ctx.Logger.WithError(err).Error("Error copying photo file")
			w.WriteHeader(http.StatusInternalServerError)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			_, err = w.Write(buf.Bytes())
			if err != nil {
				ctx.Logger.WithError(err).Error("Error writing photo file")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}
	}
}
