// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsInformationProtectionDeviceRegistrationRequestBuilder is request builder for WindowsInformationProtectionDeviceRegistration
type WindowsInformationProtectionDeviceRegistrationRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsInformationProtectionDeviceRegistrationRequest
func (b *WindowsInformationProtectionDeviceRegistrationRequestBuilder) Request() *WindowsInformationProtectionDeviceRegistrationRequest {
	return &WindowsInformationProtectionDeviceRegistrationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsInformationProtectionDeviceRegistrationRequest is request for WindowsInformationProtectionDeviceRegistration
type WindowsInformationProtectionDeviceRegistrationRequest struct{ BaseRequest }

// Get performs GET request for WindowsInformationProtectionDeviceRegistration
func (r *WindowsInformationProtectionDeviceRegistrationRequest) Get() (resObj *WindowsInformationProtectionDeviceRegistration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsInformationProtectionDeviceRegistration
func (r *WindowsInformationProtectionDeviceRegistrationRequest) Update(reqObj *WindowsInformationProtectionDeviceRegistration) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsInformationProtectionDeviceRegistration
func (r *WindowsInformationProtectionDeviceRegistrationRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
