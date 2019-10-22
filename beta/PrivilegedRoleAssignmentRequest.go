// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// PrivilegedRoleAssignmentRequestBuilder is request builder for PrivilegedRoleAssignment
type PrivilegedRoleAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns PrivilegedRoleAssignmentRequest
func (b *PrivilegedRoleAssignmentRequestBuilder) Request() *PrivilegedRoleAssignmentRequest {
	return &PrivilegedRoleAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PrivilegedRoleAssignmentRequest is request for PrivilegedRoleAssignment
type PrivilegedRoleAssignmentRequest struct{ BaseRequest }

// Get performs GET request for PrivilegedRoleAssignment
func (r *PrivilegedRoleAssignmentRequest) Get() (resObj *PrivilegedRoleAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PrivilegedRoleAssignment
func (r *PrivilegedRoleAssignmentRequest) Update(reqObj *PrivilegedRoleAssignment) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PrivilegedRoleAssignment
func (r *PrivilegedRoleAssignmentRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// RoleInfo is navigation property
func (b *PrivilegedRoleAssignmentRequestBuilder) RoleInfo() *PrivilegedRoleRequestBuilder {
	bb := &PrivilegedRoleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleInfo"
	return bb
}
