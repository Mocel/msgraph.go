// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// UnifiedRbacResourceActionRequestBuilder is request builder for UnifiedRbacResourceAction
type UnifiedRbacResourceActionRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRbacResourceActionRequest
func (b *UnifiedRbacResourceActionRequestBuilder) Request() *UnifiedRbacResourceActionRequest {
	return &UnifiedRbacResourceActionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRbacResourceActionRequest is request for UnifiedRbacResourceAction
type UnifiedRbacResourceActionRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRbacResourceAction
func (r *UnifiedRbacResourceActionRequest) Get(ctx context.Context) (resObj *UnifiedRbacResourceAction, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRbacResourceAction
func (r *UnifiedRbacResourceActionRequest) Update(ctx context.Context, reqObj *UnifiedRbacResourceAction) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRbacResourceAction
func (r *UnifiedRbacResourceActionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRbacResourceNamespaceRequestBuilder is request builder for UnifiedRbacResourceNamespace
type UnifiedRbacResourceNamespaceRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRbacResourceNamespaceRequest
func (b *UnifiedRbacResourceNamespaceRequestBuilder) Request() *UnifiedRbacResourceNamespaceRequest {
	return &UnifiedRbacResourceNamespaceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRbacResourceNamespaceRequest is request for UnifiedRbacResourceNamespace
type UnifiedRbacResourceNamespaceRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRbacResourceNamespace
func (r *UnifiedRbacResourceNamespaceRequest) Get(ctx context.Context) (resObj *UnifiedRbacResourceNamespace, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRbacResourceNamespace
func (r *UnifiedRbacResourceNamespaceRequest) Update(ctx context.Context, reqObj *UnifiedRbacResourceNamespace) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRbacResourceNamespace
func (r *UnifiedRbacResourceNamespaceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRbacResourceScopeRequestBuilder is request builder for UnifiedRbacResourceScope
type UnifiedRbacResourceScopeRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRbacResourceScopeRequest
func (b *UnifiedRbacResourceScopeRequestBuilder) Request() *UnifiedRbacResourceScopeRequest {
	return &UnifiedRbacResourceScopeRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRbacResourceScopeRequest is request for UnifiedRbacResourceScope
type UnifiedRbacResourceScopeRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRbacResourceScope
func (r *UnifiedRbacResourceScopeRequest) Get(ctx context.Context) (resObj *UnifiedRbacResourceScope, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRbacResourceScope
func (r *UnifiedRbacResourceScopeRequest) Update(ctx context.Context, reqObj *UnifiedRbacResourceScope) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRbacResourceScope
func (r *UnifiedRbacResourceScopeRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRoleAssignmentRequestBuilder is request builder for UnifiedRoleAssignment
type UnifiedRoleAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleAssignmentRequest
func (b *UnifiedRoleAssignmentRequestBuilder) Request() *UnifiedRoleAssignmentRequest {
	return &UnifiedRoleAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleAssignmentRequest is request for UnifiedRoleAssignment
type UnifiedRoleAssignmentRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRoleAssignment
func (r *UnifiedRoleAssignmentRequest) Get(ctx context.Context) (resObj *UnifiedRoleAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRoleAssignment
func (r *UnifiedRoleAssignmentRequest) Update(ctx context.Context, reqObj *UnifiedRoleAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRoleAssignment
func (r *UnifiedRoleAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRoleAssignmentMultipleRequestBuilder is request builder for UnifiedRoleAssignmentMultiple
type UnifiedRoleAssignmentMultipleRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleAssignmentMultipleRequest
func (b *UnifiedRoleAssignmentMultipleRequestBuilder) Request() *UnifiedRoleAssignmentMultipleRequest {
	return &UnifiedRoleAssignmentMultipleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleAssignmentMultipleRequest is request for UnifiedRoleAssignmentMultiple
type UnifiedRoleAssignmentMultipleRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRoleAssignmentMultiple
func (r *UnifiedRoleAssignmentMultipleRequest) Get(ctx context.Context) (resObj *UnifiedRoleAssignmentMultiple, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRoleAssignmentMultiple
func (r *UnifiedRoleAssignmentMultipleRequest) Update(ctx context.Context, reqObj *UnifiedRoleAssignmentMultiple) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRoleAssignmentMultiple
func (r *UnifiedRoleAssignmentMultipleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRoleAssignmentScheduleRequestBuilder is request builder for UnifiedRoleAssignmentSchedule
type UnifiedRoleAssignmentScheduleRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleAssignmentScheduleRequest
func (b *UnifiedRoleAssignmentScheduleRequestBuilder) Request() *UnifiedRoleAssignmentScheduleRequest {
	return &UnifiedRoleAssignmentScheduleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleAssignmentScheduleRequest is request for UnifiedRoleAssignmentSchedule
type UnifiedRoleAssignmentScheduleRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRoleAssignmentSchedule
func (r *UnifiedRoleAssignmentScheduleRequest) Get(ctx context.Context) (resObj *UnifiedRoleAssignmentSchedule, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRoleAssignmentSchedule
func (r *UnifiedRoleAssignmentScheduleRequest) Update(ctx context.Context, reqObj *UnifiedRoleAssignmentSchedule) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRoleAssignmentSchedule
func (r *UnifiedRoleAssignmentScheduleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRoleAssignmentScheduleInstanceRequestBuilder is request builder for UnifiedRoleAssignmentScheduleInstance
type UnifiedRoleAssignmentScheduleInstanceRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleAssignmentScheduleInstanceRequest
func (b *UnifiedRoleAssignmentScheduleInstanceRequestBuilder) Request() *UnifiedRoleAssignmentScheduleInstanceRequest {
	return &UnifiedRoleAssignmentScheduleInstanceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleAssignmentScheduleInstanceRequest is request for UnifiedRoleAssignmentScheduleInstance
type UnifiedRoleAssignmentScheduleInstanceRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRoleAssignmentScheduleInstance
func (r *UnifiedRoleAssignmentScheduleInstanceRequest) Get(ctx context.Context) (resObj *UnifiedRoleAssignmentScheduleInstance, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRoleAssignmentScheduleInstance
func (r *UnifiedRoleAssignmentScheduleInstanceRequest) Update(ctx context.Context, reqObj *UnifiedRoleAssignmentScheduleInstance) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRoleAssignmentScheduleInstance
func (r *UnifiedRoleAssignmentScheduleInstanceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRoleAssignmentScheduleRequestObjectRequestBuilder is request builder for UnifiedRoleAssignmentScheduleRequestObject
type UnifiedRoleAssignmentScheduleRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleAssignmentScheduleRequestObjectRequest
func (b *UnifiedRoleAssignmentScheduleRequestObjectRequestBuilder) Request() *UnifiedRoleAssignmentScheduleRequestObjectRequest {
	return &UnifiedRoleAssignmentScheduleRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleAssignmentScheduleRequestObjectRequest is request for UnifiedRoleAssignmentScheduleRequestObject
type UnifiedRoleAssignmentScheduleRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRoleAssignmentScheduleRequestObject
func (r *UnifiedRoleAssignmentScheduleRequestObjectRequest) Get(ctx context.Context) (resObj *UnifiedRoleAssignmentScheduleRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRoleAssignmentScheduleRequestObject
func (r *UnifiedRoleAssignmentScheduleRequestObjectRequest) Update(ctx context.Context, reqObj *UnifiedRoleAssignmentScheduleRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRoleAssignmentScheduleRequestObject
func (r *UnifiedRoleAssignmentScheduleRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRoleDefinitionRequestBuilder is request builder for UnifiedRoleDefinition
type UnifiedRoleDefinitionRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleDefinitionRequest
func (b *UnifiedRoleDefinitionRequestBuilder) Request() *UnifiedRoleDefinitionRequest {
	return &UnifiedRoleDefinitionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleDefinitionRequest is request for UnifiedRoleDefinition
type UnifiedRoleDefinitionRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRoleDefinition
func (r *UnifiedRoleDefinitionRequest) Get(ctx context.Context) (resObj *UnifiedRoleDefinition, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRoleDefinition
func (r *UnifiedRoleDefinitionRequest) Update(ctx context.Context, reqObj *UnifiedRoleDefinition) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRoleDefinition
func (r *UnifiedRoleDefinitionRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRoleEligibilityScheduleRequestBuilder is request builder for UnifiedRoleEligibilitySchedule
type UnifiedRoleEligibilityScheduleRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleEligibilityScheduleRequest
func (b *UnifiedRoleEligibilityScheduleRequestBuilder) Request() *UnifiedRoleEligibilityScheduleRequest {
	return &UnifiedRoleEligibilityScheduleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleEligibilityScheduleRequest is request for UnifiedRoleEligibilitySchedule
type UnifiedRoleEligibilityScheduleRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRoleEligibilitySchedule
func (r *UnifiedRoleEligibilityScheduleRequest) Get(ctx context.Context) (resObj *UnifiedRoleEligibilitySchedule, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRoleEligibilitySchedule
func (r *UnifiedRoleEligibilityScheduleRequest) Update(ctx context.Context, reqObj *UnifiedRoleEligibilitySchedule) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRoleEligibilitySchedule
func (r *UnifiedRoleEligibilityScheduleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRoleEligibilityScheduleInstanceRequestBuilder is request builder for UnifiedRoleEligibilityScheduleInstance
type UnifiedRoleEligibilityScheduleInstanceRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleEligibilityScheduleInstanceRequest
func (b *UnifiedRoleEligibilityScheduleInstanceRequestBuilder) Request() *UnifiedRoleEligibilityScheduleInstanceRequest {
	return &UnifiedRoleEligibilityScheduleInstanceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleEligibilityScheduleInstanceRequest is request for UnifiedRoleEligibilityScheduleInstance
type UnifiedRoleEligibilityScheduleInstanceRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRoleEligibilityScheduleInstance
func (r *UnifiedRoleEligibilityScheduleInstanceRequest) Get(ctx context.Context) (resObj *UnifiedRoleEligibilityScheduleInstance, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRoleEligibilityScheduleInstance
func (r *UnifiedRoleEligibilityScheduleInstanceRequest) Update(ctx context.Context, reqObj *UnifiedRoleEligibilityScheduleInstance) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRoleEligibilityScheduleInstance
func (r *UnifiedRoleEligibilityScheduleInstanceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRoleEligibilityScheduleRequestObjectRequestBuilder is request builder for UnifiedRoleEligibilityScheduleRequestObject
type UnifiedRoleEligibilityScheduleRequestObjectRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleEligibilityScheduleRequestObjectRequest
func (b *UnifiedRoleEligibilityScheduleRequestObjectRequestBuilder) Request() *UnifiedRoleEligibilityScheduleRequestObjectRequest {
	return &UnifiedRoleEligibilityScheduleRequestObjectRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleEligibilityScheduleRequestObjectRequest is request for UnifiedRoleEligibilityScheduleRequestObject
type UnifiedRoleEligibilityScheduleRequestObjectRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRoleEligibilityScheduleRequestObject
func (r *UnifiedRoleEligibilityScheduleRequestObjectRequest) Get(ctx context.Context) (resObj *UnifiedRoleEligibilityScheduleRequestObject, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRoleEligibilityScheduleRequestObject
func (r *UnifiedRoleEligibilityScheduleRequestObjectRequest) Update(ctx context.Context, reqObj *UnifiedRoleEligibilityScheduleRequestObject) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRoleEligibilityScheduleRequestObject
func (r *UnifiedRoleEligibilityScheduleRequestObjectRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRoleManagementPolicyRequestBuilder is request builder for UnifiedRoleManagementPolicy
type UnifiedRoleManagementPolicyRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleManagementPolicyRequest
func (b *UnifiedRoleManagementPolicyRequestBuilder) Request() *UnifiedRoleManagementPolicyRequest {
	return &UnifiedRoleManagementPolicyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleManagementPolicyRequest is request for UnifiedRoleManagementPolicy
type UnifiedRoleManagementPolicyRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRoleManagementPolicy
func (r *UnifiedRoleManagementPolicyRequest) Get(ctx context.Context) (resObj *UnifiedRoleManagementPolicy, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRoleManagementPolicy
func (r *UnifiedRoleManagementPolicyRequest) Update(ctx context.Context, reqObj *UnifiedRoleManagementPolicy) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRoleManagementPolicy
func (r *UnifiedRoleManagementPolicyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRoleManagementPolicyAssignmentRequestBuilder is request builder for UnifiedRoleManagementPolicyAssignment
type UnifiedRoleManagementPolicyAssignmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleManagementPolicyAssignmentRequest
func (b *UnifiedRoleManagementPolicyAssignmentRequestBuilder) Request() *UnifiedRoleManagementPolicyAssignmentRequest {
	return &UnifiedRoleManagementPolicyAssignmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleManagementPolicyAssignmentRequest is request for UnifiedRoleManagementPolicyAssignment
type UnifiedRoleManagementPolicyAssignmentRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRoleManagementPolicyAssignment
func (r *UnifiedRoleManagementPolicyAssignmentRequest) Get(ctx context.Context) (resObj *UnifiedRoleManagementPolicyAssignment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRoleManagementPolicyAssignment
func (r *UnifiedRoleManagementPolicyAssignmentRequest) Update(ctx context.Context, reqObj *UnifiedRoleManagementPolicyAssignment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRoleManagementPolicyAssignment
func (r *UnifiedRoleManagementPolicyAssignmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRoleManagementPolicyRuleRequestBuilder is request builder for UnifiedRoleManagementPolicyRule
type UnifiedRoleManagementPolicyRuleRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleManagementPolicyRuleRequest
func (b *UnifiedRoleManagementPolicyRuleRequestBuilder) Request() *UnifiedRoleManagementPolicyRuleRequest {
	return &UnifiedRoleManagementPolicyRuleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleManagementPolicyRuleRequest is request for UnifiedRoleManagementPolicyRule
type UnifiedRoleManagementPolicyRuleRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRoleManagementPolicyRule
func (r *UnifiedRoleManagementPolicyRuleRequest) Get(ctx context.Context) (resObj *UnifiedRoleManagementPolicyRule, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRoleManagementPolicyRule
func (r *UnifiedRoleManagementPolicyRuleRequest) Update(ctx context.Context, reqObj *UnifiedRoleManagementPolicyRule) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRoleManagementPolicyRule
func (r *UnifiedRoleManagementPolicyRuleRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRoleManagementPolicyRuleTargetRequestBuilder is request builder for UnifiedRoleManagementPolicyRuleTarget
type UnifiedRoleManagementPolicyRuleTargetRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleManagementPolicyRuleTargetRequest
func (b *UnifiedRoleManagementPolicyRuleTargetRequestBuilder) Request() *UnifiedRoleManagementPolicyRuleTargetRequest {
	return &UnifiedRoleManagementPolicyRuleTargetRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleManagementPolicyRuleTargetRequest is request for UnifiedRoleManagementPolicyRuleTarget
type UnifiedRoleManagementPolicyRuleTargetRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRoleManagementPolicyRuleTarget
func (r *UnifiedRoleManagementPolicyRuleTargetRequest) Get(ctx context.Context) (resObj *UnifiedRoleManagementPolicyRuleTarget, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRoleManagementPolicyRuleTarget
func (r *UnifiedRoleManagementPolicyRuleTargetRequest) Update(ctx context.Context, reqObj *UnifiedRoleManagementPolicyRuleTarget) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRoleManagementPolicyRuleTarget
func (r *UnifiedRoleManagementPolicyRuleTargetRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRoleScheduleBaseRequestBuilder is request builder for UnifiedRoleScheduleBase
type UnifiedRoleScheduleBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleScheduleBaseRequest
func (b *UnifiedRoleScheduleBaseRequestBuilder) Request() *UnifiedRoleScheduleBaseRequest {
	return &UnifiedRoleScheduleBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleScheduleBaseRequest is request for UnifiedRoleScheduleBase
type UnifiedRoleScheduleBaseRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRoleScheduleBase
func (r *UnifiedRoleScheduleBaseRequest) Get(ctx context.Context) (resObj *UnifiedRoleScheduleBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRoleScheduleBase
func (r *UnifiedRoleScheduleBaseRequest) Update(ctx context.Context, reqObj *UnifiedRoleScheduleBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRoleScheduleBase
func (r *UnifiedRoleScheduleBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// UnifiedRoleScheduleInstanceBaseRequestBuilder is request builder for UnifiedRoleScheduleInstanceBase
type UnifiedRoleScheduleInstanceBaseRequestBuilder struct{ BaseRequestBuilder }

// Request returns UnifiedRoleScheduleInstanceBaseRequest
func (b *UnifiedRoleScheduleInstanceBaseRequestBuilder) Request() *UnifiedRoleScheduleInstanceBaseRequest {
	return &UnifiedRoleScheduleInstanceBaseRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UnifiedRoleScheduleInstanceBaseRequest is request for UnifiedRoleScheduleInstanceBase
type UnifiedRoleScheduleInstanceBaseRequest struct{ BaseRequest }

// Get performs GET request for UnifiedRoleScheduleInstanceBase
func (r *UnifiedRoleScheduleInstanceBaseRequest) Get(ctx context.Context) (resObj *UnifiedRoleScheduleInstanceBase, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UnifiedRoleScheduleInstanceBase
func (r *UnifiedRoleScheduleInstanceBaseRequest) Update(ctx context.Context, reqObj *UnifiedRoleScheduleInstanceBase) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UnifiedRoleScheduleInstanceBase
func (r *UnifiedRoleScheduleInstanceBaseRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
