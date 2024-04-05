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

type ExternalConfig struct {

	// Client ID for the given authentication type. Required only for `google` and `microsoft` authentication types.
	ClientId string `json:"client_id,omitempty"`

	// Client secret for the given authentication type. Required only for `google` and `microsoft` authentication types.
	ClientSecret string `json:"client_secret,omitempty"`

	// The API key to use for Firebase Cloud Messaging
	FcmApiKey string `json:"fcm_api_key,omitempty"`

	// Public certificate to use for Apple Push Notification Service
	IosApnCertificate string `json:"ios_apn_certificate,omitempty"`

	// Private key to use for Apple Push Notification Service
	IosApnPrivate bool `json:"ios_apn_private,omitempty"`

	// Whether to use sandbox for Apple Push Notification Service
	UseSandbox bool `json:"use_sandbox,omitempty"`
}

// AssertExternalConfigRequired checks if the required fields are not zero-ed
func AssertExternalConfigRequired(obj ExternalConfig) error {
	return nil
}

// AssertExternalConfigConstraints checks if the values respects the defined constraints
func AssertExternalConfigConstraints(obj ExternalConfig) error {
	return nil
}
