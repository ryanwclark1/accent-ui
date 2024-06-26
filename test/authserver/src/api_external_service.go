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

// ExternalAPIService is a service that implements the logic for the ExternalAPIServicer
// This service should implement the business logic for every endpoint for the ExternalAPI API.
// Include any external packages or services that will be required by this service.
type ExternalAPIService struct {
}

// NewExternalAPIService creates a default api service
func NewExternalAPIService() ExternalAPIServicer {
	return &ExternalAPIService{}
}

// DeleteExternalAuthConfig - Delete the client id and client secret
func (s *ExternalAPIService) DeleteExternalAuthConfig(ctx context.Context, authType string, accentTenant string) (ImplResponse, error) {
	// TODO - update DeleteExternalAuthConfig with the required logic for this service method.
	// Add api_external_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteExternalAuthConfig method not implemented")
}

// GetExternalAuthConfig - Retrieve the client id and client secret
func (s *ExternalAPIService) GetExternalAuthConfig(ctx context.Context, authType string, accentTenant string) (ImplResponse, error) {
	// TODO - update GetExternalAuthConfig with the required logic for this service method.
	// Add api_external_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ExternalConfig{}) or use other options such as http.Ok ...
	// return Response(200, ExternalConfig{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetExternalAuthConfig method not implemented")
}

// GetExternalAuthUsers - Retrieves the list of connected users to this external source
func (s *ExternalAPIService) GetExternalAuthUsers(ctx context.Context, authType string, accentTenant string, recurse bool, limit int32, offset int32) (ImplResponse, error) {
	// TODO - update GetExternalAuthUsers with the required logic for this service method.
	// Add api_external_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ExternalAuthUserList{}) or use other options such as http.Ok ...
	// return Response(200, ExternalAuthUserList{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetExternalAuthUsers method not implemented")
}

// GetUserExternalAuth - Retrieves the list of the users external auth data
func (s *ExternalAPIService) GetUserExternalAuth(ctx context.Context, userUuid string, order string, direction string, limit int32, offset int32, search string) (ImplResponse, error) {
	// TODO - update GetUserExternalAuth with the required logic for this service method.
	// Add api_external_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ExternalAuthList{}) or use other options such as http.Ok ...
	// return Response(200, ExternalAuthList{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUserExternalAuth method not implemented")
}

// PostExternalAuthConfig - Add configuration for the given auth_type
func (s *ExternalAPIService) PostExternalAuthConfig(ctx context.Context, authType string, config ExternalConfig, accentTenant string) (ImplResponse, error) {
	// TODO - update PostExternalAuthConfig with the required logic for this service method.
	// Add api_external_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, {}) or use other options such as http.Ok ...
	// return Response(201, nil),nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(409, Error{}) or use other options such as http.Ok ...
	// return Response(409, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("PostExternalAuthConfig method not implemented")
}

// UpdateExternalAuthConfig - Update configuration for the given auth_type
func (s *ExternalAPIService) UpdateExternalAuthConfig(ctx context.Context, authType string, config ExternalConfig, accentTenant string) (ImplResponse, error) {
	// TODO - update UpdateExternalAuthConfig with the required logic for this service method.
	// Add api_external_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(201, {}) or use other options such as http.Ok ...
	// return Response(201, nil),nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateExternalAuthConfig method not implemented")
}
