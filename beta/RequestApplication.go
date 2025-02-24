// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// ApplicationRequestBuilder is request builder for Application
type ApplicationRequestBuilder struct{ BaseRequestBuilder }

// Request returns ApplicationRequest
func (b *ApplicationRequestBuilder) Request() *ApplicationRequest {
	return &ApplicationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ApplicationRequest is request for Application
type ApplicationRequest struct{ BaseRequest }

// Get performs GET request for Application
func (r *ApplicationRequest) Get(ctx context.Context) (resObj *Application, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Application
func (r *ApplicationRequest) Update(ctx context.Context, reqObj *Application) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Application
func (r *ApplicationRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ApplicationServicePrincipalRequestBuilder is request builder for ApplicationServicePrincipal
type ApplicationServicePrincipalRequestBuilder struct{ BaseRequestBuilder }

// Request returns ApplicationServicePrincipalRequest
func (b *ApplicationServicePrincipalRequestBuilder) Request() *ApplicationServicePrincipalRequest {
	return &ApplicationServicePrincipalRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ApplicationServicePrincipalRequest is request for ApplicationServicePrincipal
type ApplicationServicePrincipalRequest struct{ BaseRequest }

// Get performs GET request for ApplicationServicePrincipal
func (r *ApplicationServicePrincipalRequest) Get(ctx context.Context) (resObj *ApplicationServicePrincipal, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ApplicationServicePrincipal
func (r *ApplicationServicePrincipalRequest) Update(ctx context.Context, reqObj *ApplicationServicePrincipal) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ApplicationServicePrincipal
func (r *ApplicationServicePrincipalRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ApplicationSignInDetailedSummaryRequestBuilder is request builder for ApplicationSignInDetailedSummary
type ApplicationSignInDetailedSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns ApplicationSignInDetailedSummaryRequest
func (b *ApplicationSignInDetailedSummaryRequestBuilder) Request() *ApplicationSignInDetailedSummaryRequest {
	return &ApplicationSignInDetailedSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ApplicationSignInDetailedSummaryRequest is request for ApplicationSignInDetailedSummary
type ApplicationSignInDetailedSummaryRequest struct{ BaseRequest }

// Get performs GET request for ApplicationSignInDetailedSummary
func (r *ApplicationSignInDetailedSummaryRequest) Get(ctx context.Context) (resObj *ApplicationSignInDetailedSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ApplicationSignInDetailedSummary
func (r *ApplicationSignInDetailedSummaryRequest) Update(ctx context.Context, reqObj *ApplicationSignInDetailedSummary) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ApplicationSignInDetailedSummary
func (r *ApplicationSignInDetailedSummaryRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// ApplicationTemplateRequestBuilder is request builder for ApplicationTemplate
type ApplicationTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns ApplicationTemplateRequest
func (b *ApplicationTemplateRequestBuilder) Request() *ApplicationTemplateRequest {
	return &ApplicationTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ApplicationTemplateRequest is request for ApplicationTemplate
type ApplicationTemplateRequest struct{ BaseRequest }

// Get performs GET request for ApplicationTemplate
func (r *ApplicationTemplateRequest) Get(ctx context.Context) (resObj *ApplicationTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ApplicationTemplate
func (r *ApplicationTemplateRequest) Update(ctx context.Context, reqObj *ApplicationTemplate) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ApplicationTemplate
func (r *ApplicationTemplateRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
