package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) GetUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extracting userId from the URL path parameter.
	userIdStr := ps.ByName("userId")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userProfile, err := rt.db.GetUserProfile(userId)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	userID := r.Header.Get("Authorization")

	if exists := rt.db.CheckUserId(userID); exists {
		isFollowing, isBanned, err := rt.db.GetUserStatus(userID, userIdStr)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		userProfile.IsFollowing = isFollowing
		userProfile.IsBanned = isBanned
	}

	// Setting the content type for the response.
	w.Header().Set("Content-Type", "application/json")

	// Encoding the user profile into JSON and writing it to the response.
	err = json.NewEncoder(w).Encode(userProfile)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}
}
