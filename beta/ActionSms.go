// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// IncludeTargets returns request builder for SmsAuthenticationMethodTarget collection
func (b *SmsAuthenticationMethodConfigurationRequestBuilder) IncludeTargets() *SmsAuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder {
	bb := &SmsAuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/includeTargets"
	return bb
}

// SmsAuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder is request builder for SmsAuthenticationMethodTarget collection
type SmsAuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SmsAuthenticationMethodTarget collection
func (b *SmsAuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder) Request() *SmsAuthenticationMethodConfigurationIncludeTargetsCollectionRequest {
	return &SmsAuthenticationMethodConfigurationIncludeTargetsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SmsAuthenticationMethodTarget item
func (b *SmsAuthenticationMethodConfigurationIncludeTargetsCollectionRequestBuilder) ID(id string) *SmsAuthenticationMethodTargetRequestBuilder {
	bb := &SmsAuthenticationMethodTargetRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SmsAuthenticationMethodConfigurationIncludeTargetsCollectionRequest is request for SmsAuthenticationMethodTarget collection
type SmsAuthenticationMethodConfigurationIncludeTargetsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SmsAuthenticationMethodTarget collection
func (r *SmsAuthenticationMethodConfigurationIncludeTargetsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]SmsAuthenticationMethodTarget, error) {
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
	var values []SmsAuthenticationMethodTarget
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
			value  []SmsAuthenticationMethodTarget
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

// GetN performs GET request for SmsAuthenticationMethodTarget collection, max N pages
func (r *SmsAuthenticationMethodConfigurationIncludeTargetsCollectionRequest) GetN(ctx context.Context, n int) ([]SmsAuthenticationMethodTarget, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for SmsAuthenticationMethodTarget collection
func (r *SmsAuthenticationMethodConfigurationIncludeTargetsCollectionRequest) Get(ctx context.Context) ([]SmsAuthenticationMethodTarget, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for SmsAuthenticationMethodTarget collection
func (r *SmsAuthenticationMethodConfigurationIncludeTargetsCollectionRequest) Add(ctx context.Context, reqObj *SmsAuthenticationMethodTarget) (resObj *SmsAuthenticationMethodTarget, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
