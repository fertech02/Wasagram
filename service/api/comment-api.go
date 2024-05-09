package api

import (
	"encoding/json"
	"net/http"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var comment db.Comment
	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		logger.WithError(err).Error("error decoding request body")
		ctx.Respond(w, http.StatusBadRequest, "error decoding request body")
		return
	}

	// get the photo id from the request params
	pid := ps.ByName("pid")
	if pid == "" {
		logger.Error("missing photo id")
		ctx.Respond(w, http.StatusBadRequest, "missing photo id")
		return
	}

	// get the user id from the request params
	uid := ps.ByName("uid")
	if uid == "" {
		logger.Error("missing user id")
		ctx.Respond(w, http.StatusBadRequest, "missing user id")
		return
	}

	comment.pid = pid
	comment.uid = uid

	// get the message from the request body
	var message struct {	
		Message string `json:"message"`
	}

	err = json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		logger.WithError(err).Error("error decoding request body")
		ctx.Respond(w, http.StatusBadRequest, "error decoding request body")
		return
	}

	comment.message = message.Message

	// Add the comment in the db
	err = rt.db.AddComment(comment)
	if err != nil {
		logger.WithError(err).Error("error adding comment")
		ctx.Respond(w, http.StatusInternalServerError, "error adding comment")
		return
	}
}

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get the photo id from the request params
	pid := ps.ByName("pid")
	if pid == "" {
		logger.Error("missing photo id")
		ctx.Respond(w, http.StatusBadRequest, "missing photo id")
		return
	}

	// get the user id from the request params
	uid := ps.ByName("uid")
	if uid == "" {
		logger.Error("missing user id")
		ctx.Respond(w, http.StatusBadRequest, "missing user id")
		return
	}

	// Check if uid is authorized to delete the comment
	if !ctx.IsAuthorized(uid) {
		ctx.RespondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	// Delete the comment from the db
	err := rt.db.DeleteComment(pid, uid)
	if err != nil {
		logger.WithError(err).Error("error deleting comment")
		ctx.Respond(w, http.StatusInternalServerError, "error deleting comment")
		return
	}
}

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// get the photo id from the request params
	pid := ps.ByName("pid")
	if pid == "" {
		logger.Error("missing photo id")
		ctx.Respond(w, http.StatusBadRequest, "missing photo id")
		return
	}

	// Get the comments from the db
	comments, err := rt.db.GetComments(pid)
	if err != nil {
		logger.WithError(err).Error("error getting comments")
		ctx.Respond(w, http.StatusInternalServerError, "error getting comments")
		return
	}

	ctx.Respond(w, http.StatusOK, comments)
}
