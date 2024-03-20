package api;

import (
	"net/http"
	"service/database"
)


func (rt *_router) DeletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userid := ps.ByName("id")
	photoid := ps.ByName("photoid")
	user, err := database.userdao.GetUser(userid)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	photo := &database.Photo{pid: photoid}
	err = user.DeletePhoto(photo)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}