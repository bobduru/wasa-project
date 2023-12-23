package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteUnfollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Retrieve the logged-in user's ID from the context
	loggedInUserId := r.Context().Value("userID").(string)

	// Extracting userId from the URL path parameter.
	userId := ps.ByName("userId")

	if err := rt.db.UnfollowUser(loggedInUserId, userId); err != nil {
		// Handle specific errors as needed
		http.Error(w, err.Error(), http.StatusBadRequest) // Or other appropriate status code
		return
	}

	// If successful
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User unfollowed successfully"))
}
