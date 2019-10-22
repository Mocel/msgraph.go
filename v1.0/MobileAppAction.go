// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// MobileAppAssignRequestParameter undocumented
type MobileAppAssignRequestParameter struct {
	// MobileAppAssignments undocumented
	MobileAppAssignments []MobileAppAssignment `json:"mobileAppAssignments,omitempty"`
}

//
type MobileAppAssignRequestBuilder struct{ BaseRequestBuilder }

// Assign action undocumented
func (b *MobileAppRequestBuilder) Assign(reqObj *MobileAppAssignRequestParameter) *MobileAppAssignRequestBuilder {
	bb := &MobileAppAssignRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/assign"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MobileAppAssignRequest struct{ BaseRequest }

//
func (b *MobileAppAssignRequestBuilder) Request() *MobileAppAssignRequest {
	return &MobileAppAssignRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MobileAppAssignRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}
