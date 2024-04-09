package handler

import (
	"accent-ui/pkg/accentapi"
	"accent-ui/types"
	"context"
	"net/http"
	"strings"
	"os"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
)

// WithUser is a middleware that adds a user object to the request context
func WithUser(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		// Allows for public files to be served without the middleware
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}
		store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
		session, err := store.Get(r, sessionUserKey)
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}

		accessToken := session.Values[sessionAccessTokenKey]
		if accessToken == nil {
			next.ServeHTTP(w, r)
			return
		}

		resp, err := accentapi.Client.Auth.User(r.Context(), accessToken.(string))
		if err != nil {
			next.ServeHTTP(w, r)
			return
		}
		user := types.AuthenticatedUser{
			ID:          uuid.MustParse(resp.ID),
			Email:       resp.Email,
			LoggedIn:    true,
			// AccessToken: accessToken.(string),
		}


		ctx := context.WithValue(r.Context(), types.UserContextKey, user)
			r = r.WithContext(ctx)
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func WithAuth(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if strings.Contains(r.URL.Path, "/public") {
			next.ServeHTTP(w, r)
			return
		}
		user := getAuthenticatedUser(r)
		if !user.LoggedIn {
			path := r.URL.Path
			http.Redirect(w, r, "/login?to="+path, http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}