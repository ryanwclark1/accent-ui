/*
 * accent-auth
 *
 * Accent's authentication service
 *
 * API version: 0.1
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	authserver "github.com/ryanwclark/accent-voice/src"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	log.Printf("Server started")

	AdminAPIService := authserver.NewAdminAPIService()
	AdminAPIController := authserver.NewAdminAPIController(AdminAPIService)

	BackendsAPIService := authserver.NewBackendsAPIService()
	BackendsAPIController := authserver.NewBackendsAPIController(BackendsAPIService)

	ConfigAPIService := authserver.NewConfigAPIService()
	ConfigAPIController := authserver.NewConfigAPIController(ConfigAPIService)

	EmailsAPIService := authserver.NewEmailsAPIService()
	EmailsAPIController := authserver.NewEmailsAPIController(EmailsAPIService)

	ExternalAPIService := authserver.NewExternalAPIService()
	ExternalAPIController := authserver.NewExternalAPIController(ExternalAPIService)

	GroupsAPIService := authserver.NewGroupsAPIService()
	GroupsAPIController := authserver.NewGroupsAPIController(GroupsAPIService)

	PoliciesAPIService := authserver.NewPoliciesAPIService()
	PoliciesAPIController := authserver.NewPoliciesAPIController(PoliciesAPIService)

	SessionsAPIService := authserver.NewSessionsAPIService()
	SessionsAPIController := authserver.NewSessionsAPIController(SessionsAPIService)

	StatusAPIService := authserver.NewStatusAPIService()
	StatusAPIController := authserver.NewStatusAPIController(StatusAPIService)

	TenantsAPIService := authserver.NewTenantsAPIService()
	TenantsAPIController := authserver.NewTenantsAPIController(TenantsAPIService)

	TokenAPIService := authserver.NewTokenAPIService()
	TokenAPIController := authserver.NewTokenAPIController(TokenAPIService)

	UsersAPIService := authserver.NewUsersAPIService()
	UsersAPIController := authserver.NewUsersAPIController(UsersAPIService)

	router := authserver.NewRouter(AdminAPIController, BackendsAPIController, ConfigAPIController, EmailsAPIController, ExternalAPIController, GroupsAPIController, PoliciesAPIController, SessionsAPIController, StatusAPIController, TenantsAPIController, TokenAPIController, UsersAPIController)

	port := os.Getenv("AUTH_LISTEN_ADDR")
	slog.Info("application running", "port", port )
	log.Fatal(http.ListenAndServe(port, router))
}

