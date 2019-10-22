// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AccountRequestBuilder is request builder for Account
type AccountRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccountRequest
func (b *AccountRequestBuilder) Request() *AccountRequest {
	return &AccountRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccountRequest is request for Account
type AccountRequest struct{ BaseRequest }

// Get performs GET request for Account
func (r *AccountRequest) Get() (resObj *Account, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Account
func (r *AccountRequest) Update(reqObj *Account) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Account
func (r *AccountRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
