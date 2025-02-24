// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Configuration undocumented
type Configuration struct {
	// Object is the base model of Configuration
	Object
	// AuthorizedAppIDs undocumented
	AuthorizedAppIDs []string `json:"authorizedAppIds,omitempty"`
	// AuthorizedApps undocumented
	AuthorizedApps []string `json:"authorizedApps,omitempty"`
}

// ConfigurationManagerAction Parameter for action triggerConfigurationManagerAction
type ConfigurationManagerAction struct {
	// Object is the base model of ConfigurationManagerAction
	Object
	// Action The action type to trigger on Configuration Manager client
	Action *ConfigurationManagerActionType `json:"action,omitempty"`
}

// ConfigurationManagerActionResult Result of the ConfigurationManager action
type ConfigurationManagerActionResult struct {
	// DeviceActionResult is the base model of ConfigurationManagerActionResult
	DeviceActionResult
	// ActionDeliveryStatus State of the action being delivered to on-prem server
	ActionDeliveryStatus *ConfigurationManagerActionDeliveryStatus `json:"actionDeliveryStatus,omitempty"`
	// ErrorCode Error code of Configuration Manager action from client
	ErrorCode *int `json:"errorCode,omitempty"`
}

// ConfigurationManagerClientEnabledFeatures configuration Manager client enabled features
type ConfigurationManagerClientEnabledFeatures struct {
	// Object is the base model of ConfigurationManagerClientEnabledFeatures
	Object
	// CompliancePolicy Whether compliance policy is managed by Intune
	CompliancePolicy *bool `json:"compliancePolicy,omitempty"`
	// DeviceConfiguration Whether device configuration is managed by Intune
	DeviceConfiguration *bool `json:"deviceConfiguration,omitempty"`
	// EndpointProtection Whether Endpoint Protection is managed by Intune
	EndpointProtection *bool `json:"endpointProtection,omitempty"`
	// Inventory Whether inventory is managed by Intune
	Inventory *bool `json:"inventory,omitempty"`
	// ModernApps Whether modern application is managed by Intune
	ModernApps *bool `json:"modernApps,omitempty"`
	// OfficeApps Whether Office application is managed by Intune
	OfficeApps *bool `json:"officeApps,omitempty"`
	// ResourceAccess Whether resource access is managed by Intune
	ResourceAccess *bool `json:"resourceAccess,omitempty"`
	// WindowsUpdateForBusiness Whether Windows Update for Business is managed by Intune
	WindowsUpdateForBusiness *bool `json:"windowsUpdateForBusiness,omitempty"`
}

// ConfigurationManagerClientHealthState Configuration manager client health state
type ConfigurationManagerClientHealthState struct {
	// Object is the base model of ConfigurationManagerClientHealthState
	Object
	// ErrorCode Error code for failed state.
	ErrorCode *int `json:"errorCode,omitempty"`
	// LastSyncDateTime Datetime for last sync with configuration manager management point.
	LastSyncDateTime *time.Time `json:"lastSyncDateTime,omitempty"`
	// State Current configuration manager client state.
	State *ConfigurationManagerClientState `json:"state,omitempty"`
}

// ConfigurationManagerClientInformation Configuration Manager client information synced from SCCM
type ConfigurationManagerClientInformation struct {
	// Object is the base model of ConfigurationManagerClientInformation
	Object
	// ClientIdentifier Configuration Manager Client Id from SCCM
	ClientIdentifier *string `json:"clientIdentifier,omitempty"`
	// IsBlocked Configuration Manager Client blocked status from SCCM
	IsBlocked *bool `json:"isBlocked,omitempty"`
}

// ConfigurationManagerCollectionAssignmentTarget Represents an assignment to a Configuration Manager Collection.
type ConfigurationManagerCollectionAssignmentTarget struct {
	// DeviceAndAppManagementAssignmentTarget is the base model of ConfigurationManagerCollectionAssignmentTarget
	DeviceAndAppManagementAssignmentTarget
	// CollectionID The collection Id that is the target of the assignment.
	CollectionID *string `json:"collectionId,omitempty"`
}
