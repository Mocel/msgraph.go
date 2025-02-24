// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// ConditionalAccessApplications undocumented
type ConditionalAccessApplications struct {
	// Object is the base model of ConditionalAccessApplications
	Object
	// ExcludeApplications undocumented
	ExcludeApplications []string `json:"excludeApplications,omitempty"`
	// IncludeApplications undocumented
	IncludeApplications []string `json:"includeApplications,omitempty"`
	// IncludeAuthenticationContextClassReferences undocumented
	IncludeAuthenticationContextClassReferences []string `json:"includeAuthenticationContextClassReferences,omitempty"`
	// IncludeUserActions undocumented
	IncludeUserActions []string `json:"includeUserActions,omitempty"`
}

// ConditionalAccessClientApplications undocumented
type ConditionalAccessClientApplications struct {
	// Object is the base model of ConditionalAccessClientApplications
	Object
	// ExcludeServicePrincipals undocumented
	ExcludeServicePrincipals []string `json:"excludeServicePrincipals,omitempty"`
	// IncludeServicePrincipals undocumented
	IncludeServicePrincipals []string `json:"includeServicePrincipals,omitempty"`
}

// ConditionalAccessConditionSet undocumented
type ConditionalAccessConditionSet struct {
	// Object is the base model of ConditionalAccessConditionSet
	Object
	// Applications undocumented
	Applications *ConditionalAccessApplications `json:"applications,omitempty"`
	// ClientApplications undocumented
	ClientApplications *ConditionalAccessClientApplications `json:"clientApplications,omitempty"`
	// ClientAppTypes undocumented
	ClientAppTypes []ConditionalAccessClientApp `json:"clientAppTypes,omitempty"`
	// Devices undocumented
	Devices *ConditionalAccessDevices `json:"devices,omitempty"`
	// DeviceStates undocumented
	DeviceStates *ConditionalAccessDeviceStates `json:"deviceStates,omitempty"`
	// Locations undocumented
	Locations *ConditionalAccessLocations `json:"locations,omitempty"`
	// Platforms undocumented
	Platforms *ConditionalAccessPlatforms `json:"platforms,omitempty"`
	// SignInRiskLevels undocumented
	SignInRiskLevels []RiskLevel `json:"signInRiskLevels,omitempty"`
	// UserRiskLevels undocumented
	UserRiskLevels []RiskLevel `json:"userRiskLevels,omitempty"`
	// Users undocumented
	Users *ConditionalAccessUsers `json:"users,omitempty"`
}

// ConditionalAccessDeviceStates undocumented
type ConditionalAccessDeviceStates struct {
	// Object is the base model of ConditionalAccessDeviceStates
	Object
	// ExcludeStates undocumented
	ExcludeStates []string `json:"excludeStates,omitempty"`
	// IncludeStates undocumented
	IncludeStates []string `json:"includeStates,omitempty"`
}

// ConditionalAccessDevices undocumented
type ConditionalAccessDevices struct {
	// Object is the base model of ConditionalAccessDevices
	Object
	// DeviceFilter undocumented
	DeviceFilter *ConditionalAccessFilter `json:"deviceFilter,omitempty"`
	// ExcludeDevices undocumented
	ExcludeDevices []string `json:"excludeDevices,omitempty"`
	// ExcludeDeviceStates undocumented
	ExcludeDeviceStates []string `json:"excludeDeviceStates,omitempty"`
	// IncludeDevices undocumented
	IncludeDevices []string `json:"includeDevices,omitempty"`
	// IncludeDeviceStates undocumented
	IncludeDeviceStates []string `json:"includeDeviceStates,omitempty"`
}

// ConditionalAccessFilter undocumented
type ConditionalAccessFilter struct {
	// Object is the base model of ConditionalAccessFilter
	Object
	// Mode undocumented
	Mode *FilterMode `json:"mode,omitempty"`
	// Rule undocumented
	Rule *string `json:"rule,omitempty"`
}

// ConditionalAccessGrantControls undocumented
type ConditionalAccessGrantControls struct {
	// Object is the base model of ConditionalAccessGrantControls
	Object
	// BuiltInControls undocumented
	BuiltInControls []ConditionalAccessGrantControl `json:"builtInControls,omitempty"`
	// CustomAuthenticationFactors undocumented
	CustomAuthenticationFactors []string `json:"customAuthenticationFactors,omitempty"`
	// Operator undocumented
	Operator *string `json:"operator,omitempty"`
	// TermsOfUse undocumented
	TermsOfUse []string `json:"termsOfUse,omitempty"`
}

// ConditionalAccessLocations undocumented
type ConditionalAccessLocations struct {
	// Object is the base model of ConditionalAccessLocations
	Object
	// ExcludeLocations undocumented
	ExcludeLocations []string `json:"excludeLocations,omitempty"`
	// IncludeLocations undocumented
	IncludeLocations []string `json:"includeLocations,omitempty"`
}

// ConditionalAccessPlatforms undocumented
type ConditionalAccessPlatforms struct {
	// Object is the base model of ConditionalAccessPlatforms
	Object
	// ExcludePlatforms undocumented
	ExcludePlatforms []ConditionalAccessDevicePlatform `json:"excludePlatforms,omitempty"`
	// IncludePlatforms undocumented
	IncludePlatforms []ConditionalAccessDevicePlatform `json:"includePlatforms,omitempty"`
}

// ConditionalAccessPolicy undocumented
type ConditionalAccessPolicy struct {
	// Entity is the base model of ConditionalAccessPolicy
	Entity
	// Conditions undocumented
	Conditions *ConditionalAccessConditionSet `json:"conditions,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// GrantControls undocumented
	GrantControls *ConditionalAccessGrantControls `json:"grantControls,omitempty"`
	// ModifiedDateTime undocumented
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// SessionControls undocumented
	SessionControls *ConditionalAccessSessionControls `json:"sessionControls,omitempty"`
	// State undocumented
	State *ConditionalAccessPolicyState `json:"state,omitempty"`
}

// ConditionalAccessRoot undocumented
type ConditionalAccessRoot struct {
	// Entity is the base model of ConditionalAccessRoot
	Entity
	// AuthenticationContextClassReferences undocumented
	AuthenticationContextClassReferences []AuthenticationContextClassReference `json:"authenticationContextClassReferences,omitempty"`
	// NamedLocations undocumented
	NamedLocations []NamedLocation `json:"namedLocations,omitempty"`
	// Policies undocumented
	Policies []ConditionalAccessPolicy `json:"policies,omitempty"`
}

// ConditionalAccessRuleSatisfied undocumented
type ConditionalAccessRuleSatisfied struct {
	// Object is the base model of ConditionalAccessRuleSatisfied
	Object
	// ConditionalAccessCondition undocumented
	ConditionalAccessCondition *ConditionalAccessConditions `json:"conditionalAccessCondition,omitempty"`
	// RuleSatisfied undocumented
	RuleSatisfied *ConditionalAccessRule `json:"ruleSatisfied,omitempty"`
}

// ConditionalAccessSessionControl undocumented
type ConditionalAccessSessionControl struct {
	// Object is the base model of ConditionalAccessSessionControl
	Object
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
}

// ConditionalAccessSessionControls undocumented
type ConditionalAccessSessionControls struct {
	// Object is the base model of ConditionalAccessSessionControls
	Object
	// ApplicationEnforcedRestrictions undocumented
	ApplicationEnforcedRestrictions *ApplicationEnforcedRestrictionsSessionControl `json:"applicationEnforcedRestrictions,omitempty"`
	// CloudAppSecurity undocumented
	CloudAppSecurity *CloudAppSecuritySessionControl `json:"cloudAppSecurity,omitempty"`
	// PersistentBrowser undocumented
	PersistentBrowser *PersistentBrowserSessionControl `json:"persistentBrowser,omitempty"`
	// SignInFrequency undocumented
	SignInFrequency *SignInFrequencySessionControl `json:"signInFrequency,omitempty"`
}

// ConditionalAccessUsers undocumented
type ConditionalAccessUsers struct {
	// Object is the base model of ConditionalAccessUsers
	Object
	// ExcludeGroups undocumented
	ExcludeGroups []string `json:"excludeGroups,omitempty"`
	// ExcludeRoles undocumented
	ExcludeRoles []string `json:"excludeRoles,omitempty"`
	// ExcludeUsers undocumented
	ExcludeUsers []string `json:"excludeUsers,omitempty"`
	// IncludeGroups undocumented
	IncludeGroups []string `json:"includeGroups,omitempty"`
	// IncludeRoles undocumented
	IncludeRoles []string `json:"includeRoles,omitempty"`
	// IncludeUsers undocumented
	IncludeUsers []string `json:"includeUsers,omitempty"`
}
