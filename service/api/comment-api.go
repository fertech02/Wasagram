package api

import (
	"encoding/json"
	"net/http"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	db "github.com/fertech02/Wasa-repository/service/database"
	"github.com/google/uuid"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Create a new comment
	if !CheckValidAuth(r) {
		ctx.Logger.Error("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userId := GetIdFromBearer(r)
	photoId := ps.ByName("pid")
	author, err := rt.db.GetPhotoAuthor(photoId)
	if err != nil || author == "" {
		w.WriteHeader(http.StatusNotFound)
		ctx.Logger.Error("Photo not found")
		return
	}

	// get the message from the request body
	var message struct {
		Message string `json:"Message"`
	}

	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		ctx.Logger.WithError(err).Error("Error decoding message")
		return
	}

	commentId := uuid.New().String()
	// Create the comment
	comment := db.Comment{
		Cid:     commentId,
		Pid:     photoId,
		Uid:     userId,
		Message: message.Message,
	}

	// Add the comment in the db
	err = rt.db.Comment(comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Error commenting photo")
		return
	}
}

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get the comment id from the request params
	cid := ps.ByName("cid")
	if cid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the comment author

	// Delete the comment from the db
	err := rt.db.Uncomment(cid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	pid := ps.ByName("pid")

	if !CheckValidAuth(r) {
		ctx.Logger.Error("Unauthorized")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	myId := GetIdFromBearer(r)

	hisId, err := rt.db.GetPhotoAuthor(pid)
	if err != nil {
		ctx.Logger.Error("Error getting photo author")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	isBan, err := rt.db.CheckBan(hisId, myId)
	if err != nil {
		ctx.Logger.Error("Error checking ban status")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if isBan {
		ctx.Logger.Error("User is banned")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var comments []db.Comment
	comments, err = rt.db.GetComments(pid)
	if err != nil {
		ctx.Logger.Error("Error getting comments")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	commentsJson, err := json.Marshal(comments)
	if err != nil {
		ctx.Logger.Error("Error marshalling comments")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(commentsJson)
	if err != nil {
		ctx.Logger.Error("Error writing comments")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
