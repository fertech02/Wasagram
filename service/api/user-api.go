package api

import (
	"encoding/json"
	"net/http"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	"github.com/fertech02/Wasa-repository/service/database"
	"github.com/julienschmidt/httprouter"
)

type User struct {
	Uid      string `json:"uid"`
	Username string `json:"username"`
}

func (u *User) FromDatabase(user database.User) {
	u.Uid = user.Uid
	u.Username = user.Username
}

func (u *User) ToDatabase() database.User {
	return database.User{
		Uid:       u.Uid,
		Username: u.Username,
	}
}

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {


	var user User
	var requestData struct {
		Username string `json:"username"`
	}
	err := json.NewDecoder(r.Body).Decode(&requestData)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := requestData.Username

	userid, err := rt.db.GetUserId(username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// The user already exists
	if userid != "" {
		w.WriteHeader(http.StatusOK)
		return
	} 
	// The user does not exist
	DBuser, err := rt.db.CreateUser(username)
	if err != nil {
		http.Error(w, "Error creating user", http.StatusInternalServerError)
	}

	w.Header().Set("content-type", "application/json")
	user.FromDatabase(DBuser)
	_ = json.NewEncoder(w).Encode(user)
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