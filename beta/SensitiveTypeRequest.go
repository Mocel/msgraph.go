// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SensitiveTypeRequestBuilder is request builder for SensitiveType
type SensitiveTypeRequestBuilder struct{ BaseRequestBuilder }

// Request returns SensitiveTypeRequest
func (b *SensitiveTypeRequestBuilder) Request() *SensitiveTypeRequest {
	return &SensitiveTypeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SensitiveTypeRequest is request for SensitiveType
type SensitiveTypeRequest struct{ BaseRequest }

// Get performs GET request for SensitiveType
func (r *SensitiveTypeRequest) Get() (resObj *SensitiveType, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SensitiveType
func (r *SensitiveTypeRequest) Update(reqObj *SensitiveType) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SensitiveType
func (r *SensitiveTypeRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
