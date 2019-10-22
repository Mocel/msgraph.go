// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// SharedDriveItemRequestBuilder is request builder for SharedDriveItem
type SharedDriveItemRequestBuilder struct{ BaseRequestBuilder }

// Request returns SharedDriveItemRequest
func (b *SharedDriveItemRequestBuilder) Request() *SharedDriveItemRequest {
	return &SharedDriveItemRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SharedDriveItemRequest is request for SharedDriveItem
type SharedDriveItemRequest struct{ BaseRequest }

// Get performs GET request for SharedDriveItem
func (r *SharedDriveItemRequest) Get() (resObj *SharedDriveItem, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SharedDriveItem
func (r *SharedDriveItemRequest) Update(reqObj *SharedDriveItem) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SharedDriveItem
func (r *SharedDriveItemRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// DriveItem is navigation property
func (b *SharedDriveItemRequestBuilder) DriveItem() *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/driveItem"
	return bb
}

// Items returns request builder for DriveItem collection
func (b *SharedDriveItemRequestBuilder) Items() *SharedDriveItemItemsCollectionRequestBuilder {
	bb := &SharedDriveItemItemsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/items"
	return bb
}

// SharedDriveItemItemsCollectionRequestBuilder is request builder for DriveItem collection
type SharedDriveItemItemsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DriveItem collection
func (b *SharedDriveItemItemsCollectionRequestBuilder) Request() *SharedDriveItemItemsCollectionRequest {
	return &SharedDriveItemItemsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DriveItem item
func (b *SharedDriveItemItemsCollectionRequestBuilder) ID(id string) *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SharedDriveItemItemsCollectionRequest is request for DriveItem collection
type SharedDriveItemItemsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DriveItem collection
func (r *SharedDriveItemItemsCollectionRequest) Paging(method, path string, obj interface{}) ([]DriveItem, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DriveItem
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DriveItem
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

// Get performs GET request for DriveItem collection
func (r *SharedDriveItemItemsCollectionRequest) Get() ([]DriveItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DriveItem collection
func (r *SharedDriveItemItemsCollectionRequest) Add(reqObj *DriveItem) (resObj *DriveItem, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// List is navigation property
func (b *SharedDriveItemRequestBuilder) List() *ListRequestBuilder {
	bb := &ListRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/list"
	return bb
}

// ListItem is navigation property
func (b *SharedDriveItemRequestBuilder) ListItem() *ListItemRequestBuilder {
	bb := &ListItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/listItem"
	return bb
}

// Permission is navigation property
func (b *SharedDriveItemRequestBuilder) Permission() *PermissionRequestBuilder {
	bb := &PermissionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/permission"
	return bb
}

// Root is navigation property
func (b *SharedDriveItemRequestBuilder) Root() *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/root"
	return bb
}

// Site is navigation property
func (b *SharedDriveItemRequestBuilder) Site() *SiteRequestBuilder {
	bb := &SiteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/site"
	return bb
}
