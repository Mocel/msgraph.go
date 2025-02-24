// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// PrivilegedAccessRequestBuilder is request builder for PrivilegedAccess
type PrivilegedAccessRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivilegedAccessRequest
func (b *PrivilegedAccessRequestBuilder) Request() *PrivilegedAccessRequest {
	return &PrivilegedAccessRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivilegedAccessRequest is request for PrivilegedAccess
type PrivilegedAccessRequest struct{ BaseRequest }

// Get performs GET request for PrivilegedAccess
func (r *PrivilegedAccessRequest) Get(ctx context.Context) (resObj *PrivilegedAccess, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrivilegedAccess
func (r *PrivilegedAccessRequest) Update(ctx context.Context, reqObj *PrivilegedAccess) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrivilegedAccess
func (r *PrivilegedAccessRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrivilegedApprovalRequestBuilder is request builder for PrivilegedApproval
type PrivilegedApprovalRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivilegedApprovalRequest
func (b *PrivilegedApprovalRequestBuilder) Request() *PrivilegedApprovalRequest {
	return &PrivilegedApprovalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivilegedApprovalRequest is request for PrivilegedApproval
type PrivilegedApprovalRequest struct{ BaseRequest }

// Get performs GET request for PrivilegedApproval
func (r *PrivilegedApprovalRequest) Get(ctx context.Context) (resObj *PrivilegedApproval, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrivilegedApproval
func (r *PrivilegedApprovalRequest) Update(ctx context.Context, reqObj *PrivilegedApproval) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrivilegedApproval
func (r *PrivilegedApprovalRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrivilegedOperationEventRequestBuilder is request builder for PrivilegedOperationEvent
type PrivilegedOperationEventRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivilegedOperationEventRequest
func (b *PrivilegedOperationEventRequestBuilder) Request() *PrivilegedOperationEventRequest {
	return &PrivilegedOperationEventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivilegedOperationEventRequest is request for PrivilegedOperationEvent
type PrivilegedOperationEventRequest struct{ BaseRequest }

// Get performs GET request for PrivilegedOperationEvent
func (r *PrivilegedOperationEventRequest) Get(ctx context.Context) (resObj *PrivilegedOperationEvent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrivilegedOperationEvent
func (r *PrivilegedOperationEventRequest) Update(ctx context.Context, reqObj *PrivilegedOperationEvent) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrivilegedOperationEvent
func (r *PrivilegedOperationEventRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrivilegedRoleRequestBuilder is request builder for PrivilegedRole
type PrivilegedRoleRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivilegedRoleRequest
func (b *PrivilegedRoleRequestBuilder) Request() *PrivilegedRoleRequest {
	return &PrivilegedRoleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivilegedRoleRequest is request for PrivilegedRole
type PrivilegedRoleRequest struct{ BaseRequest }

// Get performs GET request for PrivilegedRole
func (r *PrivilegedRoleRequest) Get(ctx context.Context) (resObj *PrivilegedRole, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrivilegedRole
func (r *PrivilegedRoleRequest) Update(ctx context.Context, reqObj *PrivilegedRole) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrivilegedRole
func (r *PrivilegedRoleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrivilegedRoleAssignmentRequestBuilder is request builder for PrivilegedRoleAssignment
type PrivilegedRoleAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivilegedRoleAssignmentRequest
func (b *PrivilegedRoleAssignmentRequestBuilder) Request() *PrivilegedRoleAssignmentRequest {
	return &PrivilegedRoleAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivilegedRoleAssignmentRequest is request for PrivilegedRoleAssignment
type PrivilegedRoleAssignmentRequest struct{ BaseRequest }

// Get performs GET request for PrivilegedRoleAssignment
func (r *PrivilegedRoleAssignmentRequest) Get(ctx context.Context) (resObj *PrivilegedRoleAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrivilegedRoleAssignment
func (r *PrivilegedRoleAssignmentRequest) Update(ctx context.Context, reqObj *PrivilegedRoleAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrivilegedRoleAssignment
func (r *PrivilegedRoleAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrivilegedRoleAssignmentRequestObjectRequestBuilder is request builder for PrivilegedRoleAssignmentRequestObject
type PrivilegedRoleAssignmentRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivilegedRoleAssignmentRequestObjectRequest
func (b *PrivilegedRoleAssignmentRequestObjectRequestBuilder) Request() *PrivilegedRoleAssignmentRequestObjectRequest {
	return &PrivilegedRoleAssignmentRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivilegedRoleAssignmentRequestObjectRequest is request for PrivilegedRoleAssignmentRequestObject
type PrivilegedRoleAssignmentRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for PrivilegedRoleAssignmentRequestObject
func (r *PrivilegedRoleAssignmentRequestObjectRequest) Get(ctx context.Context) (resObj *PrivilegedRoleAssignmentRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrivilegedRoleAssignmentRequestObject
func (r *PrivilegedRoleAssignmentRequestObjectRequest) Update(ctx context.Context, reqObj *PrivilegedRoleAssignmentRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrivilegedRoleAssignmentRequestObject
func (r *PrivilegedRoleAssignmentRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrivilegedRoleSettingsRequestBuilder is request builder for PrivilegedRoleSettings
type PrivilegedRoleSettingsRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivilegedRoleSettingsRequest
func (b *PrivilegedRoleSettingsRequestBuilder) Request() *PrivilegedRoleSettingsRequest {
	return &PrivilegedRoleSettingsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivilegedRoleSettingsRequest is request for PrivilegedRoleSettings
type PrivilegedRoleSettingsRequest struct{ BaseRequest }

// Get performs GET request for PrivilegedRoleSettings
func (r *PrivilegedRoleSettingsRequest) Get(ctx context.Context) (resObj *PrivilegedRoleSettings, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrivilegedRoleSettings
func (r *PrivilegedRoleSettingsRequest) Update(ctx context.Context, reqObj *PrivilegedRoleSettings) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrivilegedRoleSettings
func (r *PrivilegedRoleSettingsRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrivilegedRoleSummaryRequestBuilder is request builder for PrivilegedRoleSummary
type PrivilegedRoleSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivilegedRoleSummaryRequest
func (b *PrivilegedRoleSummaryRequestBuilder) Request() *PrivilegedRoleSummaryRequest {
	return &PrivilegedRoleSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivilegedRoleSummaryRequest is request for PrivilegedRoleSummary
type PrivilegedRoleSummaryRequest struct{ BaseRequest }

// Get performs GET request for PrivilegedRoleSummary
func (r *PrivilegedRoleSummaryRequest) Get(ctx context.Context) (resObj *PrivilegedRoleSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrivilegedRoleSummary
func (r *PrivilegedRoleSummaryRequest) Update(ctx context.Context, reqObj *PrivilegedRoleSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrivilegedRoleSummary
func (r *PrivilegedRoleSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// PrivilegedSignupStatusRequestBuilder is request builder for PrivilegedSignupStatus
type PrivilegedSignupStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivilegedSignupStatusRequest
func (b *PrivilegedSignupStatusRequestBuilder) Request() *PrivilegedSignupStatusRequest {
	return &PrivilegedSignupStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivilegedSignupStatusRequest is request for PrivilegedSignupStatus
type PrivilegedSignupStatusRequest struct{ BaseRequest }

// Get performs GET request for PrivilegedSignupStatus
func (r *PrivilegedSignupStatusRequest) Get(ctx context.Context) (resObj *PrivilegedSignupStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrivilegedSignupStatus
func (r *PrivilegedSignupStatusRequest) Update(ctx context.Context, reqObj *PrivilegedSignupStatus) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrivilegedSignupStatus
func (r *PrivilegedSignupStatusRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
