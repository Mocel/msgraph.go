// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ActivityHistoryItemRequestBuilder is request builder for ActivityHistoryItem
type ActivityHistoryItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns ActivityHistoryItemRequest
func (b *ActivityHistoryItemRequestBuilder) Request() *ActivityHistoryItemRequest {
	return &ActivityHistoryItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ActivityHistoryItemRequest is request for ActivityHistoryItem
type ActivityHistoryItemRequest struct{ BaseRequest }

// Get performs GET request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Get() (resObj *ActivityHistoryItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Update(reqObj *ActivityHistoryItem) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ActivityHistoryItem
func (r *ActivityHistoryItemRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Activity is navigation property
func (b *ActivityHistoryItemRequestBuilder) Activity() *UserActivityRequestBuilder {
	bb := &UserActivityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/activity"
	return bb
}
