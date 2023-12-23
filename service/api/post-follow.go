package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) postFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	loggedInUserId := r.Context().Value("userID").(string)

	// Extracting userId from the URL path parameter.
	userId := ps.ByName("userId")

	if err := rt.db.FollowUser(loggedInUserId, userId); err != nil {
		// Handle specific errors as needed
		http.Error(w, err.Error(), http.StatusBadRequest) // Or other appropriate status code
		return
	}

	// If successful
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User followed successfully"))
}
