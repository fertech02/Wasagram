package api

import (
	"encoding/json"
	"net/http"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	db "github.com/fertech02/Wasa-repository/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Create a new comment
	comment := db.Comment{}
	// get the photo id from the request params
	photoId := ps.ByName("pid")

	// get the user id from the request params
	userId := ps.ByName("uid")

	comment.Pid = photoId
	comment.Uid = userId

	// get the message from the request body
	var message struct {
		Message string `json:"message"`
	}

	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	comment.Message = message.Message

	// Add the comment in the db
	err = rt.db.Comment(&comment)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get the photo id from the request params
	pid := ps.ByName("pid")
	if pid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// get the user id from the request params
	uid := ps.ByName("uid")
	if uid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Delete the comment from the db
	err := rt.db.Uncomment(pid, uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get the photo id from the request params
	pid := ps.ByName("pid")
	if pid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the comments from the db
	comments, err := rt.db.GetComments(pid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the comments
	err = json.NewEncoder(w).Encode(comments)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
