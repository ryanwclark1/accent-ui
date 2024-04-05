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

// UsersAPIService is a service that implements the logic for the UsersAPIServicer
// This service should implement the business logic for every endpoint for the UsersAPI API.
// Include any external packages or services that will be required by this service.
type UsersAPIService struct {
}

// NewUsersAPIService creates a default api service
func NewUsersAPIService() UsersAPIServicer {
	return &UsersAPIService{}
}

// AddGroupUser - Associate a group to a user
func (s *UsersAPIService) AddGroupUser(ctx context.Context, groupUuid string, userUuid string) (ImplResponse, error) {
	// TODO - update AddGroupUser with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddGroupUser method not implemented")
}

// AddUserPolicy - Associate a policy to a user
func (s *UsersAPIService) AddUserPolicy(ctx context.Context, policyUuid string, userUuid string) (ImplResponse, error) {
	// TODO - update AddUserPolicy with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("AddUserPolicy method not implemented")
}

// ChangeUserPassword - Change the user&#39;s password
func (s *UsersAPIService) ChangeUserPassword(ctx context.Context, userUuid string, body PasswordChange) (ImplResponse, error) {
	// TODO - update ChangeUserPassword with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	// return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("ChangeUserPassword method not implemented")
}

// CreateUser - Create a user
func (s *UsersAPIService) CreateUser(ctx context.Context, accentTenant string, body UserCreate) (ImplResponse, error) {
	// TODO - update CreateUser with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, UserPostResponse{}) or use other options such as http.Ok ...
	// return Response(200, UserPostResponse{}), nil

	// TODO: Uncomment the next line to return response Response(400, ApiError{}) or use other options such as http.Ok ...
	// return Response(400, ApiError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("CreateUser method not implemented")
}

// DeleteRefreshTokens - Delete a user&#39;s refresh token
func (s *UsersAPIService) DeleteRefreshTokens(ctx context.Context, userUuidOrMe string, clientId string) (ImplResponse, error) {
	// TODO - update DeleteRefreshTokens with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(401, ApiError{}) or use other options such as http.Ok ...
	// return Response(401, ApiError{}), nil

	// TODO: Uncomment the next line to return response Response(404, ApiError{}) or use other options such as http.Ok ...
	// return Response(404, ApiError{}), nil

	// TODO: Uncomment the next line to return response Response(500, ApiError{}) or use other options such as http.Ok ...
	// return Response(500, ApiError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteRefreshTokens method not implemented")
}

// DeleteUser - Delete a user
func (s *UsersAPIService) DeleteUser(ctx context.Context, userUuid string) (ImplResponse, error) {
	// TODO - update DeleteUser with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	// return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteUser method not implemented")
}

// DeleteUserPolicy - Dissociate a policy from a user
func (s *UsersAPIService) DeleteUserPolicy(ctx context.Context, policyUuid string, userUuid string) (ImplResponse, error) {
	// TODO - update DeleteUserPolicy with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("DeleteUserPolicy method not implemented")
}

// GetNewEmailConfirmation - Ask a new confirmation email
func (s *UsersAPIService) GetNewEmailConfirmation(ctx context.Context, userUuid string, emailUuid string) (ImplResponse, error) {
	// TODO - update GetNewEmailConfirmation with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(409, {}) or use other options such as http.Ok ...
	// return Response(409, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetNewEmailConfirmation method not implemented")
}

// GetUser - Retrieves the details of a user
func (s *UsersAPIService) GetUser(ctx context.Context, userUuid string) (ImplResponse, error) {
	// TODO - update GetUser with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, UserResult{}) or use other options such as http.Ok ...
	// return Response(200, UserResult{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	// return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUser method not implemented")
}

// GetUserExternalAuth - Retrieves the list of the users external auth data
func (s *UsersAPIService) GetUserExternalAuth(ctx context.Context, userUuid string, order string, direction string, limit int32, offset int32, search string) (ImplResponse, error) {
	// TODO - update GetUserExternalAuth with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, ExternalAuthList{}) or use other options such as http.Ok ...
	// return Response(200, ExternalAuthList{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUserExternalAuth method not implemented")
}

// GetUserGroups - Retrieves the list of groups associated to a user
func (s *UsersAPIService) GetUserGroups(ctx context.Context, userUuid string, order string, direction string, limit int32, offset int32, search string) (ImplResponse, error) {
	// TODO - update GetUserGroups with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetGroupsResult{}) or use other options such as http.Ok ...
	// return Response(200, GetGroupsResult{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	// return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUserGroups method not implemented")
}

// GetUserPolicies - Retrieves the list of policies associated to a user
func (s *UsersAPIService) GetUserPolicies(ctx context.Context, userUuid string, order string, direction string, limit int32, offset int32, search string) (ImplResponse, error) {
	// TODO - update GetUserPolicies with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

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

// GetUserSessions - Retrieves the list of sessions associated to a user
func (s *UsersAPIService) GetUserSessions(ctx context.Context, userUuid string, accentTenant string, limit int32, offset int32) (ImplResponse, error) {
	// TODO - update GetUserSessions with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, GetSessionsResult{}) or use other options such as http.Ok ...
	// return Response(200, GetSessionsResult{}), nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	// TODO: Uncomment the next line to return response Response(500, Error{}) or use other options such as http.Ok ...
	// return Response(500, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUserSessions method not implemented")
}

// GetUserTokens - Retrieve a user&#39;s refresh token list
func (s *UsersAPIService) GetUserTokens(ctx context.Context, userUuidOrMe string, accentTenant string, order string, direction string, limit int32, offset int32, search string) (ImplResponse, error) {
	// TODO - update GetUserTokens with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, RefreshTokenList{}) or use other options such as http.Ok ...
	// return Response(200, RefreshTokenList{}), nil

	// TODO: Uncomment the next line to return response Response(400, ApiError{}) or use other options such as http.Ok ...
	// return Response(400, ApiError{}), nil

	// TODO: Uncomment the next line to return response Response(401, ApiError{}) or use other options such as http.Ok ...
	// return Response(401, ApiError{}), nil

	// TODO: Uncomment the next line to return response Response(404, ApiError{}) or use other options such as http.Ok ...
	// return Response(404, ApiError{}), nil

	// TODO: Uncomment the next line to return response Response(500, ApiError{}) or use other options such as http.Ok ...
	// return Response(500, ApiError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUserTokens method not implemented")
}

// GetUsers - Retrieves the list of users
func (s *UsersAPIService) GetUsers(ctx context.Context, order string, direction string, limit int32, offset int32, search string, accentTenant string, recurse bool, hasPolicySlug string, hasPolicyUuid string, policySlug string, policyUuid string) (ImplResponse, error) {
	// TODO - update GetUsers with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, UserList{}) or use other options such as http.Ok ...
	// return Response(200, UserList{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("GetUsers method not implemented")
}

// RegisterUser - Create a user
func (s *UsersAPIService) RegisterUser(ctx context.Context, body UserRegister) (ImplResponse, error) {
	// TODO - update RegisterUser with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, UserPostResponse{}) or use other options such as http.Ok ...
	// return Response(200, UserPostResponse{}), nil

	// TODO: Uncomment the next line to return response Response(400, ApiError{}) or use other options such as http.Ok ...
	// return Response(400, ApiError{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("RegisterUser method not implemented")
}

// RemoveGroupUser - Dissociate a user from a group
func (s *UsersAPIService) RemoveGroupUser(ctx context.Context, groupUuid string, userUuid string) (ImplResponse, error) {
	// TODO - update RemoveGroupUser with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("RemoveGroupUser method not implemented")
}

// ResetPassword - Reset the user password
func (s *UsersAPIService) ResetPassword(ctx context.Context, username string, email string, login string) (ImplResponse, error) {
	// TODO - update ResetPassword with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ResetPassword method not implemented")
}

// ResetPasswordChange - Set the user password
func (s *UsersAPIService) ResetPasswordChange(ctx context.Context, userUuid string, body PostPasswordReset) (ImplResponse, error) {
	// TODO - update ResetPasswordChange with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	return Response(http.StatusNotImplemented, nil), errors.New("ResetPasswordChange method not implemented")
}

// UpdateAllUserEmails - Update email addresses
func (s *UsersAPIService) UpdateAllUserEmails(ctx context.Context, userUuid string, body AdminUserEmailList) (ImplResponse, error) {
	// TODO - update UpdateAllUserEmails with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateAllUserEmails method not implemented")
}

// UpdateUser - Update an existing user
func (s *UsersAPIService) UpdateUser(ctx context.Context, userUuid string, body UserEdit) (ImplResponse, error) {
	// TODO - update UpdateUser with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, UserPostResponse{}) or use other options such as http.Ok ...
	// return Response(200, UserPostResponse{}), nil

	// TODO: Uncomment the next line to return response Response(400, ApiError{}) or use other options such as http.Ok ...
	// return Response(400, ApiError{}), nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateUser method not implemented")
}

// UpdateUserEmails - Update email addresses
func (s *UsersAPIService) UpdateUserEmails(ctx context.Context, userUuid string, body UserEmailList) (ImplResponse, error) {
	// TODO - update UpdateUserEmails with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(200, {}) or use other options such as http.Ok ...
	// return Response(200, nil),nil

	// TODO: Uncomment the next line to return response Response(404, Error{}) or use other options such as http.Ok ...
	// return Response(404, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UpdateUserEmails method not implemented")
}

// UserDeleteSession - Delete a session
func (s *UsersAPIService) UserDeleteSession(ctx context.Context, userUuid string, sessionUuid string) (ImplResponse, error) {
	// TODO - update UserDeleteSession with the required logic for this service method.
	// Add api_users_service.go to the .openapi-generator-ignore to avoid overwriting this service implementation when updating open api generation.

	// TODO: Uncomment the next line to return response Response(204, {}) or use other options such as http.Ok ...
	// return Response(204, nil),nil

	// TODO: Uncomment the next line to return response Response(401, Error{}) or use other options such as http.Ok ...
	// return Response(401, Error{}), nil

	return Response(http.StatusNotImplemented, nil), errors.New("UserDeleteSession method not implemented")
}