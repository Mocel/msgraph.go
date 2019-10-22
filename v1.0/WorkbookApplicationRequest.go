// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkbookApplicationRequestBuilder is request builder for WorkbookApplication
type WorkbookApplicationRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookApplicationRequest
func (b *WorkbookApplicationRequestBuilder) Request() *WorkbookApplicationRequest {
	return &WorkbookApplicationRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookApplicationRequest is request for WorkbookApplication
type WorkbookApplicationRequest struct{ BaseRequest }

// Get performs GET request for WorkbookApplication
func (r *WorkbookApplicationRequest) Get() (resObj *WorkbookApplication, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookApplication
func (r *WorkbookApplicationRequest) Update(reqObj *WorkbookApplication) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookApplication
func (r *WorkbookApplicationRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
