// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceConfigurationUserStateSummaryRequestBuilder is request builder for DeviceConfigurationUserStateSummary
type DeviceConfigurationUserStateSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceConfigurationUserStateSummaryRequest
func (b *DeviceConfigurationUserStateSummaryRequestBuilder) Request() *DeviceConfigurationUserStateSummaryRequest {
	return &DeviceConfigurationUserStateSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceConfigurationUserStateSummaryRequest is request for DeviceConfigurationUserStateSummary
type DeviceConfigurationUserStateSummaryRequest struct{ BaseRequest }

// Get performs GET request for DeviceConfigurationUserStateSummary
func (r *DeviceConfigurationUserStateSummaryRequest) Get() (resObj *DeviceConfigurationUserStateSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceConfigurationUserStateSummary
func (r *DeviceConfigurationUserStateSummaryRequest) Update(reqObj *DeviceConfigurationUserStateSummary) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceConfigurationUserStateSummary
func (r *DeviceConfigurationUserStateSummaryRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
