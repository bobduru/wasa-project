package api

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) postBan(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	loggedInUserId, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "User ID missing from context", http.StatusUnauthorized)
		return
	}
	// Extracting userId from the URL path parameter.
	userIdToBan := ps.ByName("userId")
	if userIdToBan == "" {
		http.Error(w, "User ID to ban is required", http.StatusBadRequest)
		return
	}

	if err := rt.db.BanUser(loggedInUserId, userIdToBan); err != nil {
		// Handle specific errors as needed
		http.Error(w, err.Error(), http.StatusBadRequest) // Or other appropriate status code
		return
	}

	// If successful
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("User banned successfully"))
}
