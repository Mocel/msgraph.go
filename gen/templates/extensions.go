// +build templates

// This file is a part of msgraph.go/gen/templates.
// Anything until the first appearance of "// BEGIN" line will be ignored.

package msgraph

import (
	"context"
	"fmt"
	"net/url"
	"strings"
)

// BEGIN - everything below this line will be copied to the output

// CallRecordscallRecord is undefined in the spec, so we add a simple interface typedef
type CallRecordscallRecord interface{}

// ItemWithPath returns DriveItemRequestBuilder addressed by relative path
func (b *DriveItemRequestBuilder) ItemWithPath(path string) *DriveItemRequestBuilder {
	bb := &DriveItemRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	if len(path) == 0 || path[0] != '/' {
		path = "/" + path
	}
	bb.baseURL += ":" + path + ":"
	return bb
}

// GetByPath returns SiteRequestBuilder addressed by hostname and path
func (b *GraphServiceSitesCollectionRequestBuilder) GetByPath(hostname, path string) *SiteRequestBuilder {
	bb := &SiteRequestBuilder{BaseRequestBuilder: b.BaseRequestBuilder}
	if len(path) == 0 || path[0] != '/' {
		path = "/" + path
	}
	bb.baseURL += "/" + hostname + ":" + path + ":"
	return bb
}

// GetDriveItemByURL returns DriveItemRequestBuilder addressed by URL
func (b *GraphServiceRequestBuilder) GetDriveItemByURL(ctx context.Context, itemURL string) (*DriveItemRequestBuilder, error) {
	u, err := url.Parse(itemURL)
	if err != nil {
		return nil, err
	}
	u.RawQuery = ""
	itemURL = u.String()
	var site *Site
	segments := strings.Split(u.Path, "/")
	for i := 3; i <= len(segments); i++ {
		site, err = b.Sites().GetByPath(u.Hostname(), strings.Join(segments[:i], "/")).Request().Get(ctx)
		if err == nil {
			break
		}
	}
	if site == nil {
		return nil, fmt.Errorf("Site for %s not found", itemURL)
	}
	drives, err := b.Sites().ID(*site.ID).Drives().Request().Get(ctx)
	if err != nil {
		return nil, err
	}
	for _, drive := range drives {
		if strings.HasPrefix(itemURL, *drive.WebURL) {
			path := itemURL[len(*drive.WebURL):]
			return b.Drives().ID(*drive.ID).Root().ItemWithPath(path), nil
		}
	}
	return nil, fmt.Errorf("DriveItem for %s not found", itemURL)
}
