// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// AppleUserInitiatedEnrollmentProfileSetPriorityRequestParameter undocumented
type AppleUserInitiatedEnrollmentProfileSetPriorityRequestParameter struct {
	// Priority undocumented
	Priority *int `json:"priority,omitempty"`
}

//
type AppleUserInitiatedEnrollmentProfileSetPriorityRequestBuilder struct{ BaseRequestBuilder }

// SetPriority action undocumented
func (b *AppleUserInitiatedEnrollmentProfileRequestBuilder) SetPriority(reqObj *AppleUserInitiatedEnrollmentProfileSetPriorityRequestParameter) *AppleUserInitiatedEnrollmentProfileSetPriorityRequestBuilder {
	bb := &AppleUserInitiatedEnrollmentProfileSetPriorityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/setPriority"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type AppleUserInitiatedEnrollmentProfileSetPriorityRequest struct{ BaseRequest }

//
func (b *AppleUserInitiatedEnrollmentProfileSetPriorityRequestBuilder) Request() *AppleUserInitiatedEnrollmentProfileSetPriorityRequest {
	return &AppleUserInitiatedEnrollmentProfileSetPriorityRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *AppleUserInitiatedEnrollmentProfileSetPriorityRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}
