package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteUnfollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	loggedInUserId, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "User ID missing from context", http.StatusUnauthorized)
		return
	}
	// Extracting userId from the URL path parameter.
	userId := ps.ByName("userId")
	followers, err := rt.db.UnfollowUser(loggedInUserId, userId)
	if err != nil {
		// Handle specific errors as needed
		http.Error(w, err.Error(), http.StatusBadRequest) // Or other appropriate status code
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Encode the Image struct to JSON
	err = json.NewEncoder(w).Encode(followers)
	if err != nil {
		// Handle the error if JSON encoding fails
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
