// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SynchronizationTemplateRequestBuilder is request builder for SynchronizationTemplate
type SynchronizationTemplateRequestBuilder struct{ BaseRequestBuilder }

// Request returns SynchronizationTemplateRequest
func (b *SynchronizationTemplateRequestBuilder) Request() *SynchronizationTemplateRequest {
	return &SynchronizationTemplateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SynchronizationTemplateRequest is request for SynchronizationTemplate
type SynchronizationTemplateRequest struct{ BaseRequest }

// Get performs GET request for SynchronizationTemplate
func (r *SynchronizationTemplateRequest) Get() (resObj *SynchronizationTemplate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SynchronizationTemplate
func (r *SynchronizationTemplateRequest) Update(reqObj *SynchronizationTemplate) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SynchronizationTemplate
func (r *SynchronizationTemplateRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Schema is navigation property
func (b *SynchronizationTemplateRequestBuilder) Schema() *SynchronizationSchemaRequestBuilder {
	bb := &SynchronizationSchemaRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/schema"
	return bb
}
