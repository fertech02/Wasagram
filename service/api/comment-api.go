package api;

import (
	"net/http"
	"service/database"
	"service/models"
)

// postComment handles the POST /photo/{pid}/comment API endpoint.
func (rt *_router) postComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var comment models.Comment
	err := parseJSON(r, &comment)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err = database.commentdao.Comment(comment)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, comment)
}

// deleteComment handles the DELETE /photo/{pid}/comment API endpoint.
func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	pid := ps.ByName("pid")
	uif := ps.ByName("uid")
	err := database.commentdao.Uncomment(pid, uid)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, nil)
}

// getComment handles the GET /photo/{pid}/comment API endpoint.


