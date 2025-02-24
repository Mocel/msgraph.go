// Code generated by msgraph.go/gen DO NOT EDIT.

package msgraph

// Site undocumented
type Site struct {
	// BaseItem is the base model of Site
	BaseItem
	// Deleted undocumented
	Deleted *Deleted `json:"deleted,omitempty"`
	// DisplayName undocumented
	DisplayName *string `json:"displayName,omitempty"`
	// Root undocumented
	Root *Root `json:"root,omitempty"`
	// SharepointIDs undocumented
	SharepointIDs *SharepointIDs `json:"sharepointIds,omitempty"`
	// SiteCollection undocumented
	SiteCollection *SiteCollection `json:"siteCollection,omitempty"`
	// Analytics undocumented
	Analytics *ItemAnalytics `json:"analytics,omitempty"`
	// Columns undocumented
	Columns []ColumnDefinition `json:"columns,omitempty"`
	// ContentTypes undocumented
	ContentTypes []ContentType `json:"contentTypes,omitempty"`
	// Drive undocumented
	Drive *Drive `json:"drive,omitempty"`
	// Drives undocumented
	Drives []Drive `json:"drives,omitempty"`
	// ExternalColumns undocumented
	ExternalColumns []ColumnDefinition `json:"externalColumns,omitempty"`
	// Items undocumented
	Items []BaseItem `json:"items,omitempty"`
	// Lists undocumented
	Lists []List `json:"lists,omitempty"`
	// Pages undocumented
	Pages []SitePage `json:"pages,omitempty"`
	// Permissions undocumented
	Permissions []Permission `json:"permissions,omitempty"`
	// Sites undocumented
	Sites []Site `json:"sites,omitempty"`
	// TermStore undocumented
	TermStore *TermStorestore `json:"termStore,omitempty"`
	// Onenote undocumented
	Onenote *Onenote `json:"onenote,omitempty"`
}

// SiteActivitySummary undocumented
type SiteActivitySummary struct {
	// Entity is the base model of SiteActivitySummary
	Entity
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// SharedExternally undocumented
	SharedExternally *int `json:"sharedExternally,omitempty"`
	// SharedInternally undocumented
	SharedInternally *int `json:"sharedInternally,omitempty"`
	// Synced undocumented
	Synced *int `json:"synced,omitempty"`
	// ViewedOrEdited undocumented
	ViewedOrEdited *int `json:"viewedOrEdited,omitempty"`
}

// SiteCollection undocumented
type SiteCollection struct {
	// Object is the base model of SiteCollection
	Object
	// DataLocationCode undocumented
	DataLocationCode *string `json:"dataLocationCode,omitempty"`
	// Hostname undocumented
	Hostname *string `json:"hostname,omitempty"`
	// Root undocumented
	Root *Root `json:"root,omitempty"`
}

// SitePage undocumented
type SitePage struct {
	// BaseItem is the base model of SitePage
	BaseItem
	// ContentType undocumented
	ContentType *ContentTypeInfo `json:"contentType,omitempty"`
	// PageLayoutType undocumented
	PageLayoutType *string `json:"pageLayoutType,omitempty"`
	// PublishingState undocumented
	PublishingState *PublicationFacet `json:"publishingState,omitempty"`
	// Title undocumented
	Title *string `json:"title,omitempty"`
	// WebParts undocumented
	WebParts []WebPart `json:"webParts,omitempty"`
}

// SitePageData undocumented
type SitePageData struct {
	// Object is the base model of SitePageData
	Object
}

// SiteUsageStorage undocumented
type SiteUsageStorage struct {
	// Entity is the base model of SiteUsageStorage
	Entity
	// ReportDate undocumented
	ReportDate *Date `json:"reportDate,omitempty"`
	// ReportPeriod undocumented
	ReportPeriod *string `json:"reportPeriod,omitempty"`
	// ReportRefreshDate undocumented
	ReportRefreshDate *Date `json:"reportRefreshDate,omitempty"`
	// SiteType undocumented
	SiteType *string `json:"siteType,omitempty"`
	// StorageUsedInBytes undocumented
	StorageUsedInBytes *int `json:"storageUsedInBytes,omitempty"`
}
