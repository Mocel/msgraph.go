// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// CurrencyRequestBuilder is request builder for Currency
type CurrencyRequestBuilder struct{ BaseRequestBuilder }

// Request returns CurrencyRequest
func (b *CurrencyRequestBuilder) Request() *CurrencyRequest {
	return &CurrencyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CurrencyRequest is request for Currency
type CurrencyRequest struct{ BaseRequest }

// Get performs GET request for Currency
func (r *CurrencyRequest) Get() (resObj *Currency, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Currency
func (r *CurrencyRequest) Update(reqObj *Currency) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Currency
func (r *CurrencyRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
