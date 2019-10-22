// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/yaegashi/msgraph.go/jsonx"
)

// UserTeamworkRequestBuilder is request builder for UserTeamwork
type UserTeamworkRequestBuilder struct{ BaseRequestBuilder }

// Request returns UserTeamworkRequest
func (b *UserTeamworkRequestBuilder) Request() *UserTeamworkRequest {
	return &UserTeamworkRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// UserTeamworkRequest is request for UserTeamwork
type UserTeamworkRequest struct{ BaseRequest }

// Get performs GET request for UserTeamwork
func (r *UserTeamworkRequest) Get() (resObj *UserTeamwork, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest("GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for UserTeamwork
func (r *UserTeamworkRequest) Update(reqObj *UserTeamwork) error {
	return r.JSONRequest("PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for UserTeamwork
func (r *UserTeamworkRequest) Delete() error {
	return r.JSONRequest("DELETE", "", nil, nil)
}

// InstalledApps returns request builder for TeamsAppInstallation collection
func (b *UserTeamworkRequestBuilder) InstalledApps() *UserTeamworkInstalledAppsCollectionRequestBuilder {
	bb := &UserTeamworkInstalledAppsCollectionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/installedApps"
	return bb
}

// UserTeamworkInstalledAppsCollectionRequestBuilder is request builder for TeamsAppInstallation collection
type UserTeamworkInstalledAppsCollectionRequestBuilder struct{ BaseRequestBuilder }

// Request returns request for TeamsAppInstallation collection
func (b *UserTeamworkInstalledAppsCollectionRequestBuilder) Request() *UserTeamworkInstalledAppsCollectionRequest {
	return &UserTeamworkInstalledAppsCollectionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// ID returns request builder for TeamsAppInstallation item
func (b *UserTeamworkInstalledAppsCollectionRequestBuilder) ID(id string) *TeamsAppInstallationRequestBuilder {
	bb := &TeamsAppInstallationRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.baseURL += "/" + id
	return bb
}

// UserTeamworkInstalledAppsCollectionRequest is request for TeamsAppInstallation collection
type UserTeamworkInstalledAppsCollectionRequest struct{ BaseRequest }

// Paging perfoms paging operation for TeamsAppInstallation collection
func (r *UserTeamworkInstalledAppsCollectionRequest) Paging(method, path string, obj interface{}) ([]TeamsAppInstallation, error) {
	req, err := r.NewJSONRequest(method, path, obj)
	if err != nil {
		return nil, err
	}
	res, err := r.client.Do(req)
	if err != nil {
		return nil, err
	}
	var values []TeamsAppInstallation
	for {
		defer res.Body.Close()
		if res.StatusCode != http.StatusOK {
			b, _ := ioutil.ReadAll(res.Body)
			return nil, fmt.Errorf("%s: %s", res.Status, string(b))
		}
		var (
			paging Paging
			value  []TeamsAppInstallation
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

// Get performs GET request for TeamsAppInstallation collection
func (r *UserTeamworkInstalledAppsCollectionRequest) Get() ([]TeamsAppInstallation, error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	return r.Paging("GET", query, nil)
}

// Add performs POST request for TeamsAppInstallation collection
func (r *UserTeamworkInstalledAppsCollectionRequest) Add(reqObj *TeamsAppInstallation) (resObj *TeamsAppInstallation, err error) {
	err = r.JSONRequest("POST", "", reqObj, &resObj)
	return
}
