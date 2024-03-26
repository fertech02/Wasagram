package api;

import (
	"net/http"
	"service/database"
)

// addBan handles the PUT /user/{uid}/ban API endpoint.
func (rt *_router) addBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	bannerId := ps.ByName("bannerId")
	bannedId := ps.ByName("bannedId")
	err := database.ban-dao.AddBan(bannerId, bannedId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, nil)
}

// deleteBan handles the DELETE /user/{uid}/ban API endpoint.
func (rt *_router) deleteBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	bannerId := ps.ByName("bannerId")
	bannedId := ps.ByName("bannedId")
	err := database.ban-dao.DeleteBan(bannerId, bannedId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, nil)
}

// getBanList handles the GET /user/{uid}/ban API endpoint.

func (rt *_router) getBanList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	bannerId := ps.ByName("bannerId")
	bans, err := database.ban-dao.GetBanList(bannerId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, bans)
}
