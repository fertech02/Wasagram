package api;

import (
	"net/http"
	"service/database"
)

// Like handles the POST /photos/{pid}/like API endpoint.
func (rt *_router) Like(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	pid := ps.ByName("pid")
	uid := ps.ByName("uid")
	err := database.like-dao.Like(pid, uid)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, nil)
}

// Unlike handles the DELETE /photos/{pid}/like API endpoint.
func (rt *_router) Unlike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	pid := ps.ByName("pid")
	uid := ps.ByName("uid")
	err := database.like-dao.Unlike(pid, uid)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, nil)
}

// Get Likes handles the GET /photos/{pid}/likes API endpoint.
func (rt *_router) GetLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	pid := ps.ByName("pid")
	likes, err := database.like-dao.GetLikes(pid)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, likes)
}
