package api

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extracting userId from the context
	userId, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "User ID missing from context", http.StatusUnauthorized)
		return
	}

	// Extracting commentId from the query parameters
	commentId := r.URL.Query().Get("commentId")
	if commentId == "" {
		http.Error(w, "Comment ID is required", http.StatusBadRequest)
		return
	}

	// Call the DeleteComment method from the database layer
	err := rt.db.DeleteComment(userId, commentId)
	if err != nil {
		// Handle the error, e.g., log it or return an appropriate HTTP error response
		rt.baseLogger.Println(err)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(204)
	w.Write([]byte("Comment deleted successfully"))
}
