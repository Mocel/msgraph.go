// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// GroupPolicyDefinitionFileRequestBuilder is request builder for GroupPolicyDefinitionFile
type GroupPolicyDefinitionFileRequestBuilder struct{ BaseRequestBuilder }

// Request returns GroupPolicyDefinitionFileRequest
func (b *GroupPolicyDefinitionFileRequestBuilder) Request() *GroupPolicyDefinitionFileRequest {
	return &GroupPolicyDefinitionFileRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// GroupPolicyDefinitionFileRequest is request for GroupPolicyDefinitionFile
type GroupPolicyDefinitionFileRequest struct{ BaseRequest }

// Get performs GET request for GroupPolicyDefinitionFile
func (r *GroupPolicyDefinitionFileRequest) Get() (resObj *GroupPolicyDefinitionFile, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for GroupPolicyDefinitionFile
func (r *GroupPolicyDefinitionFileRequest) Update(reqObj *GroupPolicyDefinitionFile) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for GroupPolicyDefinitionFile
func (r *GroupPolicyDefinitionFileRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Definitions returns request builder for GroupPolicyDefinition collection
func (b *GroupPolicyDefinitionFileRequestBuilder) Definitions() *GroupPolicyDefinitionFileDefinitionsCollectionRequestBuilder {
	bb := &GroupPolicyDefinitionFileDefinitionsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/definitions"
	return bb
}

// GroupPolicyDefinitionFileDefinitionsCollectionRequestBuilder is request builder for GroupPolicyDefinition collection
type GroupPolicyDefinitionFileDefinitionsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for GroupPolicyDefinition collection
func (b *GroupPolicyDefinitionFileDefinitionsCollectionRequestBuilder) Request() *GroupPolicyDefinitionFileDefinitionsCollectionRequest {
	return &GroupPolicyDefinitionFileDefinitionsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for GroupPolicyDefinition item
func (b *GroupPolicyDefinitionFileDefinitionsCollectionRequestBuilder) ID(id string) *GroupPolicyDefinitionRequestBuilder {
	bb := &GroupPolicyDefinitionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// GroupPolicyDefinitionFileDefinitionsCollectionRequest is request for GroupPolicyDefinition collection
type GroupPolicyDefinitionFileDefinitionsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for GroupPolicyDefinition collection
func (r *GroupPolicyDefinitionFileDefinitionsCollectionRequest) Paging(method, path string, obj interface{}) ([]GroupPolicyDefinition, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []GroupPolicyDefinition
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []GroupPolicyDefinition
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

// Get performs GET request for GroupPolicyDefinition collection
func (r *GroupPolicyDefinitionFileDefinitionsCollectionRequest) Get() ([]GroupPolicyDefinition, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for GroupPolicyDefinition collection
func (r *GroupPolicyDefinitionFileDefinitionsCollectionRequest) Add(reqObj *GroupPolicyDefinition) (resObj *GroupPolicyDefinition, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
