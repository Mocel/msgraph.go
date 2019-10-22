// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// UserExperienceAnalyticsBaselineRequestBuilder is request builder for UserExperienceAnalyticsBaseline
type UserExperienceAnalyticsBaselineRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserExperienceAnalyticsBaselineRequest
func (b *UserExperienceAnalyticsBaselineRequestBuilder) Request() *UserExperienceAnalyticsBaselineRequest {
	return &UserExperienceAnalyticsBaselineRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserExperienceAnalyticsBaselineRequest is request for UserExperienceAnalyticsBaseline
type UserExperienceAnalyticsBaselineRequest struct{ BaseRequest }

// Get performs GET request for UserExperienceAnalyticsBaseline
func (r *UserExperienceAnalyticsBaselineRequest) Get() (resObj *UserExperienceAnalyticsBaseline, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserExperienceAnalyticsBaseline
func (r *UserExperienceAnalyticsBaselineRequest) Update(reqObj *UserExperienceAnalyticsBaseline) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserExperienceAnalyticsBaseline
func (r *UserExperienceAnalyticsBaselineRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// BestPracticesMetrics is navigation property
func (b *UserExperienceAnalyticsBaselineRequestBuilder) BestPracticesMetrics() *UserExperienceAnalyticsCategoryRequestBuilder {
	bb := &UserExperienceAnalyticsCategoryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/bestPracticesMetrics"
	return bb
}

// DeviceBootPerformanceMetrics is navigation property
func (b *UserExperienceAnalyticsBaselineRequestBuilder) DeviceBootPerformanceMetrics() *UserExperienceAnalyticsCategoryRequestBuilder {
	bb := &UserExperienceAnalyticsCategoryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/deviceBootPerformanceMetrics"
	return bb
}
