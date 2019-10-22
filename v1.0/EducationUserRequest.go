// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// EducationUserRequestBuilder is request builder for EducationUser
type EducationUserRequestBuilder struct{ BaseRequestBuilder }

// Request returns EducationUserRequest
func (b *EducationUserRequestBuilder) Request() *EducationUserRequest {
	return &EducationUserRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EducationUserRequest is request for EducationUser
type EducationUserRequest struct{ BaseRequest }

// Get performs GET request for EducationUser
func (r *EducationUserRequest) Get() (resObj *EducationUser, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EducationUser
func (r *EducationUserRequest) Update(reqObj *EducationUser) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EducationUser
func (r *EducationUserRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Classes returns request builder for EducationClass collection
func (b *EducationUserRequestBuilder) Classes() *EducationUserClassesCollectionRequestBuilder {
	bb := &EducationUserClassesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/classes"
	return bb
}

// EducationUserClassesCollectionRequestBuilder is request builder for EducationClass collection
type EducationUserClassesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationClass collection
func (b *EducationUserClassesCollectionRequestBuilder) Request() *EducationUserClassesCollectionRequest {
	return &EducationUserClassesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationClass item
func (b *EducationUserClassesCollectionRequestBuilder) ID(id string) *EducationClassRequestBuilder {
	bb := &EducationClassRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationUserClassesCollectionRequest is request for EducationClass collection
type EducationUserClassesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationClass collection
func (r *EducationUserClassesCollectionRequest) Paging(method, path string, obj interface{}) ([]EducationClass, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EducationClass
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []EducationClass
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

// Get performs GET request for EducationClass collection
func (r *EducationUserClassesCollectionRequest) Get() ([]EducationClass, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EducationClass collection
func (r *EducationUserClassesCollectionRequest) Add(reqObj *EducationClass) (resObj *EducationClass, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Schools returns request builder for EducationSchool collection
func (b *EducationUserRequestBuilder) Schools() *EducationUserSchoolsCollectionRequestBuilder {
	bb := &EducationUserSchoolsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/schools"
	return bb
}

// EducationUserSchoolsCollectionRequestBuilder is request builder for EducationSchool collection
type EducationUserSchoolsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationSchool collection
func (b *EducationUserSchoolsCollectionRequestBuilder) Request() *EducationUserSchoolsCollectionRequest {
	return &EducationUserSchoolsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationSchool item
func (b *EducationUserSchoolsCollectionRequestBuilder) ID(id string) *EducationSchoolRequestBuilder {
	bb := &EducationSchoolRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationUserSchoolsCollectionRequest is request for EducationSchool collection
type EducationUserSchoolsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationSchool collection
func (r *EducationUserSchoolsCollectionRequest) Paging(method, path string, obj interface{}) ([]EducationSchool, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EducationSchool
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []EducationSchool
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

// Get performs GET request for EducationSchool collection
func (r *EducationUserSchoolsCollectionRequest) Get() ([]EducationSchool, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EducationSchool collection
func (r *EducationUserSchoolsCollectionRequest) Add(reqObj *EducationSchool) (resObj *EducationSchool, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// User is navigation property
func (b *EducationUserRequestBuilder) User() *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/user"
	return bb
}
