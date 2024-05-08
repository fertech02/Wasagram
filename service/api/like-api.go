package api

import (
	"net/http"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the photo id from the request
	pid, err := ps.ByName("pid")
	if err != nil {
		ctx.RespondWithError(w, http.StatusBadRequest, "missing photo id")
		return
	}

	// Parse the user id from the request
	uid, err := ps.ByName("uid")
	if err != nil {
		ctx.RespondWithError(w, http.StatusBadRequest, "missing user id")
		return
	}

	// Check if user authorized
	if !ctx.IsAuthorized(uid) {
		ctx.RespondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	// Like the photo
	err = rt.db.LikePhoto(uid, pid)
	if err != nil {
		ctx.RespondWithError(w, http.StatusInternalServerError, "error liking photo")
		return
	}

	ctx.RespondWithJSON(w, http.StatusOK, nil)
}

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the photo id from the request
	pid, err := ps.ByName("pid")
	if err != nil {
		ctx.RespondWithError(w, http.StatusBadRequest, "missing photo id")
		return
	}

	// Parse the user id from the request
	uid, err := ps.ByName("uid")
	if err != nil {
		ctx.RespondWithError(w, http.StatusBadRequest, "missing user id")
		return
	}

	// Check if user authorized
	if !ctx.IsAuthorized(uid) {
		ctx.RespondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	// Unlike the photo
	err = rt.db.UnlikePhoto(pid, uid)
	if err != nil {
		ctx.RespondWithError(w, http.StatusInternalServerError, "error unliking photo")
		return
	}

	ctx.RespondWithJSON(w, http.StatusOK, nil)
}

func (rt *_router) checkLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the photo id from the request
	pid, err := ps.ByName("pid")
	if err != nil {
		ctx.RespondWithError(w, http.StatusBadRequest, "missing photo id")
		return
	}

	// Parse the user id from the request
	uid, err := ps.ByName("uid")
	if err != nil {
		ctx.RespondWithError(w, http.StatusBadRequest, "missing user id")
		return
	}

	// Check if user authorized
	if !ctx.IsAuthorized(uid) {
		ctx.RespondWithError(w, http.StatusUnauthorized, "unauthorized")
		return
	}

	// Check if user liked the photo
	liked, err := rt.db.CheckLike(pid, uid)
	if err != nil {
		ctx.RespondWithError(w, http.StatusInternalServerError, "error checking like")
		return
	}

	ctx.RespondWithJSON(w, http.StatusOK, liked)
}
