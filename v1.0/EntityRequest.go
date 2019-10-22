// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// EntityRequestBuilder is request builder for Entity
type EntityRequestBuilder struct{ BaseRequestBuilder }

// Request returns EntityRequest
func (b *EntityRequestBuilder) Request() *EntityRequest {
	return &EntityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EntityRequest is request for Entity
type EntityRequest struct{ BaseRequest }

// Get performs GET request for Entity
func (r *EntityRequest) Get() (resObj *Entity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Entity
func (r *EntityRequest) Update(reqObj *Entity) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Entity
func (r *EntityRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
