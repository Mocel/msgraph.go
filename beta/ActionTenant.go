// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// ManagedTenants is navigation property
func (b *TenantRelationshipRequestBuilder) ManagedTenants() *ManagedTenantsmanagedTenantRequestBuilder {
	bb := &ManagedTenantsmanagedTenantRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedTenants"
	return bb
}

// DefaultRolesSettings is navigation property
func (b *TenantSetupInfoRequestBuilder) DefaultRolesSettings() *PrivilegedRoleSettingsRequestBuilder {
	bb := &PrivilegedRoleSettingsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/defaultRolesSettings"
	return bb
}
