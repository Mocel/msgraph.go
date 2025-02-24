// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// AuthorizationPolicy undocumented
type AuthorizationPolicy struct {
	// PolicyBase is the base model of AuthorizationPolicy
	PolicyBase
	// AllowedToSignUpEmailBasedSubscriptions undocumented
	AllowedToSignUpEmailBasedSubscriptions *bool `json:"allowedToSignUpEmailBasedSubscriptions,omitempty"`
	// AllowedToUseSSPR undocumented
	AllowedToUseSSPR *bool `json:"allowedToUseSSPR,omitempty"`
	// AllowEmailVerifiedUsersToJoinOrganization undocumented
	AllowEmailVerifiedUsersToJoinOrganization *bool `json:"allowEmailVerifiedUsersToJoinOrganization,omitempty"`
	// AllowInvitesFrom undocumented
	AllowInvitesFrom *AllowInvitesFrom `json:"allowInvitesFrom,omitempty"`
	// BlockMsolPowerShell undocumented
	BlockMsolPowerShell *bool `json:"blockMsolPowerShell,omitempty"`
	// DefaultUserRolePermissions undocumented
	DefaultUserRolePermissions *DefaultUserRolePermissions `json:"defaultUserRolePermissions,omitempty"`
	// GuestUserRoleID undocumented
	GuestUserRoleID *UUID `json:"guestUserRoleId,omitempty"`
}
