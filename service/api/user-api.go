package api;

import (
	"net/http"
	"service/database"
)

// getUser handles the GET /user/:id API endpoint.
func (rt *_router) GetUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	tokenString := string.Split(authHeader, " ")[1]

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		http.Error(w, "Invalid Token", http.StatusUnauthorized)
		return
	}
	// Get the requesting user's ID from the token
	claims := token.Claims.(jwt.MapClaims)
	requestingUserID := claims["uid"].(string)

	profileUserID := ps.ByName("id")

	// Check if profileUserID is not in the banned list of userid
	bannedUsers, err := database.ban-dao.GetBanList(profileUserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, bannedUser := range bannedUsers {
		if bannedUser == profileUserID {
			http.Error(w, "You are banned from viewing this profile")
			return
		}
	}
	
	user, err := database.user-dao.GetUser(userid)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}

// updateUsername handles the PUT /user/:id/username API endpoint.
func (rt *_router) updateUsername(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Get the token from the Authorization header
	authHeader := r.Header.Get("Authorization")
	tokenString := strings.Split.(authHeader, " ")[1]

	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})

	if err != nil {
		http.Error(w,"Invalid Token",http.StatusUnauthorized)
		return
	}

	// Get the requesting user's ID from the token
	claims := token.Claims.(jwt.MapClaims)
	requestingUserID := claims["uid"].(string)

	userid := ps.ByName("uid")

	if requestingUserID != userid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var updateRequest struct {
		Username string `json:"username"`
	}

	err := json.NewDecoder(r.Body).Decode(&updateRequest)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	
	err = database.user-dao.UpdateUsername(userid, updateRequest.Username)
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

	user, err := database.user-dao.CreateUser(createRequest.Username)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, user)
}