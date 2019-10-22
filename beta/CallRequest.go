// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// CallRequestBuilder is request builder for Call
type CallRequestBuilder struct{ BaseRequestBuilder }

// Request returns CallRequest
func (b *CallRequestBuilder) Request() *CallRequest {
	return &CallRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// CallRequest is request for Call
type CallRequest struct{ BaseRequest }

// Get performs GET request for Call
func (r *CallRequest) Get() (resObj *Call, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Call
func (r *CallRequest) Update(reqObj *Call) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Call
func (r *CallRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// AudioRoutingGroups returns request builder for AudioRoutingGroup collection
func (b *CallRequestBuilder) AudioRoutingGroups() *CallAudioRoutingGroupsCollectionRequestBuilder {
	bb := &CallAudioRoutingGroupsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/audioRoutingGroups"
	return bb
}

// CallAudioRoutingGroupsCollectionRequestBuilder is request builder for AudioRoutingGroup collection
type CallAudioRoutingGroupsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for AudioRoutingGroup collection
func (b *CallAudioRoutingGroupsCollectionRequestBuilder) Request() *CallAudioRoutingGroupsCollectionRequest {
	return &CallAudioRoutingGroupsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for AudioRoutingGroup item
func (b *CallAudioRoutingGroupsCollectionRequestBuilder) ID(id string) *AudioRoutingGroupRequestBuilder {
	bb := &AudioRoutingGroupRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CallAudioRoutingGroupsCollectionRequest is request for AudioRoutingGroup collection
type CallAudioRoutingGroupsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for AudioRoutingGroup collection
func (r *CallAudioRoutingGroupsCollectionRequest) Paging(method, path string, obj interface{}) ([]AudioRoutingGroup, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []AudioRoutingGroup
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []AudioRoutingGroup
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

// Get performs GET request for AudioRoutingGroup collection
func (r *CallAudioRoutingGroupsCollectionRequest) Get() ([]AudioRoutingGroup, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for AudioRoutingGroup collection
func (r *CallAudioRoutingGroupsCollectionRequest) Add(reqObj *AudioRoutingGroup) (resObj *AudioRoutingGroup, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Operations returns request builder for CommsOperation collection
func (b *CallRequestBuilder) Operations() *CallOperationsCollectionRequestBuilder {
	bb := &CallOperationsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/operations"
	return bb
}

// CallOperationsCollectionRequestBuilder is request builder for CommsOperation collection
type CallOperationsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for CommsOperation collection
func (b *CallOperationsCollectionRequestBuilder) Request() *CallOperationsCollectionRequest {
	return &CallOperationsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for CommsOperation item
func (b *CallOperationsCollectionRequestBuilder) ID(id string) *CommsOperationRequestBuilder {
	bb := &CommsOperationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CallOperationsCollectionRequest is request for CommsOperation collection
type CallOperationsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for CommsOperation collection
func (r *CallOperationsCollectionRequest) Paging(method, path string, obj interface{}) ([]CommsOperation, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []CommsOperation
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []CommsOperation
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

// Get performs GET request for CommsOperation collection
func (r *CallOperationsCollectionRequest) Get() ([]CommsOperation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for CommsOperation collection
func (r *CallOperationsCollectionRequest) Add(reqObj *CommsOperation) (resObj *CommsOperation, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}

// Participants returns request builder for Participant collection
func (b *CallRequestBuilder) Participants() *CallParticipantsCollectionRequestBuilder {
	bb := &CallParticipantsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/participants"
	return bb
}

// CallParticipantsCollectionRequestBuilder is request builder for Participant collection
type CallParticipantsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for Participant collection
func (b *CallParticipantsCollectionRequestBuilder) Request() *CallParticipantsCollectionRequest {
	return &CallParticipantsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for Participant item
func (b *CallParticipantsCollectionRequestBuilder) ID(id string) *ParticipantRequestBuilder {
	bb := &ParticipantRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// CallParticipantsCollectionRequest is request for Participant collection
type CallParticipantsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for Participant collection
func (r *CallParticipantsCollectionRequest) Paging(method, path string, obj interface{}) ([]Participant, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []Participant
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []Participant
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

// Get performs GET request for Participant collection
func (r *CallParticipantsCollectionRequest) Get() ([]Participant, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for Participant collection
func (r *CallParticipantsCollectionRequest) Add(reqObj *Participant) (resObj *Participant, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
