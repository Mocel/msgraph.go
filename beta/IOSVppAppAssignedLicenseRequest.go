// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// IOSVppAppAssignedLicenseRequestBuilder is request builder for IOSVppAppAssignedLicense
type IOSVppAppAssignedLicenseRequestBuilder struct{ BaseRequestBuilder }

// Request returns IOSVppAppAssignedLicenseRequest
func (b *IOSVppAppAssignedLicenseRequestBuilder) Request() *IOSVppAppAssignedLicenseRequest {
	return &IOSVppAppAssignedLicenseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// IOSVppAppAssignedLicenseRequest is request for IOSVppAppAssignedLicense
type IOSVppAppAssignedLicenseRequest struct{ BaseRequest }

// Get performs GET request for IOSVppAppAssignedLicense
func (r *IOSVppAppAssignedLicenseRequest) Get() (resObj *IOSVppAppAssignedLicense, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for IOSVppAppAssignedLicense
func (r *IOSVppAppAssignedLicenseRequest) Update(reqObj *IOSVppAppAssignedLicense) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for IOSVppAppAssignedLicense
func (r *IOSVppAppAssignedLicenseRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
