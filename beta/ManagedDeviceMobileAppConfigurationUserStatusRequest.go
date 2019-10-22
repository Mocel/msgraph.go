// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ManagedDeviceMobileAppConfigurationUserStatusRequestBuilder is request builder for ManagedDeviceMobileAppConfigurationUserStatus
type ManagedDeviceMobileAppConfigurationUserStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns ManagedDeviceMobileAppConfigurationUserStatusRequest
func (b *ManagedDeviceMobileAppConfigurationUserStatusRequestBuilder) Request() *ManagedDeviceMobileAppConfigurationUserStatusRequest {
	return &ManagedDeviceMobileAppConfigurationUserStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ManagedDeviceMobileAppConfigurationUserStatusRequest is request for ManagedDeviceMobileAppConfigurationUserStatus
type ManagedDeviceMobileAppConfigurationUserStatusRequest struct{ BaseRequest }

// Get performs GET request for ManagedDeviceMobileAppConfigurationUserStatus
func (r *ManagedDeviceMobileAppConfigurationUserStatusRequest) Get() (resObj *ManagedDeviceMobileAppConfigurationUserStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ManagedDeviceMobileAppConfigurationUserStatus
func (r *ManagedDeviceMobileAppConfigurationUserStatusRequest) Update(reqObj *ManagedDeviceMobileAppConfigurationUserStatus) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ManagedDeviceMobileAppConfigurationUserStatus
func (r *ManagedDeviceMobileAppConfigurationUserStatusRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
