package accentapi

import (
	"os"
	"accent-ui/pkg/restclient"
)

const BaseAuthURL = "https://api.accent.io/api/auth/0.1"

var Client *restclient.Client

//
func Init() error {
	accentHost := os.Getenv("ACCENT_HOST")
	accentToken := os.Getenv("ACCENT_TOKEN")
	Client = restclient.CreateClient(accentHost, accentToken)
	return nil
}