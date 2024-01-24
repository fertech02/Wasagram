package api;

import (
	"net/http"
	"service/database"
)

// Like handles the POST /pictures/{pid}/like API endpoint.
func (rt *_router) Like(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// get the user id from the URL
	pid := ps.ByName("pid")
	uid := ps.ByName("uid")
	err := database.likedao.Like(pid, uid)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, nil)
}

// Unlike handles the DELETE /pictures/{pid}/like API endpoint.
func (rt *_router) Unlike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// get the user id from the URL
	pid := ps.ByName("pid")
	uid := ps.ByName("uid")
	err := database.likedao.Unlike(pid, uid)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, nil)
}

