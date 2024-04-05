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

type RefreshToken struct {

	// The `client_id` that was used to create this refresh token
	ClientId string `json:"client_id,omitempty"`

	// The time at which this token was created
	CreatedAt string `json:"created_at,omitempty"`

	// Indicate if that refresh token was created with a mobile session type
	Mobile bool `json:"mobile,omitempty"`

	// The tenant UUID of the user which created this refresh token
	TenantUuid string `json:"tenant_uuid,omitempty"`

	// The UUID of the user which created this refresh token
	UserUuid string `json:"user_uuid,omitempty"`
}

// AssertRefreshTokenRequired checks if the required fields are not zero-ed
func AssertRefreshTokenRequired(obj RefreshToken) error {
	return nil
}

// AssertRefreshTokenConstraints checks if the values respects the defined constraints
func AssertRefreshTokenConstraints(obj RefreshToken) error {
	return nil
}
