package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteUnban(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Retrieve the logged-in user's ID from the context
	loggedInUserId := r.Context().Value("userID").(string)

	// Decode the JSON body
	var requestData map[string]string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	userIdToUnban := requestData["userIdToUnban"]
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
