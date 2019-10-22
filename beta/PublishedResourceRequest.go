// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// PublishedResourceRequestBuilder is request builder for PublishedResource
type PublishedResourceRequestBuilder struct{ BaseRequestBuilder }

// Request returns PublishedResourceRequest
func (b *PublishedResourceRequestBuilder) Request() *PublishedResourceRequest {
	return &PublishedResourceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// PublishedResourceRequest is request for PublishedResource
type PublishedResourceRequest struct{ BaseRequest }

// Get performs GET request for PublishedResource
func (r *PublishedResourceRequest) Get() (resObj *PublishedResource, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for PublishedResource
func (r *PublishedResourceRequest) Update(reqObj *PublishedResource) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for PublishedResource
func (r *PublishedResourceRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// AgentGroups returns request builder for OnPremisesAgentGroup collection
func (b *PublishedResourceRequestBuilder) AgentGroups() *PublishedResourceAgentGroupsCollectionRequestBuilder {
	bb := &PublishedResourceAgentGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/agentGroups"
	return bb
}

// PublishedResourceAgentGroupsCollectionRequestBuilder is request builder for OnPremisesAgentGroup collection
type PublishedResourceAgentGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnPremisesAgentGroup collection
func (b *PublishedResourceAgentGroupsCollectionRequestBuilder) Request() *PublishedResourceAgentGroupsCollectionRequest {
	return &PublishedResourceAgentGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnPremisesAgentGroup item
func (b *PublishedResourceAgentGroupsCollectionRequestBuilder) ID(id string) *OnPremisesAgentGroupRequestBuilder {
	bb := &OnPremisesAgentGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// PublishedResourceAgentGroupsCollectionRequest is request for OnPremisesAgentGroup collection
type PublishedResourceAgentGroupsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnPremisesAgentGroup collection
func (r *PublishedResourceAgentGroupsCollectionRequest) Paging(method, path string, obj interface{}) ([]OnPremisesAgentGroup, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OnPremisesAgentGroup
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []OnPremisesAgentGroup
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

// Get performs GET request for OnPremisesAgentGroup collection
func (r *PublishedResourceAgentGroupsCollectionRequest) Get() ([]OnPremisesAgentGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for OnPremisesAgentGroup collection
func (r *PublishedResourceAgentGroupsCollectionRequest) Add(reqObj *OnPremisesAgentGroup) (resObj *OnPremisesAgentGroup, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
