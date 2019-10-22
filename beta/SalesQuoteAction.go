// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// SalesQuoteMakeInvoiceRequestParameter undocumented
type SalesQuoteMakeInvoiceRequestParameter struct {
}

// SalesQuoteSendRequestParameter undocumented
type SalesQuoteSendRequestParameter struct {
}

//
type SalesQuoteMakeInvoiceRequestBuilder struct{ BaseRequestBuilder }

// MakeInvoice action undocumented
func (b *SalesQuoteRequestBuilder) MakeInvoice(reqObj *SalesQuoteMakeInvoiceRequestParameter) *SalesQuoteMakeInvoiceRequestBuilder {
	bb := &SalesQuoteMakeInvoiceRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/makeInvoice"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SalesQuoteMakeInvoiceRequest struct{ BaseRequest }

//
func (b *SalesQuoteMakeInvoiceRequestBuilder) Request() *SalesQuoteMakeInvoiceRequest {
	return &SalesQuoteMakeInvoiceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SalesQuoteMakeInvoiceRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type SalesQuoteSendRequestBuilder struct{ BaseRequestBuilder }

// Send action undocumented
func (b *SalesQuoteRequestBuilder) Send(reqObj *SalesQuoteSendRequestParameter) *SalesQuoteSendRequestBuilder {
	bb := &SalesQuoteSendRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/send"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type SalesQuoteSendRequest struct{ BaseRequest }

//
func (b *SalesQuoteSendRequestBuilder) Request() *SalesQuoteSendRequest {
	return &SalesQuoteSendRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *SalesQuoteSendRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}
