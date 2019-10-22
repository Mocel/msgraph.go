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

// Assignments returns request builder for EducationAssignment collection
func (b *EducationUserRequestBuilder) Assignments() *EducationUserAssignmentsCollectionRequestBuilder {
	bb := &EducationUserAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// EducationUserAssignmentsCollectionRequestBuilder is request builder for EducationAssignment collection
type EducationUserAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationAssignment collection
func (b *EducationUserAssignmentsCollectionRequestBuilder) Request() *EducationUserAssignmentsCollectionRequest {
	return &EducationUserAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationAssignment item
func (b *EducationUserAssignmentsCollectionRequestBuilder) ID(id string) *EducationAssignmentRequestBuilder {
	bb := &EducationAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationUserAssignmentsCollectionRequest is request for EducationAssignment collection
type EducationUserAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationAssignment collection
func (r *EducationUserAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]EducationAssignment, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EducationAssignment
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []EducationAssignment
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

// Get performs GET request for EducationAssignment collection
func (r *EducationUserAssignmentsCollectionRequest) Get() ([]EducationAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EducationAssignment collection
func (r *EducationUserAssignmentsCollectionRequest) Add(reqObj *EducationAssignment) (resObj *EducationAssignment, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
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

// Rubrics returns request builder for EducationRubric collection
func (b *EducationUserRequestBuilder) Rubrics() *EducationUserRubricsCollectionRequestBuilder {
	bb := &EducationUserRubricsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/rubrics"
	return bb
}

// EducationUserRubricsCollectionRequestBuilder is request builder for EducationRubric collection
type EducationUserRubricsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationRubric collection
func (b *EducationUserRubricsCollectionRequestBuilder) Request() *EducationUserRubricsCollectionRequest {
	return &EducationUserRubricsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationRubric item
func (b *EducationUserRubricsCollectionRequestBuilder) ID(id string) *EducationRubricRequestBuilder {
	bb := &EducationRubricRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationUserRubricsCollectionRequest is request for EducationRubric collection
type EducationUserRubricsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationRubric collection
func (r *EducationUserRubricsCollectionRequest) Paging(method, path string, obj interface{}) ([]EducationRubric, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EducationRubric
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []EducationRubric
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

// Get performs GET request for EducationRubric collection
func (r *EducationUserRubricsCollectionRequest) Get() ([]EducationRubric, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EducationRubric collection
func (r *EducationUserRubricsCollectionRequest) Add(reqObj *EducationRubric) (resObj *EducationRubric, err error) {
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

// TaughtClasses returns request builder for EducationClass collection
func (b *EducationUserRequestBuilder) TaughtClasses() *EducationUserTaughtClassesCollectionRequestBuilder {
	bb := &EducationUserTaughtClassesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/taughtClasses"
	return bb
}

// EducationUserTaughtClassesCollectionRequestBuilder is request builder for EducationClass collection
type EducationUserTaughtClassesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationClass collection
func (b *EducationUserTaughtClassesCollectionRequestBuilder) Request() *EducationUserTaughtClassesCollectionRequest {
	return &EducationUserTaughtClassesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationClass item
func (b *EducationUserTaughtClassesCollectionRequestBuilder) ID(id string) *EducationClassRequestBuilder {
	bb := &EducationClassRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationUserTaughtClassesCollectionRequest is request for EducationClass collection
type EducationUserTaughtClassesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationClass collection
func (r *EducationUserTaughtClassesCollectionRequest) Paging(method, path string, obj interface{}) ([]EducationClass, error) {
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
func (r *EducationUserTaughtClassesCollectionRequest) Get() ([]EducationClass, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EducationClass collection
func (r *EducationUserTaughtClassesCollectionRequest) Add(reqObj *EducationClass) (resObj *EducationClass, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// User is navigation property
func (b *EducationUserRequestBuilder) User() *UserRequestBuilder {
	bb := &UserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/user"
	return bb
}
