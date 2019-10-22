// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// GovernanceRoleSettingRequestBuilder is request builder for GovernanceRoleSetting
type GovernanceRoleSettingRequestBuilder struct{ BaseRequestBuilder }

// Request returns GovernanceRoleSettingRequest
func (b *GovernanceRoleSettingRequestBuilder) Request() *GovernanceRoleSettingRequest {
	return &GovernanceRoleSettingRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GovernanceRoleSettingRequest is request for GovernanceRoleSetting
type GovernanceRoleSettingRequest struct{ BaseRequest }

// Get performs GET request for GovernanceRoleSetting
func (r *GovernanceRoleSettingRequest) Get() (resObj *GovernanceRoleSetting, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for GovernanceRoleSetting
func (r *GovernanceRoleSettingRequest) Update(reqObj *GovernanceRoleSetting) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for GovernanceRoleSetting
func (r *GovernanceRoleSettingRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Resource is navigation property
func (b *GovernanceRoleSettingRequestBuilder) Resource() *GovernanceResourceRequestBuilder {
	bb := &GovernanceResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resource"
	return bb
}

// RoleDefinition is navigation property
func (b *GovernanceRoleSettingRequestBuilder) RoleDefinition() *GovernanceRoleDefinitionRequestBuilder {
	bb := &GovernanceRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinition"
	return bb
}
