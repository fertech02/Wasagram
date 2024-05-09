package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	
	// User routes
	rt.router.POST("/session", rt.wrap(rt.doLogin))
	rt.router.PUT("/user/:uid/usrename", rt.wrap(rt.setMyUserName))
	rt.router.GET("/user/:uid/stream", rt.wrap(rt.getMyStream))
	rt.router.GET("/user/:uid/profile", rt.wrap(rt.getUserProfile))

	// Photo routes
	rt.router.POST("/photo", rt.wrap(rt.uploadPhoto))
	rt.router.DELETE("/photo/:pid", rt.wrap(rt.deletePhoto))
	rt.router.GET("/photo/:pid", rt.wrap(rt.getPhoto))

	// Follow routes
	rt.router.PUT("/user/:uid/follow", rt.wrap(rt.followUser))
	rt.router.DELETE("/user/:uid/follow", rt.wrap(rt.unfollowUser))
	// rt.router.GET("/user/:uid/follow", rt.wrap(rt.getFollowedUsers))

	// Comment routes
	rt.router.POST("/photo/:pid/comment", rt.wrap(rt.commentPhoto))
	rt.router.DELETE("/photo/:pid/comment/:cid", rt.wrap(rt.uncommentPhoto))
	rt.router.GET("/photo/:pid/comment", rt.wrap(rt.getComments))

	// Like routes
	rt.router.PUT("/photo/:pid/like", rt.wrap(rt.likePhoto))
	rt.router.DELETE("/photo/:pid/like", rt.wrap(rt.unlikePhoto))
	rt.router.GET("/photo/:pid/like", rt.wrap(rt.getLikes))

	// Ban routes
	rt.router.PUT("/user/:uid/ban", rt.wrap(rt.banUser))
	rt.router.DELETE("/user/:uid/ban", rt.wrap(rt.unbanUser))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
