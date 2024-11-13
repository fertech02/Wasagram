package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// User routes
	rt.router.POST("/session/", rt.wrap(rt.doLogin))
	rt.router.PUT("/users/:uid/username", rt.wrap(rt.setMyUserName))
	rt.router.GET("/users/:uid/stream", rt.wrap(rt.getMyStream))
	rt.router.GET("/users/:uid/profile", rt.wrap(rt.getUserProfile))
	rt.router.GET("/users", rt.wrap(rt.searchUser))

	// Photo routes
	rt.router.POST("/photos", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/photos/:pid", rt.wrap(rt.deletePhoto))
	rt.router.GET("/photos/:pid", rt.wrap(rt.getPhoto))
	rt.router.GET("/photos", rt.wrap(rt.getPhotos))

	// Follow routes
	rt.router.PUT("/users/:uid/follow/:fid", rt.wrap(rt.followUser))
	rt.router.DELETE("/users/:uid/follow/:fid", rt.wrap(rt.unfollowUser))
	rt.router.GET("/users/:uid/follow", rt.wrap(rt.getFollowedUsers))

	// Comment routes
	rt.router.POST("/photos/:pid/comments/:uid", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/photos/:pid/comments/:uid", rt.wrap(rt.uncommentPhoto))
	rt.router.GET("/photos/:pid/comments", rt.wrap(rt.getComments))

	// Like routes
	rt.router.PUT("/photos/:pid/likes/:uid", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/photos/:pid/likes/:uid", rt.wrap(rt.unlikePhoto))
	rt.router.GET("/photos/:pid/likes/:uid", rt.wrap(rt.checkLike))
	rt.router.GET("/photos/:pid/likes", rt.wrap(rt.getLikes))

	// Ban routes
	rt.router.PUT("/users/:uid/ban/:bid", rt.wrap(rt.banUser))
	rt.router.DELETE("/users/:uid/ban/:bid", rt.wrap(rt.unbanUser))
	rt.router.GET("/users/:uid/ban", rt.wrap(rt.getBannedUsers))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
