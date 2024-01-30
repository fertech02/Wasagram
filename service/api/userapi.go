package api;

import (
	"net/http"
	"service/database"
)

// getUser handles the GET /user/:id API endpoint.
func (rt *_router) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

// setUsername handles the PUT /user/:id/username API endpoint.
func (rt *_router) setUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

