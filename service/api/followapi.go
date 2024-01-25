package api;

import (
	"net/http"
	"service/database"
)

// addFollow handles the POST /user/{uid}/follow API endpoint.
func (rt *_router) addFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

// deleteFollow handles the DELETE /user/{uid}/follow API endpoint.
func (rt *_router) deleteFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

// getFollows handles the GET /user/{uid}/follows API endpoint.
func (rt *_router) getFollows(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}
