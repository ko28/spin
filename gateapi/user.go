/*
 * Spinnaker API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type User struct {

	AllowedAccounts []string `json:"allowedAccounts,omitempty"`

	AccountNonExpired bool `json:"accountNonExpired,omitempty"`

	LastName string `json:"lastName,omitempty"`

	Email string `json:"email,omitempty"`

	FirstName string `json:"firstName,omitempty"`

	Roles []string `json:"roles,omitempty"`

	Username string `json:"username,omitempty"`

	AccountNonLocked bool `json:"accountNonLocked,omitempty"`

	CredentialsNonExpired bool `json:"credentialsNonExpired,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	Authorities []GrantedAuthority `json:"authorities,omitempty"`
}
