// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "time"

// Approval undocumented
type Approval struct {
	// Entity is the base model of Approval
	Entity
	// RequestNavigation undocumented
	RequestNavigation *RequestObject `json:"request,omitempty"`
	// Steps undocumented
	Steps []ApprovalStep `json:"steps,omitempty"`
}

// ApprovalSettings undocumented
type ApprovalSettings struct {
	// Object is the base model of ApprovalSettings
	Object
	// ApprovalMode undocumented
	ApprovalMode *string `json:"approvalMode,omitempty"`
	// ApprovalStages undocumented
	ApprovalStages []ApprovalStage `json:"approvalStages,omitempty"`
	// IsApprovalRequired undocumented
	IsApprovalRequired *bool `json:"isApprovalRequired,omitempty"`
	// IsApprovalRequiredForExtension undocumented
	IsApprovalRequiredForExtension *bool `json:"isApprovalRequiredForExtension,omitempty"`
	// IsRequestorJustificationRequired undocumented
	IsRequestorJustificationRequired *bool `json:"isRequestorJustificationRequired,omitempty"`
}

// ApprovalStage undocumented
type ApprovalStage struct {
	// Object is the base model of ApprovalStage
	Object
	// ApprovalStageTimeOutInDays undocumented
	ApprovalStageTimeOutInDays *int `json:"approvalStageTimeOutInDays,omitempty"`
	// EscalationApprovers undocumented
	EscalationApprovers []UserSet `json:"escalationApprovers,omitempty"`
	// EscalationTimeInMinutes undocumented
	EscalationTimeInMinutes *int `json:"escalationTimeInMinutes,omitempty"`
	// IsApproverJustificationRequired undocumented
	IsApproverJustificationRequired *bool `json:"isApproverJustificationRequired,omitempty"`
	// IsEscalationEnabled undocumented
	IsEscalationEnabled *bool `json:"isEscalationEnabled,omitempty"`
	// PrimaryApprovers undocumented
	PrimaryApprovers []UserSet `json:"primaryApprovers,omitempty"`
}

// ApprovalStep undocumented
type ApprovalStep struct {
	// Entity is the base model of ApprovalStep
	Entity
	// AssignedToMe undocumented
	AssignedToMe *bool `json:"assignedToMe,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Justification undocumented
	Justification *string `json:"justification,omitempty"`
	// ReviewedBy undocumented
	ReviewedBy *Identity `json:"reviewedBy,omitempty"`
	// ReviewedDateTime undocumented
	ReviewedDateTime *time.Time `json:"reviewedDateTime,omitempty"`
	// ReviewResult undocumented
	ReviewResult *string `json:"reviewResult,omitempty"`
	// Status undocumented
	Status *string `json:"status,omitempty"`
}

// ApprovalWorkflowProvider undocumented
type ApprovalWorkflowProvider struct {
	// Entity is the base model of ApprovalWorkflowProvider
	Entity
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// BusinessFlows undocumented
	BusinessFlows []BusinessFlow `json:"businessFlows,omitempty"`
	// BusinessFlowsWithRequestsAwaitingMyDecision undocumented
	BusinessFlowsWithRequestsAwaitingMyDecision []BusinessFlow `json:"businessFlowsWithRequestsAwaitingMyDecision,omitempty"`
	// PolicyTemplates undocumented
	PolicyTemplates []GovernancePolicyTemplate `json:"policyTemplates,omitempty"`
}
