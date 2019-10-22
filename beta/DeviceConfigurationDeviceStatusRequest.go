// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceConfigurationDeviceStatusRequestBuilder is request builder for DeviceConfigurationDeviceStatus
type DeviceConfigurationDeviceStatusRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceConfigurationDeviceStatusRequest
func (b *DeviceConfigurationDeviceStatusRequestBuilder) Request() *DeviceConfigurationDeviceStatusRequest {
	return &DeviceConfigurationDeviceStatusRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceConfigurationDeviceStatusRequest is request for DeviceConfigurationDeviceStatus
type DeviceConfigurationDeviceStatusRequest struct{ BaseRequest }

// Get performs GET request for DeviceConfigurationDeviceStatus
func (r *DeviceConfigurationDeviceStatusRequest) Get() (resObj *DeviceConfigurationDeviceStatus, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceConfigurationDeviceStatus
func (r *DeviceConfigurationDeviceStatusRequest) Update(reqObj *DeviceConfigurationDeviceStatus) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceConfigurationDeviceStatus
func (r *DeviceConfigurationDeviceStatusRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
