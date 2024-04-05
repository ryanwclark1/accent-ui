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

type AdminUserEmailList struct {
	Emails []AdminUserEmailListEmailsInner `json:"emails,omitempty"`
}

// AssertAdminUserEmailListRequired checks if the required fields are not zero-ed
func AssertAdminUserEmailListRequired(obj AdminUserEmailList) error {
	for _, el := range obj.Emails {
		if err := AssertAdminUserEmailListEmailsInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertAdminUserEmailListConstraints checks if the values respects the defined constraints
func AssertAdminUserEmailListConstraints(obj AdminUserEmailList) error {
	return nil
}