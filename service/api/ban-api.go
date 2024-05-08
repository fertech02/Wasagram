package api

import (
	"net/http"

	"github.com/fertech02/Wasa-directory/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	bannerId, err := ps.ByName("bannerId")
	if err != nil {
		ctx.Error(w, http.StatusBadRequest, "missing bannerId")
		return
	}

	bannedId, err := ps.ByName("bannedId")
	if err != nil {
		ctx.Error(w, http.StatusBadRequest, "missing bannedId")
		return
	}

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

	// Ban the user
	err = rt.db.BanUser(bannerId, bannedId)
	if err != nil {
		ctx.Error(w, http.StatusInternalServerError, "error banning user")
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the bannerId and bannedId
	bannerId, err := ps.ByName("bannerId")
	if err != nil {
		ctx.Error(w, http.StatusBadRequest, "missing bannerId")
		return
	}

	bannedId, err := ps.ByName("bannedId")
	if err != nil {
		ctx.Error(w, http.StatusBadRequest, "missing bannedId")
		return
	}

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

	// Unban the user
	err = rt.db.UnbanUser(bannerId, bannedId)
	if err != nil {
		ctx.Error(w, http.StatusInternalServerError, "error unbanning user")
		return
	}

	w.WriteHeader(http.StatusOK)
}
