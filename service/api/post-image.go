package api

import (
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) postImage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	userID, ok := r.Context().Value("userID").(string)
	if !ok {
		http.Error(w, "User ID missing from context", http.StatusUnauthorized)
		return
	}

	path := "C:/Users/Asus/Documents/UM/Erasmus/Wasa/wasa-project/service/images/"

	// Maximum upload of 10 MB files
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Error parsing the image", http.StatusInternalServerError)
		return
	}

	// Retrieve the file from form data
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error Retrieving the File", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	image, err := rt.db.UploadImage(userID)

	if err != nil {
		// Handle the error, e.g., log it or return an appropriate HTTP error response
		rt.baseLogger.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create a file in the static directory
	dst, err := os.Create(path + image.FileName)
	if err != nil {
		http.Error(w, "Error Creating the File", http.StatusInternalServerError)
		rt.baseLogger.Println(err)
		return
	}
	defer dst.Close()

	// Copy the uploaded file to the destination file
	_, err = io.Copy(dst, file)
	if err != nil {
		http.Error(w, "Error Saving the File", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Encode the Image struct to JSON
	err = json.NewEncoder(w).Encode(image)
	if err != nil {
		// Handle the error if JSON encoding fails
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
