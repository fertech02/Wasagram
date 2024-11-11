package api

import (
	"encoding/json"
	"net/http"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the photo id from the request
	pid := ps.ByName("pid")
	if pid == "" {
		ctx.Logger.Error("Failed to parse photo id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse the user id from the request
	uid := ps.ByName("uid")
	if uid == "" {
		ctx.Logger.Error("Failed to parse user id")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Like the photo
	err := rt.db.Like(uid, pid)
	if err != nil {
		ctx.Logger.WithField("error", err).Error("Failed to like photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the photo id from the request
	pid := ps.ByName("pid")
	if pid != "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse the user id from the request
	uid := ps.ByName("uid")
	if uid != "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Unlike the photo
	err := rt.db.Unlike(pid, uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

/**
func (rt *_router) checkLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the photo id from the request
	pid := ps.ByName("pid")
	if pid != "" {
		w.WriteHeader(http.StatusBadRequest);
		return
	}

	// Parse the user id from the request
	uid := ps.ByName("uid")
	if uid != "" {
		w.WriteHeader(http.StatusBadRequest);
		return
	}

	// Check if user authorized
	unauthorized, err := CheckAuthorizedId(r, uid)
	if err != nil {
		ctx.Logger.WithField("error", err).Error("Failed to check authorization")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if unauthorized {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Check if user liked the photo
	_, err = rt.db.CheckLike(pid, uid)
	if err != nil {
		ctx.Logger.WithField("error", err).Error("Failed to check if user liked the photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
*/

// Get Likes
func (rt *_router) getLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the photo id from the request
	pid := ps.ByName("pid")
	if pid != "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Parse the user id from the request
	uid := ps.ByName("uid")
	if uid != "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the likes
	likes, err := rt.db.GetLikes(pid)
	if err != nil {
		ctx.Logger.WithField("error", err).Error("Failed to get likes")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the likes
	err = json.NewEncoder(w).Encode(likes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}
