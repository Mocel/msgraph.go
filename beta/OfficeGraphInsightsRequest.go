// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// OfficeGraphInsightsRequestBuilder is request builder for OfficeGraphInsights
type OfficeGraphInsightsRequestBuilder struct{ BaseRequestBuilder }

// Request returns OfficeGraphInsightsRequest
func (b *OfficeGraphInsightsRequestBuilder) Request() *OfficeGraphInsightsRequest {
	return &OfficeGraphInsightsRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OfficeGraphInsightsRequest is request for OfficeGraphInsights
type OfficeGraphInsightsRequest struct{ BaseRequest }

// Get performs GET request for OfficeGraphInsights
func (r *OfficeGraphInsightsRequest) Get() (resObj *OfficeGraphInsights, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OfficeGraphInsights
func (r *OfficeGraphInsightsRequest) Update(reqObj *OfficeGraphInsights) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OfficeGraphInsights
func (r *OfficeGraphInsightsRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Shared returns request builder for SharedInsight collection
func (b *OfficeGraphInsightsRequestBuilder) Shared() *OfficeGraphInsightsSharedCollectionRequestBuilder {
	bb := &OfficeGraphInsightsSharedCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/shared"
	return bb
}

// OfficeGraphInsightsSharedCollectionRequestBuilder is request builder for SharedInsight collection
type OfficeGraphInsightsSharedCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SharedInsight collection
func (b *OfficeGraphInsightsSharedCollectionRequestBuilder) Request() *OfficeGraphInsightsSharedCollectionRequest {
	return &OfficeGraphInsightsSharedCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SharedInsight item
func (b *OfficeGraphInsightsSharedCollectionRequestBuilder) ID(id string) *SharedInsightRequestBuilder {
	bb := &SharedInsightRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OfficeGraphInsightsSharedCollectionRequest is request for SharedInsight collection
type OfficeGraphInsightsSharedCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SharedInsight collection
func (r *OfficeGraphInsightsSharedCollectionRequest) Paging(method, path string, obj interface{}) ([]SharedInsight, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SharedInsight
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []SharedInsight
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

// Get performs GET request for SharedInsight collection
func (r *OfficeGraphInsightsSharedCollectionRequest) Get() ([]SharedInsight, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SharedInsight collection
func (r *OfficeGraphInsightsSharedCollectionRequest) Add(reqObj *SharedInsight) (resObj *SharedInsight, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Trending returns request builder for Trending collection
func (b *OfficeGraphInsightsRequestBuilder) Trending() *OfficeGraphInsightsTrendingCollectionRequestBuilder {
	bb := &OfficeGraphInsightsTrendingCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/trending"
	return bb
}

// OfficeGraphInsightsTrendingCollectionRequestBuilder is request builder for Trending collection
type OfficeGraphInsightsTrendingCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Trending collection
func (b *OfficeGraphInsightsTrendingCollectionRequestBuilder) Request() *OfficeGraphInsightsTrendingCollectionRequest {
	return &OfficeGraphInsightsTrendingCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Trending item
func (b *OfficeGraphInsightsTrendingCollectionRequestBuilder) ID(id string) *TrendingRequestBuilder {
	bb := &TrendingRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OfficeGraphInsightsTrendingCollectionRequest is request for Trending collection
type OfficeGraphInsightsTrendingCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Trending collection
func (r *OfficeGraphInsightsTrendingCollectionRequest) Paging(method, path string, obj interface{}) ([]Trending, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Trending
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Trending
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

// Get performs GET request for Trending collection
func (r *OfficeGraphInsightsTrendingCollectionRequest) Get() ([]Trending, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Trending collection
func (r *OfficeGraphInsightsTrendingCollectionRequest) Add(reqObj *Trending) (resObj *Trending, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Used returns request builder for UsedInsight collection
func (b *OfficeGraphInsightsRequestBuilder) Used() *OfficeGraphInsightsUsedCollectionRequestBuilder {
	bb := &OfficeGraphInsightsUsedCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/used"
	return bb
}

// OfficeGraphInsightsUsedCollectionRequestBuilder is request builder for UsedInsight collection
type OfficeGraphInsightsUsedCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UsedInsight collection
func (b *OfficeGraphInsightsUsedCollectionRequestBuilder) Request() *OfficeGraphInsightsUsedCollectionRequest {
	return &OfficeGraphInsightsUsedCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UsedInsight item
func (b *OfficeGraphInsightsUsedCollectionRequestBuilder) ID(id string) *UsedInsightRequestBuilder {
	bb := &UsedInsightRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OfficeGraphInsightsUsedCollectionRequest is request for UsedInsight collection
type OfficeGraphInsightsUsedCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UsedInsight collection
func (r *OfficeGraphInsightsUsedCollectionRequest) Paging(method, path string, obj interface{}) ([]UsedInsight, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []UsedInsight
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []UsedInsight
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

// Get performs GET request for UsedInsight collection
func (r *OfficeGraphInsightsUsedCollectionRequest) Get() ([]UsedInsight, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for UsedInsight collection
func (r *OfficeGraphInsightsUsedCollectionRequest) Add(reqObj *UsedInsight) (resObj *UsedInsight, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
