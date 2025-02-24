// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// AssignmentFilterEvaluationStatusDetails A class containing information about the payloads on which filter has been applied.
type AssignmentFilterEvaluationStatusDetails struct {
	// Entity is the base model of AssignmentFilterEvaluationStatusDetails
	Entity
	// PayloadID PayloadId on which filter has been applied.
	PayloadID *string `json:"payloadId,omitempty"`
}

// AssignmentFilterEvaluationSummary Represent result summary for assignment filter evaluation
type AssignmentFilterEvaluationSummary struct {
	// Object is the base model of AssignmentFilterEvaluationSummary
	Object
	// AssignmentFilterDisplayName The admin defined name for assignment filter.
	AssignmentFilterDisplayName *string `json:"assignmentFilterDisplayName,omitempty"`
	// AssignmentFilterID Unique identifier for the assignment filter object
	AssignmentFilterID *string `json:"assignmentFilterId,omitempty"`
	// AssignmentFilterLastModifiedDateTime The time the assignment filter was last modified.
	AssignmentFilterLastModifiedDateTime *time.Time `json:"assignmentFilterLastModifiedDateTime,omitempty"`
	// AssignmentFilterPlatform The platform for which this assignment filter is created.
	AssignmentFilterPlatform *DevicePlatformType `json:"assignmentFilterPlatform,omitempty"`
	// AssignmentFilterType Indicate filter type either include or exclude.
	AssignmentFilterType *DeviceAndAppManagementAssignmentFilterType `json:"assignmentFilterType,omitempty"`
	// AssignmentFilterTypeAndEvaluationResults A collection of filter types and their corresponding evaluation results.
	AssignmentFilterTypeAndEvaluationResults []AssignmentFilterTypeAndEvaluationResult `json:"assignmentFilterTypeAndEvaluationResults,omitempty"`
	// EvaluationDateTime The time assignment filter was evaluated.
	EvaluationDateTime *time.Time `json:"evaluationDateTime,omitempty"`
	// EvaluationResult Assignment filter evaluation result.
	EvaluationResult *AssignmentFilterEvaluationResult `json:"evaluationResult,omitempty"`
}

// AssignmentFilterState Represents result of GetState API.
type AssignmentFilterState struct {
	// Object is the base model of AssignmentFilterState
	Object
	// Enabled Indicator to if AssignmentFilter is enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`
}

// AssignmentFilterStatusDetails Represent status details for device and payload and all associated applied filters.
type AssignmentFilterStatusDetails struct {
	// Object is the base model of AssignmentFilterStatusDetails
	Object
	// DeviceProperties Device properties used for filter evaluation during device check-in time.
	DeviceProperties []KeyValuePair `json:"deviceProperties,omitempty"`
	// EvalutionSummaries Evaluation result summaries for each filter associated to device and payload
	EvalutionSummaries []AssignmentFilterEvaluationSummary `json:"evalutionSummaries,omitempty"`
	// ManagedDeviceID Unique identifier for the device object.
	ManagedDeviceID *string `json:"managedDeviceId,omitempty"`
	// PayloadID Unique identifier for payload object.
	PayloadID *string `json:"payloadId,omitempty"`
	// UserID Unique identifier for UserId object. Can be null
	UserID *string `json:"userId,omitempty"`
}

// AssignmentFilterTypeAndEvaluationResult Represents the filter type and evalaution result of the filter.
type AssignmentFilterTypeAndEvaluationResult struct {
	// Object is the base model of AssignmentFilterTypeAndEvaluationResult
	Object
	// AssignmentFilterType Represents the filter type.
	AssignmentFilterType *DeviceAndAppManagementAssignmentFilterType `json:"assignmentFilterType,omitempty"`
	// EvaluationResult Represents the evalaution result of the filter.
	EvaluationResult *AssignmentFilterEvaluationResult `json:"evaluationResult,omitempty"`
}

// AssignmentFilterValidationResult Represents result of Validation API.
type AssignmentFilterValidationResult struct {
	// Object is the base model of AssignmentFilterValidationResult
	Object
	// IsValidRule Indicator to valid or invalid rule.
	IsValidRule *bool `json:"isValidRule,omitempty"`
}

// AssignmentOrder undocumented
type AssignmentOrder struct {
	// Object is the base model of AssignmentOrder
	Object
	// Order undocumented
	Order []string `json:"order,omitempty"`
}

// AssignmentReviewSettings undocumented
type AssignmentReviewSettings struct {
	// Object is the base model of AssignmentReviewSettings
	Object
	// DurationInDays undocumented
	DurationInDays *int `json:"durationInDays,omitempty"`
	// IsEnabled undocumented
	IsEnabled *bool `json:"isEnabled,omitempty"`
	// RecurrenceType undocumented
	RecurrenceType *string `json:"recurrenceType,omitempty"`
	// Reviewers undocumented
	Reviewers []UserSet `json:"reviewers,omitempty"`
	// ReviewerType undocumented
	ReviewerType *string `json:"reviewerType,omitempty"`
	// StartDateTime undocumented
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
}
