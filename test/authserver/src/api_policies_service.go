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

// PoliciesAPIService is a service that implements the logic for the PoliciesAPIServicer
// This service should implement the business logic for every endpoint for the PoliciesAPI API.
// Include any external packages or services that will be required by this service.
type PoliciesAPIService struct {
}

// NewPoliciesAPIService creates a default api service
func NewPoliciesAPIService() PoliciesAPIServicer {
	return &PoliciesAPIService{}
}

// AddGroupPolicy - Associate a group to a policy
func (s *PoliciesAPIService) AddGroupPolicy(ctx context.Context, groupUuid string, policyUuid string) (ImplResponse, error) {
	// TODO - update AddGroupPolicy with the required logic for this service method.
	// Add api_policies_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddGroupPolicy method not implemented")
}

// AddPolicyAccess - Associate an access to a policy
func (s *PoliciesAPIService) AddPolicyAccess(ctx context.Context, policyUuid string, access string, accentTenant string) (ImplResponse, error) {
	// TODO - update AddPolicyAccess with the required logic for this service method.
	// Add api_policies_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddPolicyAccess method not implemented")
}

// AddUserPolicy - Associate a policy to a user
func (s *PoliciesAPIService) AddUserPolicy(ctx context.Context, policyUuid string, userUuid string) (ImplResponse, error) {
	// TODO - update AddUserPolicy with the required logic for this service method.
	// Add api_policies_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddUserPolicy method not implemented")
}

// CreatePolicies - Create a new ACL policy
func (s *PoliciesAPIService) CreatePolicies(ctx context.Context, body Policy, accentTenant string) (ImplResponse, error) {
	// TODO - update CreatePolicies with the required logic for this service method.
	// Add api_policies_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, PolicyResult{}) or use other options such as http.Ok ...
	// return Response(200, PolicyResult{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(409, Error{}) or use other options such as http.Ok ...
	// return Response(409, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreatePolicies method not implemented")
}

// DeleteGroupPolicy - Dissociate a policy from a group
func (s *PoliciesAPIService) DeleteGroupPolicy(ctx context.Context, groupUuid string, policyUuid string) (ImplResponse, error) {
	// TODO - update DeleteGroupPolicy with the required logic for this service method.
	// Add api_policies_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteGroupPolicy method not implemented")
}

// DeletePolicy - Delete a policy
func (s *PoliciesAPIService) DeletePolicy(ctx context.Context, policyUuid string, accentTenant string) (ImplResponse, error) {
	// TODO - update DeletePolicy with the required logic for this service method.
	// Add api_policies_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	// return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeletePolicy method not implemented")
}

// DeletePolicyAccess - Dissociate an access from a policy
func (s *PoliciesAPIService) DeletePolicyAccess(ctx context.Context, policyUuid string, access string, accentTenant string) (ImplResponse, error) {
	// TODO - update DeletePolicyAccess with the required logic for this service method.
	// Add api_policies_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeletePolicyAccess method not implemented")
}

// DeleteUserPolicy - Dissociate a policy from a user
func (s *PoliciesAPIService) DeleteUserPolicy(ctx context.Context, policyUuid string, userUuid string) (ImplResponse, error) {
	// TODO - update DeleteUserPolicy with the required logic for this service method.
	// Add api_policies_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteUserPolicy method not implemented")
}

// EditPolicy - Modify an ACL policy
func (s *PoliciesAPIService) EditPolicy(ctx context.Context, policyUuid string, body Policy, accentTenant string) (ImplResponse, error) {
	// TODO - update EditPolicy with the required logic for this service method.
	// Add api_policies_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, PolicyResult{}) or use other options such as http.Ok ...
	// return Response(200, PolicyResult{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(409, Error{}) or use other options such as http.Ok ...
	// return Response(409, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("EditPolicy method not implemented")
}

// GetPolicies - List ACL policies
func (s *PoliciesAPIService) GetPolicies(ctx context.Context, order string, direction string, limit int32, offset int32, search string, accentTenant string, recurse bool) (ImplResponse, error) {
	// TODO - update GetPolicies with the required logic for this service method.
	// Add api_policies_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetPoliciesResult{}) or use other options such as http.Ok ...
	// return Response(200, GetPoliciesResult{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetPolicies method not implemented")
}

// GetPolicy - Retrieves the details of a policy
func (s *PoliciesAPIService) GetPolicy(ctx context.Context, policyUuid string, accentTenant string) (ImplResponse, error) {
	// TODO - update GetPolicy with the required logic for this service method.
	// Add api_policies_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, PolicyResult{}) or use other options such as http.Ok ...
	// return Response(200, PolicyResult{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	// return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetPolicy method not implemented")
}

// GetUserPolicies - Retrieves the list of policies associated to a user
func (s *PoliciesAPIService) GetUserPolicies(ctx context.Context, userUuid string, order string, direction string, limit int32, offset int32, search string) (ImplResponse, error) {
	// TODO - update GetUserPolicies with the required logic for this service method.
	// Add api_policies_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetPoliciesResult{}) or use other options such as http.Ok ...
	// return Response(200, GetPoliciesResult{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	// return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUserPolicies method not implemented")
}
