// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AndroidTrustedRootCertificateRequestBuilder is request builder for AndroidTrustedRootCertificate
type AndroidTrustedRootCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns AndroidTrustedRootCertificateRequest
func (b *AndroidTrustedRootCertificateRequestBuilder) Request() *AndroidTrustedRootCertificateRequest {
	return &AndroidTrustedRootCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AndroidTrustedRootCertificateRequest is request for AndroidTrustedRootCertificate
type AndroidTrustedRootCertificateRequest struct{ BaseRequest }

// Get performs GET request for AndroidTrustedRootCertificate
func (r *AndroidTrustedRootCertificateRequest) Get() (resObj *AndroidTrustedRootCertificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AndroidTrustedRootCertificate
func (r *AndroidTrustedRootCertificateRequest) Update(reqObj *AndroidTrustedRootCertificate) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AndroidTrustedRootCertificate
func (r *AndroidTrustedRootCertificateRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
