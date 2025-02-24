// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ChatMessageImportance undocumented
type ChatMessageImportance string

const (
	// ChatMessageImportanceVNormal undocumented
	ChatMessageImportanceVNormal ChatMessageImportance = "normal"
	// ChatMessageImportanceVHigh undocumented
	ChatMessageImportanceVHigh ChatMessageImportance = "high"
	// ChatMessageImportanceVUrgent undocumented
	ChatMessageImportanceVUrgent ChatMessageImportance = "urgent"
)

var (
	// ChatMessageImportancePNormal is a pointer to ChatMessageImportanceVNormal
	ChatMessageImportancePNormal = &_ChatMessageImportancePNormal
	// ChatMessageImportancePHigh is a pointer to ChatMessageImportanceVHigh
	ChatMessageImportancePHigh = &_ChatMessageImportancePHigh
	// ChatMessageImportancePUrgent is a pointer to ChatMessageImportanceVUrgent
	ChatMessageImportancePUrgent = &_ChatMessageImportancePUrgent
)

var (
	_ChatMessageImportancePNormal = ChatMessageImportanceVNormal
	_ChatMessageImportancePHigh   = ChatMessageImportanceVHigh
	_ChatMessageImportancePUrgent = ChatMessageImportanceVUrgent
)

// ChatMessagePolicyViolationDlpActionTypes undocumented
type ChatMessagePolicyViolationDlpActionTypes string

const (
	// ChatMessagePolicyViolationDlpActionTypesVNone undocumented
	ChatMessagePolicyViolationDlpActionTypesVNone ChatMessagePolicyViolationDlpActionTypes = "none"
	// ChatMessagePolicyViolationDlpActionTypesVNotifySender undocumented
	ChatMessagePolicyViolationDlpActionTypesVNotifySender ChatMessagePolicyViolationDlpActionTypes = "notifySender"
	// ChatMessagePolicyViolationDlpActionTypesVBlockAccess undocumented
	ChatMessagePolicyViolationDlpActionTypesVBlockAccess ChatMessagePolicyViolationDlpActionTypes = "blockAccess"
	// ChatMessagePolicyViolationDlpActionTypesVBlockAccessExternal undocumented
	ChatMessagePolicyViolationDlpActionTypesVBlockAccessExternal ChatMessagePolicyViolationDlpActionTypes = "blockAccessExternal"
)

var (
	// ChatMessagePolicyViolationDlpActionTypesPNone is a pointer to ChatMessagePolicyViolationDlpActionTypesVNone
	ChatMessagePolicyViolationDlpActionTypesPNone = &_ChatMessagePolicyViolationDlpActionTypesPNone
	// ChatMessagePolicyViolationDlpActionTypesPNotifySender is a pointer to ChatMessagePolicyViolationDlpActionTypesVNotifySender
	ChatMessagePolicyViolationDlpActionTypesPNotifySender = &_ChatMessagePolicyViolationDlpActionTypesPNotifySender
	// ChatMessagePolicyViolationDlpActionTypesPBlockAccess is a pointer to ChatMessagePolicyViolationDlpActionTypesVBlockAccess
	ChatMessagePolicyViolationDlpActionTypesPBlockAccess = &_ChatMessagePolicyViolationDlpActionTypesPBlockAccess
	// ChatMessagePolicyViolationDlpActionTypesPBlockAccessExternal is a pointer to ChatMessagePolicyViolationDlpActionTypesVBlockAccessExternal
	ChatMessagePolicyViolationDlpActionTypesPBlockAccessExternal = &_ChatMessagePolicyViolationDlpActionTypesPBlockAccessExternal
)

var (
	_ChatMessagePolicyViolationDlpActionTypesPNone                = ChatMessagePolicyViolationDlpActionTypesVNone
	_ChatMessagePolicyViolationDlpActionTypesPNotifySender        = ChatMessagePolicyViolationDlpActionTypesVNotifySender
	_ChatMessagePolicyViolationDlpActionTypesPBlockAccess         = ChatMessagePolicyViolationDlpActionTypesVBlockAccess
	_ChatMessagePolicyViolationDlpActionTypesPBlockAccessExternal = ChatMessagePolicyViolationDlpActionTypesVBlockAccessExternal
)

// ChatMessagePolicyViolationUserActionTypes undocumented
type ChatMessagePolicyViolationUserActionTypes string

const (
	// ChatMessagePolicyViolationUserActionTypesVNone undocumented
	ChatMessagePolicyViolationUserActionTypesVNone ChatMessagePolicyViolationUserActionTypes = "none"
	// ChatMessagePolicyViolationUserActionTypesVOverride undocumented
	ChatMessagePolicyViolationUserActionTypesVOverride ChatMessagePolicyViolationUserActionTypes = "override"
	// ChatMessagePolicyViolationUserActionTypesVReportFalsePositive undocumented
	ChatMessagePolicyViolationUserActionTypesVReportFalsePositive ChatMessagePolicyViolationUserActionTypes = "reportFalsePositive"
)

var (
	// ChatMessagePolicyViolationUserActionTypesPNone is a pointer to ChatMessagePolicyViolationUserActionTypesVNone
	ChatMessagePolicyViolationUserActionTypesPNone = &_ChatMessagePolicyViolationUserActionTypesPNone
	// ChatMessagePolicyViolationUserActionTypesPOverride is a pointer to ChatMessagePolicyViolationUserActionTypesVOverride
	ChatMessagePolicyViolationUserActionTypesPOverride = &_ChatMessagePolicyViolationUserActionTypesPOverride
	// ChatMessagePolicyViolationUserActionTypesPReportFalsePositive is a pointer to ChatMessagePolicyViolationUserActionTypesVReportFalsePositive
	ChatMessagePolicyViolationUserActionTypesPReportFalsePositive = &_ChatMessagePolicyViolationUserActionTypesPReportFalsePositive
)

var (
	_ChatMessagePolicyViolationUserActionTypesPNone                = ChatMessagePolicyViolationUserActionTypesVNone
	_ChatMessagePolicyViolationUserActionTypesPOverride            = ChatMessagePolicyViolationUserActionTypesVOverride
	_ChatMessagePolicyViolationUserActionTypesPReportFalsePositive = ChatMessagePolicyViolationUserActionTypesVReportFalsePositive
)

// ChatMessagePolicyViolationVerdictDetailsTypes undocumented
type ChatMessagePolicyViolationVerdictDetailsTypes string

const (
	// ChatMessagePolicyViolationVerdictDetailsTypesVNone undocumented
	ChatMessagePolicyViolationVerdictDetailsTypesVNone ChatMessagePolicyViolationVerdictDetailsTypes = "none"
	// ChatMessagePolicyViolationVerdictDetailsTypesVAllowFalsePositiveOverride undocumented
	ChatMessagePolicyViolationVerdictDetailsTypesVAllowFalsePositiveOverride ChatMessagePolicyViolationVerdictDetailsTypes = "allowFalsePositiveOverride"
	// ChatMessagePolicyViolationVerdictDetailsTypesVAllowOverrideWithoutJustification undocumented
	ChatMessagePolicyViolationVerdictDetailsTypesVAllowOverrideWithoutJustification ChatMessagePolicyViolationVerdictDetailsTypes = "allowOverrideWithoutJustification"
	// ChatMessagePolicyViolationVerdictDetailsTypesVAllowOverrideWithJustification undocumented
	ChatMessagePolicyViolationVerdictDetailsTypesVAllowOverrideWithJustification ChatMessagePolicyViolationVerdictDetailsTypes = "allowOverrideWithJustification"
)

var (
	// ChatMessagePolicyViolationVerdictDetailsTypesPNone is a pointer to ChatMessagePolicyViolationVerdictDetailsTypesVNone
	ChatMessagePolicyViolationVerdictDetailsTypesPNone = &_ChatMessagePolicyViolationVerdictDetailsTypesPNone
	// ChatMessagePolicyViolationVerdictDetailsTypesPAllowFalsePositiveOverride is a pointer to ChatMessagePolicyViolationVerdictDetailsTypesVAllowFalsePositiveOverride
	ChatMessagePolicyViolationVerdictDetailsTypesPAllowFalsePositiveOverride = &_ChatMessagePolicyViolationVerdictDetailsTypesPAllowFalsePositiveOverride
	// ChatMessagePolicyViolationVerdictDetailsTypesPAllowOverrideWithoutJustification is a pointer to ChatMessagePolicyViolationVerdictDetailsTypesVAllowOverrideWithoutJustification
	ChatMessagePolicyViolationVerdictDetailsTypesPAllowOverrideWithoutJustification = &_ChatMessagePolicyViolationVerdictDetailsTypesPAllowOverrideWithoutJustification
	// ChatMessagePolicyViolationVerdictDetailsTypesPAllowOverrideWithJustification is a pointer to ChatMessagePolicyViolationVerdictDetailsTypesVAllowOverrideWithJustification
	ChatMessagePolicyViolationVerdictDetailsTypesPAllowOverrideWithJustification = &_ChatMessagePolicyViolationVerdictDetailsTypesPAllowOverrideWithJustification
)

var (
	_ChatMessagePolicyViolationVerdictDetailsTypesPNone                              = ChatMessagePolicyViolationVerdictDetailsTypesVNone
	_ChatMessagePolicyViolationVerdictDetailsTypesPAllowFalsePositiveOverride        = ChatMessagePolicyViolationVerdictDetailsTypesVAllowFalsePositiveOverride
	_ChatMessagePolicyViolationVerdictDetailsTypesPAllowOverrideWithoutJustification = ChatMessagePolicyViolationVerdictDetailsTypesVAllowOverrideWithoutJustification
	_ChatMessagePolicyViolationVerdictDetailsTypesPAllowOverrideWithJustification    = ChatMessagePolicyViolationVerdictDetailsTypesVAllowOverrideWithJustification
)

// ChatMessageType undocumented
type ChatMessageType string

const (
	// ChatMessageTypeVMessage undocumented
	ChatMessageTypeVMessage ChatMessageType = "message"
	// ChatMessageTypeVChatEvent undocumented
	ChatMessageTypeVChatEvent ChatMessageType = "chatEvent"
	// ChatMessageTypeVTyping undocumented
	ChatMessageTypeVTyping ChatMessageType = "typing"
	// ChatMessageTypeVUnknownFutureValue undocumented
	ChatMessageTypeVUnknownFutureValue ChatMessageType = "unknownFutureValue"
	// ChatMessageTypeVSystemEventMessage undocumented
	ChatMessageTypeVSystemEventMessage ChatMessageType = "systemEventMessage"
)

var (
	// ChatMessageTypePMessage is a pointer to ChatMessageTypeVMessage
	ChatMessageTypePMessage = &_ChatMessageTypePMessage
	// ChatMessageTypePChatEvent is a pointer to ChatMessageTypeVChatEvent
	ChatMessageTypePChatEvent = &_ChatMessageTypePChatEvent
	// ChatMessageTypePTyping is a pointer to ChatMessageTypeVTyping
	ChatMessageTypePTyping = &_ChatMessageTypePTyping
	// ChatMessageTypePUnknownFutureValue is a pointer to ChatMessageTypeVUnknownFutureValue
	ChatMessageTypePUnknownFutureValue = &_ChatMessageTypePUnknownFutureValue
	// ChatMessageTypePSystemEventMessage is a pointer to ChatMessageTypeVSystemEventMessage
	ChatMessageTypePSystemEventMessage = &_ChatMessageTypePSystemEventMessage
)

var (
	_ChatMessageTypePMessage            = ChatMessageTypeVMessage
	_ChatMessageTypePChatEvent          = ChatMessageTypeVChatEvent
	_ChatMessageTypePTyping             = ChatMessageTypeVTyping
	_ChatMessageTypePUnknownFutureValue = ChatMessageTypeVUnknownFutureValue
	_ChatMessageTypePSystemEventMessage = ChatMessageTypeVSystemEventMessage
)

// ChatType undocumented
type ChatType string

const (
	// ChatTypeVOneOnOne undocumented
	ChatTypeVOneOnOne ChatType = "oneOnOne"
	// ChatTypeVGroup undocumented
	ChatTypeVGroup ChatType = "group"
	// ChatTypeVMeeting undocumented
	ChatTypeVMeeting ChatType = "meeting"
	// ChatTypeVUnknownFutureValue undocumented
	ChatTypeVUnknownFutureValue ChatType = "unknownFutureValue"
)

var (
	// ChatTypePOneOnOne is a pointer to ChatTypeVOneOnOne
	ChatTypePOneOnOne = &_ChatTypePOneOnOne
	// ChatTypePGroup is a pointer to ChatTypeVGroup
	ChatTypePGroup = &_ChatTypePGroup
	// ChatTypePMeeting is a pointer to ChatTypeVMeeting
	ChatTypePMeeting = &_ChatTypePMeeting
	// ChatTypePUnknownFutureValue is a pointer to ChatTypeVUnknownFutureValue
	ChatTypePUnknownFutureValue = &_ChatTypePUnknownFutureValue
)

var (
	_ChatTypePOneOnOne           = ChatTypeVOneOnOne
	_ChatTypePGroup              = ChatTypeVGroup
	_ChatTypePMeeting            = ChatTypeVMeeting
	_ChatTypePUnknownFutureValue = ChatTypeVUnknownFutureValue
)
