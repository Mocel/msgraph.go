// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// SalesOrderRequestBuilder is request builder for SalesOrder
type SalesOrderRequestBuilder struct{ BaseRequestBuilder }

// Request returns SalesOrderRequest
func (b *SalesOrderRequestBuilder) Request() *SalesOrderRequest {
	return &SalesOrderRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// SalesOrderRequest is request for SalesOrder
type SalesOrderRequest struct{ BaseRequest }

// Get performs GET request for SalesOrder
func (r *SalesOrderRequest) Get() (resObj *SalesOrder, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for SalesOrder
func (r *SalesOrderRequest) Update(reqObj *SalesOrder) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for SalesOrder
func (r *SalesOrderRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// Currency is navigation property
func (b *SalesOrderRequestBuilder) Currency() *CurrencyRequestBuilder {
	bb := &CurrencyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/currency"
	return bb
}

// Customer is navigation property
func (b *SalesOrderRequestBuilder) Customer() *CustomerRequestBuilder {
	bb := &CustomerRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/customer"
	return bb
}

// PaymentTerm is navigation property
func (b *SalesOrderRequestBuilder) PaymentTerm() *PaymentTermRequestBuilder {
	bb := &PaymentTermRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/paymentTerm"
	return bb
}

// SalesOrderLines returns request builder for SalesOrderLine collection
func (b *SalesOrderRequestBuilder) SalesOrderLines() *SalesOrderSalesOrderLinesCollectionRequestBuilder {
	bb := &SalesOrderSalesOrderLinesCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/salesOrderLines"
	return bb
}

// SalesOrderSalesOrderLinesCollectionRequestBuilder is request builder for SalesOrderLine collection
type SalesOrderSalesOrderLinesCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for SalesOrderLine collection
func (b *SalesOrderSalesOrderLinesCollectionRequestBuilder) Request() *SalesOrderSalesOrderLinesCollectionRequest {
	return &SalesOrderSalesOrderLinesCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for SalesOrderLine item
func (b *SalesOrderSalesOrderLinesCollectionRequestBuilder) ID(id string) *SalesOrderLineRequestBuilder {
	bb := &SalesOrderLineRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// SalesOrderSalesOrderLinesCollectionRequest is request for SalesOrderLine collection
type SalesOrderSalesOrderLinesCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for SalesOrderLine collection
func (r *SalesOrderSalesOrderLinesCollectionRequest) Paging(method, path string, obj interface{}) ([]SalesOrderLine, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []SalesOrderLine
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []SalesOrderLine
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

// Get performs GET request for SalesOrderLine collection
func (r *SalesOrderSalesOrderLinesCollectionRequest) Get() ([]SalesOrderLine, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for SalesOrderLine collection
func (r *SalesOrderSalesOrderLinesCollectionRequest) Add(reqObj *SalesOrderLine) (resObj *SalesOrderLine, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
