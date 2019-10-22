// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GroupPolicyPresentationValueRequestBuilder is request builder for GroupPolicyPresentationValue
type GroupPolicyPresentationValueRequestBuilder struct{ BaseRequestBuilder }

// Request returns GroupPolicyPresentationValueRequest
func (b *GroupPolicyPresentationValueRequestBuilder) Request() *GroupPolicyPresentationValueRequest {
	return &GroupPolicyPresentationValueRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GroupPolicyPresentationValueRequest is request for GroupPolicyPresentationValue
type GroupPolicyPresentationValueRequest struct{ BaseRequest }

// Get performs GET request for GroupPolicyPresentationValue
func (r *GroupPolicyPresentationValueRequest) Get() (resObj *GroupPolicyPresentationValue, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for GroupPolicyPresentationValue
func (r *GroupPolicyPresentationValueRequest) Update(reqObj *GroupPolicyPresentationValue) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for GroupPolicyPresentationValue
func (r *GroupPolicyPresentationValueRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// DefinitionValue is navigation property
func (b *GroupPolicyPresentationValueRequestBuilder) DefinitionValue() *GroupPolicyDefinitionValueRequestBuilder {
	bb := &GroupPolicyDefinitionValueRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/definitionValue"
	return bb
}

// Presentation is navigation property
func (b *GroupPolicyPresentationValueRequestBuilder) Presentation() *GroupPolicyPresentationRequestBuilder {
	bb := &GroupPolicyPresentationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/presentation"
	return bb
}
