// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// CalendarGetScheduleRequestParameter undocumented
type CalendarGetScheduleRequestParameter struct {
	// Schedules undocumented
	Schedules []string `json:"Schedules,omitempty"`
	// EndTime undocumented
	EndTime *DateTimeTimeZone `json:"EndTime,omitempty"`
	// StartTime undocumented
	StartTime *DateTimeTimeZone `json:"StartTime,omitempty"`
	// AvailabilityViewInterval undocumented
	AvailabilityViewInterval *int `json:"AvailabilityViewInterval,omitempty"`
}

//
type CalendarGetScheduleRequestBuilder struct{ BaseRequestBuilder }

// GetSchedule action undocumented
func (b *CalendarRequestBuilder) GetSchedule(reqObj *CalendarGetScheduleRequestParameter) *CalendarGetScheduleRequestBuilder {
	bb := &CalendarGetScheduleRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getSchedule"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type CalendarGetScheduleRequest struct{ BaseRequest }

//
func (b *CalendarGetScheduleRequestBuilder) Request() *CalendarGetScheduleRequest {
	return &CalendarGetScheduleRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *CalendarGetScheduleRequest) Paging(method, path string, obj interface{}) ([][]ScheduleInformation, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values [][]ScheduleInformation
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  [][]ScheduleInformation
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
func (r *CalendarGetScheduleRequest) Get() ([][]ScheduleInformation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}
