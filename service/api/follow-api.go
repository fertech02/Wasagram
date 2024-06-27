package api

import (
	"encoding/json"
	"net/http"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	followeeId := ps.ByName("followeeId")
	if followeeId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	followerId := ps.ByName("followerId")
	if followerId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if followerId is authorized to follow followeeId
	unauthorized, err := CheckAuthorizedId(r, followerId)
	if err != nil {
		ctx.Logger.WithField("error", err).Error("Failed to check authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if unauthorized {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Follow the user
	err = rt.db.Follow(followeeId, followerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	followeeId := ps.ByName("followeeId")
	if followeeId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	followerId := ps.ByName("followerId")
	if followerId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if followerId is authorized to unfollow followeeId
	unauthorized, err := CheckAuthorizedId(r, followerId)
	if err != nil {
		ctx.Logger.WithField("error", err).Error("Failed to check authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if unauthorized {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Unfollow the user
	err = rt.db.Unfollow(followeeId, followerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

/*
func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	uid := ps.ByName("uid")
	if uid == "" {
		w.WriteHeader(http.StatusBadRequest);
		return
	}

	// Get the followers from the db
	followers, err := rt.db.GetFollowers(uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError);
		return
	}

	// Return the followers
	err = json.NewEncoder(w).Encode(followers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError);
		return
	}
}
*/

func (rt *_router) getFollowedUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	uid := ps.ByName("uid")
	if uid == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the followees from the db
	// -- DA VEDERE
	followees, err := rt.db.GetFollowees(uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the followees
	err = json.NewEncoder(w).Encode(followees)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
