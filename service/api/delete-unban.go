package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteUnban(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	loggedInUserId, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "User ID missing from context", http.StatusUnauthorized)
		return
	}
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
