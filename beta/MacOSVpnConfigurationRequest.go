// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MacOSVpnConfigurationRequestBuilder is request builder for MacOSVpnConfiguration
type MacOSVpnConfigurationRequestBuilder struct{ BaseRequestBuilder }

// Request returns MacOSVpnConfigurationRequest
func (b *MacOSVpnConfigurationRequestBuilder) Request() *MacOSVpnConfigurationRequest {
	return &MacOSVpnConfigurationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MacOSVpnConfigurationRequest is request for MacOSVpnConfiguration
type MacOSVpnConfigurationRequest struct{ BaseRequest }

// Get performs GET request for MacOSVpnConfiguration
func (r *MacOSVpnConfigurationRequest) Get() (resObj *MacOSVpnConfiguration, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MacOSVpnConfiguration
func (r *MacOSVpnConfigurationRequest) Update(reqObj *MacOSVpnConfiguration) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MacOSVpnConfiguration
func (r *MacOSVpnConfigurationRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// IdentityCertificate is navigation property
func (b *MacOSVpnConfigurationRequestBuilder) IdentityCertificate() *MacOSCertificateProfileBaseRequestBuilder {
	bb := &MacOSCertificateProfileBaseRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/identityCertificate"
	return bb
}
