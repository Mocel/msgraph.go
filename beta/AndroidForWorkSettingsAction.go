// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidForWorkSettingsRequestSignupURLRequestParameter undocumented
type AndroidForWorkSettingsRequestSignupURLRequestParameter struct {
	// HostName undocumented
	HostName *string `json:"hostName,omitempty"`
}

// AndroidForWorkSettingsCompleteSignupRequestParameter undocumented
type AndroidForWorkSettingsCompleteSignupRequestParameter struct {
	// EnterpriseToken undocumented
	EnterpriseToken *string `json:"enterpriseToken,omitempty"`
}

// AndroidForWorkSettingsSyncAppsRequestParameter undocumented
type AndroidForWorkSettingsSyncAppsRequestParameter struct {
}

// AndroidForWorkSettingsUnbindRequestParameter undocumented
type AndroidForWorkSettingsUnbindRequestParameter struct {
}

//
type AndroidForWorkSettingsRequestSignupURLRequestBuilder struct{ BaseRequestBuilder }

// RequestSignupURL action undocumented
func (b *AndroidForWorkSettingsRequestBuilder) RequestSignupURL(reqObj *AndroidForWorkSettingsRequestSignupURLRequestParameter) *AndroidForWorkSettingsRequestSignupURLRequestBuilder {
	bb := &AndroidForWorkSettingsRequestSignupURLRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/requestSignupUrl"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AndroidForWorkSettingsRequestSignupURLRequest struct{ BaseRequest }

//
func (b *AndroidForWorkSettingsRequestSignupURLRequestBuilder) Request() *AndroidForWorkSettingsRequestSignupURLRequest {
	return &AndroidForWorkSettingsRequestSignupURLRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AndroidForWorkSettingsRequestSignupURLRequest) Post() (resObj *string, err error) {
	err = r.JSONRequest("POST", "", r.requestObject, &resObj)
	return
}

//
type AndroidForWorkSettingsCompleteSignupRequestBuilder struct{ BaseRequestBuilder }

// CompleteSignup action undocumented
func (b *AndroidForWorkSettingsRequestBuilder) CompleteSignup(reqObj *AndroidForWorkSettingsCompleteSignupRequestParameter) *AndroidForWorkSettingsCompleteSignupRequestBuilder {
	bb := &AndroidForWorkSettingsCompleteSignupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/completeSignup"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AndroidForWorkSettingsCompleteSignupRequest struct{ BaseRequest }

//
func (b *AndroidForWorkSettingsCompleteSignupRequestBuilder) Request() *AndroidForWorkSettingsCompleteSignupRequest {
	return &AndroidForWorkSettingsCompleteSignupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AndroidForWorkSettingsCompleteSignupRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type AndroidForWorkSettingsSyncAppsRequestBuilder struct{ BaseRequestBuilder }

// SyncApps action undocumented
func (b *AndroidForWorkSettingsRequestBuilder) SyncApps(reqObj *AndroidForWorkSettingsSyncAppsRequestParameter) *AndroidForWorkSettingsSyncAppsRequestBuilder {
	bb := &AndroidForWorkSettingsSyncAppsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/syncApps"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AndroidForWorkSettingsSyncAppsRequest struct{ BaseRequest }

//
func (b *AndroidForWorkSettingsSyncAppsRequestBuilder) Request() *AndroidForWorkSettingsSyncAppsRequest {
	return &AndroidForWorkSettingsSyncAppsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AndroidForWorkSettingsSyncAppsRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type AndroidForWorkSettingsUnbindRequestBuilder struct{ BaseRequestBuilder }

// Unbind action undocumented
func (b *AndroidForWorkSettingsRequestBuilder) Unbind(reqObj *AndroidForWorkSettingsUnbindRequestParameter) *AndroidForWorkSettingsUnbindRequestBuilder {
	bb := &AndroidForWorkSettingsUnbindRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/unbind"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AndroidForWorkSettingsUnbindRequest struct{ BaseRequest }

//
func (b *AndroidForWorkSettingsUnbindRequestBuilder) Request() *AndroidForWorkSettingsUnbindRequest {
	return &AndroidForWorkSettingsUnbindRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AndroidForWorkSettingsUnbindRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}
