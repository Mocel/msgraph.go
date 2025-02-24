// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// BookingAppointmentRequestBuilder is request builder for BookingAppointment
type BookingAppointmentRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingAppointmentRequest
func (b *BookingAppointmentRequestBuilder) Request() *BookingAppointmentRequest {
	return &BookingAppointmentRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingAppointmentRequest is request for BookingAppointment
type BookingAppointmentRequest struct{ BaseRequest }

// Get performs GET request for BookingAppointment
func (r *BookingAppointmentRequest) Get(ctx context.Context) (resObj *BookingAppointment, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingAppointment
func (r *BookingAppointmentRequest) Update(ctx context.Context, reqObj *BookingAppointment) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingAppointment
func (r *BookingAppointmentRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingBusinessRequestBuilder is request builder for BookingBusiness
type BookingBusinessRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingBusinessRequest
func (b *BookingBusinessRequestBuilder) Request() *BookingBusinessRequest {
	return &BookingBusinessRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingBusinessRequest is request for BookingBusiness
type BookingBusinessRequest struct{ BaseRequest }

// Get performs GET request for BookingBusiness
func (r *BookingBusinessRequest) Get(ctx context.Context) (resObj *BookingBusiness, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingBusiness
func (r *BookingBusinessRequest) Update(ctx context.Context, reqObj *BookingBusiness) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingBusiness
func (r *BookingBusinessRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingCurrencyRequestBuilder is request builder for BookingCurrency
type BookingCurrencyRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingCurrencyRequest
func (b *BookingCurrencyRequestBuilder) Request() *BookingCurrencyRequest {
	return &BookingCurrencyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingCurrencyRequest is request for BookingCurrency
type BookingCurrencyRequest struct{ BaseRequest }

// Get performs GET request for BookingCurrency
func (r *BookingCurrencyRequest) Get(ctx context.Context) (resObj *BookingCurrency, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingCurrency
func (r *BookingCurrencyRequest) Update(ctx context.Context, reqObj *BookingCurrency) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingCurrency
func (r *BookingCurrencyRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingCustomerRequestBuilder is request builder for BookingCustomer
type BookingCustomerRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingCustomerRequest
func (b *BookingCustomerRequestBuilder) Request() *BookingCustomerRequest {
	return &BookingCustomerRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingCustomerRequest is request for BookingCustomer
type BookingCustomerRequest struct{ BaseRequest }

// Get performs GET request for BookingCustomer
func (r *BookingCustomerRequest) Get(ctx context.Context) (resObj *BookingCustomer, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingCustomer
func (r *BookingCustomerRequest) Update(ctx context.Context, reqObj *BookingCustomer) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingCustomer
func (r *BookingCustomerRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingServiceRequestBuilder is request builder for BookingService
type BookingServiceRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingServiceRequest
func (b *BookingServiceRequestBuilder) Request() *BookingServiceRequest {
	return &BookingServiceRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingServiceRequest is request for BookingService
type BookingServiceRequest struct{ BaseRequest }

// Get performs GET request for BookingService
func (r *BookingServiceRequest) Get(ctx context.Context) (resObj *BookingService, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingService
func (r *BookingServiceRequest) Update(ctx context.Context, reqObj *BookingService) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingService
func (r *BookingServiceRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

// BookingStaffMemberRequestBuilder is request builder for BookingStaffMember
type BookingStaffMemberRequestBuilder struct{ BaseRequestBuilder }

// Request returns BookingStaffMemberRequest
func (b *BookingStaffMemberRequestBuilder) Request() *BookingStaffMemberRequest {
	return &BookingStaffMemberRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// BookingStaffMemberRequest is request for BookingStaffMember
type BookingStaffMemberRequest struct{ BaseRequest }

// Get performs GET request for BookingStaffMember
func (r *BookingStaffMemberRequest) Get(ctx context.Context) (resObj *BookingStaffMember, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for BookingStaffMember
func (r *BookingStaffMemberRequest) Update(ctx context.Context, reqObj *BookingStaffMember) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for BookingStaffMember
func (r *BookingStaffMemberRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}
