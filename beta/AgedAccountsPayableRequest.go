// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AgedAccountsPayableRequestBuilder is request builder for AgedAccountsPayable
type AgedAccountsPayableRequestBuilder struct{ BaseRequestBuilder }

// Request returns AgedAccountsPayableRequest
func (b *AgedAccountsPayableRequestBuilder) Request() *AgedAccountsPayableRequest {
	return &AgedAccountsPayableRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AgedAccountsPayableRequest is request for AgedAccountsPayable
type AgedAccountsPayableRequest struct{ BaseRequest }

// Get performs GET request for AgedAccountsPayable
func (r *AgedAccountsPayableRequest) Get() (resObj *AgedAccountsPayable, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AgedAccountsPayable
func (r *AgedAccountsPayableRequest) Update(reqObj *AgedAccountsPayable) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AgedAccountsPayable
func (r *AgedAccountsPayableRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
