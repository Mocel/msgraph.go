// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// DriveItemVersionRestoreVersionRequestParameter undocumented
type DriveItemVersionRestoreVersionRequestParameter struct {
}

//
type DriveItemVersionRestoreVersionRequestBuilder struct{ BaseRequestBuilder }

// RestoreVersion action undocumented
func (b *DriveItemVersionRequestBuilder) RestoreVersion(reqObj *DriveItemVersionRestoreVersionRequestParameter) *DriveItemVersionRestoreVersionRequestBuilder {
	bb := &DriveItemVersionRestoreVersionRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/restoreVersion"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type DriveItemVersionRestoreVersionRequest struct{ BaseRequest }

//
func (b *DriveItemVersionRestoreVersionRequestBuilder) Request() *DriveItemVersionRestoreVersionRequest {
	return &DriveItemVersionRestoreVersionRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *DriveItemVersionRestoreVersionRequest) Post() error {
	return r.JSONRequest("POST", "", r.requestObject, nil)
}
