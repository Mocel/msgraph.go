// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AccessReviewReviewerRequestBuilder is request builder for AccessReviewReviewer
type AccessReviewReviewerRequestBuilder struct{ BaseRequestBuilder }

// Request returns AccessReviewReviewerRequest
func (b *AccessReviewReviewerRequestBuilder) Request() *AccessReviewReviewerRequest {
	return &AccessReviewReviewerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AccessReviewReviewerRequest is request for AccessReviewReviewer
type AccessReviewReviewerRequest struct{ BaseRequest }

// Get performs GET request for AccessReviewReviewer
func (r *AccessReviewReviewerRequest) Get() (resObj *AccessReviewReviewer, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AccessReviewReviewer
func (r *AccessReviewReviewerRequest) Update(reqObj *AccessReviewReviewer) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AccessReviewReviewer
func (r *AccessReviewReviewerRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
