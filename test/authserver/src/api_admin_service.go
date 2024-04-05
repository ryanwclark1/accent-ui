/*
 * accent-auth
 *
 * Accent's authentication service
 *
 * API version: 0.1
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package authserver

import (
	"context"
	"errors"
	"net/http"
)

// AdminAPIService is a service that implements the logic for the AdminAPIServicer
// This service should implement the business logic for every endpoint for the AdminAPI API.
// Include any external packages or services that will be required by this service.
type AdminAPIService struct {
}

// NewAdminAPIService creates a default api service
func NewAdminAPIService() AdminAPIServicer {
	return &AdminAPIService{}
}

// UpdateAllUserEmails - Update email addresses
func (s *AdminAPIService) UpdateAllUserEmails(ctx context.Context, userUuid string, body AdminUserEmailList) (ImplResponse, error) {
	// TODO - update UpdateAllUserEmails with the required logic for this service method.
	// Add api_admin_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateAllUserEmails method not implemented")
}
