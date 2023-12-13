package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) postLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	// Decode the JSON body
	var data map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	// Check if the "name" key exists in the JSON
	name, ok := data["name"].(string)
	if !ok {
		http.Error(w, "Invalid JSON format, missing 'name' field or it's not a string", http.StatusBadRequest)
		return
	}

	// Check if the name is in the database
	identifier, err := rt.db.FindName(name)
	if err != nil {
		// Handle the error, e.g., log it or return an appropriate HTTP error response
		fmt.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if identifier == "" {
		// Name not found in the database, create a new entry
		identifier, err = rt.db.SetName(name)
		if err != nil {
			fmt.Println(err)
			// Handle the error, e.g., log it or return an appropriate HTTP error response
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

	}

	// Respond with the identifier
	response := map[string]string{"identifier": identifier}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	_, _ = w.Write(jsonResponse)
}
