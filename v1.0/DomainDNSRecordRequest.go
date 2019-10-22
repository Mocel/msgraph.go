// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DomainDNSRecordRequestBuilder is request builder for DomainDNSRecord
type DomainDNSRecordRequestBuilder struct{ BaseRequestBuilder }

// Request returns DomainDNSRecordRequest
func (b *DomainDNSRecordRequestBuilder) Request() *DomainDNSRecordRequest {
	return &DomainDNSRecordRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DomainDNSRecordRequest is request for DomainDNSRecord
type DomainDNSRecordRequest struct{ BaseRequest }

// Get performs GET request for DomainDNSRecord
func (r *DomainDNSRecordRequest) Get() (resObj *DomainDNSRecord, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for DomainDNSRecord
func (r *DomainDNSRecordRequest) Update(reqObj *DomainDNSRecord) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for DomainDNSRecord
func (r *DomainDNSRecordRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
