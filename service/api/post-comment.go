package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Assuming your request body structure looks like this
type CommentRequest struct {
	Comment string `json:"comment"`
}

func (rt *_router) postComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extracting userId from the context
	userId, ok := r.Context().Value(userIDKey).(string)
	if !ok {
		http.Error(w, "User ID missing from context", http.StatusUnauthorized)
		return
	}

	// Extracting photoId and commentId from the query parameters
	photoId := r.URL.Query().Get("photoId")
	if photoId == "" {
		http.Error(w, "Photo ID is required", http.StatusBadRequest)
		return
	}

	// Decode the JSON body
	var data map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Check if the "comment" key exists in the JSON
	comment, ok := data["comment"].(string)
	if !ok {
		http.Error(w, "Invalid JSON format, missing 'comment' field or it's not a string", http.StatusBadRequest)
		return
	}

	commentObj, err := rt.db.AddComment(userId, photoId, comment)

	if err != nil {
		// Log the error and return an appropriate HTTP error response
		rt.baseLogger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Encode the Image struct to JSON
	err = json.NewEncoder(w).Encode(commentObj)
	if err != nil {
		// Handle the error if JSON encoding fails
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
