// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import "encoding/json"

// WorkbookFilterApplyRequestParameter undocumented
type WorkbookFilterApplyRequestParameter struct {
	// Criteria undocumented
	Criteria *WorkbookFilterCriteria `json:"criteria,omitempty"`
}

// WorkbookFilterApplyBottomItemsFilterRequestParameter undocumented
type WorkbookFilterApplyBottomItemsFilterRequestParameter struct {
	// Count undocumented
	Count *int `json:"count,omitempty"`
}

// WorkbookFilterApplyBottomPercentFilterRequestParameter undocumented
type WorkbookFilterApplyBottomPercentFilterRequestParameter struct {
	// Percent undocumented
	Percent *int `json:"percent,omitempty"`
}

// WorkbookFilterApplyCellColorFilterRequestParameter undocumented
type WorkbookFilterApplyCellColorFilterRequestParameter struct {
	// Color undocumented
	Color *string `json:"color,omitempty"`
}

// WorkbookFilterApplyCustomFilterRequestParameter undocumented
type WorkbookFilterApplyCustomFilterRequestParameter struct {
	// Criteria1 undocumented
	Criteria1 *string `json:"criteria1,omitempty"`
	// Criteria2 undocumented
	Criteria2 *string `json:"criteria2,omitempty"`
	// Oper undocumented
	Oper *string `json:"oper,omitempty"`
}

// WorkbookFilterApplyDynamicFilterRequestParameter undocumented
type WorkbookFilterApplyDynamicFilterRequestParameter struct {
	// Criteria undocumented
	Criteria *string `json:"criteria,omitempty"`
}

// WorkbookFilterApplyFontColorFilterRequestParameter undocumented
type WorkbookFilterApplyFontColorFilterRequestParameter struct {
	// Color undocumented
	Color *string `json:"color,omitempty"`
}

// WorkbookFilterApplyIconFilterRequestParameter undocumented
type WorkbookFilterApplyIconFilterRequestParameter struct {
	// Icon undocumented
	Icon *WorkbookIcon `json:"icon,omitempty"`
}

// WorkbookFilterApplyTopItemsFilterRequestParameter undocumented
type WorkbookFilterApplyTopItemsFilterRequestParameter struct {
	// Count undocumented
	Count *int `json:"count,omitempty"`
}

// WorkbookFilterApplyTopPercentFilterRequestParameter undocumented
type WorkbookFilterApplyTopPercentFilterRequestParameter struct {
	// Percent undocumented
	Percent *int `json:"percent,omitempty"`
}

// WorkbookFilterApplyValuesFilterRequestParameter undocumented
type WorkbookFilterApplyValuesFilterRequestParameter struct {
	// Values undocumented
	Values json.RawMessage `json:"values,omitempty"`
}

// WorkbookFilterClearRequestParameter undocumented
type WorkbookFilterClearRequestParameter struct {
}

//
type WorkbookFilterApplyRequestBuilder struct{ BaseRequestBuilder }

// Apply action undocumented
func (b *WorkbookFilterRequestBuilder) Apply(reqObj *WorkbookFilterApplyRequestParameter) *WorkbookFilterApplyRequestBuilder {
	bb := &WorkbookFilterApplyRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/apply"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFilterApplyRequest struct{ BaseRequest }

//
func (b *WorkbookFilterApplyRequestBuilder) Request() *WorkbookFilterApplyRequest {
	return &WorkbookFilterApplyRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFilterApplyRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type WorkbookFilterApplyBottomItemsFilterRequestBuilder struct{ BaseRequestBuilder }

// ApplyBottomItemsFilter action undocumented
func (b *WorkbookFilterRequestBuilder) ApplyBottomItemsFilter(reqObj *WorkbookFilterApplyBottomItemsFilterRequestParameter) *WorkbookFilterApplyBottomItemsFilterRequestBuilder {
	bb := &WorkbookFilterApplyBottomItemsFilterRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/applyBottomItemsFilter"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFilterApplyBottomItemsFilterRequest struct{ BaseRequest }

//
func (b *WorkbookFilterApplyBottomItemsFilterRequestBuilder) Request() *WorkbookFilterApplyBottomItemsFilterRequest {
	return &WorkbookFilterApplyBottomItemsFilterRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFilterApplyBottomItemsFilterRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type WorkbookFilterApplyBottomPercentFilterRequestBuilder struct{ BaseRequestBuilder }

// ApplyBottomPercentFilter action undocumented
func (b *WorkbookFilterRequestBuilder) ApplyBottomPercentFilter(reqObj *WorkbookFilterApplyBottomPercentFilterRequestParameter) *WorkbookFilterApplyBottomPercentFilterRequestBuilder {
	bb := &WorkbookFilterApplyBottomPercentFilterRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/applyBottomPercentFilter"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFilterApplyBottomPercentFilterRequest struct{ BaseRequest }

//
func (b *WorkbookFilterApplyBottomPercentFilterRequestBuilder) Request() *WorkbookFilterApplyBottomPercentFilterRequest {
	return &WorkbookFilterApplyBottomPercentFilterRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFilterApplyBottomPercentFilterRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type WorkbookFilterApplyCellColorFilterRequestBuilder struct{ BaseRequestBuilder }

// ApplyCellColorFilter action undocumented
func (b *WorkbookFilterRequestBuilder) ApplyCellColorFilter(reqObj *WorkbookFilterApplyCellColorFilterRequestParameter) *WorkbookFilterApplyCellColorFilterRequestBuilder {
	bb := &WorkbookFilterApplyCellColorFilterRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/applyCellColorFilter"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFilterApplyCellColorFilterRequest struct{ BaseRequest }

//
func (b *WorkbookFilterApplyCellColorFilterRequestBuilder) Request() *WorkbookFilterApplyCellColorFilterRequest {
	return &WorkbookFilterApplyCellColorFilterRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFilterApplyCellColorFilterRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type WorkbookFilterApplyCustomFilterRequestBuilder struct{ BaseRequestBuilder }

// ApplyCustomFilter action undocumented
func (b *WorkbookFilterRequestBuilder) ApplyCustomFilter(reqObj *WorkbookFilterApplyCustomFilterRequestParameter) *WorkbookFilterApplyCustomFilterRequestBuilder {
	bb := &WorkbookFilterApplyCustomFilterRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/applyCustomFilter"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFilterApplyCustomFilterRequest struct{ BaseRequest }

//
func (b *WorkbookFilterApplyCustomFilterRequestBuilder) Request() *WorkbookFilterApplyCustomFilterRequest {
	return &WorkbookFilterApplyCustomFilterRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFilterApplyCustomFilterRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type WorkbookFilterApplyDynamicFilterRequestBuilder struct{ BaseRequestBuilder }

// ApplyDynamicFilter action undocumented
func (b *WorkbookFilterRequestBuilder) ApplyDynamicFilter(reqObj *WorkbookFilterApplyDynamicFilterRequestParameter) *WorkbookFilterApplyDynamicFilterRequestBuilder {
	bb := &WorkbookFilterApplyDynamicFilterRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/applyDynamicFilter"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFilterApplyDynamicFilterRequest struct{ BaseRequest }

//
func (b *WorkbookFilterApplyDynamicFilterRequestBuilder) Request() *WorkbookFilterApplyDynamicFilterRequest {
	return &WorkbookFilterApplyDynamicFilterRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFilterApplyDynamicFilterRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type WorkbookFilterApplyFontColorFilterRequestBuilder struct{ BaseRequestBuilder }

// ApplyFontColorFilter action undocumented
func (b *WorkbookFilterRequestBuilder) ApplyFontColorFilter(reqObj *WorkbookFilterApplyFontColorFilterRequestParameter) *WorkbookFilterApplyFontColorFilterRequestBuilder {
	bb := &WorkbookFilterApplyFontColorFilterRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/applyFontColorFilter"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFilterApplyFontColorFilterRequest struct{ BaseRequest }

//
func (b *WorkbookFilterApplyFontColorFilterRequestBuilder) Request() *WorkbookFilterApplyFontColorFilterRequest {
	return &WorkbookFilterApplyFontColorFilterRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFilterApplyFontColorFilterRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type WorkbookFilterApplyIconFilterRequestBuilder struct{ BaseRequestBuilder }

// ApplyIconFilter action undocumented
func (b *WorkbookFilterRequestBuilder) ApplyIconFilter(reqObj *WorkbookFilterApplyIconFilterRequestParameter) *WorkbookFilterApplyIconFilterRequestBuilder {
	bb := &WorkbookFilterApplyIconFilterRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/applyIconFilter"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFilterApplyIconFilterRequest struct{ BaseRequest }

//
func (b *WorkbookFilterApplyIconFilterRequestBuilder) Request() *WorkbookFilterApplyIconFilterRequest {
	return &WorkbookFilterApplyIconFilterRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFilterApplyIconFilterRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type WorkbookFilterApplyTopItemsFilterRequestBuilder struct{ BaseRequestBuilder }

// ApplyTopItemsFilter action undocumented
func (b *WorkbookFilterRequestBuilder) ApplyTopItemsFilter(reqObj *WorkbookFilterApplyTopItemsFilterRequestParameter) *WorkbookFilterApplyTopItemsFilterRequestBuilder {
	bb := &WorkbookFilterApplyTopItemsFilterRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/applyTopItemsFilter"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFilterApplyTopItemsFilterRequest struct{ BaseRequest }

//
func (b *WorkbookFilterApplyTopItemsFilterRequestBuilder) Request() *WorkbookFilterApplyTopItemsFilterRequest {
	return &WorkbookFilterApplyTopItemsFilterRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFilterApplyTopItemsFilterRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type WorkbookFilterApplyTopPercentFilterRequestBuilder struct{ BaseRequestBuilder }

// ApplyTopPercentFilter action undocumented
func (b *WorkbookFilterRequestBuilder) ApplyTopPercentFilter(reqObj *WorkbookFilterApplyTopPercentFilterRequestParameter) *WorkbookFilterApplyTopPercentFilterRequestBuilder {
	bb := &WorkbookFilterApplyTopPercentFilterRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/applyTopPercentFilter"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFilterApplyTopPercentFilterRequest struct{ BaseRequest }

//
func (b *WorkbookFilterApplyTopPercentFilterRequestBuilder) Request() *WorkbookFilterApplyTopPercentFilterRequest {
	return &WorkbookFilterApplyTopPercentFilterRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFilterApplyTopPercentFilterRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type WorkbookFilterApplyValuesFilterRequestBuilder struct{ BaseRequestBuilder }

// ApplyValuesFilter action undocumented
func (b *WorkbookFilterRequestBuilder) ApplyValuesFilter(reqObj *WorkbookFilterApplyValuesFilterRequestParameter) *WorkbookFilterApplyValuesFilterRequestBuilder {
	bb := &WorkbookFilterApplyValuesFilterRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/applyValuesFilter"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFilterApplyValuesFilterRequest struct{ BaseRequest }

//
func (b *WorkbookFilterApplyValuesFilterRequestBuilder) Request() *WorkbookFilterApplyValuesFilterRequest {
	return &WorkbookFilterApplyValuesFilterRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFilterApplyValuesFilterRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}

//
type WorkbookFilterClearRequestBuilder struct{ BaseRequestBuilder }

// Clear action undocumented
func (b *WorkbookFilterRequestBuilder) Clear(reqObj *WorkbookFilterClearRequestParameter) *WorkbookFilterClearRequestBuilder {
	bb := &WorkbookFilterClearRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/clear"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type WorkbookFilterClearRequest struct{ BaseRequest }

//
func (b *WorkbookFilterClearRequestBuilder) Request() *WorkbookFilterClearRequest {
	return &WorkbookFilterClearRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *WorkbookFilterClearRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}
