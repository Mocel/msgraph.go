// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// TrustFrameworkRequestBuilder is request builder for TrustFramework
type TrustFrameworkRequestBuilder struct{ BaseRequestBuilder }

// Request returns TrustFrameworkRequest
func (b *TrustFrameworkRequestBuilder) Request() *TrustFrameworkRequest {
	return &TrustFrameworkRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// TrustFrameworkRequest is request for TrustFramework
type TrustFrameworkRequest struct{ BaseRequest }

// Get performs GET request for TrustFramework
func (r *TrustFrameworkRequest) Get() (resObj *TrustFramework, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for TrustFramework
func (r *TrustFrameworkRequest) Update(reqObj *TrustFramework) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for TrustFramework
func (r *TrustFrameworkRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// KeySets returns request builder for TrustFrameworkKeySet collection
func (b *TrustFrameworkRequestBuilder) KeySets() *TrustFrameworkKeySetsCollectionRequestBuilder {
	bb := &TrustFrameworkKeySetsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/keySets"
	return bb
}

// TrustFrameworkKeySetsCollectionRequestBuilder is request builder for TrustFrameworkKeySet collection
type TrustFrameworkKeySetsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TrustFrameworkKeySet collection
func (b *TrustFrameworkKeySetsCollectionRequestBuilder) Request() *TrustFrameworkKeySetsCollectionRequest {
	return &TrustFrameworkKeySetsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TrustFrameworkKeySet item
func (b *TrustFrameworkKeySetsCollectionRequestBuilder) ID(id string) *TrustFrameworkKeySetRequestBuilder {
	bb := &TrustFrameworkKeySetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TrustFrameworkKeySetsCollectionRequest is request for TrustFrameworkKeySet collection
type TrustFrameworkKeySetsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TrustFrameworkKeySet collection
func (r *TrustFrameworkKeySetsCollectionRequest) Paging(method, path string, obj interface{}) ([]TrustFrameworkKeySet, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TrustFrameworkKeySet
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []TrustFrameworkKeySet
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

// Get performs GET request for TrustFrameworkKeySet collection
func (r *TrustFrameworkKeySetsCollectionRequest) Get() ([]TrustFrameworkKeySet, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for TrustFrameworkKeySet collection
func (r *TrustFrameworkKeySetsCollectionRequest) Add(reqObj *TrustFrameworkKeySet) (resObj *TrustFrameworkKeySet, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Policies returns request builder for TrustFrameworkPolicy collection
func (b *TrustFrameworkRequestBuilder) Policies() *TrustFrameworkPoliciesCollectionRequestBuilder {
	bb := &TrustFrameworkPoliciesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/policies"
	return bb
}

// TrustFrameworkPoliciesCollectionRequestBuilder is request builder for TrustFrameworkPolicy collection
type TrustFrameworkPoliciesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TrustFrameworkPolicy collection
func (b *TrustFrameworkPoliciesCollectionRequestBuilder) Request() *TrustFrameworkPoliciesCollectionRequest {
	return &TrustFrameworkPoliciesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TrustFrameworkPolicy item
func (b *TrustFrameworkPoliciesCollectionRequestBuilder) ID(id string) *TrustFrameworkPolicyRequestBuilder {
	bb := &TrustFrameworkPolicyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// TrustFrameworkPoliciesCollectionRequest is request for TrustFrameworkPolicy collection
type TrustFrameworkPoliciesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TrustFrameworkPolicy collection
func (r *TrustFrameworkPoliciesCollectionRequest) Paging(method, path string, obj interface{}) ([]TrustFrameworkPolicy, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TrustFrameworkPolicy
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []TrustFrameworkPolicy
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

// Get performs GET request for TrustFrameworkPolicy collection
func (r *TrustFrameworkPoliciesCollectionRequest) Get() ([]TrustFrameworkPolicy, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for TrustFrameworkPolicy collection
func (r *TrustFrameworkPoliciesCollectionRequest) Add(reqObj *TrustFrameworkPolicy) (resObj *TrustFrameworkPolicy, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
