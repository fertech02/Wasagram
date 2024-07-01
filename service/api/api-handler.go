package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {

	// User routes
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.PUT("/user/:uid/username", rt.wrap(rt.setMyUserName))
	rt.router.GET("/user/:uid/stream", rt.wrap(rt.getMyStream))
	rt.router.GET("/user/:uid/profile", rt.wrap(rt.getUserProfile))

	// Photo routes
	rt.router.POST("/photos", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/photos/:pid", rt.wrap(rt.deletePhoto))
	rt.router.GET("/photos/:pid", rt.wrap(rt.getPhoto))
	rt.router.GET("/photos", rt.wrap(rt.getPhotos))

	// Follow routes
	rt.router.PUT("/user/:uid/follow/:fid", rt.wrap(rt.followUser))
	rt.router.DELETE("/user/:uid/follow/:fid", rt.wrap(rt.unfollowUser))
	rt.router.GET("/user/:uid/follow", rt.wrap(rt.getFollowedUsers))

	// Comment routes
	rt.router.POST("/photo/:pid/comments/:uid", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/photo/:pid/comments/:uid", rt.wrap(rt.uncommentPhoto))
	rt.router.GET("/photo/:pid/comments", rt.wrap(rt.getComments))

	// Like routes
	rt.router.PUT("/photo/:pid/likes/:uid", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/photo/:pid/likes/:uid", rt.wrap(rt.unlikePhoto))
	rt.router.GET("/photo/:pid/likes", rt.wrap(rt.getLikes))

	// Ban routes
	rt.router.PUT("/user/:uid/ban/:bid", rt.wrap(rt.banUser))
	rt.router.DELETE("/user/:uid/ban/:bid", rt.wrap(rt.unbanUser))
	rt.router.GET("/user/:uid/ban", rt.wrap(rt.getBannedUsers))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
