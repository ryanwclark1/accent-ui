package handler

import (
	"net/http"
	"accent-ui/view/home"
)

func HandleHomeIndex(w http.ResponseWriter, r *http.Request) error {
	return home.Index().Render(r.Context(), w)
	// Code
}