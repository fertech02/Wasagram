package api

import (
	"net/http"

	"github.com/fertech02/Wasa-repository/service/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	
	followeeId, err := ps.ByName("followeeId")
	if err != nil {
		ctx.RespondWithError(w, http.StatusBadRequest, "missing followeeId")
		return
	}

	followerId, err := ps.ByName("followerId")
	if err != nil {
		ctx.RespondWithError(w, http.StatusBadRequest, "missing followerId")
		return
	}

	// Check if followerId is authorized to follow followeeId
	if !ctx.IsAuthorized(followerId) {
		ctx.RespondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	// Follow the user
	err = rt.db.FollowUser(followeeId, followerId)
	if err != nil {
		ctx.RespondWithError(w, http.StatusInternalServerError, "error following user")
		return
	}
}

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	followeeId, err := ps.ByName("followeeId")
	if err != nil {
		ctx.RespondWithError(w, http.StatusBadRequest, "missing followeeId")
		return
	}

	followerId, err := ps.ByName("followerId")
	if err != nil {
		ctx.RespondWithError(w, http.StatusBadRequest, "missing followerId")
		return
	}

	// Check if followerId is authorized to unfollow followeeId
	if !ctx.IsAuthorized(followerId) {
		ctx.RespondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	// Unfollow the user
	err = rt.db.UnfollowUser(followeeId, followerId)
	if err != nil {
		ctx.RespondWithError(w, http.StatusInternalServerError, "error unfollowing user")
		return
	}
}

func (rt *_router) getFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	uid, err := ps.ByName("uid")
	if err != nil {
		ctx.RespondWithError(w, http.StatusBadRequest, "missing uid")
		return
	}

	// Get the followers from the db
	followers, err := rt.db.GetFollowers(uid)
	if err != nil {
		ctx.RespondWithError(w, http.StatusInternalServerError, "error getting followers")
		return
	}

	ctx.Respond(w, http.StatusOK, followers)
}