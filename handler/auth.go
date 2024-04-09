package handler

import (
	"accent-ui/view/auth"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
)


const (
	sessionUserKey        = "user"
	sessionAccessTokenKey = "accessToken"
)

func HandleLoginIndex(w http.ResponseWriter, r *http.Request) error {
	return auth.Loginin().Render(r.Context(), w)
}

func HandleLoginCreate(w http.ResponseWriter, r *http.Request) error {

	Username := r.FormValue("username")
	Password := r.FormValue("password")
	// credentials := restclient.UserCredentials{

 // Pass the *http.Client instance to NewClient

	fmt.Println(Username, Password)
	return auth.Loginin().Render(r.Context(), w)
	// fmt.Println(Email, Password)
	// return nil
	// return auth.Loginin().Render(r.Context(), w)
}

func HandleAuthCallback(w http.ResponseWriter, r *http.Request) error {
	accessToken := r.URL.Query().Get("access_token")
	if len(accessToken) == 0 {
		return render(r, w, auth.CallbackScript())
	}
	if err := setAuthSession(w, r, accessToken); err != nil {
		return err
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
	return nil
}

func HandleLogoutCreate(w http.ResponseWriter, r *http.Request) error {
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	session, _ := store.Get(r, sessionUserKey)
	session.Values[sessionAccessTokenKey] = ""
	session.Save(r, w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
	return nil
}

func setAuthSession(w http.ResponseWriter, r *http.Request, accessToken string) error {
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	session, _ := store.Get(r, sessionUserKey)
	session.Values[sessionAccessTokenKey] = accessToken
	return session.Save(r, w)
}