// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DeviceManagementAutopilotEventRequestBuilder is request builder for DeviceManagementAutopilotEvent
type DeviceManagementAutopilotEventRequestBuilder struct{ BaseRequestBuilder }

// Request returns DeviceManagementAutopilotEventRequest
func (b *DeviceManagementAutopilotEventRequestBuilder) Request() *DeviceManagementAutopilotEventRequest {
	return &DeviceManagementAutopilotEventRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DeviceManagementAutopilotEventRequest is request for DeviceManagementAutopilotEvent
type DeviceManagementAutopilotEventRequest struct{ BaseRequest }

// Get performs GET request for DeviceManagementAutopilotEvent
func (r *DeviceManagementAutopilotEventRequest) Get() (resObj *DeviceManagementAutopilotEvent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DeviceManagementAutopilotEvent
func (r *DeviceManagementAutopilotEventRequest) Update(reqObj *DeviceManagementAutopilotEvent) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DeviceManagementAutopilotEvent
func (r *DeviceManagementAutopilotEventRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
