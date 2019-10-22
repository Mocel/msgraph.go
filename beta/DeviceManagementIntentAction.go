// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementIntentUpdateSettingsRequestParameter undocumented
type DeviceManagementIntentUpdateSettingsRequestParameter struct {
	// Settings undocumented
	Settings []DeviceManagementSettingInstance `json:"settings,omitempty"`
}

// DeviceManagementIntentMigrateToTemplateRequestParameter undocumented
type DeviceManagementIntentMigrateToTemplateRequestParameter struct {
	// NewTemplateID undocumented
	NewTemplateID *string `json:"newTemplateId,omitempty"`
	// PreserveCustomValues undocumented
	PreserveCustomValues *bool `json:"preserveCustomValues,omitempty"`
}

// DeviceManagementIntentAssignRequestParameter undocumented
type DeviceManagementIntentAssignRequestParameter struct {
	// Assignments undocumented
	Assignments []DeviceManagementIntentAssignment `json:"assignments,omitempty"`
}

//
type DeviceManagementIntentUpdateSettingsRequestBuilder struct{ BaseRequestBuilder }

// UpdateSettings action undocumented
func (b *DeviceManagementIntentRequestBuilder) UpdateSettings(reqObj *DeviceManagementIntentUpdateSettingsRequestParameter) *DeviceManagementIntentUpdateSettingsRequestBuilder {
	bb := &DeviceManagementIntentUpdateSettingsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/updateSettings"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceManagementIntentUpdateSettingsRequest struct{ BaseRequest }

//
func (b *DeviceManagementIntentUpdateSettingsRequestBuilder) Request() *DeviceManagementIntentUpdateSettingsRequest {
	return &DeviceManagementIntentUpdateSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceManagementIntentUpdateSettingsRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type DeviceManagementIntentMigrateToTemplateRequestBuilder struct{ BaseRequestBuilder }

// MigrateToTemplate action undocumented
func (b *DeviceManagementIntentRequestBuilder) MigrateToTemplate(reqObj *DeviceManagementIntentMigrateToTemplateRequestParameter) *DeviceManagementIntentMigrateToTemplateRequestBuilder {
	bb := &DeviceManagementIntentMigrateToTemplateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/migrateToTemplate"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceManagementIntentMigrateToTemplateRequest struct{ BaseRequest }

//
func (b *DeviceManagementIntentMigrateToTemplateRequestBuilder) Request() *DeviceManagementIntentMigrateToTemplateRequest {
	return &DeviceManagementIntentMigrateToTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceManagementIntentMigrateToTemplateRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type DeviceManagementIntentAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *DeviceManagementIntentRequestBuilder) Assign(reqObj *DeviceManagementIntentAssignRequestParameter) *DeviceManagementIntentAssignRequestBuilder {
	bb := &DeviceManagementIntentAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DeviceManagementIntentAssignRequest struct{ BaseRequest }

//
func (b *DeviceManagementIntentAssignRequestBuilder) Request() *DeviceManagementIntentAssignRequest {
	return &DeviceManagementIntentAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DeviceManagementIntentAssignRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}
