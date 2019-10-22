// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// UserExperienceAnalyticsCategoryRequestBuilder is request builder for UserExperienceAnalyticsCategory
type UserExperienceAnalyticsCategoryRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserExperienceAnalyticsCategoryRequest
func (b *UserExperienceAnalyticsCategoryRequestBuilder) Request() *UserExperienceAnalyticsCategoryRequest {
	return &UserExperienceAnalyticsCategoryRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserExperienceAnalyticsCategoryRequest is request for UserExperienceAnalyticsCategory
type UserExperienceAnalyticsCategoryRequest struct{ BaseRequest }

// Get performs GET request for UserExperienceAnalyticsCategory
func (r *UserExperienceAnalyticsCategoryRequest) Get() (resObj *UserExperienceAnalyticsCategory, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserExperienceAnalyticsCategory
func (r *UserExperienceAnalyticsCategoryRequest) Update(reqObj *UserExperienceAnalyticsCategory) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserExperienceAnalyticsCategory
func (r *UserExperienceAnalyticsCategoryRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// MetricValues returns request builder for UserExperienceAnalyticsMetric collection
func (b *UserExperienceAnalyticsCategoryRequestBuilder) MetricValues() *UserExperienceAnalyticsCategoryMetricValuesCollectionRequestBuilder {
	bb := &UserExperienceAnalyticsCategoryMetricValuesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/metricValues"
	return bb
}

// UserExperienceAnalyticsCategoryMetricValuesCollectionRequestBuilder is request builder for UserExperienceAnalyticsMetric collection
type UserExperienceAnalyticsCategoryMetricValuesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for UserExperienceAnalyticsMetric collection
func (b *UserExperienceAnalyticsCategoryMetricValuesCollectionRequestBuilder) Request() *UserExperienceAnalyticsCategoryMetricValuesCollectionRequest {
	return &UserExperienceAnalyticsCategoryMetricValuesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for UserExperienceAnalyticsMetric item
func (b *UserExperienceAnalyticsCategoryMetricValuesCollectionRequestBuilder) ID(id string) *UserExperienceAnalyticsMetricRequestBuilder {
	bb := &UserExperienceAnalyticsMetricRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UserExperienceAnalyticsCategoryMetricValuesCollectionRequest is request for UserExperienceAnalyticsMetric collection
type UserExperienceAnalyticsCategoryMetricValuesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for UserExperienceAnalyticsMetric collection
func (r *UserExperienceAnalyticsCategoryMetricValuesCollectionRequest) Paging(method, path string, obj interface{}) ([]UserExperienceAnalyticsMetric, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []UserExperienceAnalyticsMetric
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []UserExperienceAnalyticsMetric
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

// Get performs GET request for UserExperienceAnalyticsMetric collection
func (r *UserExperienceAnalyticsCategoryMetricValuesCollectionRequest) Get() ([]UserExperienceAnalyticsMetric, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for UserExperienceAnalyticsMetric collection
func (r *UserExperienceAnalyticsCategoryMetricValuesCollectionRequest) Add(reqObj *UserExperienceAnalyticsMetric) (resObj *UserExperienceAnalyticsMetric, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
