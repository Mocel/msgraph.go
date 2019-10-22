// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// WindowsAutopilotDeploymentProfileRequestBuilder is request builder for WindowsAutopilotDeploymentProfile
type WindowsAutopilotDeploymentProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsAutopilotDeploymentProfileRequest
func (b *WindowsAutopilotDeploymentProfileRequestBuilder) Request() *WindowsAutopilotDeploymentProfileRequest {
	return &WindowsAutopilotDeploymentProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsAutopilotDeploymentProfileRequest is request for WindowsAutopilotDeploymentProfile
type WindowsAutopilotDeploymentProfileRequest struct{ BaseRequest }

// Get performs GET request for WindowsAutopilotDeploymentProfile
func (r *WindowsAutopilotDeploymentProfileRequest) Get() (resObj *WindowsAutopilotDeploymentProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsAutopilotDeploymentProfile
func (r *WindowsAutopilotDeploymentProfileRequest) Update(reqObj *WindowsAutopilotDeploymentProfile) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsAutopilotDeploymentProfile
func (r *WindowsAutopilotDeploymentProfileRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// AssignedDevices returns request builder for WindowsAutopilotDeviceIdentity collection
func (b *WindowsAutopilotDeploymentProfileRequestBuilder) AssignedDevices() *WindowsAutopilotDeploymentProfileAssignedDevicesCollectionRequestBuilder {
	bb := &WindowsAutopilotDeploymentProfileAssignedDevicesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignedDevices"
	return bb
}

// WindowsAutopilotDeploymentProfileAssignedDevicesCollectionRequestBuilder is request builder for WindowsAutopilotDeviceIdentity collection
type WindowsAutopilotDeploymentProfileAssignedDevicesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WindowsAutopilotDeviceIdentity collection
func (b *WindowsAutopilotDeploymentProfileAssignedDevicesCollectionRequestBuilder) Request() *WindowsAutopilotDeploymentProfileAssignedDevicesCollectionRequest {
	return &WindowsAutopilotDeploymentProfileAssignedDevicesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WindowsAutopilotDeviceIdentity item
func (b *WindowsAutopilotDeploymentProfileAssignedDevicesCollectionRequestBuilder) ID(id string) *WindowsAutopilotDeviceIdentityRequestBuilder {
	bb := &WindowsAutopilotDeviceIdentityRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsAutopilotDeploymentProfileAssignedDevicesCollectionRequest is request for WindowsAutopilotDeviceIdentity collection
type WindowsAutopilotDeploymentProfileAssignedDevicesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WindowsAutopilotDeviceIdentity collection
func (r *WindowsAutopilotDeploymentProfileAssignedDevicesCollectionRequest) Paging(method, path string, obj interface{}) ([]WindowsAutopilotDeviceIdentity, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WindowsAutopilotDeviceIdentity
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []WindowsAutopilotDeviceIdentity
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

// Get performs GET request for WindowsAutopilotDeviceIdentity collection
func (r *WindowsAutopilotDeploymentProfileAssignedDevicesCollectionRequest) Get() ([]WindowsAutopilotDeviceIdentity, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for WindowsAutopilotDeviceIdentity collection
func (r *WindowsAutopilotDeploymentProfileAssignedDevicesCollectionRequest) Add(reqObj *WindowsAutopilotDeviceIdentity) (resObj *WindowsAutopilotDeviceIdentity, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Assignments returns request builder for WindowsAutopilotDeploymentProfileAssignment collection
func (b *WindowsAutopilotDeploymentProfileRequestBuilder) Assignments() *WindowsAutopilotDeploymentProfileAssignmentsCollectionRequestBuilder {
	bb := &WindowsAutopilotDeploymentProfileAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// WindowsAutopilotDeploymentProfileAssignmentsCollectionRequestBuilder is request builder for WindowsAutopilotDeploymentProfileAssignment collection
type WindowsAutopilotDeploymentProfileAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for WindowsAutopilotDeploymentProfileAssignment collection
func (b *WindowsAutopilotDeploymentProfileAssignmentsCollectionRequestBuilder) Request() *WindowsAutopilotDeploymentProfileAssignmentsCollectionRequest {
	return &WindowsAutopilotDeploymentProfileAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for WindowsAutopilotDeploymentProfileAssignment item
func (b *WindowsAutopilotDeploymentProfileAssignmentsCollectionRequestBuilder) ID(id string) *WindowsAutopilotDeploymentProfileAssignmentRequestBuilder {
	bb := &WindowsAutopilotDeploymentProfileAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsAutopilotDeploymentProfileAssignmentsCollectionRequest is request for WindowsAutopilotDeploymentProfileAssignment collection
type WindowsAutopilotDeploymentProfileAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for WindowsAutopilotDeploymentProfileAssignment collection
func (r *WindowsAutopilotDeploymentProfileAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]WindowsAutopilotDeploymentProfileAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []WindowsAutopilotDeploymentProfileAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []WindowsAutopilotDeploymentProfileAssignment
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

// Get performs GET request for WindowsAutopilotDeploymentProfileAssignment collection
func (r *WindowsAutopilotDeploymentProfileAssignmentsCollectionRequest) Get() ([]WindowsAutopilotDeploymentProfileAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for WindowsAutopilotDeploymentProfileAssignment collection
func (r *WindowsAutopilotDeploymentProfileAssignmentsCollectionRequest) Add(reqObj *WindowsAutopilotDeploymentProfileAssignment) (resObj *WindowsAutopilotDeploymentProfileAssignment, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
