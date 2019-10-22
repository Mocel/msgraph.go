// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WindowsPrivacyDataAccessControlItemRequestBuilder is request builder for WindowsPrivacyDataAccessControlItem
type WindowsPrivacyDataAccessControlItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsPrivacyDataAccessControlItemRequest
func (b *WindowsPrivacyDataAccessControlItemRequestBuilder) Request() *WindowsPrivacyDataAccessControlItemRequest {
	return &WindowsPrivacyDataAccessControlItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsPrivacyDataAccessControlItemRequest is request for WindowsPrivacyDataAccessControlItem
type WindowsPrivacyDataAccessControlItemRequest struct{ BaseRequest }

// Get performs GET request for WindowsPrivacyDataAccessControlItem
func (r *WindowsPrivacyDataAccessControlItemRequest) Get() (resObj *WindowsPrivacyDataAccessControlItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsPrivacyDataAccessControlItem
func (r *WindowsPrivacyDataAccessControlItemRequest) Update(reqObj *WindowsPrivacyDataAccessControlItem) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsPrivacyDataAccessControlItem
func (r *WindowsPrivacyDataAccessControlItemRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
