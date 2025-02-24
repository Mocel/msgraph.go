// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// MeetingActivityStatistics undocumented
type MeetingActivityStatistics struct {
	// ActivityStatistics is the base model of MeetingActivityStatistics
	ActivityStatistics
	// AfterHours undocumented
	AfterHours *Duration `json:"afterHours,omitempty"`
	// Conflicting undocumented
	Conflicting *Duration `json:"conflicting,omitempty"`
	// Long undocumented
	Long *Duration `json:"long,omitempty"`
	// Multitasking undocumented
	Multitasking *Duration `json:"multitasking,omitempty"`
	// Organized undocumented
	Organized *Duration `json:"organized,omitempty"`
	// Recurring undocumented
	Recurring *Duration `json:"recurring,omitempty"`
}

// MeetingAttendanceReport undocumented
type MeetingAttendanceReport struct {
	// Entity is the base model of MeetingAttendanceReport
	Entity
	// AttendanceRecords undocumented
	AttendanceRecords []AttendanceRecord `json:"attendanceRecords,omitempty"`
}

// MeetingCapability undocumented
type MeetingCapability struct {
	// Object is the base model of MeetingCapability
	Object
	// AllowAnonymousUsersToDialOut undocumented
	AllowAnonymousUsersToDialOut *bool `json:"allowAnonymousUsersToDialOut,omitempty"`
	// AllowAnonymousUsersToStartMeeting undocumented
	AllowAnonymousUsersToStartMeeting *bool `json:"allowAnonymousUsersToStartMeeting,omitempty"`
	// AutoAdmittedUsers undocumented
	AutoAdmittedUsers *AutoAdmittedUsersType `json:"autoAdmittedUsers,omitempty"`
}

// MeetingInfo undocumented
type MeetingInfo struct {
	// Object is the base model of MeetingInfo
	Object
	// AllowConversationWithoutHost undocumented
	AllowConversationWithoutHost *bool `json:"allowConversationWithoutHost,omitempty"`
}

// MeetingParticipantInfo undocumented
type MeetingParticipantInfo struct {
	// Object is the base model of MeetingParticipantInfo
	Object
	// Identity undocumented
	Identity *IdentitySet `json:"identity,omitempty"`
	// Role undocumented
	Role *OnlineMeetingRole `json:"role,omitempty"`
	// Upn undocumented
	Upn *string `json:"upn,omitempty"`
}

// MeetingParticipants undocumented
type MeetingParticipants struct {
	// Object is the base model of MeetingParticipants
	Object
	// Attendees undocumented
	Attendees []MeetingParticipantInfo `json:"attendees,omitempty"`
	// Contributors undocumented
	Contributors []MeetingParticipantInfo `json:"contributors,omitempty"`
	// Organizer undocumented
	Organizer *MeetingParticipantInfo `json:"organizer,omitempty"`
	// Producers undocumented
	Producers []MeetingParticipantInfo `json:"producers,omitempty"`
}

// MeetingPolicyUpdatedEventMessageDetail undocumented
type MeetingPolicyUpdatedEventMessageDetail struct {
	// EventMessageDetail is the base model of MeetingPolicyUpdatedEventMessageDetail
	EventMessageDetail
	// Initiator undocumented
	Initiator *IdentitySet `json:"initiator,omitempty"`
	// MeetingChatEnabled undocumented
	MeetingChatEnabled *bool `json:"meetingChatEnabled,omitempty"`
	// MeetingChatID undocumented
	MeetingChatID *string `json:"meetingChatId,omitempty"`
}

// MeetingTimeSuggestion undocumented
type MeetingTimeSuggestion struct {
	// Object is the base model of MeetingTimeSuggestion
	Object
	// AttendeeAvailability undocumented
	AttendeeAvailability []AttendeeAvailability `json:"attendeeAvailability,omitempty"`
	// Confidence undocumented
	Confidence *float64 `json:"confidence,omitempty"`
	// Locations undocumented
	Locations []Location `json:"locations,omitempty"`
	// MeetingTimeSlot undocumented
	MeetingTimeSlot *TimeSlot `json:"meetingTimeSlot,omitempty"`
	// Order undocumented
	Order *int `json:"order,omitempty"`
	// OrganizerAvailability undocumented
	OrganizerAvailability *FreeBusyStatus `json:"organizerAvailability,omitempty"`
	// SuggestionReason undocumented
	SuggestionReason *string `json:"suggestionReason,omitempty"`
}

// MeetingTimeSuggestionsResult undocumented
type MeetingTimeSuggestionsResult struct {
	// Object is the base model of MeetingTimeSuggestionsResult
	Object
	// EmptySuggestionsReason undocumented
	EmptySuggestionsReason *string `json:"emptySuggestionsReason,omitempty"`
	// MeetingTimeSuggestions undocumented
	MeetingTimeSuggestions []MeetingTimeSuggestion `json:"meetingTimeSuggestions,omitempty"`
}
