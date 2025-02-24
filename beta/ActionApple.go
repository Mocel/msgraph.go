// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestActionRequestParameter undocumented
type ApplePushNotificationCertificateGenerateApplePushNotificationCertificateSigningRequestActionRequestParameter struct {
}

// AppleUserInitiatedEnrollmentProfileSetPriorityRequestParameter undocumented
type AppleUserInitiatedEnrollmentProfileSetPriorityRequestParameter struct {
	// Priority undocumented
	Priority *int `json:"priority,omitempty"`
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
func (r *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]AppleEnrollmentProfileAssignment, error) {
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
	var values []AppleEnrollmentProfileAssignment
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
			value  []AppleEnrollmentProfileAssignment
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

// GetN performs GET request for AppleEnrollmentProfileAssignment collection, max N pages
func (r *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest) GetN(ctx context.Context, n int) ([]AppleEnrollmentProfileAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for AppleEnrollmentProfileAssignment collection
func (r *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest) Get(ctx context.Context) ([]AppleEnrollmentProfileAssignment, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for AppleEnrollmentProfileAssignment collection
func (r *AppleUserInitiatedEnrollmentProfileAssignmentsCollectionRequest) Add(ctx context.Context, reqObj *AppleEnrollmentProfileAssignment) (resObj *AppleEnrollmentProfileAssignment, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
