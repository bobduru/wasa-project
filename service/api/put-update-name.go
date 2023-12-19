package api

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// putUpdateName handles the HTTP PUT request for updating the name.
func (rt *_router) putUpdateName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extract user ID from the Authorization header
	userID := r.Header.Get("Authorization")

	// Decode the JSON body to get the new name
	var requestData map[string]string
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&requestData); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Get the new name from the JSON body
	newName, ok := requestData["name"]
	if !ok {
		http.Error(w, "Missing 'name' field in the JSON body", http.StatusBadRequest)
		return
	}

	// Call the database function to update the name
	updatedName, err := rt.db.UpdateName(userID, newName)
	if err != nil {
		http.Error(w, "Error updating name in the database", http.StatusInternalServerError)
		return
	}

	// Respond with the updated name
	response := map[string]string{"updated_name": updatedName}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(jsonResponse)
}
