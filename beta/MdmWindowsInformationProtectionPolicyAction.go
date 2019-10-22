// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// MdmWindowsInformationProtectionPolicyCollectionHasPayloadLinksRequestParameter undocumented
type MdmWindowsInformationProtectionPolicyCollectionHasPayloadLinksRequestParameter struct {
	// PayloadIDs undocumented
	PayloadIDs []string `json:"payloadIds,omitempty"`
}

//
type MdmWindowsInformationProtectionPolicyCollectionHasPayloadLinksRequestBuilder struct{ BaseRequestBuilder }

// HasPayloadLinks action undocumented
func (b *DeviceAppManagementMdmWindowsInformationProtectionPoliciesCollectionRequestBuilder) HasPayloadLinks(reqObj *MdmWindowsInformationProtectionPolicyCollectionHasPayloadLinksRequestParameter) *MdmWindowsInformationProtectionPolicyCollectionHasPayloadLinksRequestBuilder {
	bb := &MdmWindowsInformationProtectionPolicyCollectionHasPayloadLinksRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/hasPayloadLinks"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type MdmWindowsInformationProtectionPolicyCollectionHasPayloadLinksRequest struct{ BaseRequest }

//
func (b *MdmWindowsInformationProtectionPolicyCollectionHasPayloadLinksRequestBuilder) Request() *MdmWindowsInformationProtectionPolicyCollectionHasPayloadLinksRequest {
	return &MdmWindowsInformationProtectionPolicyCollectionHasPayloadLinksRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *MdmWindowsInformationProtectionPolicyCollectionHasPayloadLinksRequest) Paging(method, path string, obj interface{}) ([][]HasPayloadLinkResultItem, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]HasPayloadLinkResultItem
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]HasPayloadLinkResultItem
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

//
func (r *MdmWindowsInformationProtectionPolicyCollectionHasPayloadLinksRequest) Get() ([][]HasPayloadLinkResultItem, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}
