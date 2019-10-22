// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

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
func (r *ApplicationTemplateRequest) Get() (resObj *ApplicationTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ApplicationTemplate
func (r *ApplicationTemplateRequest) Update(reqObj *ApplicationTemplate) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ApplicationTemplate
func (r *ApplicationTemplateRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
