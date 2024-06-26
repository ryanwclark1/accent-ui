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

type Token struct {
	Data TokenData `json:"data,omitempty"`
}

// AssertTokenRequired checks if the required fields are not zero-ed
func AssertTokenRequired(obj Token) error {
	if err := AssertTokenDataRequired(obj.Data); err != nil {
		return err
	}
	return nil
}

// AssertTokenConstraints checks if the values respects the defined constraints
func AssertTokenConstraints(obj Token) error {
	return nil
}

type TokenData struct {

	// The list of allowed accesses for this token
	Acl []string `json:"acl,omitempty"`

	// The unique identifier retrieved from the backend
	AuthId string `json:"auth_id,omitempty"`

	ExpiresAt string `json:"expires_at,omitempty"`

	IssuedAt string `json:"issued_at,omitempty"`

	// Information owned by accent-auth about this user
	Metadata      struct {
        UUID       string `json:"uuid,omitempty"`
        TenantUUID string `json:"tenant_uuid,omitempty"`
        AuthID     string `json:"auth_id,omitempty"`
        PBXUserUUID string `json:"pbx_user_uuid,omitempty"`
        AccentUUID string `json:"accent_uuid,omitempty"`
        Purpose    string `json:"purpose,omitempty"`
        Admin      bool   `json:"admin,omitempty"`
    } `json:"metadata,omitempty"`

	SessionUuid string `json:"session_uuid,omitempty"`

	Token string `json:"token,omitempty"`

	UtcExpiresAt string `json:"utc_expires_at,omitempty"`

	UtcIssuedAt string `json:"utc_issued_at,omitempty"`

	// The UUID of the matching accent-confd user if there is one. This field can be null. This field should NOT be used anymore, the \"pbx_user_uuid\" in the metadata field is the prefered method to access this information.
	AccentUserUuid string `json:"accent_user_uuid,omitempty"`

	AccentUuid string `json:"accent_uuid,omitempty"`
}

// AssertTokenDataRequired checks if the required fields are not zero-ed
func AssertTokenDataRequired(obj TokenData) error {
	return nil
}

// AssertTokenDataConstraints checks if the values respects the defined constraints
func AssertTokenDataConstraints(obj TokenData) error {
	return nil
}
