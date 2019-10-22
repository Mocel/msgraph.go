// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// SiteCollectionAddRequestParameter undocumented
type SiteCollectionAddRequestParameter struct {
	// Value undocumented
	Value []Site `json:"value,omitempty"`
}

// SiteCollectionRemoveRequestParameter undocumented
type SiteCollectionRemoveRequestParameter struct {
	// Value undocumented
	Value []Site `json:"value,omitempty"`
}

//
type SiteCollectionAddRequestBuilder struct{ BaseRequestBuilder }

// Add action undocumented
func (b *GroupSitesCollectionRequestBuilder) Add(reqObj *SiteCollectionAddRequestParameter) *SiteCollectionAddRequestBuilder {
	bb := &SiteCollectionAddRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/add"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// Add action undocumented
func (b *SiteSitesCollectionRequestBuilder) Add(reqObj *SiteCollectionAddRequestParameter) *SiteCollectionAddRequestBuilder {
	bb := &SiteCollectionAddRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/add"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// Add action undocumented
func (b *UserFollowedSitesCollectionRequestBuilder) Add(reqObj *SiteCollectionAddRequestParameter) *SiteCollectionAddRequestBuilder {
	bb := &SiteCollectionAddRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/add"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SiteCollectionAddRequest struct{ BaseRequest }

//
func (b *SiteCollectionAddRequestBuilder) Request() *SiteCollectionAddRequest {
	return &SiteCollectionAddRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SiteCollectionAddRequest) Paging(method, path string, obj interface{}) ([][]Site, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]Site
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]Site
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

//
func (r *SiteCollectionAddRequest) Get() ([][]Site, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

//
type SiteCollectionRemoveRequestBuilder struct{ BaseRequestBuilder }

// Remove action undocumented
func (b *GroupSitesCollectionRequestBuilder) Remove(reqObj *SiteCollectionRemoveRequestParameter) *SiteCollectionRemoveRequestBuilder {
	bb := &SiteCollectionRemoveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/remove"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// Remove action undocumented
func (b *SiteSitesCollectionRequestBuilder) Remove(reqObj *SiteCollectionRemoveRequestParameter) *SiteCollectionRemoveRequestBuilder {
	bb := &SiteCollectionRemoveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/remove"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

// Remove action undocumented
func (b *UserFollowedSitesCollectionRequestBuilder) Remove(reqObj *SiteCollectionRemoveRequestParameter) *SiteCollectionRemoveRequestBuilder {
	bb := &SiteCollectionRemoveRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/remove"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SiteCollectionRemoveRequest struct{ BaseRequest }

//
func (b *SiteCollectionRemoveRequestBuilder) Request() *SiteCollectionRemoveRequest {
	return &SiteCollectionRemoveRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SiteCollectionRemoveRequest) Paging(method, path string, obj interface{}) ([][]Site, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]Site
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]Site
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

//
func (r *SiteCollectionRemoveRequest) Get() ([][]Site, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}
