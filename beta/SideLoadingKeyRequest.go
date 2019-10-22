// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SideLoadingKeyRequestBuilder is request builder for SideLoadingKey
type SideLoadingKeyRequestBuilder struct{ BaseRequestBuilder }

// Request returns SideLoadingKeyRequest
func (b *SideLoadingKeyRequestBuilder) Request() *SideLoadingKeyRequest {
	return &SideLoadingKeyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SideLoadingKeyRequest is request for SideLoadingKey
type SideLoadingKeyRequest struct{ BaseRequest }

// Get performs GET request for SideLoadingKey
func (r *SideLoadingKeyRequest) Get() (resObj *SideLoadingKey, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SideLoadingKey
func (r *SideLoadingKeyRequest) Update(reqObj *SideLoadingKey) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SideLoadingKey
func (r *SideLoadingKeyRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
