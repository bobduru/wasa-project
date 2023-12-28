package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) postLike(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extracting userId from the context
	userId, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "User ID missing from context", http.StatusUnauthorized)
		return
	}

	// Extracting photoId from the URL path parameters
	photoId := ps.ByName("photoId")
	if photoId == "" {
		http.Error(w, "Photo ID is required", http.StatusBadRequest)
		return
	}

	// Call the AddLike method from the database layer
	likes, err := rt.db.AddLike(userId, photoId)
	if err != nil {
		// Handle the error, e.g., log it or return an appropriate HTTP error response
		rt.baseLogger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(likes)
}
