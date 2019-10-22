// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// OnenoteSectionRequestBuilder is request builder for OnenoteSection
type OnenoteSectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns OnenoteSectionRequest
func (b *OnenoteSectionRequestBuilder) Request() *OnenoteSectionRequest {
	return &OnenoteSectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// OnenoteSectionRequest is request for OnenoteSection
type OnenoteSectionRequest struct{ BaseRequest }

// Get performs GET request for OnenoteSection
func (r *OnenoteSectionRequest) Get() (resObj *OnenoteSection, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for OnenoteSection
func (r *OnenoteSectionRequest) Update(reqObj *OnenoteSection) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for OnenoteSection
func (r *OnenoteSectionRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Pages returns request builder for OnenotePage collection
func (b *OnenoteSectionRequestBuilder) Pages() *OnenoteSectionPagesCollectionRequestBuilder {
	bb := &OnenoteSectionPagesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/pages"
	return bb
}

// OnenoteSectionPagesCollectionRequestBuilder is request builder for OnenotePage collection
type OnenoteSectionPagesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for OnenotePage collection
func (b *OnenoteSectionPagesCollectionRequestBuilder) Request() *OnenoteSectionPagesCollectionRequest {
	return &OnenoteSectionPagesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for OnenotePage item
func (b *OnenoteSectionPagesCollectionRequestBuilder) ID(id string) *OnenotePageRequestBuilder {
	bb := &OnenotePageRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// OnenoteSectionPagesCollectionRequest is request for OnenotePage collection
type OnenoteSectionPagesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for OnenotePage collection
func (r *OnenoteSectionPagesCollectionRequest) Paging(method, path string, obj interface{}) ([]OnenotePage, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []OnenotePage
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []OnenotePage
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

// Get performs GET request for OnenotePage collection
func (r *OnenoteSectionPagesCollectionRequest) Get() ([]OnenotePage, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for OnenotePage collection
func (r *OnenoteSectionPagesCollectionRequest) Add(reqObj *OnenotePage) (resObj *OnenotePage, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// ParentNotebook is navigation property
func (b *OnenoteSectionRequestBuilder) ParentNotebook() *NotebookRequestBuilder {
	bb := &NotebookRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/parentNotebook"
	return bb
}

// ParentSectionGroup is navigation property
func (b *OnenoteSectionRequestBuilder) ParentSectionGroup() *SectionGroupRequestBuilder {
	bb := &SectionGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/parentSectionGroup"
	return bb
}
