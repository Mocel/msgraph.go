// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// SecurityRequestBuilder is request builder for Security
type SecurityRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecurityRequest
func (b *SecurityRequestBuilder) Request() *SecurityRequest {
	return &SecurityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SecurityRequest is request for Security
type SecurityRequest struct{ BaseRequest }

// Get performs GET request for Security
func (r *SecurityRequest) Get(ctx context.Context) (resObj *Security, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Security
func (r *SecurityRequest) Update(ctx context.Context, reqObj *Security) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Security
func (r *SecurityRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SecurityActionRequestBuilder is request builder for SecurityAction
type SecurityActionRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecurityActionRequest
func (b *SecurityActionRequestBuilder) Request() *SecurityActionRequest {
	return &SecurityActionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SecurityActionRequest is request for SecurityAction
type SecurityActionRequest struct{ BaseRequest }

// Get performs GET request for SecurityAction
func (r *SecurityActionRequest) Get(ctx context.Context) (resObj *SecurityAction, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SecurityAction
func (r *SecurityActionRequest) Update(ctx context.Context, reqObj *SecurityAction) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SecurityAction
func (r *SecurityActionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SecurityBaselineCategoryStateSummaryRequestBuilder is request builder for SecurityBaselineCategoryStateSummary
type SecurityBaselineCategoryStateSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecurityBaselineCategoryStateSummaryRequest
func (b *SecurityBaselineCategoryStateSummaryRequestBuilder) Request() *SecurityBaselineCategoryStateSummaryRequest {
	return &SecurityBaselineCategoryStateSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SecurityBaselineCategoryStateSummaryRequest is request for SecurityBaselineCategoryStateSummary
type SecurityBaselineCategoryStateSummaryRequest struct{ BaseRequest }

// Get performs GET request for SecurityBaselineCategoryStateSummary
func (r *SecurityBaselineCategoryStateSummaryRequest) Get(ctx context.Context) (resObj *SecurityBaselineCategoryStateSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SecurityBaselineCategoryStateSummary
func (r *SecurityBaselineCategoryStateSummaryRequest) Update(ctx context.Context, reqObj *SecurityBaselineCategoryStateSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SecurityBaselineCategoryStateSummary
func (r *SecurityBaselineCategoryStateSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SecurityBaselineDeviceStateRequestBuilder is request builder for SecurityBaselineDeviceState
type SecurityBaselineDeviceStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecurityBaselineDeviceStateRequest
func (b *SecurityBaselineDeviceStateRequestBuilder) Request() *SecurityBaselineDeviceStateRequest {
	return &SecurityBaselineDeviceStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SecurityBaselineDeviceStateRequest is request for SecurityBaselineDeviceState
type SecurityBaselineDeviceStateRequest struct{ BaseRequest }

// Get performs GET request for SecurityBaselineDeviceState
func (r *SecurityBaselineDeviceStateRequest) Get(ctx context.Context) (resObj *SecurityBaselineDeviceState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SecurityBaselineDeviceState
func (r *SecurityBaselineDeviceStateRequest) Update(ctx context.Context, reqObj *SecurityBaselineDeviceState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SecurityBaselineDeviceState
func (r *SecurityBaselineDeviceStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SecurityBaselineSettingStateRequestBuilder is request builder for SecurityBaselineSettingState
type SecurityBaselineSettingStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecurityBaselineSettingStateRequest
func (b *SecurityBaselineSettingStateRequestBuilder) Request() *SecurityBaselineSettingStateRequest {
	return &SecurityBaselineSettingStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SecurityBaselineSettingStateRequest is request for SecurityBaselineSettingState
type SecurityBaselineSettingStateRequest struct{ BaseRequest }

// Get performs GET request for SecurityBaselineSettingState
func (r *SecurityBaselineSettingStateRequest) Get(ctx context.Context) (resObj *SecurityBaselineSettingState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SecurityBaselineSettingState
func (r *SecurityBaselineSettingStateRequest) Update(ctx context.Context, reqObj *SecurityBaselineSettingState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SecurityBaselineSettingState
func (r *SecurityBaselineSettingStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SecurityBaselineStateRequestBuilder is request builder for SecurityBaselineState
type SecurityBaselineStateRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecurityBaselineStateRequest
func (b *SecurityBaselineStateRequestBuilder) Request() *SecurityBaselineStateRequest {
	return &SecurityBaselineStateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SecurityBaselineStateRequest is request for SecurityBaselineState
type SecurityBaselineStateRequest struct{ BaseRequest }

// Get performs GET request for SecurityBaselineState
func (r *SecurityBaselineStateRequest) Get(ctx context.Context) (resObj *SecurityBaselineState, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SecurityBaselineState
func (r *SecurityBaselineStateRequest) Update(ctx context.Context, reqObj *SecurityBaselineState) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SecurityBaselineState
func (r *SecurityBaselineStateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SecurityBaselineStateSummaryRequestBuilder is request builder for SecurityBaselineStateSummary
type SecurityBaselineStateSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecurityBaselineStateSummaryRequest
func (b *SecurityBaselineStateSummaryRequestBuilder) Request() *SecurityBaselineStateSummaryRequest {
	return &SecurityBaselineStateSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SecurityBaselineStateSummaryRequest is request for SecurityBaselineStateSummary
type SecurityBaselineStateSummaryRequest struct{ BaseRequest }

// Get performs GET request for SecurityBaselineStateSummary
func (r *SecurityBaselineStateSummaryRequest) Get(ctx context.Context) (resObj *SecurityBaselineStateSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SecurityBaselineStateSummary
func (r *SecurityBaselineStateSummaryRequest) Update(ctx context.Context, reqObj *SecurityBaselineStateSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SecurityBaselineStateSummary
func (r *SecurityBaselineStateSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SecurityBaselineTemplateRequestBuilder is request builder for SecurityBaselineTemplate
type SecurityBaselineTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecurityBaselineTemplateRequest
func (b *SecurityBaselineTemplateRequestBuilder) Request() *SecurityBaselineTemplateRequest {
	return &SecurityBaselineTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SecurityBaselineTemplateRequest is request for SecurityBaselineTemplate
type SecurityBaselineTemplateRequest struct{ BaseRequest }

// Get performs GET request for SecurityBaselineTemplate
func (r *SecurityBaselineTemplateRequest) Get(ctx context.Context) (resObj *SecurityBaselineTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SecurityBaselineTemplate
func (r *SecurityBaselineTemplateRequest) Update(ctx context.Context, reqObj *SecurityBaselineTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SecurityBaselineTemplate
func (r *SecurityBaselineTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// SecurityConfigurationTaskRequestBuilder is request builder for SecurityConfigurationTask
type SecurityConfigurationTaskRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecurityConfigurationTaskRequest
func (b *SecurityConfigurationTaskRequestBuilder) Request() *SecurityConfigurationTaskRequest {
	return &SecurityConfigurationTaskRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SecurityConfigurationTaskRequest is request for SecurityConfigurationTask
type SecurityConfigurationTaskRequest struct{ BaseRequest }

// Get performs GET request for SecurityConfigurationTask
func (r *SecurityConfigurationTaskRequest) Get(ctx context.Context) (resObj *SecurityConfigurationTask, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SecurityConfigurationTask
func (r *SecurityConfigurationTaskRequest) Update(ctx context.Context, reqObj *SecurityConfigurationTask) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SecurityConfigurationTask
func (r *SecurityConfigurationTaskRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
