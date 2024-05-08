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
		ctx.Error(w, http.StatusBadRequest, "invalid request")
		return
	}

	if requestData.Username == "" {
		ctx.Error(w, http.StatusBadRequest, "username is required")
		return
	}

	username := requestData.Username

	// check if the user exists
	uid, err := rt.db.GetUserByUsername(username)
	if err != nil {
		ctx.Error(w, http.StatusInternalServerError, "error checking user")
		return
	}

	if uid != "" {
		response := map[string]string{"uid": uid}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			ctx.Error(w, http.StatusInternalServerError, "error encoding response")
			return
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	nuid, err := rt.db.CreateUser(username)
	if err != nil {
		ctx.Error(w, http.StatusInternalServerError, "error creating user")
		return
	}

	response := map[string]string{"uid": nuid}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		ctx.Error(w, http.StatusInternalServerError, "error encoding response")
		return
	}

}

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		ctx.Error(w, http.StatusUnauthorized, "missing Authorization header")
		return
	}

	// Validate the token
	isValid, err := rt.auth.validateToken(authHeader)
	if err != nil || !isValid {
		ctx.Error(w, http.StatusUnauthorized, "invalid token")
		return
	}

	var requestData struct {
		Username string `json:"username"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		ctx.Error(w, http.StatusBadRequest, "invalid request")
		return
	}

	if requestData.Username == "" {
		ctx.Error(w, http.StatusBadRequest, "username is required")
		return
	}

	username := requestData.Username

	err = rt.db.SetUserName(ctx.UID(), username)
	if err != nil {
		ctx.Error(w, http.StatusInternalServerError, "error setting username")
		return
	}

	w.WriteHeader(http.StatusOK)

}

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		ctx.Error(w, http.StatusUnauthorized, "missing Authorization header")
		return
	}

	// Validate the token
	isValid, err := rt.auth.validateToken(authHeader)
	if err != nil || !isValid {
		ctx.Error(w, http.StatusUnauthorized, "invalid token")
		return
	}

	uid := ps.ByName("uid")

	profile, err := rt.db.GetUserProfile(uid)
	if err != nil {
		ctx.Error(w, http.StatusInternalServerError, "error getting user profile")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(profile)
	if err != nil {
		ctx.Error(w, http.StatusInternalServerError, "error encoding response")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the Authorization header
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		ctx.Error(w, http.StatusUnauthorized, "missing Authorization header")
		return
	}

	// Validate the token
	isValid, err := rt.auth.validateToken(authHeader)
	if err != nil || !isValid {
		ctx.Error(w, http.StatusUnauthorized, "invalid token")
		return
	}

	photos, err := rt.db.GetMyStream(ctx.UID())
	if err != nil {
		ctx.Error(w, http.StatusInternalServerError, "error getting stream")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(photos)
	if err != nil {
		ctx.Error(w, http.StatusInternalServerError, "error encoding response")
		return
	}

	w.WriteHeader(http.StatusOK)
}
