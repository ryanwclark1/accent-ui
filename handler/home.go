package handler

import (
	"fmt"
	"net/http"
	"accent-ui/view/home"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	user := getAuthenticatedUser(r)

	fmt.Printf("%+v\n", user.Account)
	return home.Index().Render(r.Context(), w)
	// Code
}