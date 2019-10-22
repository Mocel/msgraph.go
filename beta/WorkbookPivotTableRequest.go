// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// WorkbookPivotTableRequestBuilder is request builder for WorkbookPivotTable
type WorkbookPivotTableRequestBuilder struct{ BaseRequestBuilder }

// Request returns WorkbookPivotTableRequest
func (b *WorkbookPivotTableRequestBuilder) Request() *WorkbookPivotTableRequest {
	return &WorkbookPivotTableRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WorkbookPivotTableRequest is request for WorkbookPivotTable
type WorkbookPivotTableRequest struct{ BaseRequest }

// Get performs GET request for WorkbookPivotTable
func (r *WorkbookPivotTableRequest) Get() (resObj *WorkbookPivotTable, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WorkbookPivotTable
func (r *WorkbookPivotTableRequest) Update(reqObj *WorkbookPivotTable) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WorkbookPivotTable
func (r *WorkbookPivotTableRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Worksheet is navigation property
func (b *WorkbookPivotTableRequestBuilder) Worksheet() *WorkbookWorksheetRequestBuilder {
	bb := &WorkbookWorksheetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/worksheet"
	return bb
}
