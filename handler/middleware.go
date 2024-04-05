package handler

import (
	"accent-ui/types"
	"net/http"
	"strings"
	"context"
)

// WithUser is a middleware that adds a user object to the request context
func WithUser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Allows for public files to be served without the middleware
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}
		// Create a new user object and add it to the context
		user := types.AuthenticatedUser{
			Email: "admin@accentvoice.io",
			LoggedIn: true,

		}
		ctx := context.WithValue(r.Context(), types.UserContextKey, user)
			r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}