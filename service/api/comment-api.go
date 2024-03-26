package api;

import (
	"net/http"
	"service/database"
	"service/models"
)

// postComment handles the POST /photo/{pid}/comment API endpoint.
func (rt *_router) postComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var commentRequest struct {
		Comment string `json:"comment"`
	}

	err := hs.decodeJSON(r.Body, &commentRequest)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	comment := models.Comment{
		uid: ps.ByName("uid"),
		pid: ps.ByName("pid"),
		message: commentRequest.Comment
	}
	
	err = database.comment-dao.Comment(comment)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

}

// deleteComment handles the DELETE /photo/{pid}/comment API endpoint.
func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	uid := ps.ByName("uid")
	err := database.comment-dao.Uncomment(pid, uid)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, nil)
}

// getComment handles the GET /photo/{pid}/comment API endpoint.
func (rt *_router) getComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	comments, err := database.comment-dao.GetComments(pid)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, comments)
}