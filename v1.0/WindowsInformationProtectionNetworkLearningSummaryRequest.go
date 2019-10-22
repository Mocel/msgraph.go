// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsInformationProtectionNetworkLearningSummaryRequestBuilder is request builder for WindowsInformationProtectionNetworkLearningSummary
type WindowsInformationProtectionNetworkLearningSummaryRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsInformationProtectionNetworkLearningSummaryRequest
func (b *WindowsInformationProtectionNetworkLearningSummaryRequestBuilder) Request() *WindowsInformationProtectionNetworkLearningSummaryRequest {
	return &WindowsInformationProtectionNetworkLearningSummaryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsInformationProtectionNetworkLearningSummaryRequest is request for WindowsInformationProtectionNetworkLearningSummary
type WindowsInformationProtectionNetworkLearningSummaryRequest struct{ BaseRequest }

// Get performs GET request for WindowsInformationProtectionNetworkLearningSummary
func (r *WindowsInformationProtectionNetworkLearningSummaryRequest) Get() (resObj *WindowsInformationProtectionNetworkLearningSummary, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsInformationProtectionNetworkLearningSummary
func (r *WindowsInformationProtectionNetworkLearningSummaryRequest) Update(reqObj *WindowsInformationProtectionNetworkLearningSummary) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsInformationProtectionNetworkLearningSummary
func (r *WindowsInformationProtectionNetworkLearningSummaryRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
