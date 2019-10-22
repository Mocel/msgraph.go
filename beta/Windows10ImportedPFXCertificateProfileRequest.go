// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// Windows10ImportedPFXCertificateProfileRequestBuilder is request builder for Windows10ImportedPFXCertificateProfile
type Windows10ImportedPFXCertificateProfileRequestBuilder struct{ BaseRequestBuilder }

// Request returns Windows10ImportedPFXCertificateProfileRequest
func (b *Windows10ImportedPFXCertificateProfileRequestBuilder) Request() *Windows10ImportedPFXCertificateProfileRequest {
	return &Windows10ImportedPFXCertificateProfileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// Windows10ImportedPFXCertificateProfileRequest is request for Windows10ImportedPFXCertificateProfile
type Windows10ImportedPFXCertificateProfileRequest struct{ BaseRequest }

// Get performs GET request for Windows10ImportedPFXCertificateProfile
func (r *Windows10ImportedPFXCertificateProfileRequest) Get() (resObj *Windows10ImportedPFXCertificateProfile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Windows10ImportedPFXCertificateProfile
func (r *Windows10ImportedPFXCertificateProfileRequest) Update(reqObj *Windows10ImportedPFXCertificateProfile) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Windows10ImportedPFXCertificateProfile
func (r *Windows10ImportedPFXCertificateProfileRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// ManagedDeviceCertificateStates returns request builder for ManagedDeviceCertificateState collection
func (b *Windows10ImportedPFXCertificateProfileRequestBuilder) ManagedDeviceCertificateStates() *Windows10ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder {
	bb := &Windows10ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/managedDeviceCertificateStates"
	return bb
}

// Windows10ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder is request builder for ManagedDeviceCertificateState collection
type Windows10ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for ManagedDeviceCertificateState collection
func (b *Windows10ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) Request() *Windows10ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest {
	return &Windows10ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for ManagedDeviceCertificateState item
func (b *Windows10ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequestBuilder) ID(id string) *ManagedDeviceCertificateStateRequestBuilder {
	bb := &ManagedDeviceCertificateStateRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// Windows10ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest is request for ManagedDeviceCertificateState collection
type Windows10ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for ManagedDeviceCertificateState collection
func (r *Windows10ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Paging(method, path string, obj interface{}) ([]ManagedDeviceCertificateState, error) {
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
func (r *Windows10ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Get() ([]ManagedDeviceCertificateState, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for ManagedDeviceCertificateState collection
func (r *Windows10ImportedPFXCertificateProfileManagedDeviceCertificateStatesCollectionRequest) Add(reqObj *ManagedDeviceCertificateState) (resObj *ManagedDeviceCertificateState, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
