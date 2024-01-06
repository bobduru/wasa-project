package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) deleteImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// Extracting userId from the context
	userId, ok := r.Context().Value(userIDKey).(string)
	if !ok {
		http.Error(w, "User ID missing from context", http.StatusUnauthorized)
		return
	}

	path := "C:/Users/Asus/Documents/UM/Erasmus/Wasa/wasa-project/service/images/"

	// Extract the photoId from the query parameters
	photoId := r.URL.Query().Get("photoId")
	if photoId == "" {
		http.Error(w, "Photo ID is required", http.StatusBadRequest)
		return
	}

	// Call the DeleteImage method
	fileName, err := rt.db.DeleteImage(userId, photoId)
	if err != nil {
		// Handle the error, e.g., log it or return an appropriate HTTP error response
		rt.baseLogger.Println(err)
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	// Construct the full file path
	fullPath := filepath.Join(path, fileName)

	// Delete the file from the filesystem
	err = os.Remove(fullPath)
	if err != nil {
		rt.baseLogger.Println(err)
		http.Error(w, "Error Deleting the File from the Filesystem", http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Image deleted successfully"))

}
