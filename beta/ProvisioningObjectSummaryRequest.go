// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ProvisioningObjectSummaryRequestBuilder is request builder for ProvisioningObjectSummary
type ProvisioningObjectSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns ProvisioningObjectSummaryRequest
func (b *ProvisioningObjectSummaryRequestBuilder) Request() *ProvisioningObjectSummaryRequest {
	return &ProvisioningObjectSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ProvisioningObjectSummaryRequest is request for ProvisioningObjectSummary
type ProvisioningObjectSummaryRequest struct{ BaseRequest }

// Get performs GET request for ProvisioningObjectSummary
func (r *ProvisioningObjectSummaryRequest) Get() (resObj *ProvisioningObjectSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ProvisioningObjectSummary
func (r *ProvisioningObjectSummaryRequest) Update(reqObj *ProvisioningObjectSummary) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ProvisioningObjectSummary
func (r *ProvisioningObjectSummaryRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
