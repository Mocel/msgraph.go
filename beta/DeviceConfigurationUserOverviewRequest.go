// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceConfigurationUserOverviewRequestBuilder is request builder for DeviceConfigurationUserOverview
type DeviceConfigurationUserOverviewRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceConfigurationUserOverviewRequest
func (b *DeviceConfigurationUserOverviewRequestBuilder) Request() *DeviceConfigurationUserOverviewRequest {
	return &DeviceConfigurationUserOverviewRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceConfigurationUserOverviewRequest is request for DeviceConfigurationUserOverview
type DeviceConfigurationUserOverviewRequest struct{ BaseRequest }

// Get performs GET request for DeviceConfigurationUserOverview
func (r *DeviceConfigurationUserOverviewRequest) Get() (resObj *DeviceConfigurationUserOverview, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceConfigurationUserOverview
func (r *DeviceConfigurationUserOverviewRequest) Update(reqObj *DeviceConfigurationUserOverview) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceConfigurationUserOverview
func (r *DeviceConfigurationUserOverviewRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
