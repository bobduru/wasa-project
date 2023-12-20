package api

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rt *_router) AuthenticateMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		userID := r.Header.Get("Authorization")

		// Assuming CheckUserId is a method of rt that checks the user ID
		if exists := rt.db.CheckUserId(userID); !exists {
			http.Error(w, "User ID not found", http.StatusNotFound)
			return
		}

		// Add the userID to the request context
		ctx := context.WithValue(r.Context(), "userID", userID)
		r = r.WithContext(ctx)

		// Call the next handler if authentication is successful
		next(w, r, ps)
	}
}
