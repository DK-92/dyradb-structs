package api

// AlbumList
// Describes a list of albums, together with pagination information.
type AlbumList struct {
	TotalAlbums int         `json:"total_albums"`
	PageCount   int         `json:"page_count,omitempty"`
	ExportMeta  *ExportMeta `json:"export_meta,omitempty"`

	Albums []*Album `json:"albums"`
}
