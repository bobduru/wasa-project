package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) postFollow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	loggedInUserId := r.Context().Value("userID").(string)

	// Decode the JSON body
	var requestData map[string]string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Call the business logic function
	if err := rt.db.FollowUser(loggedInUserId, requestData["userIdToFollow"]); err != nil {
		// Handle specific errors as needed
		http.Error(w, err.Error(), http.StatusBadRequest) // Or other appropriate status code
		return
	}

	// If successful
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User followed successfully"))
}
