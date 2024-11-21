package api

import (
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
