// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// DomainRequestBuilder is request builder for Domain
type DomainRequestBuilder struct{ BaseRequestBuilder }

// Request returns DomainRequest
func (b *DomainRequestBuilder) Request() *DomainRequest {
	return &DomainRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// DomainRequest is request for Domain
type DomainRequest struct{ BaseRequest }

// Get performs GET request for Domain
func (r *DomainRequest) Get() (resObj *Domain, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Domain
func (r *DomainRequest) Update(reqObj *Domain) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Domain
func (r *DomainRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// DomainNameReferences returns request builder for DirectoryObject collection
func (b *DomainRequestBuilder) DomainNameReferences() *DomainDomainNameReferencesCollectionRequestBuilder {
	bb := &DomainDomainNameReferencesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/domainNameReferences"
	return bb
}

// DomainDomainNameReferencesCollectionRequestBuilder is request builder for DirectoryObject collection
type DomainDomainNameReferencesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DirectoryObject collection
func (b *DomainDomainNameReferencesCollectionRequestBuilder) Request() *DomainDomainNameReferencesCollectionRequest {
	return &DomainDomainNameReferencesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DirectoryObject item
func (b *DomainDomainNameReferencesCollectionRequestBuilder) ID(id string) *DirectoryObjectRequestBuilder {
	bb := &DirectoryObjectRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DomainDomainNameReferencesCollectionRequest is request for DirectoryObject collection
type DomainDomainNameReferencesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DirectoryObject collection
func (r *DomainDomainNameReferencesCollectionRequest) Paging(method, path string, obj interface{}) ([]DirectoryObject, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DirectoryObject
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DirectoryObject
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

// Get performs GET request for DirectoryObject collection
func (r *DomainDomainNameReferencesCollectionRequest) Get() ([]DirectoryObject, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DirectoryObject collection
func (r *DomainDomainNameReferencesCollectionRequest) Add(reqObj *DirectoryObject) (resObj *DirectoryObject, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// ServiceConfigurationRecords returns request builder for DomainDNSRecord collection
func (b *DomainRequestBuilder) ServiceConfigurationRecords() *DomainServiceConfigurationRecordsCollectionRequestBuilder {
	bb := &DomainServiceConfigurationRecordsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/serviceConfigurationRecords"
	return bb
}

// DomainServiceConfigurationRecordsCollectionRequestBuilder is request builder for DomainDNSRecord collection
type DomainServiceConfigurationRecordsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DomainDNSRecord collection
func (b *DomainServiceConfigurationRecordsCollectionRequestBuilder) Request() *DomainServiceConfigurationRecordsCollectionRequest {
	return &DomainServiceConfigurationRecordsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DomainDNSRecord item
func (b *DomainServiceConfigurationRecordsCollectionRequestBuilder) ID(id string) *DomainDNSRecordRequestBuilder {
	bb := &DomainDNSRecordRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DomainServiceConfigurationRecordsCollectionRequest is request for DomainDNSRecord collection
type DomainServiceConfigurationRecordsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DomainDNSRecord collection
func (r *DomainServiceConfigurationRecordsCollectionRequest) Paging(method, path string, obj interface{}) ([]DomainDNSRecord, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DomainDNSRecord
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DomainDNSRecord
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

// Get performs GET request for DomainDNSRecord collection
func (r *DomainServiceConfigurationRecordsCollectionRequest) Get() ([]DomainDNSRecord, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DomainDNSRecord collection
func (r *DomainServiceConfigurationRecordsCollectionRequest) Add(reqObj *DomainDNSRecord) (resObj *DomainDNSRecord, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// VerificationDNSRecords returns request builder for DomainDNSRecord collection
func (b *DomainRequestBuilder) VerificationDNSRecords() *DomainVerificationDNSRecordsCollectionRequestBuilder {
	bb := &DomainVerificationDNSRecordsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/verificationDnsRecords"
	return bb
}

// DomainVerificationDNSRecordsCollectionRequestBuilder is request builder for DomainDNSRecord collection
type DomainVerificationDNSRecordsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for DomainDNSRecord collection
func (b *DomainVerificationDNSRecordsCollectionRequestBuilder) Request() *DomainVerificationDNSRecordsCollectionRequest {
	return &DomainVerificationDNSRecordsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for DomainDNSRecord item
func (b *DomainVerificationDNSRecordsCollectionRequestBuilder) ID(id string) *DomainDNSRecordRequestBuilder {
	bb := &DomainDNSRecordRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// DomainVerificationDNSRecordsCollectionRequest is request for DomainDNSRecord collection
type DomainVerificationDNSRecordsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for DomainDNSRecord collection
func (r *DomainVerificationDNSRecordsCollectionRequest) Paging(method, path string, obj interface{}) ([]DomainDNSRecord, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []DomainDNSRecord
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []DomainDNSRecord
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

// Get performs GET request for DomainDNSRecord collection
func (r *DomainVerificationDNSRecordsCollectionRequest) Get() ([]DomainDNSRecord, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for DomainDNSRecord collection
func (r *DomainVerificationDNSRecordsCollectionRequest) Add(reqObj *DomainDNSRecord) (resObj *DomainDNSRecord, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
