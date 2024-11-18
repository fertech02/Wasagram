package api

import (
	"encoding/json"
	"io"
	"net/http"
	"time"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// check auth
	if !CheckValidAuth(r) {
		ctx.Logger.Error("Auth header invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	var photo Photo
	var err error

	photo.Uid = GetIdFromBearer(r)
	photo.Pid = uuid.New().String()
	photo.File, err = io.ReadAll(r.Body)
	if err != nil {
		ctx.Logger.Error("Error reading photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	photo.Date = time.Now().Format("2006-01-02 15:04:05")

	// add photo info to database
	PhotoData, err := rt.db.PostPhoto(photo.PhotoToDatabase())
	if err != nil {
		ctx.Logger.WithError(err).Error("Error posting photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	photo.PhotoFromDatabase(PhotoData)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(photo)
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

	// get the photo
	dbphoto, present, err := rt.db.GetPhoto(pid)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error getting photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if !present {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Da aggiustare
	var photo Photo
	photo.PhotoFromDatabase(dbphoto)
	mimeType := http.DetectContentType(photo.File)
	w.Header().Set("Content-Type", mimeType)
	_, err = w.Write(photo.File)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error writing photo")
		http.Error(w, "Error writing photo", http.StatusInternalServerError)
		return
	}
}
