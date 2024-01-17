package api

import (
	"net/http"
	// "github.com/julienschmidt/httprouter"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes

	// path := "C:/Users/Asus/Documents/UM/Erasmus/Wasa/wasa-project/service/images/"
	path := "/home/wasa/Desktop/wasa-project/service/images"
	// path := "../service/images"

	rt.router.ServeFiles("/images/*filepath", http.Dir(path))

	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	rt.router.POST("/session", rt.postLogin)

	rt.router.GET("/users", rt.getAllUsers)

	rt.router.GET("/stream", rt.AuthenticateMiddleware(rt.GetStream))
	rt.router.PUT("/users/name", rt.AuthenticateMiddleware(rt.putUpdateName))

	rt.router.GET("/users/:userId", rt.GetUserProfile)

	rt.router.POST("/users/:userId/followers", rt.AuthenticateMiddleware(rt.postFollow))
	rt.router.DELETE("/users/:userId/followers", rt.AuthenticateMiddleware(rt.deleteUnfollow))

	rt.router.POST("/users/:userId/bans", rt.AuthenticateMiddleware(rt.postBan))
	rt.router.DELETE("/users/:userId/bans", rt.AuthenticateMiddleware(rt.deleteUnban))

	rt.router.POST("/photos", rt.AuthenticateMiddleware(rt.postImage))
	rt.router.DELETE("/photos", rt.AuthenticateMiddleware(rt.deleteImage))

	rt.router.POST("/comments", rt.AuthenticateMiddleware(rt.postComment))
	rt.router.DELETE("/comments", rt.AuthenticateMiddleware(rt.deleteComment))

	rt.router.POST("/photos/:photoId/likes", rt.AuthenticateMiddleware(rt.postLike))
	rt.router.DELETE("/photos/:photoId/likes", rt.AuthenticateMiddleware(rt.deleteLike))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
