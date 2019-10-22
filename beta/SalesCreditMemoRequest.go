// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// SalesCreditMemoRequestBuilder is request builder for SalesCreditMemo
type SalesCreditMemoRequestBuilder struct{ BaseRequestBuilder }

// Request returns SalesCreditMemoRequest
func (b *SalesCreditMemoRequestBuilder) Request() *SalesCreditMemoRequest {
	return &SalesCreditMemoRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SalesCreditMemoRequest is request for SalesCreditMemo
type SalesCreditMemoRequest struct{ BaseRequest }

// Get performs GET request for SalesCreditMemo
func (r *SalesCreditMemoRequest) Get() (resObj *SalesCreditMemo, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SalesCreditMemo
func (r *SalesCreditMemoRequest) Update(reqObj *SalesCreditMemo) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SalesCreditMemo
func (r *SalesCreditMemoRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Currency is navigation property
func (b *SalesCreditMemoRequestBuilder) Currency() *CurrencyRequestBuilder {
	bb := &CurrencyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/currency"
	return bb
}

// Customer is navigation property
func (b *SalesCreditMemoRequestBuilder) Customer() *CustomerRequestBuilder {
	bb := &CustomerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/customer"
	return bb
}

// PaymentTerm is navigation property
func (b *SalesCreditMemoRequestBuilder) PaymentTerm() *PaymentTermRequestBuilder {
	bb := &PaymentTermRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/paymentTerm"
	return bb
}

// SalesCreditMemoLines returns request builder for SalesCreditMemoLine collection
func (b *SalesCreditMemoRequestBuilder) SalesCreditMemoLines() *SalesCreditMemoSalesCreditMemoLinesCollectionRequestBuilder {
	bb := &SalesCreditMemoSalesCreditMemoLinesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/salesCreditMemoLines"
	return bb
}

// SalesCreditMemoSalesCreditMemoLinesCollectionRequestBuilder is request builder for SalesCreditMemoLine collection
type SalesCreditMemoSalesCreditMemoLinesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SalesCreditMemoLine collection
func (b *SalesCreditMemoSalesCreditMemoLinesCollectionRequestBuilder) Request() *SalesCreditMemoSalesCreditMemoLinesCollectionRequest {
	return &SalesCreditMemoSalesCreditMemoLinesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SalesCreditMemoLine item
func (b *SalesCreditMemoSalesCreditMemoLinesCollectionRequestBuilder) ID(id string) *SalesCreditMemoLineRequestBuilder {
	bb := &SalesCreditMemoLineRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SalesCreditMemoSalesCreditMemoLinesCollectionRequest is request for SalesCreditMemoLine collection
type SalesCreditMemoSalesCreditMemoLinesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SalesCreditMemoLine collection
func (r *SalesCreditMemoSalesCreditMemoLinesCollectionRequest) Paging(method, path string, obj interface{}) ([]SalesCreditMemoLine, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SalesCreditMemoLine
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []SalesCreditMemoLine
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

// Get performs GET request for SalesCreditMemoLine collection
func (r *SalesCreditMemoSalesCreditMemoLinesCollectionRequest) Get() ([]SalesCreditMemoLine, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SalesCreditMemoLine collection
func (r *SalesCreditMemoSalesCreditMemoLinesCollectionRequest) Add(reqObj *SalesCreditMemoLine) (resObj *SalesCreditMemoLine, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
