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
	rt.router.PUT("/update-name", rt.putUpdateName)
	rt.router.POST("/image", rt.postImage)
	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
