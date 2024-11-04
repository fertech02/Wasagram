package api

import (
	"encoding/json"
	"net/http"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var requestData struct {
		Username string `json:"username"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		ctx.Logger.Error("Error decoding request body: %v", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := requestData.Username

	userId, err := rt.db.GetUserId(username)

	if err != nil {
		ctx.Logger.WithError(err).WithField("username", username).Error("Error getting user id")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userId != "" {
		response := map[string]string{"userId": userId}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			ctx.Logger.WithError(err).WithField("username", username).Error("Error encoding response")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	newUser, err := rt.db.CreateUser(username)
	if err != nil {
		ctx.Logger.WithError(err).WithField("username", username).Error("Error creating user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := map[string]string{"userId": newUser.Uid}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		ctx.Logger.WithError(err).WithField("username", username).Error("Error encoding response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Validate the token
	isValid, err := validateToken(authHeader)
	if err != nil || !isValid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var requestData struct {
		Username string `json:"username"`
	}

	err = json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if requestData.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := requestData.Username
	userId := ps.ByName("uid")
	err = rt.db.UpdateUsername(userId, username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Validate the token
	isValid, err := validateToken(authHeader)
	if err != nil || !isValid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	uid := ps.ByName("uid")

	profile, err := rt.db.GetUserProfile(uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Validate the token
	isValid, err := validateToken(authHeader)
	if err != nil || !isValid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	userId := ps.ByName("uid")
	photos, err := rt.db.GetMyStream(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(photos)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Validate the token
	isValid, err := validateToken(authHeader)
	if err != nil || !isValid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	usernameToSearch := r.URL.Query().Get("username")
	usersList, err := rt.db.SearchUser(usernameToSearch)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(usersList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
