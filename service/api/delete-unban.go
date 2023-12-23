package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteUnban(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Retrieve the logged-in user's ID from the context
	loggedInUserId := r.Context().Value("userID").(string)

	// Extracting userId from the URL path parameter.
	userIdToUnban := ps.ByName("userId")

	if userIdToUnban == "" {
		http.Error(w, "User ID to unban is required", http.StatusBadRequest)
		return
	}

	if err := rt.db.UnbanUser(loggedInUserId, userIdToUnban); err != nil {
		// Handle specific errors as needed
		http.Error(w, err.Error(), http.StatusBadRequest) // Or other appropriate status code
		return
	}

	// If successful
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User unbanned successfully"))
}
