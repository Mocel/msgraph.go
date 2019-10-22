// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// AppleUserInitiatedEnrollmentProfileRequestBuilder is request builder for AppleUserInitiatedEnrollmentProfile
type AppleUserInitiatedEnrollmentProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns AppleUserInitiatedEnrollmentProfileRequest
func (b *AppleUserInitiatedEnrollmentProfileRequestBuilder) Request() *AppleUserInitiatedEnrollmentProfileRequest {
	return &AppleUserInitiatedEnrollmentProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// AppleUserInitiatedEnrollmentProfileRequest is request for AppleUserInitiatedEnrollmentProfile
type AppleUserInitiatedEnrollmentProfileRequest struct{ BaseRequest }

// Get performs GET request for AppleUserInitiatedEnrollmentProfile
func (r *AppleUserInitiatedEnrollmentProfileRequest) Get() (resObj *AppleUserInitiatedEnrollmentProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for AppleUserInitiatedEnrollmentProfile
func (r *AppleUserInitiatedEnrollmentProfileRequest) Update(reqObj *AppleUserInitiatedEnrollmentProfile) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for AppleUserInitiatedEnrollmentProfile
func (r *AppleUserInitiatedEnrollmentProfileRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Assignments returns request builder for AppleEnrollmentProfileAssignment collection
func (b *AppleUserInitiatedEnrollmentProfileRequestBuilder) Assignments() *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequestBuilder {
	bb := &AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequestBuilder is request builder for AppleEnrollmentProfileAssignment collection
type AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AppleEnrollmentProfileAssignment collection
func (b *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequestBuilder) Request() *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest {
	return &AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AppleEnrollmentProfileAssignment item
func (b *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequestBuilder) ID(id string) *AppleEnrollmentProfileAssignmentRequestBuilder {
	bb := &AppleEnrollmentProfileAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest is request for AppleEnrollmentProfileAssignment collection
type AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AppleEnrollmentProfileAssignment collection
func (r *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]AppleEnrollmentProfileAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []AppleEnrollmentProfileAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []AppleEnrollmentProfileAssignment
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

// Get performs GET request for AppleEnrollmentProfileAssignment collection
func (r *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest) Get() ([]AppleEnrollmentProfileAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for AppleEnrollmentProfileAssignment collection
func (r *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest) Add(reqObj *AppleEnrollmentProfileAssignment) (resObj *AppleEnrollmentProfileAssignment, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
