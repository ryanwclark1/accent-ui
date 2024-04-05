package handler

import (
	"accent-ui/view/auth"
	"fmt"
	"net/http"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return auth.Loginin().Render(r.Context(), w)
}

func HandleLoginCreate(w http.ResponseWriter, r *http.Request) error {

	Email := r.FormValue("email")
	Password := r.FormValue("password")
	fmt.Println(Email, Password)
	return auth.Loginin().Render(r.Context(), w)
	// fmt.Println(Email, Password)
	// return nil
	// return auth.Loginin().Render(r.Context(), w)
}