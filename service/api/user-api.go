package api;

import (
	"net/http"
	"service/database"
)

// getUser handles the GET /user/:id API endpoint.
func (rt *_router) GetUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userid := ps.ByName("id")
	user, err := database.userdao.GetUser(userid)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

// updateUsername handles the PUT /user/:id/username API endpoint.
func (rt *_router) updateUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userid := ps.ByName("id")

	var updateRequest struct {
		Username string `json:"username"`
	}

	err := json.NewDecoder(r.Body).Decode(&updateRequest)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	err = database.userdao.UpdateUsername(userid, updateRequest.Username)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// createUser handles the POST /session API endpoint.
func (rt *_router) createUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var createRequest struct {
		Username string `json:"username"`
	}

	err := json.NewDecoder(r.Body).Decode(&createRequest)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}

	user, err := database.userdao.CreateUser(createRequest.Username)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}