package api;

import (
	"net/http"
	"service/database"
	"github.com/google/uuid"
)


// Post handles the POST /photos API endpoint.
func (rt *_router) Post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
	uid = ps.ByName("uid")
	pid = uuid.New().String()
	photo, err := database.photo-dao.PostPhoto()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusCreated, photo)
}

// deletePhoto handles the DELETE /photos/{pid} API endpoint.
func (rt *_router) Delete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uid = ps.ByName("uid")
	pid = ps.ByName("pid")

	photo, err := database.photo-dao.deletePhoto(uid, pid)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, photo)
}

// getPhotos handles the GET /photos API endpoint.
func (rt *_router) getPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	uid = ps.ByName("uid")

	photos, err := database.photo-dao.getPhotos(uid)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, photos)
}