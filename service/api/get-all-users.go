package api

import (
	"encoding/json"
	"net/http"


	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	users, err := rt.db.GetAllUsers()
	if err != nil {
		http.Error(w, "Failed to retrieve users: "+err.Error(), http.StatusInternalServerError)
		return
	}
	// Set content type as JSON for the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Encode the users slice into JSON and send it in the response
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, "Failed to encode users: "+err.Error(), http.StatusInternalServerError)
	}
}
