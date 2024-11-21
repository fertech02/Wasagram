package api

import (
	"net/http"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	followeeId := ps.ByName("uid")
	if followeeId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	followerId := ps.ByName("fid")
	if followerId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Follow the user
	err := rt.db.Follow(followeeId, followerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Error following user")
		return
	}
}

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	followeeId := ps.ByName("uid")
	if followeeId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	followerId := ps.ByName("fid")
	if followerId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Unfollow the user
	err := rt.db.Unfollow(followeeId, followerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
