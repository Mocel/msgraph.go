// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MobileAppInstallSummaryRequestBuilder is request builder for MobileAppInstallSummary
type MobileAppInstallSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppInstallSummaryRequest
func (b *MobileAppInstallSummaryRequestBuilder) Request() *MobileAppInstallSummaryRequest {
	return &MobileAppInstallSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppInstallSummaryRequest is request for MobileAppInstallSummary
type MobileAppInstallSummaryRequest struct{ BaseRequest }

// Get performs GET request for MobileAppInstallSummary
func (r *MobileAppInstallSummaryRequest) Get() (resObj *MobileAppInstallSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileAppInstallSummary
func (r *MobileAppInstallSummaryRequest) Update(reqObj *MobileAppInstallSummary) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileAppInstallSummary
func (r *MobileAppInstallSummaryRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
