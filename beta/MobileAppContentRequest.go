// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// MobileAppContentRequestBuilder is request builder for MobileAppContent
type MobileAppContentRequestBuilder struct{ BaseRequestBuilder }

// Request returns MobileAppContentRequest
func (b *MobileAppContentRequestBuilder) Request() *MobileAppContentRequest {
	return &MobileAppContentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// MobileAppContentRequest is request for MobileAppContent
type MobileAppContentRequest struct{ BaseRequest }

// Get performs GET request for MobileAppContent
func (r *MobileAppContentRequest) Get() (resObj *MobileAppContent, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for MobileAppContent
func (r *MobileAppContentRequest) Update(reqObj *MobileAppContent) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for MobileAppContent
func (r *MobileAppContentRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// ContainedApps returns request builder for MobileContainedApp collection
func (b *MobileAppContentRequestBuilder) ContainedApps() *MobileAppContentContainedAppsCollectionRequestBuilder {
	bb := &MobileAppContentContainedAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/containedApps"
	return bb
}

// MobileAppContentContainedAppsCollectionRequestBuilder is request builder for MobileContainedApp collection
type MobileAppContentContainedAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileContainedApp collection
func (b *MobileAppContentContainedAppsCollectionRequestBuilder) Request() *MobileAppContentContainedAppsCollectionRequest {
	return &MobileAppContentContainedAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileContainedApp item
func (b *MobileAppContentContainedAppsCollectionRequestBuilder) ID(id string) *MobileContainedAppRequestBuilder {
	bb := &MobileContainedAppRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppContentContainedAppsCollectionRequest is request for MobileContainedApp collection
type MobileAppContentContainedAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileContainedApp collection
func (r *MobileAppContentContainedAppsCollectionRequest) Paging(method, path string, obj interface{}) ([]MobileContainedApp, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MobileContainedApp
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []MobileContainedApp
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

// Get performs GET request for MobileContainedApp collection
func (r *MobileAppContentContainedAppsCollectionRequest) Get() ([]MobileContainedApp, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for MobileContainedApp collection
func (r *MobileAppContentContainedAppsCollectionRequest) Add(reqObj *MobileContainedApp) (resObj *MobileContainedApp, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Files returns request builder for MobileAppContentFile collection
func (b *MobileAppContentRequestBuilder) Files() *MobileAppContentFilesCollectionRequestBuilder {
	bb := &MobileAppContentFilesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/files"
	return bb
}

// MobileAppContentFilesCollectionRequestBuilder is request builder for MobileAppContentFile collection
type MobileAppContentFilesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for MobileAppContentFile collection
func (b *MobileAppContentFilesCollectionRequestBuilder) Request() *MobileAppContentFilesCollectionRequest {
	return &MobileAppContentFilesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for MobileAppContentFile item
func (b *MobileAppContentFilesCollectionRequestBuilder) ID(id string) *MobileAppContentFileRequestBuilder {
	bb := &MobileAppContentFileRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// MobileAppContentFilesCollectionRequest is request for MobileAppContentFile collection
type MobileAppContentFilesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for MobileAppContentFile collection
func (r *MobileAppContentFilesCollectionRequest) Paging(method, path string, obj interface{}) ([]MobileAppContentFile, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []MobileAppContentFile
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []MobileAppContentFile
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

// Get performs GET request for MobileAppContentFile collection
func (r *MobileAppContentFilesCollectionRequest) Get() ([]MobileAppContentFile, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for MobileAppContentFile collection
func (r *MobileAppContentFilesCollectionRequest) Add(reqObj *MobileAppContentFile) (resObj *MobileAppContentFile, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
