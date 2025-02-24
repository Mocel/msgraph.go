// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

import "context"

// NotebookRequestBuilder is request builder for Notebook
type NotebookRequestBuilder struct{ BaseRequestBuilder }

// Request returns NotebookRequest
func (b *NotebookRequestBuilder) Request() *NotebookRequest {
	return &NotebookRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client},
	}
}

// NotebookRequest is request for Notebook
type NotebookRequest struct{ BaseRequest }

// Get performs GET request for Notebook
func (r *NotebookRequest) Get(ctx context.Context) (resObj *Notebook, err error) {
	var query string
	if r.query != nil {
		query = "?" + r.query.Encode()
	}
	err = r.JSONRequest(ctx, "GET", query, nil, &resObj)
	return
}

// Update performs PATCH request for Notebook
func (r *NotebookRequest) Update(ctx context.Context, reqObj *Notebook) error {
	return r.JSONRequest(ctx, "PATCH", "", reqObj, nil)
}

// Delete performs DELETE request for Notebook
func (r *NotebookRequest) Delete(ctx context.Context) error {
	return r.JSONRequest(ctx, "DELETE", "", nil, nil)
}

//
type NotebookCollectionGetNotebookFromWebURLRequestBuilder struct{ BaseRequestBuilder }

// GetNotebookFromWebURL action undocumented
func (b *OnenoteNotebooksCollectionRequestBuilder) GetNotebookFromWebURL(reqObj *NotebookCollectionGetNotebookFromWebURLRequestParameter) *NotebookCollectionGetNotebookFromWebURLRequestBuilder {
	bb := &NotebookCollectionGetNotebookFromWebURLRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	bb.BaseRequestBuilder.baseURL += "/getNotebookFromWebUrl"
	bb.BaseRequestBuilder.requestObject = reqObj
	return bb
}

//
type NotebookCollectionGetNotebookFromWebURLRequest struct{ BaseRequest }

//
func (b *NotebookCollectionGetNotebookFromWebURLRequestBuilder) Request() *NotebookCollectionGetNotebookFromWebURLRequest {
	return &NotebookCollectionGetNotebookFromWebURLRequest{
		BaseRequest: BaseRequest{baseURL: b.baseURL, client: b.client, requestObject: b.requestObject},
	}
}

//
func (r *NotebookCollectionGetNotebookFromWebURLRequest) Post(ctx context.Context) (resObj *CopyNotebookModel, err error) {
	err = r.JSONRequest(ctx, "POST", "", r.requestObject, &resObj)
	return
}
