// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// EducationClassRequestBuilder is request builder for EducationClass
type EducationClassRequestBuilder struct{ BaseRequestBuilder }

// Request returns EducationClassRequest
func (b *EducationClassRequestBuilder) Request() *EducationClassRequest {
	return &EducationClassRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// EducationClassRequest is request for EducationClass
type EducationClassRequest struct{ BaseRequest }

// Get performs GET request for EducationClass
func (r *EducationClassRequest) Get() (resObj *EducationClass, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for EducationClass
func (r *EducationClassRequest) Update(reqObj *EducationClass) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for EducationClass
func (r *EducationClassRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// AssignmentCategories returns request builder for EducationCategory collection
func (b *EducationClassRequestBuilder) AssignmentCategories() *EducationClassAssignmentCategoriesCollectionRequestBuilder {
	bb := &EducationClassAssignmentCategoriesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignmentCategories"
	return bb
}

// EducationClassAssignmentCategoriesCollectionRequestBuilder is request builder for EducationCategory collection
type EducationClassAssignmentCategoriesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationCategory collection
func (b *EducationClassAssignmentCategoriesCollectionRequestBuilder) Request() *EducationClassAssignmentCategoriesCollectionRequest {
	return &EducationClassAssignmentCategoriesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationCategory item
func (b *EducationClassAssignmentCategoriesCollectionRequestBuilder) ID(id string) *EducationCategoryRequestBuilder {
	bb := &EducationCategoryRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationClassAssignmentCategoriesCollectionRequest is request for EducationCategory collection
type EducationClassAssignmentCategoriesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationCategory collection
func (r *EducationClassAssignmentCategoriesCollectionRequest) Paging(method, path string, obj interface{}) ([]EducationCategory, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EducationCategory
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []EducationCategory
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

// Get performs GET request for EducationCategory collection
func (r *EducationClassAssignmentCategoriesCollectionRequest) Get() ([]EducationCategory, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EducationCategory collection
func (r *EducationClassAssignmentCategoriesCollectionRequest) Add(reqObj *EducationCategory) (resObj *EducationCategory, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Assignments returns request builder for EducationAssignment collection
func (b *EducationClassRequestBuilder) Assignments() *EducationClassAssignmentsCollectionRequestBuilder {
	bb := &EducationClassAssignmentsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/assignments"
	return bb
}

// EducationClassAssignmentsCollectionRequestBuilder is request builder for EducationAssignment collection
type EducationClassAssignmentsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationAssignment collection
func (b *EducationClassAssignmentsCollectionRequestBuilder) Request() *EducationClassAssignmentsCollectionRequest {
	return &EducationClassAssignmentsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationAssignment item
func (b *EducationClassAssignmentsCollectionRequestBuilder) ID(id string) *EducationAssignmentRequestBuilder {
	bb := &EducationAssignmentRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationClassAssignmentsCollectionRequest is request for EducationAssignment collection
type EducationClassAssignmentsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationAssignment collection
func (r *EducationClassAssignmentsCollectionRequest) Paging(method, path string, obj interface{}) ([]EducationAssignment, error) {
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
func (r *EducationClassAssignmentsCollectionRequest) Get() ([]EducationAssignment, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EducationAssignment collection
func (r *EducationClassAssignmentsCollectionRequest) Add(reqObj *EducationAssignment) (resObj *EducationAssignment, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Group is navigation property
func (b *EducationClassRequestBuilder) Group() *GroupRequestBuilder {
	bb := &GroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/group"
	return bb
}

// Members returns request builder for EducationUser collection
func (b *EducationClassRequestBuilder) Members() *EducationClassMembersCollectionRequestBuilder {
	bb := &EducationClassMembersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/members"
	return bb
}

// EducationClassMembersCollectionRequestBuilder is request builder for EducationUser collection
type EducationClassMembersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationUser collection
func (b *EducationClassMembersCollectionRequestBuilder) Request() *EducationClassMembersCollectionRequest {
	return &EducationClassMembersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationUser item
func (b *EducationClassMembersCollectionRequestBuilder) ID(id string) *EducationUserRequestBuilder {
	bb := &EducationUserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationClassMembersCollectionRequest is request for EducationUser collection
type EducationClassMembersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationUser collection
func (r *EducationClassMembersCollectionRequest) Paging(method, path string, obj interface{}) ([]EducationUser, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EducationUser
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []EducationUser
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

// Get performs GET request for EducationUser collection
func (r *EducationClassMembersCollectionRequest) Get() ([]EducationUser, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EducationUser collection
func (r *EducationClassMembersCollectionRequest) Add(reqObj *EducationUser) (resObj *EducationUser, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Schools returns request builder for EducationSchool collection
func (b *EducationClassRequestBuilder) Schools() *EducationClassSchoolsCollectionRequestBuilder {
	bb := &EducationClassSchoolsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/schools"
	return bb
}

// EducationClassSchoolsCollectionRequestBuilder is request builder for EducationSchool collection
type EducationClassSchoolsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationSchool collection
func (b *EducationClassSchoolsCollectionRequestBuilder) Request() *EducationClassSchoolsCollectionRequest {
	return &EducationClassSchoolsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationSchool item
func (b *EducationClassSchoolsCollectionRequestBuilder) ID(id string) *EducationSchoolRequestBuilder {
	bb := &EducationSchoolRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationClassSchoolsCollectionRequest is request for EducationSchool collection
type EducationClassSchoolsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationSchool collection
func (r *EducationClassSchoolsCollectionRequest) Paging(method, path string, obj interface{}) ([]EducationSchool, error) {
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
func (r *EducationClassSchoolsCollectionRequest) Get() ([]EducationSchool, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EducationSchool collection
func (r *EducationClassSchoolsCollectionRequest) Add(reqObj *EducationSchool) (resObj *EducationSchool, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Teachers returns request builder for EducationUser collection
func (b *EducationClassRequestBuilder) Teachers() *EducationClassTeachersCollectionRequestBuilder {
	bb := &EducationClassTeachersCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/teachers"
	return bb
}

// EducationClassTeachersCollectionRequestBuilder is request builder for EducationUser collection
type EducationClassTeachersCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for EducationUser collection
func (b *EducationClassTeachersCollectionRequestBuilder) Request() *EducationClassTeachersCollectionRequest {
	return &EducationClassTeachersCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for EducationUser item
func (b *EducationClassTeachersCollectionRequestBuilder) ID(id string) *EducationUserRequestBuilder {
	bb := &EducationUserRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// EducationClassTeachersCollectionRequest is request for EducationUser collection
type EducationClassTeachersCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for EducationUser collection
func (r *EducationClassTeachersCollectionRequest) Paging(method, path string, obj interface{}) ([]EducationUser, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []EducationUser
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []EducationUser
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

// Get performs GET request for EducationUser collection
func (r *EducationClassTeachersCollectionRequest) Get() ([]EducationUser, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for EducationUser collection
func (r *EducationClassTeachersCollectionRequest) Add(reqObj *EducationUser) (resObj *EducationUser, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
