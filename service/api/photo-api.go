package api;

import (
	"net/http"
	"service/database"
	"time"
)


// Post handles the POST /photos API endpoint.
func (rt *_router) Post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	err := r.ParseMultipartForm(10 << 20) // Max memory 10MB
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not parse multipart form")
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Could not get uploaded file")
		return
	}
	defer file.Close()

	dst, err := os.Create(handler.Filename)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not create file")
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Could not write file")
		return
	}

	uid := ps.ByName("uid")
	pid := uuid.New().String()
	url := "/photos/" + handler.Filename
	err = database.photo-dao.postPhoto(uid, pid, url)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(w, http.StatusOK, map[string]string{"url": url})
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