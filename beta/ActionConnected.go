// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// ExternalSponsors returns request builder for DirectoryObject collection
func (b *ConnectedOrganizationRequestBuilder) ExternalSponsors() *ConnectedOrganizationExternalSponsorsCollectionRequestBuilder {
	bb := &ConnectedOrganizationExternalSponsorsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/externalSponsors"
	return bb
}

// ConnectedOrganizationExternalSponsorsCollectionRequestBuilder is request builder for DirectoryObject collection
type ConnectedOrganizationExternalSponsorsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ConnectedOrganizationExternalSponsorsCollectionRequestBuilder) Request() *ConnectedOrganizationExternalSponsorsCollectionRequest {
	return &ConnectedOrganizationExternalSponsorsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ConnectedOrganizationExternalSponsorsCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ConnectedOrganizationExternalSponsorsCollectionRequest is request for DirectoryObject collection
type ConnectedOrganizationExternalSponsorsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ConnectedOrganizationExternalSponsorsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *ConnectedOrganizationExternalSponsorsCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *ConnectedOrganizationExternalSponsorsCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *ConnectedOrganizationExternalSponsorsCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}

// InternalSponsors returns request builder for DirectoryObject collection
func (b *ConnectedOrganizationRequestBuilder) InternalSponsors() *ConnectedOrganizationInternalSponsorsCollectionRequestBuilder {
	bb := &ConnectedOrganizationInternalSponsorsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/internalSponsors"
	return bb
}

// ConnectedOrganizationInternalSponsorsCollectionRequestBuilder is request builder for DirectoryObject collection
type ConnectedOrganizationInternalSponsorsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *ConnectedOrganizationInternalSponsorsCollectionRequestBuilder) Request() *ConnectedOrganizationInternalSponsorsCollectionRequest {
	return &ConnectedOrganizationInternalSponsorsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *ConnectedOrganizationInternalSponsorsCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// ConnectedOrganizationInternalSponsorsCollectionRequest is request for DirectoryObject collection
type ConnectedOrganizationInternalSponsorsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *ConnectedOrganizationInternalSponsorsCollectionRequest) Paging(ctx context.Context, method, path string, obj interface{}, n int) ([]DirectoryObject, error) {
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
	var values []DirectoryObject
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
			value  []DirectoryObject
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

// GetN performs GET request for DirectoryObject collection, max N pages
func (r *ConnectedOrganizationInternalSponsorsCollectionRequest) GetN(ctx context.Context, n int) ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging(ctx, "GET", query, nil, n)
}

// Get performs GET request for DirectoryObject collection
func (r *ConnectedOrganizationInternalSponsorsCollectionRequest) Get(ctx context.Context) ([]DirectoryObject, error) {
	return r.GetN(ctx, 0)
}

// Add performs POST request for DirectoryObject collection
func (r *ConnectedOrganizationInternalSponsorsCollectionRequest) Add(ctx context.Context, reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest(ctx, "POST", "", reqObj, &resObj)
	return
}
