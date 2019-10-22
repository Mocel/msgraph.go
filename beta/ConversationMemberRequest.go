// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ConversationMemberRequestBuilder is request builder for ConversationMember
type ConversationMemberRequestBuilder struct{ BaseRequestBuilder }

// Request returns ConversationMemberRequest
func (b *ConversationMemberRequestBuilder) Request() *ConversationMemberRequest {
	return &ConversationMemberRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ConversationMemberRequest is request for ConversationMember
type ConversationMemberRequest struct{ BaseRequest }

// Get performs GET request for ConversationMember
func (r *ConversationMemberRequest) Get() (resObj *ConversationMember, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for ConversationMember
func (r *ConversationMemberRequest) Update(reqObj *ConversationMember) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for ConversationMember
func (r *ConversationMemberRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}
