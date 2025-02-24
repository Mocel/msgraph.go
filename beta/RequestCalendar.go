// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// CalendarRequestBuilder is request builder for Calendar
type CalendarRequestBuilder struct{ BaseRequestBuilder }

// Request returns CalendarRequest
func (b *CalendarRequestBuilder) Request() *CalendarRequest {
	return &CalendarRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CalendarRequest is request for Calendar
type CalendarRequest struct{ BaseRequest }

// Get performs GET request for Calendar
func (r *CalendarRequest) Get(ctx context.Context) (resObj *Calendar, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Calendar
func (r *CalendarRequest) Update(ctx context.Context, reqObj *Calendar) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Calendar
func (r *CalendarRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CalendarGroupRequestBuilder is request builder for CalendarGroup
type CalendarGroupRequestBuilder struct{ BaseRequestBuilder }

// Request returns CalendarGroupRequest
func (b *CalendarGroupRequestBuilder) Request() *CalendarGroupRequest {
	return &CalendarGroupRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CalendarGroupRequest is request for CalendarGroup
type CalendarGroupRequest struct{ BaseRequest }

// Get performs GET request for CalendarGroup
func (r *CalendarGroupRequest) Get(ctx context.Context) (resObj *CalendarGroup, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CalendarGroup
func (r *CalendarGroupRequest) Update(ctx context.Context, reqObj *CalendarGroup) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CalendarGroup
func (r *CalendarGroupRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CalendarPermissionRequestBuilder is request builder for CalendarPermission
type CalendarPermissionRequestBuilder struct{ BaseRequestBuilder }

// Request returns CalendarPermissionRequest
func (b *CalendarPermissionRequestBuilder) Request() *CalendarPermissionRequest {
	return &CalendarPermissionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CalendarPermissionRequest is request for CalendarPermission
type CalendarPermissionRequest struct{ BaseRequest }

// Get performs GET request for CalendarPermission
func (r *CalendarPermissionRequest) Get(ctx context.Context) (resObj *CalendarPermission, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CalendarPermission
func (r *CalendarPermissionRequest) Update(ctx context.Context, reqObj *CalendarPermission) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CalendarPermission
func (r *CalendarPermissionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// CalendarSharingMessageRequestBuilder is request builder for CalendarSharingMessage
type CalendarSharingMessageRequestBuilder struct{ BaseRequestBuilder }

// Request returns CalendarSharingMessageRequest
func (b *CalendarSharingMessageRequestBuilder) Request() *CalendarSharingMessageRequest {
	return &CalendarSharingMessageRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CalendarSharingMessageRequest is request for CalendarSharingMessage
type CalendarSharingMessageRequest struct{ BaseRequest }

// Get performs GET request for CalendarSharingMessage
func (r *CalendarSharingMessageRequest) Get(ctx context.Context) (resObj *CalendarSharingMessage, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for CalendarSharingMessage
func (r *CalendarSharingMessageRequest) Update(ctx context.Context, reqObj *CalendarSharingMessage) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for CalendarSharingMessage
func (r *CalendarSharingMessageRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
