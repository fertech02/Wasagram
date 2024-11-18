package api

import (
	"encoding/json"
	"net/http"

	"github.com/fertech02/Wasa-repository/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	bannerId := ps.ByName("uid")
	if bannerId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bannedId := ps.ByName("bid")
	if bannedId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Ban the user
	err := rt.db.Ban(bannerId, bannedId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (rt *_router) unbanUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the bannerId and bannedId
	bannerId := ps.ByName("uid")
	if bannerId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	bannedId := ps.ByName("bid")
	if bannedId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Unban the user
	err := rt.db.Unban(bannerId, bannedId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// Get Banned users
func (rt *_router) getBannedUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Get the bannerId
	bannerId := ps.ByName("uid")
	if bannerId == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the banned users
	bannedUsers, err := rt.db.GetBannedUsers(bannerId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return the banned users
	err = json.NewEncoder(w).Encode(bannedUsers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
