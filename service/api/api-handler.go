package api

import (
	"net/http"
	// "github.com/julienschmidt/httprouter"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes

	path := "C:/Users/Asus/Documents/UM/Erasmus/Wasa/wasa-project/service/images/"
	rt.router.ServeFiles("/images/*filepath", http.Dir(path))

	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	rt.router.POST("/login", rt.postLogin)

	rt.router.GET("/user/profile/:userId", rt.GetUserProfile)

	rt.router.PUT("/user/name", rt.AuthenticateMiddleware(rt.putUpdateName))
	rt.router.POST("/photo", rt.AuthenticateMiddleware(rt.postImage))
	rt.router.POST("/user/follow/:userId", rt.AuthenticateMiddleware(rt.postFollow))
	rt.router.DELETE("/user/follow/:userId", rt.AuthenticateMiddleware(rt.deleteUnfollow))
	rt.router.POST("/user/ban/:userId", rt.AuthenticateMiddleware(rt.postBan))
	rt.router.DELETE("/user/ban/:userId", rt.AuthenticateMiddleware(rt.deleteUnban))
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
