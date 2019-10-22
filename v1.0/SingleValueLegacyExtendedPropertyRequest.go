// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SingleValueLegacyExtendedPropertyRequestBuilder is request builder for SingleValueLegacyExtendedProperty
type SingleValueLegacyExtendedPropertyRequestBuilder struct{ BaseRequestBuilder }

// Request returns SingleValueLegacyExtendedPropertyRequest
func (b *SingleValueLegacyExtendedPropertyRequestBuilder) Request() *SingleValueLegacyExtendedPropertyRequest {
	return &SingleValueLegacyExtendedPropertyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SingleValueLegacyExtendedPropertyRequest is request for SingleValueLegacyExtendedProperty
type SingleValueLegacyExtendedPropertyRequest struct{ BaseRequest }

// Get performs GET request for SingleValueLegacyExtendedProperty
func (r *SingleValueLegacyExtendedPropertyRequest) Get() (resObj *SingleValueLegacyExtendedProperty, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SingleValueLegacyExtendedProperty
func (r *SingleValueLegacyExtendedPropertyRequest) Update(reqObj *SingleValueLegacyExtendedProperty) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SingleValueLegacyExtendedProperty
func (r *SingleValueLegacyExtendedPropertyRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
