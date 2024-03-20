package api;

import (
	"net/http"
	"service/database"
)

// addBan handles the PUT /user/{uid}/ban API endpoint.
func (rt *_router) addBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	bannerId := ps.ByName("bannerId")
	bannedId := ps.ByName("bannedId")
	err := database.bandao.AddBan(bannerId, bannedId)
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
	err := database.bandao.DeleteBan(bannerId, bannedId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, nil)
}
