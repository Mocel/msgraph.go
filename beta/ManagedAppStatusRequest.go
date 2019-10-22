// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedAppStatusRequestBuilder is request builder for ManagedAppStatus
type ManagedAppStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedAppStatusRequest
func (b *ManagedAppStatusRequestBuilder) Request() *ManagedAppStatusRequest {
	return &ManagedAppStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedAppStatusRequest is request for ManagedAppStatus
type ManagedAppStatusRequest struct{ BaseRequest }

// Get performs GET request for ManagedAppStatus
func (r *ManagedAppStatusRequest) Get() (resObj *ManagedAppStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedAppStatus
func (r *ManagedAppStatusRequest) Update(reqObj *ManagedAppStatus) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedAppStatus
func (r *ManagedAppStatusRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
