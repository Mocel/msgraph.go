// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// WindowsInformationProtectionRequestBuilder is request builder for WindowsInformationProtection
type WindowsInformationProtectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsInformationProtectionRequest
func (b *WindowsInformationProtectionRequestBuilder) Request() *WindowsInformationProtectionRequest {
	return &WindowsInformationProtectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsInformationProtectionRequest is request for WindowsInformationProtection
type WindowsInformationProtectionRequest struct{ BaseRequest }

// Get performs GET request for WindowsInformationProtection
func (r *WindowsInformationProtectionRequest) Get() (resObj *WindowsInformationProtection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsInformationProtection
func (r *WindowsInformationProtectionRequest) Update(reqObj *WindowsInformationProtection) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsInformationProtection
func (r *WindowsInformationProtectionRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Assignments returns request builder for TargetedManagedAppPolicyAssignment collection
func (b *WindowsInformationProtectionRequestBuilder) Assignments() *WindowsInformationProtectionAssignmentsCollectionRequestBuilder {
	bb := &WindowsInformationProtectionAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// WindowsInformationProtectionAssignmentsCollectionRequestBuilder is request builder for TargetedManagedAppPolicyAssignment collection
type WindowsInformationProtectionAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TargetedManagedAppPolicyAssignment collection
func (b *WindowsInformationProtectionAssignmentsCollectionRequestBuilder) Request() *WindowsInformationProtectionAssignmentsCollectionRequest {
	return &WindowsInformationProtectionAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TargetedManagedAppPolicyAssignment item
func (b *WindowsInformationProtectionAssignmentsCollectionRequestBuilder) ID(id string) *TargetedManagedAppPolicyAssignmentRequestBuilder {
	bb := &TargetedManagedAppPolicyAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsInformationProtectionAssignmentsCollectionRequest is request for TargetedManagedAppPolicyAssignment collection
type WindowsInformationProtectionAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TargetedManagedAppPolicyAssignment collection
func (r *WindowsInformationProtectionAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]TargetedManagedAppPolicyAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TargetedManagedAppPolicyAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []TargetedManagedAppPolicyAssignment
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for TargetedManagedAppPolicyAssignment collection
func (r *WindowsInformationProtectionAssignmentsCollectionRequest) Get() ([]TargetedManagedAppPolicyAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for TargetedManagedAppPolicyAssignment collection
func (r *WindowsInformationProtectionAssignmentsCollectionRequest) Add(reqObj *TargetedManagedAppPolicyAssignment) (resObj *TargetedManagedAppPolicyAssignment, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// ExemptAppLockerFiles returns request builder for WindowsInformationProtectionAppLockerFile collection
func (b *WindowsInformationProtectionRequestBuilder) ExemptAppLockerFiles() *WindowsInformationProtectionExemptAppLockerFilesCollectionRequestBuilder {
	bb := &WindowsInformationProtectionExemptAppLockerFilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/exemptAppLockerFiles"
	return bb
}

// WindowsInformationProtectionExemptAppLockerFilesCollectionRequestBuilder is request builder for WindowsInformationProtectionAppLockerFile collection
type WindowsInformationProtectionExemptAppLockerFilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WindowsInformationProtectionAppLockerFile collection
func (b *WindowsInformationProtectionExemptAppLockerFilesCollectionRequestBuilder) Request() *WindowsInformationProtectionExemptAppLockerFilesCollectionRequest {
	return &WindowsInformationProtectionExemptAppLockerFilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WindowsInformationProtectionAppLockerFile item
func (b *WindowsInformationProtectionExemptAppLockerFilesCollectionRequestBuilder) ID(id string) *WindowsInformationProtectionAppLockerFileRequestBuilder {
	bb := &WindowsInformationProtectionAppLockerFileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsInformationProtectionExemptAppLockerFilesCollectionRequest is request for WindowsInformationProtectionAppLockerFile collection
type WindowsInformationProtectionExemptAppLockerFilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WindowsInformationProtectionAppLockerFile collection
func (r *WindowsInformationProtectionExemptAppLockerFilesCollectionRequest) Paging(method, path string, obj interface{}) ([]WindowsInformationProtectionAppLockerFile, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WindowsInformationProtectionAppLockerFile
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []WindowsInformationProtectionAppLockerFile
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for WindowsInformationProtectionAppLockerFile collection
func (r *WindowsInformationProtectionExemptAppLockerFilesCollectionRequest) Get() ([]WindowsInformationProtectionAppLockerFile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for WindowsInformationProtectionAppLockerFile collection
func (r *WindowsInformationProtectionExemptAppLockerFilesCollectionRequest) Add(reqObj *WindowsInformationProtectionAppLockerFile) (resObj *WindowsInformationProtectionAppLockerFile, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// ProtectedAppLockerFiles returns request builder for WindowsInformationProtectionAppLockerFile collection
func (b *WindowsInformationProtectionRequestBuilder) ProtectedAppLockerFiles() *WindowsInformationProtectionProtectedAppLockerFilesCollectionRequestBuilder {
	bb := &WindowsInformationProtectionProtectedAppLockerFilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/protectedAppLockerFiles"
	return bb
}

// WindowsInformationProtectionProtectedAppLockerFilesCollectionRequestBuilder is request builder for WindowsInformationProtectionAppLockerFile collection
type WindowsInformationProtectionProtectedAppLockerFilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WindowsInformationProtectionAppLockerFile collection
func (b *WindowsInformationProtectionProtectedAppLockerFilesCollectionRequestBuilder) Request() *WindowsInformationProtectionProtectedAppLockerFilesCollectionRequest {
	return &WindowsInformationProtectionProtectedAppLockerFilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WindowsInformationProtectionAppLockerFile item
func (b *WindowsInformationProtectionProtectedAppLockerFilesCollectionRequestBuilder) ID(id string) *WindowsInformationProtectionAppLockerFileRequestBuilder {
	bb := &WindowsInformationProtectionAppLockerFileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsInformationProtectionProtectedAppLockerFilesCollectionRequest is request for WindowsInformationProtectionAppLockerFile collection
type WindowsInformationProtectionProtectedAppLockerFilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WindowsInformationProtectionAppLockerFile collection
func (r *WindowsInformationProtectionProtectedAppLockerFilesCollectionRequest) Paging(method, path string, obj interface{}) ([]WindowsInformationProtectionAppLockerFile, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WindowsInformationProtectionAppLockerFile
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []WindowsInformationProtectionAppLockerFile
		)
		err := jsonx.NewDecoder(res.Body).Decode(&paging)
		if err != nil {
			return nil, err
		}
		err = jsonx.Unmarshal(paging.Value, &value)
		if err != nil {
			return nil, err
		}
		values = append(values, value...)
		if len(paging.NextLink) == 0 {
			return values, nil
		}
		res, err = r.client.Get(paging.NextLink)
		if err != nil {
			return nil, err
		}
	}
}

// Get performs GET request for WindowsInformationProtectionAppLockerFile collection
func (r *WindowsInformationProtectionProtectedAppLockerFilesCollectionRequest) Get() ([]WindowsInformationProtectionAppLockerFile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for WindowsInformationProtectionAppLockerFile collection
func (r *WindowsInformationProtectionProtectedAppLockerFilesCollectionRequest) Add(reqObj *WindowsInformationProtectionAppLockerFile) (resObj *WindowsInformationProtectionAppLockerFile, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
