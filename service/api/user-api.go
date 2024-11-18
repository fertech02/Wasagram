package api

import (
	"encoding/json"
	"net/http"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	"github.com/fertech02/Wasa-repository/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	var requestData struct {
		Username string `json:"username"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		ctx.Logger.WithError(err).WithField("username", requestData.Username).Error("Can't login user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	username := requestData.Username

	userID, err := rt.db.GetUserId(username)
	if err != nil {
		ctx.Logger.WithError(err).WithField("username", username).Error("Can't operate database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if userID != "" {
		response := map[string]string{"userId": userID}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			ctx.Logger.WithError(err).WithField("username", username).Error("Can't login user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	newUser, err := rt.db.CreateUser(username)
	if err != nil {
		ctx.Logger.WithError(err).WithField("username", username).Error("Can't login user")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := map[string]string{"userId": newUser.Uid}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		ctx.Logger.WithError(err).WithField("username", username).Error("Can't login user")
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

	var requestData struct {
		Username string `json:"username"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if requestData.Username == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	username := requestData.Username
	uid := ps.ByName("uid")
	err = rt.db.UpdateUsername(uid, username)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the user id from the URL
	uid := ps.ByName("uid")
	if uid == "" {
		ctx.Logger.Error("No user id")
		return
	}

	var photoStream []database.Photo
	photoStream, err := rt.db.GetProfilePhotos(uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Error during photo getting")
		return
	}

	if !CheckValidAuth(r) {
		ctx.Logger.Error("Auth header invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	myId := GetIdFromBearer(r)
	hisId := uid

	userName, err := rt.db.GetUsername(hisId)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during name getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if userName == "" {
		ctx.Logger.WithError(err).Error("Error user not found")
		w.WriteHeader(http.StatusNotFound)
		return

	}

	isBan, err := rt.db.CheckBan(myId, hisId)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during ban getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if isBan {
		ctx.Logger.Error("Banned")
		w.WriteHeader(http.StatusForbidden)
		return
	}

	var followCount int
	followCount, err = rt.db.GetFollowersCount(uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during follow count getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var followingCount int
	followingCount, err = rt.db.GetFolloweesCount(uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during following count getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var isBanned bool
	isBanned, err = rt.db.CheckBan(uid, myId)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during ban bool getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var isFollowed bool
	isFollowed, err = rt.db.CheckFollow(uid, myId)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during follow bool getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var photoCount int
	photoCount, err = rt.db.GetPhotoCount(uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during photo count getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	response := database.Profile{
		PhotoList:     photoStream,
		Username:      userName,
		FollowCount:   followCount,
		FollowedCount: followingCount,
		PhotoCount:    photoCount,
		IsBanned:      isBanned,
		IsFollowed:    isFollowed,
	}

	profileJSON, err := json.Marshal(response)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json writing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(profileJSON)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json sending")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (rt *_router) getMyStream(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	uid := ps.ByName("uid")

	ans := CheckIdAuthorized(r, uid)
	if ans != 0 {
		if ans == 2 {
			ctx.Logger.WithField("uid", uid).Error("Unauthorized")
			w.WriteHeader(http.StatusUnauthorized)
		} else {
			ctx.Logger.WithField("uid", uid).Error("Auth header invalid")
			w.WriteHeader(http.StatusUnauthorized)
		}
		return
	}

	var photoStream []database.Photo
	photoStream, err := rt.db.GetMyStream(uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during photo getting")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	photoStreamJSON, err := json.Marshal(photoStream)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json writing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(photoStreamJSON)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json sending")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	queryParams := r.URL.Query()
	username := queryParams.Get("username")

	if !CheckValidAuth(r) {
		ctx.Logger.Error("Auth header invalid")
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	var usersList []database.User
	usersList, err := rt.db.SearchUser(username)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during search")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	usersListJSON, err := json.Marshal(usersList)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json writing")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(usersListJSON)
	if err != nil {
		ctx.Logger.WithError(err).Error("Error during json sending")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
