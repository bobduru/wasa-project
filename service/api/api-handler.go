package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	// "github.com/julienschmidt/httprouter"
)

func (rt *_router) addCORSHeader(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if next != nil {
			// Call the next handler if authentication is successful
			next(w, r, ps)
		} else {
			w.WriteHeader(http.StatusOK)
		}
	}
}

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// Register routes

	path := "C:/Users/Asus/Documents/UM/Erasmus/Wasa/wasa-project/service/images/"

	rt.router.GlobalOPTIONS = http.HandlerFunc(rt.optionsCors)

	rt.router.ServeFiles("/images/*filepath", http.Dir(path))

	rt.router.GET("/", rt.getHelloWorld)
	rt.router.GET("/users", rt.getAllUsers)
	rt.router.GET("/context", rt.wrap(rt.getContextReply))

	rt.router.POST("/login", rt.postLogin)

	rt.router.GET("/user/profile/:userId", rt.GetUserProfile)
	rt.router.GET("/user/stream", rt.AuthenticateMiddleware(rt.GetStream))

	rt.router.PUT("/user/name", rt.AuthenticateMiddleware(rt.putUpdateName))

	rt.router.POST("/photo", rt.AuthenticateMiddleware(rt.postImage))
	rt.router.DELETE("/photo", rt.AuthenticateMiddleware(rt.deleteImage))
	rt.router.POST("/photo/comment", rt.AuthenticateMiddleware(rt.postComment))
	rt.router.DELETE("/photo/comment", rt.AuthenticateMiddleware(rt.deleteComment))

	rt.router.POST("/photo/like/:photoId", rt.AuthenticateMiddleware(rt.postLike))
	rt.router.DELETE("/photo/like/:photoId", rt.AuthenticateMiddleware(rt.deleteLike))

	rt.router.POST("/user/follow/:userId", rt.AuthenticateMiddleware(rt.postFollow))
	rt.router.DELETE("/user/follow/:userId", rt.AuthenticateMiddleware(rt.deleteUnfollow))

	rt.router.POST("/user/ban/:userId", rt.AuthenticateMiddleware(rt.postBan))
	rt.router.DELETE("/user/ban/:userId", rt.AuthenticateMiddleware(rt.deleteUnban))

	// Special routes
	rt.router.GET("/liveness", rt.liveness)

	return rt.router
}
