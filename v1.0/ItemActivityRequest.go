// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ItemActivityRequestBuilder is request builder for ItemActivity
type ItemActivityRequestBuilder struct{ BaseRequestBuilder }

// Request returns ItemActivityRequest
func (b *ItemActivityRequestBuilder) Request() *ItemActivityRequest {
	return &ItemActivityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ItemActivityRequest is request for ItemActivity
type ItemActivityRequest struct{ BaseRequest }

// Get performs GET request for ItemActivity
func (r *ItemActivityRequest) Get() (resObj *ItemActivity, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ItemActivity
func (r *ItemActivityRequest) Update(reqObj *ItemActivity) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ItemActivity
func (r *ItemActivityRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// DriveItem is navigation property
func (b *ItemActivityRequestBuilder) DriveItem() *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/driveItem"
	return bb
}
