package api

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type contextKey string

const userIDKey contextKey = "userID"

func (rt *_router) AuthenticateMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		userID := r.Header.Get("Authorization")

		if exists := rt.db.CheckUserId(userID); !exists {
			http.Error(w, "User ID not found", http.StatusNotFound)
			return
		}

		// Add the userID to the request context
		ctx := context.WithValue(r.Context(), userIDKey, userID)
		r = r.WithContext(ctx)

		// Call the next handler if authentication is successful
		next(w, r, ps)
	}
}
