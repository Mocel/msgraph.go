// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// AccessAction undocumented
type AccessAction struct {
	// Object is the base model of AccessAction
	Object
}

// AccessPackage undocumented
type AccessPackage struct {
	// Entity is the base model of AccessPackage
	Entity
	// CatalogID undocumented
	CatalogID *string `json:"catalogId,omitempty"`
	// CreatedBy undocumented
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsHidden undocumented
	IsHidden *bool `json:"isHidden,omitempty"`
	// IsRoleScopesVisible undocumented
	IsRoleScopesVisible *bool `json:"isRoleScopesVisible,omitempty"`
	// ModifiedBy undocumented
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// ModifiedDateTime undocumented
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// AccessPackageAssignmentPolicies undocumented
	AccessPackageAssignmentPolicies []AccessPackageAssignmentPolicy `json:"accessPackageAssignmentPolicies,omitempty"`
	// AccessPackageCatalog undocumented
	AccessPackageCatalog *AccessPackageCatalog `json:"accessPackageCatalog,omitempty"`
	// AccessPackageResourceRoleScopes undocumented
	AccessPackageResourceRoleScopes []AccessPackageResourceRoleScope `json:"accessPackageResourceRoleScopes,omitempty"`
}

// AccessPackageAnswer undocumented
type AccessPackageAnswer struct {
	// Object is the base model of AccessPackageAnswer
	Object
	// AnsweredQuestion undocumented
	AnsweredQuestion *AccessPackageQuestion `json:"answeredQuestion,omitempty"`
	// DisplayValue undocumented
	DisplayValue *string `json:"displayValue,omitempty"`
}

// AccessPackageAnswerChoice undocumented
type AccessPackageAnswerChoice struct {
	// Object is the base model of AccessPackageAnswerChoice
	Object
	// ActualValue undocumented
	ActualValue *string `json:"actualValue,omitempty"`
	// DisplayValue undocumented
	DisplayValue *AccessPackageLocalizedContent `json:"displayValue,omitempty"`
}

// AccessPackageAnswerString undocumented
type AccessPackageAnswerString struct {
	// AccessPackageAnswer is the base model of AccessPackageAnswerString
	AccessPackageAnswer
	// Value undocumented
	Value *string `json:"value,omitempty"`
}

// AccessPackageAssignment undocumented
type AccessPackageAssignment struct {
	// Entity is the base model of AccessPackageAssignment
	Entity
	// AccessPackageID undocumented
	AccessPackageID *string `json:"accessPackageId,omitempty"`
	// AssignmentPolicyID undocumented
	AssignmentPolicyID *string `json:"assignmentPolicyId,omitempty"`
	// AssignmentState undocumented
	AssignmentState *string `json:"assignmentState,omitempty"`
	// AssignmentStatus undocumented
	AssignmentStatus *string `json:"assignmentStatus,omitempty"`
	// CatalogID undocumented
	CatalogID *string `json:"catalogId,omitempty"`
	// ExpiredDateTime undocumented
	ExpiredDateTime *time.Time `json:"expiredDateTime,omitempty"`
	// IsExtended undocumented
	IsExtended *bool `json:"isExtended,omitempty"`
	// Schedule undocumented
	Schedule *RequestSchedule `json:"schedule,omitempty"`
	// TargetID undocumented
	TargetID *string `json:"targetId,omitempty"`
	// AccessPackage undocumented
	AccessPackage *AccessPackage `json:"accessPackage,omitempty"`
	// AccessPackageAssignmentPolicy undocumented
	AccessPackageAssignmentPolicy *AccessPackageAssignmentPolicy `json:"accessPackageAssignmentPolicy,omitempty"`
	// AccessPackageAssignmentRequests undocumented
	AccessPackageAssignmentRequests []AccessPackageAssignmentRequestObject `json:"accessPackageAssignmentRequests,omitempty"`
	// AccessPackageAssignmentResourceRoles undocumented
	AccessPackageAssignmentResourceRoles []AccessPackageAssignmentResourceRole `json:"accessPackageAssignmentResourceRoles,omitempty"`
	// Target undocumented
	Target *AccessPackageSubject `json:"target,omitempty"`
}

// AccessPackageAssignmentPolicy undocumented
type AccessPackageAssignmentPolicy struct {
	// Entity is the base model of AccessPackageAssignmentPolicy
	Entity
	// AccessPackageID undocumented
	AccessPackageID *string `json:"accessPackageId,omitempty"`
	// AccessReviewSettings undocumented
	AccessReviewSettings *AssignmentReviewSettings `json:"accessReviewSettings,omitempty"`
	// CanExtend undocumented
	CanExtend *bool `json:"canExtend,omitempty"`
	// CreatedBy undocumented
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// DurationInDays undocumented
	DurationInDays *int `json:"durationInDays,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// ModifiedBy undocumented
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// ModifiedDateTime undocumented
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// Questions undocumented
	Questions []AccessPackageQuestion `json:"questions,omitempty"`
	// RequestApprovalSettings undocumented
	RequestApprovalSettings *ApprovalSettings `json:"requestApprovalSettings,omitempty"`
	// RequestorSettings undocumented
	RequestorSettings *RequestorSettings `json:"requestorSettings,omitempty"`
	// AccessPackage undocumented
	AccessPackage *AccessPackage `json:"accessPackage,omitempty"`
	// AccessPackageCatalog undocumented
	AccessPackageCatalog *AccessPackageCatalog `json:"accessPackageCatalog,omitempty"`
}

// AccessPackageAssignmentRequestObject undocumented
type AccessPackageAssignmentRequestObject struct {
	// Entity is the base model of AccessPackageAssignmentRequestObject
	Entity
	// Answers undocumented
	Answers []AccessPackageAnswer `json:"answers,omitempty"`
	// CompletedDate undocumented
	CompletedDate *time.Time `json:"completedDate,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// IsValidationOnly undocumented
	IsValidationOnly *bool `json:"isValidationOnly,omitempty"`
	// Justification undocumented
	Justification *string `json:"justification,omitempty"`
	// RequestState undocumented
	RequestState *string `json:"requestState,omitempty"`
	// RequestStatus undocumented
	RequestStatus *string `json:"requestStatus,omitempty"`
	// RequestType undocumented
	RequestType *string `json:"requestType,omitempty"`
	// Schedule undocumented
	Schedule *RequestSchedule `json:"schedule,omitempty"`
	// AccessPackage undocumented
	AccessPackage *AccessPackage `json:"accessPackage,omitempty"`
	// AccessPackageAssignment undocumented
	AccessPackageAssignment *AccessPackageAssignment `json:"accessPackageAssignment,omitempty"`
	// Requestor undocumented
	Requestor *AccessPackageSubject `json:"requestor,omitempty"`
}

// AccessPackageAssignmentRequestRequirements undocumented
type AccessPackageAssignmentRequestRequirements struct {
	// Object is the base model of AccessPackageAssignmentRequestRequirements
	Object
	// AllowCustomAssignmentSchedule undocumented
	AllowCustomAssignmentSchedule *bool `json:"allowCustomAssignmentSchedule,omitempty"`
	// ExistingAnswers undocumented
	ExistingAnswers []AccessPackageAnswer `json:"existingAnswers,omitempty"`
	// IsApprovalRequiredForAdd undocumented
	IsApprovalRequiredForAdd *bool `json:"isApprovalRequiredForAdd,omitempty"`
	// IsApprovalRequiredForUpdate undocumented
	IsApprovalRequiredForUpdate *bool `json:"isApprovalRequiredForUpdate,omitempty"`
	// PolicyDescription undocumented
	PolicyDescription *string `json:"policyDescription,omitempty"`
	// PolicyDisplayName undocumented
	PolicyDisplayName *string `json:"policyDisplayName,omitempty"`
	// PolicyID undocumented
	PolicyID *string `json:"policyId,omitempty"`
	// Questions undocumented
	Questions []AccessPackageQuestion `json:"questions,omitempty"`
	// Schedule undocumented
	Schedule *RequestSchedule `json:"schedule,omitempty"`
}

// AccessPackageAssignmentResourceRole undocumented
type AccessPackageAssignmentResourceRole struct {
	// Entity is the base model of AccessPackageAssignmentResourceRole
	Entity
	// OriginID undocumented
	OriginID *string `json:"originId,omitempty"`
	// OriginSystem undocumented
	OriginSystem *string `json:"originSystem,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// AccessPackageAssignments undocumented
	AccessPackageAssignments []AccessPackageAssignment `json:"accessPackageAssignments,omitempty"`
	// AccessPackageResourceRole undocumented
	AccessPackageResourceRole *AccessPackageResourceRole `json:"accessPackageResourceRole,omitempty"`
	// AccessPackageResourceScope undocumented
	AccessPackageResourceScope *AccessPackageResourceScope `json:"accessPackageResourceScope,omitempty"`
	// AccessPackageSubject undocumented
	AccessPackageSubject *AccessPackageSubject `json:"accessPackageSubject,omitempty"`
}

// AccessPackageCatalog undocumented
type AccessPackageCatalog struct {
	// Entity is the base model of AccessPackageCatalog
	Entity
	// CatalogStatus undocumented
	CatalogStatus *string `json:"catalogStatus,omitempty"`
	// CatalogType undocumented
	CatalogType *string `json:"catalogType,omitempty"`
	// CreatedBy undocumented
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsExternallyVisible undocumented
	IsExternallyVisible *bool `json:"isExternallyVisible,omitempty"`
	// ModifiedBy undocumented
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// ModifiedDateTime undocumented
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// AccessPackageResourceRoles undocumented
	AccessPackageResourceRoles []AccessPackageResourceRole `json:"accessPackageResourceRoles,omitempty"`
	// AccessPackageResources undocumented
	AccessPackageResources []AccessPackageResource `json:"accessPackageResources,omitempty"`
	// AccessPackageResourceScopes undocumented
	AccessPackageResourceScopes []AccessPackageResourceScope `json:"accessPackageResourceScopes,omitempty"`
	// AccessPackages undocumented
	AccessPackages []AccessPackage `json:"accessPackages,omitempty"`
}

// AccessPackageLocalizedContent undocumented
type AccessPackageLocalizedContent struct {
	// Object is the base model of AccessPackageLocalizedContent
	Object
	// DefaultText undocumented
	DefaultText *string `json:"defaultText,omitempty"`
	// LocalizedTexts undocumented
	LocalizedTexts []AccessPackageLocalizedText `json:"localizedTexts,omitempty"`
}

// AccessPackageLocalizedText undocumented
type AccessPackageLocalizedText struct {
	// Object is the base model of AccessPackageLocalizedText
	Object
	// LanguageCode undocumented
	LanguageCode *string `json:"languageCode,omitempty"`
	// Text undocumented
	Text *string `json:"text,omitempty"`
}

// AccessPackageMultipleChoiceQuestion undocumented
type AccessPackageMultipleChoiceQuestion struct {
	// AccessPackageQuestion is the base model of AccessPackageMultipleChoiceQuestion
	AccessPackageQuestion
	// AllowsMultipleSelection undocumented
	AllowsMultipleSelection *bool `json:"allowsMultipleSelection,omitempty"`
	// Choices undocumented
	Choices []AccessPackageAnswerChoice `json:"choices,omitempty"`
}

// AccessPackageQuestion undocumented
type AccessPackageQuestion struct {
	// Object is the base model of AccessPackageQuestion
	Object
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// IsRequired undocumented
	IsRequired *bool `json:"isRequired,omitempty"`
	// Sequence undocumented
	Sequence *int `json:"sequence,omitempty"`
	// Text undocumented
	Text *AccessPackageLocalizedContent `json:"text,omitempty"`
}

// AccessPackageResource undocumented
type AccessPackageResource struct {
	// Entity is the base model of AccessPackageResource
	Entity
	// AddedBy undocumented
	AddedBy *string `json:"addedBy,omitempty"`
	// AddedOn undocumented
	AddedOn *time.Time `json:"addedOn,omitempty"`
	// Attributes undocumented
	Attributes []AccessPackageResourceAttribute `json:"attributes,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsPendingOnboarding undocumented
	IsPendingOnboarding *bool `json:"isPendingOnboarding,omitempty"`
	// OriginID undocumented
	OriginID *string `json:"originId,omitempty"`
	// OriginSystem undocumented
	OriginSystem *string `json:"originSystem,omitempty"`
	// ResourceType undocumented
	ResourceType *string `json:"resourceType,omitempty"`
	// URL undocumented
	URL *string `json:"url,omitempty"`
	// AccessPackageResourceEnvironment undocumented
	AccessPackageResourceEnvironment *AccessPackageResourceEnvironment `json:"accessPackageResourceEnvironment,omitempty"`
	// AccessPackageResourceRoles undocumented
	AccessPackageResourceRoles []AccessPackageResourceRole `json:"accessPackageResourceRoles,omitempty"`
	// AccessPackageResourceScopes undocumented
	AccessPackageResourceScopes []AccessPackageResourceScope `json:"accessPackageResourceScopes,omitempty"`
}

// AccessPackageResourceAttribute undocumented
type AccessPackageResourceAttribute struct {
	// Object is the base model of AccessPackageResourceAttribute
	Object
	// AttributeDestination undocumented
	AttributeDestination *AccessPackageResourceAttributeDestination `json:"attributeDestination,omitempty"`
	// AttributeName undocumented
	AttributeName *string `json:"attributeName,omitempty"`
	// AttributeSource undocumented
	AttributeSource *AccessPackageResourceAttributeSource `json:"attributeSource,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
}

// AccessPackageResourceAttributeDestination undocumented
type AccessPackageResourceAttributeDestination struct {
	// Object is the base model of AccessPackageResourceAttributeDestination
	Object
}

// AccessPackageResourceAttributeQuestion undocumented
type AccessPackageResourceAttributeQuestion struct {
	// AccessPackageResourceAttributeSource is the base model of AccessPackageResourceAttributeQuestion
	AccessPackageResourceAttributeSource
	// Question undocumented
	Question *AccessPackageQuestion `json:"question,omitempty"`
}

// AccessPackageResourceAttributeSource undocumented
type AccessPackageResourceAttributeSource struct {
	// Object is the base model of AccessPackageResourceAttributeSource
	Object
}

// AccessPackageResourceEnvironment undocumented
type AccessPackageResourceEnvironment struct {
	// Entity is the base model of AccessPackageResourceEnvironment
	Entity
	// ConnectionInfo undocumented
	ConnectionInfo *ConnectionInfo `json:"connectionInfo,omitempty"`
	// CreatedBy undocumented
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsDefaultEnvironment undocumented
	IsDefaultEnvironment *bool `json:"isDefaultEnvironment,omitempty"`
	// ModifiedBy undocumented
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// ModifiedDateTime undocumented
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// OriginID undocumented
	OriginID *string `json:"originId,omitempty"`
	// OriginSystem undocumented
	OriginSystem *string `json:"originSystem,omitempty"`
	// AccessPackageResources undocumented
	AccessPackageResources []AccessPackageResource `json:"accessPackageResources,omitempty"`
}

// AccessPackageResourceRequestObject undocumented
type AccessPackageResourceRequestObject struct {
	// Entity is the base model of AccessPackageResourceRequestObject
	Entity
	// CatalogID undocumented
	CatalogID *string `json:"catalogId,omitempty"`
	// ExecuteImmediately undocumented
	ExecuteImmediately *bool `json:"executeImmediately,omitempty"`
	// ExpirationDateTime undocumented
	ExpirationDateTime *time.Time `json:"expirationDateTime,omitempty"`
	// IsValidationOnly undocumented
	IsValidationOnly *bool `json:"isValidationOnly,omitempty"`
	// Justification undocumented
	Justification *string `json:"justification,omitempty"`
	// RequestState undocumented
	RequestState *string `json:"requestState,omitempty"`
	// RequestStatus undocumented
	RequestStatus *string `json:"requestStatus,omitempty"`
	// RequestType undocumented
	RequestType *string `json:"requestType,omitempty"`
	// AccessPackageResource undocumented
	AccessPackageResource *AccessPackageResource `json:"accessPackageResource,omitempty"`
	// Requestor undocumented
	Requestor *AccessPackageSubject `json:"requestor,omitempty"`
}

// AccessPackageResourceRole undocumented
type AccessPackageResourceRole struct {
	// Entity is the base model of AccessPackageResourceRole
	Entity
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// OriginID undocumented
	OriginID *string `json:"originId,omitempty"`
	// OriginSystem undocumented
	OriginSystem *string `json:"originSystem,omitempty"`
	// AccessPackageResource undocumented
	AccessPackageResource *AccessPackageResource `json:"accessPackageResource,omitempty"`
}

// AccessPackageResourceRoleScope undocumented
type AccessPackageResourceRoleScope struct {
	// Entity is the base model of AccessPackageResourceRoleScope
	Entity
	// CreatedBy undocumented
	CreatedBy *string `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// ModifiedBy undocumented
	ModifiedBy *string `json:"modifiedBy,omitempty"`
	// ModifiedDateTime undocumented
	ModifiedDateTime *time.Time `json:"modifiedDateTime,omitempty"`
	// AccessPackageResourceRole undocumented
	AccessPackageResourceRole *AccessPackageResourceRole `json:"accessPackageResourceRole,omitempty"`
	// AccessPackageResourceScope undocumented
	AccessPackageResourceScope *AccessPackageResourceScope `json:"accessPackageResourceScope,omitempty"`
}

// AccessPackageResourceScope undocumented
type AccessPackageResourceScope struct {
	// Entity is the base model of AccessPackageResourceScope
	Entity
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsRootScope undocumented
	IsRootScope *bool `json:"isRootScope,omitempty"`
	// OriginID undocumented
	OriginID *string `json:"originId,omitempty"`
	// OriginSystem undocumented
	OriginSystem *string `json:"originSystem,omitempty"`
	// RoleOriginID undocumented
	RoleOriginID *string `json:"roleOriginId,omitempty"`
	// URL undocumented
	URL *string `json:"url,omitempty"`
	// AccessPackageResource undocumented
	AccessPackageResource *AccessPackageResource `json:"accessPackageResource,omitempty"`
}

// AccessPackageSubject undocumented
type AccessPackageSubject struct {
	// Entity is the base model of AccessPackageSubject
	Entity
	// AltSecID undocumented
	AltSecID *string `json:"altSecId,omitempty"`
	// ConnectedOrganizationID undocumented
	ConnectedOrganizationID *string `json:"connectedOrganizationId,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Email undocumented
	Email *string `json:"email,omitempty"`
	// ObjectID undocumented
	ObjectID *string `json:"objectId,omitempty"`
	// OnPremisesSecurityIdentifier undocumented
	OnPremisesSecurityIdentifier *string `json:"onPremisesSecurityIdentifier,omitempty"`
	// PrincipalName undocumented
	PrincipalName *string `json:"principalName,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
	// ConnectedOrganization undocumented
	ConnectedOrganization *ConnectedOrganization `json:"connectedOrganization,omitempty"`
}

// AccessPackageTextInputQuestion undocumented
type AccessPackageTextInputQuestion struct {
	// AccessPackageQuestion is the base model of AccessPackageTextInputQuestion
	AccessPackageQuestion
	// IsSingleLineQuestion undocumented
	IsSingleLineQuestion *bool `json:"isSingleLineQuestion,omitempty"`
}

// AccessPackageUserDirectoryAttributeStore undocumented
type AccessPackageUserDirectoryAttributeStore struct {
	// AccessPackageResourceAttributeDestination is the base model of AccessPackageUserDirectoryAttributeStore
	AccessPackageResourceAttributeDestination
}

// AccessReview undocumented
type AccessReview struct {
	// Entity is the base model of AccessReview
	Entity
	// BusinessFlowTemplateID undocumented
	BusinessFlowTemplateID *string `json:"businessFlowTemplateId,omitempty"`
	// CreatedBy undocumented
	CreatedBy *UserIdentity `json:"createdBy,omitempty"`
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// ReviewedEntity undocumented
	ReviewedEntity *Identity `json:"reviewedEntity,omitempty"`
	// ReviewerType undocumented
	ReviewerType *string `json:"reviewerType,omitempty"`
	// Settings undocumented
	Settings *AccessReviewSettings `json:"settings,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// Decisions undocumented
	Decisions []AccessReviewDecision `json:"decisions,omitempty"`
	// Instances undocumented
	Instances []AccessReview `json:"instances,omitempty"`
	// MyDecisions undocumented
	MyDecisions []AccessReviewDecision `json:"myDecisions,omitempty"`
	// Reviewers undocumented
	Reviewers []AccessReviewReviewer `json:"reviewers,omitempty"`
}

// AccessReviewApplyAction undocumented
type AccessReviewApplyAction struct {
	// Object is the base model of AccessReviewApplyAction
	Object
}

// AccessReviewDecision undocumented
type AccessReviewDecision struct {
	// Entity is the base model of AccessReviewDecision
	Entity
	// AccessRecommendation undocumented
	AccessRecommendation *string `json:"accessRecommendation,omitempty"`
	// AccessReviewID undocumented
	AccessReviewID *string `json:"accessReviewId,omitempty"`
	// AppliedBy undocumented
	AppliedBy *UserIdentity `json:"appliedBy,omitempty"`
	// AppliedDateTime undocumented
	AppliedDateTime *time.Time `json:"appliedDateTime,omitempty"`
	// ApplyResult undocumented
	ApplyResult *string `json:"applyResult,omitempty"`
	// Justification undocumented
	Justification *string `json:"justification,omitempty"`
	// ReviewedBy undocumented
	ReviewedBy *UserIdentity `json:"reviewedBy,omitempty"`
	// ReviewedDateTime undocumented
	ReviewedDateTime *time.Time `json:"reviewedDateTime,omitempty"`
	// ReviewResult undocumented
	ReviewResult *string `json:"reviewResult,omitempty"`
}

// AccessReviewHistoryDefinition undocumented
type AccessReviewHistoryDefinition struct {
	// Entity is the base model of AccessReviewHistoryDefinition
	Entity
	// CreatedBy undocumented
	CreatedBy *UserIdentity `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// Decisions undocumented
	Decisions []AccessReviewHistoryDecisionFilter `json:"decisions,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// DownloadURI undocumented
	DownloadURI *string `json:"downloadUri,omitempty"`
	// FulfilledDateTime undocumented
	FulfilledDateTime *time.Time `json:"fulfilledDateTime,omitempty"`
	// ReviewHistoryPeriodEndDateTime undocumented
	ReviewHistoryPeriodEndDateTime *time.Time `json:"reviewHistoryPeriodEndDateTime,omitempty"`
	// ReviewHistoryPeriodStartDateTime undocumented
	ReviewHistoryPeriodStartDateTime *time.Time `json:"reviewHistoryPeriodStartDateTime,omitempty"`
	// Scopes undocumented
	Scopes []AccessReviewScope `json:"scopes,omitempty"`
	// Status undocumented
	Status *AccessReviewHistoryStatus `json:"status,omitempty"`
}

// AccessReviewInactiveUsersQueryScope undocumented
type AccessReviewInactiveUsersQueryScope struct {
	// AccessReviewQueryScope is the base model of AccessReviewInactiveUsersQueryScope
	AccessReviewQueryScope
	// InactiveDuration undocumented
	InactiveDuration *Duration `json:"inactiveDuration,omitempty"`
}

// AccessReviewInstance undocumented
type AccessReviewInstance struct {
	// Entity is the base model of AccessReviewInstance
	Entity
	// EndDateTime undocumented
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
	// Scope undocumented
	Scope *AccessReviewScope `json:"scope,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// Decisions undocumented
	Decisions []AccessReviewInstanceDecisionItem `json:"decisions,omitempty"`
	// Definition undocumented
	Definition *AccessReviewScheduleDefinition `json:"definition,omitempty"`
}

// AccessReviewInstanceDecisionItem undocumented
type AccessReviewInstanceDecisionItem struct {
	// Entity is the base model of AccessReviewInstanceDecisionItem
	Entity
	// AccessReviewID undocumented
	AccessReviewID *string `json:"accessReviewId,omitempty"`
	// AppliedBy undocumented
	AppliedBy *UserIdentity `json:"appliedBy,omitempty"`
	// AppliedDateTime undocumented
	AppliedDateTime *time.Time `json:"appliedDateTime,omitempty"`
	// ApplyResult undocumented
	ApplyResult *string `json:"applyResult,omitempty"`
	// Decision undocumented
	Decision *string `json:"decision,omitempty"`
	// Justification undocumented
	Justification *string `json:"justification,omitempty"`
	// Principal undocumented
	Principal *Identity `json:"principal,omitempty"`
	// PrincipalLink undocumented
	PrincipalLink *string `json:"principalLink,omitempty"`
	// Recommendation undocumented
	Recommendation *string `json:"recommendation,omitempty"`
	// Resource undocumented
	Resource *AccessReviewInstanceDecisionItemResource `json:"resource,omitempty"`
	// ResourceLink undocumented
	ResourceLink *string `json:"resourceLink,omitempty"`
	// ReviewedBy undocumented
	ReviewedBy *UserIdentity `json:"reviewedBy,omitempty"`
	// ReviewedDateTime undocumented
	ReviewedDateTime *time.Time `json:"reviewedDateTime,omitempty"`
	// Target undocumented
	Target *AccessReviewInstanceDecisionItemTarget `json:"target,omitempty"`
}

// AccessReviewInstanceDecisionItemResource undocumented
type AccessReviewInstanceDecisionItemResource struct {
	// Object is the base model of AccessReviewInstanceDecisionItemResource
	Object
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// ID undocumented
	ID *string `json:"id,omitempty"`
	// Type undocumented
	Type *string `json:"type,omitempty"`
}

// AccessReviewInstanceDecisionItemServicePrincipalTarget undocumented
type AccessReviewInstanceDecisionItemServicePrincipalTarget struct {
	// AccessReviewInstanceDecisionItemTarget is the base model of AccessReviewInstanceDecisionItemServicePrincipalTarget
	AccessReviewInstanceDecisionItemTarget
	// AppID undocumented
	AppID *string `json:"appId,omitempty"`
	// ServicePrincipalDisplayName undocumented
	ServicePrincipalDisplayName *string `json:"servicePrincipalDisplayName,omitempty"`
	// ServicePrincipalID undocumented
	ServicePrincipalID *string `json:"servicePrincipalId,omitempty"`
}

// AccessReviewInstanceDecisionItemTarget undocumented
type AccessReviewInstanceDecisionItemTarget struct {
	// Object is the base model of AccessReviewInstanceDecisionItemTarget
	Object
}

// AccessReviewInstanceDecisionItemUserTarget undocumented
type AccessReviewInstanceDecisionItemUserTarget struct {
	// AccessReviewInstanceDecisionItemTarget is the base model of AccessReviewInstanceDecisionItemUserTarget
	AccessReviewInstanceDecisionItemTarget
	// UserDisplayName undocumented
	UserDisplayName *string `json:"userDisplayName,omitempty"`
	// UserID undocumented
	UserID *string `json:"userId,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

// AccessReviewNotificationRecipientItem undocumented
type AccessReviewNotificationRecipientItem struct {
	// Object is the base model of AccessReviewNotificationRecipientItem
	Object
	// NotificationRecipientScope undocumented
	NotificationRecipientScope *AccessReviewNotificationRecipientScope `json:"notificationRecipientScope,omitempty"`
	// NotificationTemplateType undocumented
	NotificationTemplateType *string `json:"notificationTemplateType,omitempty"`
}

// AccessReviewNotificationRecipientQueryScope undocumented
type AccessReviewNotificationRecipientQueryScope struct {
	// AccessReviewNotificationRecipientScope is the base model of AccessReviewNotificationRecipientQueryScope
	AccessReviewNotificationRecipientScope
	// Query undocumented
	Query *string `json:"query,omitempty"`
	// QueryRoot undocumented
	QueryRoot *string `json:"queryRoot,omitempty"`
	// QueryType undocumented
	QueryType *string `json:"queryType,omitempty"`
}

// AccessReviewNotificationRecipientScope undocumented
type AccessReviewNotificationRecipientScope struct {
	// Object is the base model of AccessReviewNotificationRecipientScope
	Object
}

// AccessReviewPolicy undocumented
type AccessReviewPolicy struct {
	// Entity is the base model of AccessReviewPolicy
	Entity
	// Description undocumented
	Description *string `json:"description,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// IsGroupOwnerManagementEnabled undocumented
	IsGroupOwnerManagementEnabled *bool `json:"isGroupOwnerManagementEnabled,omitempty"`
}

// AccessReviewQueryScope undocumented
type AccessReviewQueryScope struct {
	// AccessReviewScope is the base model of AccessReviewQueryScope
	AccessReviewScope
	// Query undocumented
	Query *string `json:"query,omitempty"`
	// QueryRoot undocumented
	QueryRoot *string `json:"queryRoot,omitempty"`
	// QueryType undocumented
	QueryType *string `json:"queryType,omitempty"`
}

// AccessReviewRecurrenceSettings undocumented
type AccessReviewRecurrenceSettings struct {
	// Object is the base model of AccessReviewRecurrenceSettings
	Object
	// DurationInDays undocumented
	DurationInDays *int `json:"durationInDays,omitempty"`
	// RecurrenceCount undocumented
	RecurrenceCount *int `json:"recurrenceCount,omitempty"`
	// RecurrenceEndType undocumented
	RecurrenceEndType *string `json:"recurrenceEndType,omitempty"`
	// RecurrenceType undocumented
	RecurrenceType *string `json:"recurrenceType,omitempty"`
}

// AccessReviewReviewer undocumented
type AccessReviewReviewer struct {
	// Entity is the base model of AccessReviewReviewer
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// UserPrincipalName undocumented
	UserPrincipalName *string `json:"userPrincipalName,omitempty"`
}

// AccessReviewReviewerScope undocumented
type AccessReviewReviewerScope struct {
	// AccessReviewScope is the base model of AccessReviewReviewerScope
	AccessReviewScope
	// Query undocumented
	Query *string `json:"query,omitempty"`
	// QueryRoot undocumented
	QueryRoot *string `json:"queryRoot,omitempty"`
	// QueryType undocumented
	QueryType *string `json:"queryType,omitempty"`
}

// AccessReviewScheduleDefinition undocumented
type AccessReviewScheduleDefinition struct {
	// Entity is the base model of AccessReviewScheduleDefinition
	Entity
	// AdditionalNotificationRecipients undocumented
	AdditionalNotificationRecipients []AccessReviewNotificationRecipientItem `json:"additionalNotificationRecipients,omitempty"`
	// BackupReviewers undocumented
	BackupReviewers []AccessReviewReviewerScope `json:"backupReviewers,omitempty"`
	// CreatedBy undocumented
	CreatedBy *UserIdentity `json:"createdBy,omitempty"`
	// CreatedDateTime undocumented
	CreatedDateTime *time.Time `json:"createdDateTime,omitempty"`
	// DescriptionForAdmins undocumented
	DescriptionForAdmins *string `json:"descriptionForAdmins,omitempty"`
	// DescriptionForReviewers undocumented
	DescriptionForReviewers *string `json:"descriptionForReviewers,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// FallbackReviewers undocumented
	FallbackReviewers []AccessReviewReviewerScope `json:"fallbackReviewers,omitempty"`
	// InstanceEnumerationScope undocumented
	InstanceEnumerationScope *AccessReviewScope `json:"instanceEnumerationScope,omitempty"`
	// LastModifiedDateTime undocumented
	LastModifiedDateTime *time.Time `json:"lastModifiedDateTime,omitempty"`
	// Reviewers undocumented
	Reviewers []AccessReviewReviewerScope `json:"reviewers,omitempty"`
	// Scope undocumented
	Scope *AccessReviewScope `json:"scope,omitempty"`
	// Settings undocumented
	Settings *AccessReviewScheduleSettings `json:"settings,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
	// Instances undocumented
	Instances []AccessReviewInstance `json:"instances,omitempty"`
}

// AccessReviewScheduleSettings undocumented
type AccessReviewScheduleSettings struct {
	// Object is the base model of AccessReviewScheduleSettings
	Object
	// ApplyActions undocumented
	ApplyActions []AccessReviewApplyAction `json:"applyActions,omitempty"`
	// AutoApplyDecisionsEnabled undocumented
	AutoApplyDecisionsEnabled *bool `json:"autoApplyDecisionsEnabled,omitempty"`
	// DefaultDecision undocumented
	DefaultDecision *string `json:"defaultDecision,omitempty"`
	// DefaultDecisionEnabled undocumented
	DefaultDecisionEnabled *bool `json:"defaultDecisionEnabled,omitempty"`
	// InstanceDurationInDays undocumented
	InstanceDurationInDays *int `json:"instanceDurationInDays,omitempty"`
	// JustificationRequiredOnApproval undocumented
	JustificationRequiredOnApproval *bool `json:"justificationRequiredOnApproval,omitempty"`
	// MailNotificationsEnabled undocumented
	MailNotificationsEnabled *bool `json:"mailNotificationsEnabled,omitempty"`
	// RecommendationLookBackDuration undocumented
	RecommendationLookBackDuration *Duration `json:"recommendationLookBackDuration,omitempty"`
	// RecommendationsEnabled undocumented
	RecommendationsEnabled *bool `json:"recommendationsEnabled,omitempty"`
	// Recurrence undocumented
	Recurrence *PatternedRecurrence `json:"recurrence,omitempty"`
	// ReminderNotificationsEnabled undocumented
	ReminderNotificationsEnabled *bool `json:"reminderNotificationsEnabled,omitempty"`
}

// AccessReviewScope undocumented
type AccessReviewScope struct {
	// Object is the base model of AccessReviewScope
	Object
}

// AccessReviewSet undocumented
type AccessReviewSet struct {
	// Entity is the base model of AccessReviewSet
	Entity
	// Definitions undocumented
	Definitions []AccessReviewScheduleDefinition `json:"definitions,omitempty"`
	// HistoryDefinitions undocumented
	HistoryDefinitions []AccessReviewHistoryDefinition `json:"historyDefinitions,omitempty"`
	// Policy undocumented
	Policy *AccessReviewPolicy `json:"policy,omitempty"`
}

// AccessReviewSettings undocumented
type AccessReviewSettings struct {
	// Object is the base model of AccessReviewSettings
	Object
	// AccessRecommendationsEnabled undocumented
	AccessRecommendationsEnabled *bool `json:"accessRecommendationsEnabled,omitempty"`
	// ActivityDurationInDays undocumented
	ActivityDurationInDays *int `json:"activityDurationInDays,omitempty"`
	// AutoApplyReviewResultsEnabled undocumented
	AutoApplyReviewResultsEnabled *bool `json:"autoApplyReviewResultsEnabled,omitempty"`
	// AutoReviewEnabled undocumented
	AutoReviewEnabled *bool `json:"autoReviewEnabled,omitempty"`
	// AutoReviewSettings undocumented
	AutoReviewSettings *AutoReviewSettings `json:"autoReviewSettings,omitempty"`
	// JustificationRequiredOnApproval undocumented
	JustificationRequiredOnApproval *bool `json:"justificationRequiredOnApproval,omitempty"`
	// MailNotificationsEnabled undocumented
	MailNotificationsEnabled *bool `json:"mailNotificationsEnabled,omitempty"`
	// RecurrenceSettings undocumented
	RecurrenceSettings *AccessReviewRecurrenceSettings `json:"recurrenceSettings,omitempty"`
	// RemindersEnabled undocumented
	RemindersEnabled *bool `json:"remindersEnabled,omitempty"`
}
