// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MobileAppProvisioningConfigGroupAssignmentRequestBuilder is request builder for MobileAppProvisioningConfigGroupAssignment
type MobileAppProvisioningConfigGroupAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppProvisioningConfigGroupAssignmentRequest
func (b *MobileAppProvisioningConfigGroupAssignmentRequestBuilder) Request() *MobileAppProvisioningConfigGroupAssignmentRequest {
	return &MobileAppProvisioningConfigGroupAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppProvisioningConfigGroupAssignmentRequest is request for MobileAppProvisioningConfigGroupAssignment
type MobileAppProvisioningConfigGroupAssignmentRequest struct{ BaseRequest }

// Get performs GET request for MobileAppProvisioningConfigGroupAssignment
func (r *MobileAppProvisioningConfigGroupAssignmentRequest) Get() (resObj *MobileAppProvisioningConfigGroupAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileAppProvisioningConfigGroupAssignment
func (r *MobileAppProvisioningConfigGroupAssignmentRequest) Update(reqObj *MobileAppProvisioningConfigGroupAssignment) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileAppProvisioningConfigGroupAssignment
func (r *MobileAppProvisioningConfigGroupAssignmentRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
