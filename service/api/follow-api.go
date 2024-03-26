package api;

import (
	"net/http"
	"service/database"
)

// followUser handles the PUT /user/{uid}/follow API endpoint.
func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uid := ps.ByName("uid")
	followeeId := ps.ByName("followeeId")
	err := database.follow-dao.FollowUser(uid, followeeId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, nil)

}

// deleteFollow handles the DELETE /user/{uid}/follow API endpoint.
func (rt *_router) deleteFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uid := ps.ByName("uid")
	followeeId := ps.ByName("followeeId")
	err := database.follow-dao.UnfollowUser(uid, followeeId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, nil)
}

// getFollowee handles the GET /user/{uid}/follows API endpoint.
func (rt *_router) getFollows(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uid := ps.ByName("uid")
	follows, err := database.follow-dao.GetFollowees(uid)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, follows)
}
