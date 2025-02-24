// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// PrivilegedSignupStatusCollectionCompleteSetupRequestParameter undocumented
type PrivilegedSignupStatusCollectionCompleteSetupRequestParameter struct {
	// TenantSetupInfo undocumented
	TenantSetupInfo *TenantSetupInfo `json:"tenantSetupInfo,omitempty"`
}

// PrivilegedSignupStatusCollectionSignUpRequestParameter undocumented
type PrivilegedSignupStatusCollectionSignUpRequestParameter struct {
}

// PrivilegedRoleSelfActivateRequestParameter undocumented
type PrivilegedRoleSelfActivateRequestParameter struct {
	// Reason undocumented
	Reason *string `json:"reason,omitempty"`
	// Duration undocumented
	Duration *string `json:"duration,omitempty"`
	// TicketNumber undocumented
	TicketNumber *string `json:"ticketNumber,omitempty"`
	// TicketSystem undocumented
	TicketSystem *string `json:"ticketSystem,omitempty"`
}

// PrivilegedRoleSelfDeactivateRequestParameter undocumented
type PrivilegedRoleSelfDeactivateRequestParameter struct {
}

// PrivilegedRoleAssignmentMakeEligibleRequestParameter undocumented
type PrivilegedRoleAssignmentMakeEligibleRequestParameter struct {
}

// PrivilegedRoleAssignmentMakePermanentRequestParameter undocumented
type PrivilegedRoleAssignmentMakePermanentRequestParameter struct {
	// Reason undocumented
	Reason *string `json:"reason,omitempty"`
	// TicketNumber undocumented
	TicketNumber *string `json:"ticketNumber,omitempty"`
	// TicketSystem undocumented
	TicketSystem *string `json:"ticketSystem,omitempty"`
}

// PrivilegedRoleAssignmentRequestObjectCancelRequestParameter undocumented
type PrivilegedRoleAssignmentRequestObjectCancelRequestParameter struct {
}

// Resources returns request builder for GovernanceResource collection
func (b *PrivilegedAccessRequestBuilder) Resources() *PrivilegedAccessResourcesCollectionRequestBuilder {
	bb := &PrivilegedAccessResourcesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/resources"
	return bb
}

// PrivilegedAccessResourcesCollectionRequestBuilder is request builder for GovernanceResource collection
type PrivilegedAccessResourcesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceResource collection
func (b *PrivilegedAccessResourcesCollectionRequestBuilder) Request() *PrivilegedAccessResourcesCollectionRequest {
	return &PrivilegedAccessResourcesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceResource item
func (b *PrivilegedAccessResourcesCollectionRequestBuilder) ID(id string) *GovernanceResourceRequestBuilder {
	bb := &GovernanceResourceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PrivilegedAccessResourcesCollectionRequest is request for GovernanceResource collection
type PrivilegedAccessResourcesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceResource collection
func (r *PrivilegedAccessResourcesCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]GovernanceResource, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []GovernanceResource
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []GovernanceResource
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for GovernanceResource collection, max N pages
func (r *PrivilegedAccessResourcesCollectionRequest) GetN(ctx context.Context, n int) ([]GovernanceResource, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for GovernanceResource collection
func (r *PrivilegedAccessResourcesCollectionRequest) Get(ctx context.Context) ([]GovernanceResource, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for GovernanceResource collection
func (r *PrivilegedAccessResourcesCollectionRequest) Add(ctx context.Context, reqObj *GovernanceResource) (resObj *GovernanceResource, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RoleAssignmentRequests returns request builder for GovernanceRoleAssignmentRequestObject collection
func (b *PrivilegedAccessRequestBuilder) RoleAssignmentRequests() *PrivilegedAccessRoleAssignmentRequestsCollectionRequestBuilder {
	bb := &PrivilegedAccessRoleAssignmentRequestsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleAssignmentRequests"
	return bb
}

// PrivilegedAccessRoleAssignmentRequestsCollectionRequestBuilder is request builder for GovernanceRoleAssignmentRequestObject collection
type PrivilegedAccessRoleAssignmentRequestsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceRoleAssignmentRequestObject collection
func (b *PrivilegedAccessRoleAssignmentRequestsCollectionRequestBuilder) Request() *PrivilegedAccessRoleAssignmentRequestsCollectionRequest {
	return &PrivilegedAccessRoleAssignmentRequestsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceRoleAssignmentRequestObject item
func (b *PrivilegedAccessRoleAssignmentRequestsCollectionRequestBuilder) ID(id string) *GovernanceRoleAssignmentRequestObjectRequestBuilder {
	bb := &GovernanceRoleAssignmentRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PrivilegedAccessRoleAssignmentRequestsCollectionRequest is request for GovernanceRoleAssignmentRequestObject collection
type PrivilegedAccessRoleAssignmentRequestsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceRoleAssignmentRequestObject collection
func (r *PrivilegedAccessRoleAssignmentRequestsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]GovernanceRoleAssignmentRequestObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []GovernanceRoleAssignmentRequestObject
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []GovernanceRoleAssignmentRequestObject
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for GovernanceRoleAssignmentRequestObject collection, max N pages
func (r *PrivilegedAccessRoleAssignmentRequestsCollectionRequest) GetN(ctx context.Context, n int) ([]GovernanceRoleAssignmentRequestObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for GovernanceRoleAssignmentRequestObject collection
func (r *PrivilegedAccessRoleAssignmentRequestsCollectionRequest) Get(ctx context.Context) ([]GovernanceRoleAssignmentRequestObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for GovernanceRoleAssignmentRequestObject collection
func (r *PrivilegedAccessRoleAssignmentRequestsCollectionRequest) Add(ctx context.Context, reqObj *GovernanceRoleAssignmentRequestObject) (resObj *GovernanceRoleAssignmentRequestObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RoleAssignments returns request builder for GovernanceRoleAssignment collection
func (b *PrivilegedAccessRequestBuilder) RoleAssignments() *PrivilegedAccessRoleAssignmentsCollectionRequestBuilder {
	bb := &PrivilegedAccessRoleAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleAssignments"
	return bb
}

// PrivilegedAccessRoleAssignmentsCollectionRequestBuilder is request builder for GovernanceRoleAssignment collection
type PrivilegedAccessRoleAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceRoleAssignment collection
func (b *PrivilegedAccessRoleAssignmentsCollectionRequestBuilder) Request() *PrivilegedAccessRoleAssignmentsCollectionRequest {
	return &PrivilegedAccessRoleAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceRoleAssignment item
func (b *PrivilegedAccessRoleAssignmentsCollectionRequestBuilder) ID(id string) *GovernanceRoleAssignmentRequestBuilder {
	bb := &GovernanceRoleAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PrivilegedAccessRoleAssignmentsCollectionRequest is request for GovernanceRoleAssignment collection
type PrivilegedAccessRoleAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceRoleAssignment collection
func (r *PrivilegedAccessRoleAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]GovernanceRoleAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []GovernanceRoleAssignment
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []GovernanceRoleAssignment
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for GovernanceRoleAssignment collection, max N pages
func (r *PrivilegedAccessRoleAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]GovernanceRoleAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for GovernanceRoleAssignment collection
func (r *PrivilegedAccessRoleAssignmentsCollectionRequest) Get(ctx context.Context) ([]GovernanceRoleAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for GovernanceRoleAssignment collection
func (r *PrivilegedAccessRoleAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *GovernanceRoleAssignment) (resObj *GovernanceRoleAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RoleDefinitions returns request builder for GovernanceRoleDefinition collection
func (b *PrivilegedAccessRequestBuilder) RoleDefinitions() *PrivilegedAccessRoleDefinitionsCollectionRequestBuilder {
	bb := &PrivilegedAccessRoleDefinitionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleDefinitions"
	return bb
}

// PrivilegedAccessRoleDefinitionsCollectionRequestBuilder is request builder for GovernanceRoleDefinition collection
type PrivilegedAccessRoleDefinitionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceRoleDefinition collection
func (b *PrivilegedAccessRoleDefinitionsCollectionRequestBuilder) Request() *PrivilegedAccessRoleDefinitionsCollectionRequest {
	return &PrivilegedAccessRoleDefinitionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceRoleDefinition item
func (b *PrivilegedAccessRoleDefinitionsCollectionRequestBuilder) ID(id string) *GovernanceRoleDefinitionRequestBuilder {
	bb := &GovernanceRoleDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PrivilegedAccessRoleDefinitionsCollectionRequest is request for GovernanceRoleDefinition collection
type PrivilegedAccessRoleDefinitionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceRoleDefinition collection
func (r *PrivilegedAccessRoleDefinitionsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]GovernanceRoleDefinition, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []GovernanceRoleDefinition
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []GovernanceRoleDefinition
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for GovernanceRoleDefinition collection, max N pages
func (r *PrivilegedAccessRoleDefinitionsCollectionRequest) GetN(ctx context.Context, n int) ([]GovernanceRoleDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for GovernanceRoleDefinition collection
func (r *PrivilegedAccessRoleDefinitionsCollectionRequest) Get(ctx context.Context) ([]GovernanceRoleDefinition, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for GovernanceRoleDefinition collection
func (r *PrivilegedAccessRoleDefinitionsCollectionRequest) Add(ctx context.Context, reqObj *GovernanceRoleDefinition) (resObj *GovernanceRoleDefinition, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RoleSettings returns request builder for GovernanceRoleSetting collection
func (b *PrivilegedAccessRequestBuilder) RoleSettings() *PrivilegedAccessRoleSettingsCollectionRequestBuilder {
	bb := &PrivilegedAccessRoleSettingsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleSettings"
	return bb
}

// PrivilegedAccessRoleSettingsCollectionRequestBuilder is request builder for GovernanceRoleSetting collection
type PrivilegedAccessRoleSettingsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GovernanceRoleSetting collection
func (b *PrivilegedAccessRoleSettingsCollectionRequestBuilder) Request() *PrivilegedAccessRoleSettingsCollectionRequest {
	return &PrivilegedAccessRoleSettingsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GovernanceRoleSetting item
func (b *PrivilegedAccessRoleSettingsCollectionRequestBuilder) ID(id string) *GovernanceRoleSettingRequestBuilder {
	bb := &GovernanceRoleSettingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PrivilegedAccessRoleSettingsCollectionRequest is request for GovernanceRoleSetting collection
type PrivilegedAccessRoleSettingsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GovernanceRoleSetting collection
func (r *PrivilegedAccessRoleSettingsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]GovernanceRoleSetting, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []GovernanceRoleSetting
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []GovernanceRoleSetting
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for GovernanceRoleSetting collection, max N pages
func (r *PrivilegedAccessRoleSettingsCollectionRequest) GetN(ctx context.Context, n int) ([]GovernanceRoleSetting, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for GovernanceRoleSetting collection
func (r *PrivilegedAccessRoleSettingsCollectionRequest) Get(ctx context.Context) ([]GovernanceRoleSetting, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for GovernanceRoleSetting collection
func (r *PrivilegedAccessRoleSettingsCollectionRequest) Add(ctx context.Context, reqObj *GovernanceRoleSetting) (resObj *GovernanceRoleSetting, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// RequestNavigation is navigation property
func (b *PrivilegedApprovalRequestBuilder) RequestNavigation() *PrivilegedRoleAssignmentRequestObjectRequestBuilder {
	bb := &PrivilegedRoleAssignmentRequestObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/request"
	return bb
}

// RoleInfo is navigation property
func (b *PrivilegedApprovalRequestBuilder) RoleInfo() *PrivilegedRoleRequestBuilder {
	bb := &PrivilegedRoleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleInfo"
	return bb
}

// Assignments returns request builder for PrivilegedRoleAssignment collection
func (b *PrivilegedRoleRequestBuilder) Assignments() *PrivilegedRoleAssignmentsCollectionRequestBuilder {
	bb := &PrivilegedRoleAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// PrivilegedRoleAssignmentsCollectionRequestBuilder is request builder for PrivilegedRoleAssignment collection
type PrivilegedRoleAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for PrivilegedRoleAssignment collection
func (b *PrivilegedRoleAssignmentsCollectionRequestBuilder) Request() *PrivilegedRoleAssignmentsCollectionRequest {
	return &PrivilegedRoleAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for PrivilegedRoleAssignment item
func (b *PrivilegedRoleAssignmentsCollectionRequestBuilder) ID(id string) *PrivilegedRoleAssignmentRequestBuilder {
	bb := &PrivilegedRoleAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PrivilegedRoleAssignmentsCollectionRequest is request for PrivilegedRoleAssignment collection
type PrivilegedRoleAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for PrivilegedRoleAssignment collection
func (r *PrivilegedRoleAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]PrivilegedRoleAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []PrivilegedRoleAssignment
	for {
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			res.Body.Close()
			errRes := &ErrorResponse{Response: res}
			err := jsonx.Unmarshal(b, errRes)
			if err != nil {
				return nil, fmt.Errorf("%s: %s", res.Status, string(b))
			}
			return nil, errRes
		}
		var (
			paging Paging
			value  []PrivilegedRoleAssignment
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		res.Body.Close()
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if n >= 0 {
			n--
		}
		if n == 0 || len(paging.NextLink) == 0 {
			return values, nil
		}
		req, err = http.NewRequest("GET", paging.NextLink, nil)
		if ctx != nil {
			req = req.WithContext(ctx)
		}
		res, err = r.client.Do(req)
		if err != nil {
			return nil, err
		}
	}
}

// GetN performs GET request for PrivilegedRoleAssignment collection, max N pages
func (r *PrivilegedRoleAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]PrivilegedRoleAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for PrivilegedRoleAssignment collection
func (r *PrivilegedRoleAssignmentsCollectionRequest) Get(ctx context.Context) ([]PrivilegedRoleAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for PrivilegedRoleAssignment collection
func (r *PrivilegedRoleAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *PrivilegedRoleAssignment) (resObj *PrivilegedRoleAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// Settings is navigation property
func (b *PrivilegedRoleRequestBuilder) Settings() *PrivilegedRoleSettingsRequestBuilder {
	bb := &PrivilegedRoleSettingsRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/settings"
	return bb
}

// Summary is navigation property
func (b *PrivilegedRoleRequestBuilder) Summary() *PrivilegedRoleSummaryRequestBuilder {
	bb := &PrivilegedRoleSummaryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/summary"
	return bb
}

// RoleInfo is navigation property
func (b *PrivilegedRoleAssignmentRequestBuilder) RoleInfo() *PrivilegedRoleRequestBuilder {
	bb := &PrivilegedRoleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleInfo"
	return bb
}

// RoleInfo is navigation property
func (b *PrivilegedRoleAssignmentRequestObjectRequestBuilder) RoleInfo() *PrivilegedRoleRequestBuilder {
	bb := &PrivilegedRoleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/roleInfo"
	return bb
}
