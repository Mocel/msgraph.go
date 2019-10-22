// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SecureScoreControlProfileRequestBuilder is request builder for SecureScoreControlProfile
type SecureScoreControlProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns SecureScoreControlProfileRequest
func (b *SecureScoreControlProfileRequestBuilder) Request() *SecureScoreControlProfileRequest {
	return &SecureScoreControlProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SecureScoreControlProfileRequest is request for SecureScoreControlProfile
type SecureScoreControlProfileRequest struct{ BaseRequest }

// Get performs GET request for SecureScoreControlProfile
func (r *SecureScoreControlProfileRequest) Get() (resObj *SecureScoreControlProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SecureScoreControlProfile
func (r *SecureScoreControlProfileRequest) Update(reqObj *SecureScoreControlProfile) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SecureScoreControlProfile
func (r *SecureScoreControlProfileRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
