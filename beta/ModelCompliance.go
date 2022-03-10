// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Compliance undocumented
type Compliance struct {
	// Object is the base model of Compliance
	Object
	// Ediscovery undocumented
	Ediscovery *EdiscoveryEdiscoveryroot `json:"ediscovery,omitempty"`
}

// ComplianceInformation undocumented
type ComplianceInformation struct {
	// Object is the base model of ComplianceInformation
	Object
	// CertificationControls undocumented
	CertificationControls []CertificationControl `json:"certificationControls,omitempty"`
	// CertificationName undocumented
	CertificationName *string `json:"certificationName,omitempty"`
}

// ComplianceManagementPartner Compliance management partner for all platforms
type ComplianceManagementPartner struct {
	// Entity is the base model of ComplianceManagementPartner
	Entity
	// AndroidEnrollmentAssignments User groups which enroll Android devices through partner.
	AndroidEnrollmentAssignments []ComplianceManagementPartnerAssignment `json:"androidEnrollmentAssignments,omitempty"`
	// AndroidOnboarded Partner onboarded for Android devices.
	AndroidOnboarded *bool `json:"androidOnboarded,omitempty"`
	// DisplayName Partner display name
	DisplayName *string `json:"displayName,omitempty"`
	// IOSEnrollmentAssignments User groups which enroll ios devices through partner.
	IOSEnrollmentAssignments []ComplianceManagementPartnerAssignment `json:"iosEnrollmentAssignments,omitempty"`
	// IOSOnboarded Partner onboarded for ios devices.
	IOSOnboarded *bool `json:"iosOnboarded,omitempty"`
	// LastHeartbeatDateTime Timestamp of last heartbeat after admin onboarded to the compliance management partner
	LastHeartbeatDateTime *time.Time `json:"lastHeartbeatDateTime,omitempty"`
	// MacOsEnrollmentAssignments User groups which enroll Mac devices through partner.
	MacOsEnrollmentAssignments []ComplianceManagementPartnerAssignment `json:"macOsEnrollmentAssignments,omitempty"`
	// MacOsOnboarded Partner onboarded for Mac devices.
	MacOsOnboarded *bool `json:"macOsOnboarded,omitempty"`
	// PartnerState Partner state of this tenant
	PartnerState *DeviceManagementPartnerTenantState `json:"partnerState,omitempty"`
	// WindowsEnrollmentAssignments User groups which enroll Windows devices through partner.
	WindowsEnrollmentAssignments []ComplianceManagementPartnerAssignment `json:"windowsEnrollmentAssignments,omitempty"`
	// WindowsOnboarded Partner onboarded for Windows devices.
	WindowsOnboarded *bool `json:"windowsOnboarded,omitempty"`
}

// ComplianceManagementPartnerAssignment User group targeting for Compliance Management Partner
type ComplianceManagementPartnerAssignment struct {
	// Object is the base model of ComplianceManagementPartnerAssignment
	Object
	// Target Group assignment target.
	Target *DeviceAndAppManagementAssignmentTarget `json:"target,omitempty"`
}
