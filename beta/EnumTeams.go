// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// TeamsAppDistributionMethod undocumented
type TeamsAppDistributionMethod string

const (
	// TeamsAppDistributionMethodVStore undocumented
	TeamsAppDistributionMethodVStore TeamsAppDistributionMethod = "store"
	// TeamsAppDistributionMethodVOrganization undocumented
	TeamsAppDistributionMethodVOrganization TeamsAppDistributionMethod = "organization"
	// TeamsAppDistributionMethodVSideloaded undocumented
	TeamsAppDistributionMethodVSideloaded TeamsAppDistributionMethod = "sideloaded"
	// TeamsAppDistributionMethodVUnknownFutureValue undocumented
	TeamsAppDistributionMethodVUnknownFutureValue TeamsAppDistributionMethod = "unknownFutureValue"
)

var (
	// TeamsAppDistributionMethodPStore is a pointer to TeamsAppDistributionMethodVStore
	TeamsAppDistributionMethodPStore = &_TeamsAppDistributionMethodPStore
	// TeamsAppDistributionMethodPOrganization is a pointer to TeamsAppDistributionMethodVOrganization
	TeamsAppDistributionMethodPOrganization = &_TeamsAppDistributionMethodPOrganization
	// TeamsAppDistributionMethodPSideloaded is a pointer to TeamsAppDistributionMethodVSideloaded
	TeamsAppDistributionMethodPSideloaded = &_TeamsAppDistributionMethodPSideloaded
	// TeamsAppDistributionMethodPUnknownFutureValue is a pointer to TeamsAppDistributionMethodVUnknownFutureValue
	TeamsAppDistributionMethodPUnknownFutureValue = &_TeamsAppDistributionMethodPUnknownFutureValue
)

var (
	_TeamsAppDistributionMethodPStore              = TeamsAppDistributionMethodVStore
	_TeamsAppDistributionMethodPOrganization       = TeamsAppDistributionMethodVOrganization
	_TeamsAppDistributionMethodPSideloaded         = TeamsAppDistributionMethodVSideloaded
	_TeamsAppDistributionMethodPUnknownFutureValue = TeamsAppDistributionMethodVUnknownFutureValue
)

// TeamsAppInstallationScopes undocumented
type TeamsAppInstallationScopes string

const (
	// TeamsAppInstallationScopesVTeam undocumented
	TeamsAppInstallationScopesVTeam TeamsAppInstallationScopes = "team"
	// TeamsAppInstallationScopesVGroupChat undocumented
	TeamsAppInstallationScopesVGroupChat TeamsAppInstallationScopes = "groupChat"
	// TeamsAppInstallationScopesVPersonal undocumented
	TeamsAppInstallationScopesVPersonal TeamsAppInstallationScopes = "personal"
	// TeamsAppInstallationScopesVUnknownFutureValue undocumented
	TeamsAppInstallationScopesVUnknownFutureValue TeamsAppInstallationScopes = "unknownFutureValue"
)

var (
	// TeamsAppInstallationScopesPTeam is a pointer to TeamsAppInstallationScopesVTeam
	TeamsAppInstallationScopesPTeam = &_TeamsAppInstallationScopesPTeam
	// TeamsAppInstallationScopesPGroupChat is a pointer to TeamsAppInstallationScopesVGroupChat
	TeamsAppInstallationScopesPGroupChat = &_TeamsAppInstallationScopesPGroupChat
	// TeamsAppInstallationScopesPPersonal is a pointer to TeamsAppInstallationScopesVPersonal
	TeamsAppInstallationScopesPPersonal = &_TeamsAppInstallationScopesPPersonal
	// TeamsAppInstallationScopesPUnknownFutureValue is a pointer to TeamsAppInstallationScopesVUnknownFutureValue
	TeamsAppInstallationScopesPUnknownFutureValue = &_TeamsAppInstallationScopesPUnknownFutureValue
)

var (
	_TeamsAppInstallationScopesPTeam               = TeamsAppInstallationScopesVTeam
	_TeamsAppInstallationScopesPGroupChat          = TeamsAppInstallationScopesVGroupChat
	_TeamsAppInstallationScopesPPersonal           = TeamsAppInstallationScopesVPersonal
	_TeamsAppInstallationScopesPUnknownFutureValue = TeamsAppInstallationScopesVUnknownFutureValue
)

// TeamsAppPublishingState undocumented
type TeamsAppPublishingState string

const (
	// TeamsAppPublishingStateVSubmitted undocumented
	TeamsAppPublishingStateVSubmitted TeamsAppPublishingState = "submitted"
	// TeamsAppPublishingStateVRejected undocumented
	TeamsAppPublishingStateVRejected TeamsAppPublishingState = "rejected"
	// TeamsAppPublishingStateVPublished undocumented
	TeamsAppPublishingStateVPublished TeamsAppPublishingState = "published"
	// TeamsAppPublishingStateVUnknownFutureValue undocumented
	TeamsAppPublishingStateVUnknownFutureValue TeamsAppPublishingState = "unknownFutureValue"
)

var (
	// TeamsAppPublishingStatePSubmitted is a pointer to TeamsAppPublishingStateVSubmitted
	TeamsAppPublishingStatePSubmitted = &_TeamsAppPublishingStatePSubmitted
	// TeamsAppPublishingStatePRejected is a pointer to TeamsAppPublishingStateVRejected
	TeamsAppPublishingStatePRejected = &_TeamsAppPublishingStatePRejected
	// TeamsAppPublishingStatePPublished is a pointer to TeamsAppPublishingStateVPublished
	TeamsAppPublishingStatePPublished = &_TeamsAppPublishingStatePPublished
	// TeamsAppPublishingStatePUnknownFutureValue is a pointer to TeamsAppPublishingStateVUnknownFutureValue
	TeamsAppPublishingStatePUnknownFutureValue = &_TeamsAppPublishingStatePUnknownFutureValue
)

var (
	_TeamsAppPublishingStatePSubmitted          = TeamsAppPublishingStateVSubmitted
	_TeamsAppPublishingStatePRejected           = TeamsAppPublishingStateVRejected
	_TeamsAppPublishingStatePPublished          = TeamsAppPublishingStateVPublished
	_TeamsAppPublishingStatePUnknownFutureValue = TeamsAppPublishingStateVUnknownFutureValue
)

// TeamsAsyncOperationStatus undocumented
type TeamsAsyncOperationStatus string

const (
	// TeamsAsyncOperationStatusVInvalid undocumented
	TeamsAsyncOperationStatusVInvalid TeamsAsyncOperationStatus = "invalid"
	// TeamsAsyncOperationStatusVNotStarted undocumented
	TeamsAsyncOperationStatusVNotStarted TeamsAsyncOperationStatus = "notStarted"
	// TeamsAsyncOperationStatusVInProgress undocumented
	TeamsAsyncOperationStatusVInProgress TeamsAsyncOperationStatus = "inProgress"
	// TeamsAsyncOperationStatusVSucceeded undocumented
	TeamsAsyncOperationStatusVSucceeded TeamsAsyncOperationStatus = "succeeded"
	// TeamsAsyncOperationStatusVFailed undocumented
	TeamsAsyncOperationStatusVFailed TeamsAsyncOperationStatus = "failed"
	// TeamsAsyncOperationStatusVUnknownFutureValue undocumented
	TeamsAsyncOperationStatusVUnknownFutureValue TeamsAsyncOperationStatus = "unknownFutureValue"
)

var (
	// TeamsAsyncOperationStatusPInvalid is a pointer to TeamsAsyncOperationStatusVInvalid
	TeamsAsyncOperationStatusPInvalid = &_TeamsAsyncOperationStatusPInvalid
	// TeamsAsyncOperationStatusPNotStarted is a pointer to TeamsAsyncOperationStatusVNotStarted
	TeamsAsyncOperationStatusPNotStarted = &_TeamsAsyncOperationStatusPNotStarted
	// TeamsAsyncOperationStatusPInProgress is a pointer to TeamsAsyncOperationStatusVInProgress
	TeamsAsyncOperationStatusPInProgress = &_TeamsAsyncOperationStatusPInProgress
	// TeamsAsyncOperationStatusPSucceeded is a pointer to TeamsAsyncOperationStatusVSucceeded
	TeamsAsyncOperationStatusPSucceeded = &_TeamsAsyncOperationStatusPSucceeded
	// TeamsAsyncOperationStatusPFailed is a pointer to TeamsAsyncOperationStatusVFailed
	TeamsAsyncOperationStatusPFailed = &_TeamsAsyncOperationStatusPFailed
	// TeamsAsyncOperationStatusPUnknownFutureValue is a pointer to TeamsAsyncOperationStatusVUnknownFutureValue
	TeamsAsyncOperationStatusPUnknownFutureValue = &_TeamsAsyncOperationStatusPUnknownFutureValue
)

var (
	_TeamsAsyncOperationStatusPInvalid            = TeamsAsyncOperationStatusVInvalid
	_TeamsAsyncOperationStatusPNotStarted         = TeamsAsyncOperationStatusVNotStarted
	_TeamsAsyncOperationStatusPInProgress         = TeamsAsyncOperationStatusVInProgress
	_TeamsAsyncOperationStatusPSucceeded          = TeamsAsyncOperationStatusVSucceeded
	_TeamsAsyncOperationStatusPFailed             = TeamsAsyncOperationStatusVFailed
	_TeamsAsyncOperationStatusPUnknownFutureValue = TeamsAsyncOperationStatusVUnknownFutureValue
)

// TeamsAsyncOperationType undocumented
type TeamsAsyncOperationType string

const (
	// TeamsAsyncOperationTypeVInvalid undocumented
	TeamsAsyncOperationTypeVInvalid TeamsAsyncOperationType = "invalid"
	// TeamsAsyncOperationTypeVCloneTeam undocumented
	TeamsAsyncOperationTypeVCloneTeam TeamsAsyncOperationType = "cloneTeam"
	// TeamsAsyncOperationTypeVArchiveTeam undocumented
	TeamsAsyncOperationTypeVArchiveTeam TeamsAsyncOperationType = "archiveTeam"
	// TeamsAsyncOperationTypeVUnarchiveTeam undocumented
	TeamsAsyncOperationTypeVUnarchiveTeam TeamsAsyncOperationType = "unarchiveTeam"
	// TeamsAsyncOperationTypeVCreateTeam undocumented
	TeamsAsyncOperationTypeVCreateTeam TeamsAsyncOperationType = "createTeam"
	// TeamsAsyncOperationTypeVUnknownFutureValue undocumented
	TeamsAsyncOperationTypeVUnknownFutureValue TeamsAsyncOperationType = "unknownFutureValue"
	// TeamsAsyncOperationTypeVTeamifyGroup undocumented
	TeamsAsyncOperationTypeVTeamifyGroup TeamsAsyncOperationType = "teamifyGroup"
	// TeamsAsyncOperationTypeVCreateChannel undocumented
	TeamsAsyncOperationTypeVCreateChannel TeamsAsyncOperationType = "createChannel"
	// TeamsAsyncOperationTypeVCreateChat undocumented
	TeamsAsyncOperationTypeVCreateChat TeamsAsyncOperationType = "createChat"
)

var (
	// TeamsAsyncOperationTypePInvalid is a pointer to TeamsAsyncOperationTypeVInvalid
	TeamsAsyncOperationTypePInvalid = &_TeamsAsyncOperationTypePInvalid
	// TeamsAsyncOperationTypePCloneTeam is a pointer to TeamsAsyncOperationTypeVCloneTeam
	TeamsAsyncOperationTypePCloneTeam = &_TeamsAsyncOperationTypePCloneTeam
	// TeamsAsyncOperationTypePArchiveTeam is a pointer to TeamsAsyncOperationTypeVArchiveTeam
	TeamsAsyncOperationTypePArchiveTeam = &_TeamsAsyncOperationTypePArchiveTeam
	// TeamsAsyncOperationTypePUnarchiveTeam is a pointer to TeamsAsyncOperationTypeVUnarchiveTeam
	TeamsAsyncOperationTypePUnarchiveTeam = &_TeamsAsyncOperationTypePUnarchiveTeam
	// TeamsAsyncOperationTypePCreateTeam is a pointer to TeamsAsyncOperationTypeVCreateTeam
	TeamsAsyncOperationTypePCreateTeam = &_TeamsAsyncOperationTypePCreateTeam
	// TeamsAsyncOperationTypePUnknownFutureValue is a pointer to TeamsAsyncOperationTypeVUnknownFutureValue
	TeamsAsyncOperationTypePUnknownFutureValue = &_TeamsAsyncOperationTypePUnknownFutureValue
	// TeamsAsyncOperationTypePTeamifyGroup is a pointer to TeamsAsyncOperationTypeVTeamifyGroup
	TeamsAsyncOperationTypePTeamifyGroup = &_TeamsAsyncOperationTypePTeamifyGroup
	// TeamsAsyncOperationTypePCreateChannel is a pointer to TeamsAsyncOperationTypeVCreateChannel
	TeamsAsyncOperationTypePCreateChannel = &_TeamsAsyncOperationTypePCreateChannel
	// TeamsAsyncOperationTypePCreateChat is a pointer to TeamsAsyncOperationTypeVCreateChat
	TeamsAsyncOperationTypePCreateChat = &_TeamsAsyncOperationTypePCreateChat
)

var (
	_TeamsAsyncOperationTypePInvalid            = TeamsAsyncOperationTypeVInvalid
	_TeamsAsyncOperationTypePCloneTeam          = TeamsAsyncOperationTypeVCloneTeam
	_TeamsAsyncOperationTypePArchiveTeam        = TeamsAsyncOperationTypeVArchiveTeam
	_TeamsAsyncOperationTypePUnarchiveTeam      = TeamsAsyncOperationTypeVUnarchiveTeam
	_TeamsAsyncOperationTypePCreateTeam         = TeamsAsyncOperationTypeVCreateTeam
	_TeamsAsyncOperationTypePUnknownFutureValue = TeamsAsyncOperationTypeVUnknownFutureValue
	_TeamsAsyncOperationTypePTeamifyGroup       = TeamsAsyncOperationTypeVTeamifyGroup
	_TeamsAsyncOperationTypePCreateChannel      = TeamsAsyncOperationTypeVCreateChannel
	_TeamsAsyncOperationTypePCreateChat         = TeamsAsyncOperationTypeVCreateChat
)
