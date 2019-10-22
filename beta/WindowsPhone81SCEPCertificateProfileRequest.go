// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// WindowsPhone81SCEPCertificateProfileRequestBuilder is request builder for WindowsPhone81SCEPCertificateProfile
type WindowsPhone81SCEPCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns WindowsPhone81SCEPCertificateProfileRequest
func (b *WindowsPhone81SCEPCertificateProfileRequestBuilder) Request() *WindowsPhone81SCEPCertificateProfileRequest {
	return &WindowsPhone81SCEPCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// WindowsPhone81SCEPCertificateProfileRequest is request for WindowsPhone81SCEPCertificateProfile
type WindowsPhone81SCEPCertificateProfileRequest struct{ BaseRequest }

// Get performs GET request for WindowsPhone81SCEPCertificateProfile
func (r *WindowsPhone81SCEPCertificateProfileRequest) Get() (resObj *WindowsPhone81SCEPCertificateProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for WindowsPhone81SCEPCertificateProfile
func (r *WindowsPhone81SCEPCertificateProfileRequest) Update(reqObj *WindowsPhone81SCEPCertificateProfile) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for WindowsPhone81SCEPCertificateProfile
func (r *WindowsPhone81SCEPCertificateProfileRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// ManagedDeviceCertificateStates returns request builder for ManagedDeviceCertificateState collection
func (b *WindowsPhone81SCEPCertificateProfileRequestBuilder) ManagedDeviceCertificateStates() *WindowsPhone81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder {
	bb := &WindowsPhone81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDeviceCertificateStates"
	return bb
}

// WindowsPhone81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder is request builder for ManagedDeviceCertificateState collection
type WindowsPhone81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDeviceCertificateState collection
func (b *WindowsPhone81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) Request() *WindowsPhone81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequest {
	return &WindowsPhone81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDeviceCertificateState item
func (b *WindowsPhone81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) ID(id string) *ManagedDeviceCertificateStateRequestBuilder {
	bb := &ManagedDeviceCertificateStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// WindowsPhone81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequest is request for ManagedDeviceCertificateState collection
type WindowsPhone81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedDeviceCertificateState collection
func (r *WindowsPhone81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Paging(method, path string, obj interface{}) ([]ManagedDeviceCertificateState, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []ManagedDeviceCertificateState
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []ManagedDeviceCertificateState
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

// Get performs GET request for ManagedDeviceCertificateState collection
func (r *WindowsPhone81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Get() ([]ManagedDeviceCertificateState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ManagedDeviceCertificateState collection
func (r *WindowsPhone81SCEPCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Add(reqObj *ManagedDeviceCertificateState) (resObj *ManagedDeviceCertificateState, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// RootCertificate is navigation property
func (b *WindowsPhone81SCEPCertificateProfileRequestBuilder) RootCertificate() *WindowsPhone81TrustedRootCertificateRequestBuilder {
	bb := &WindowsPhone81TrustedRootCertificateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rootCertificate"
	return bb
}
