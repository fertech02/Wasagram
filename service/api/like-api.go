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

	// Parse the user id from the request
	uid := ps.ByName("uid")

	// Like the photo
	err := rt.db.Like(pid, uid)
	if err != nil {
		ctx.Logger.WithField("error", err).Error("Failed to like photo")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

}

func (rt *_router) unlikePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	// Parse the photo id from the request
	pid := ps.ByName("pid")

	// Parse the user id from the request
	uid := ps.ByName("uid")

	// Unlike the photo
	err := rt.db.Unlike(pid, uid)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithField("error", err).Error("Failed to unlike photo")
		return
	}

}

func (rt *_router) checkLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	pid := ps.ByName("pid")

	if !CheckValidAuth(r) {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	uid := ps.ByName("uid")

	isLiked, err := rt.db.CheckLike(pid, uid)
	if err != nil {
		ctx.Logger.WithField("error", err).Error("Failed to check like")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData := struct {
		IsLiked bool `json:"is_liked"`
	}{
		IsLiked: isLiked,
	}

	jsonData, err := json.Marshal(responseData)
	if err != nil {
		ctx.Logger.WithField("error", err).Error("Failed to marshal response data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		ctx.Logger.WithField("error", err).Error("Failed to write response data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// Get Likes
func (rt *_router) getLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

	pid := ps.ByName("pid")

	likes, err := rt.db.GetLikeCount(pid)
	if err != nil {
		ctx.Logger.WithField("error", err).Error("Failed to get likes")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	responseData := struct {
		LikeCount int `json:"like_count"`
	}{
		LikeCount: likes,
	}

	jsonData, err := json.Marshal(responseData)
	if err != nil {
		ctx.Logger.WithField("error", err).Error("Failed to marshal response data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonData)
	if err != nil {
		ctx.Logger.WithField("error", err).Error("Failed to write response data")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
