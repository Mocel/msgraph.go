// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UserPFXCertificateRequestBuilder is request builder for UserPFXCertificate
type UserPFXCertificateRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserPFXCertificateRequest
func (b *UserPFXCertificateRequestBuilder) Request() *UserPFXCertificateRequest {
	return &UserPFXCertificateRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserPFXCertificateRequest is request for UserPFXCertificate
type UserPFXCertificateRequest struct{ BaseRequest }

// Get performs GET request for UserPFXCertificate
func (r *UserPFXCertificateRequest) Get() (resObj *UserPFXCertificate, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserPFXCertificate
func (r *UserPFXCertificateRequest) Update(reqObj *UserPFXCertificate) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserPFXCertificate
func (r *UserPFXCertificateRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
