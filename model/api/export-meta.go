package api

// ExportMeta
// Contains information about the exported albums.
type ExportMeta struct {
	ExportDate int64  `json:"export_date"`
	Version    string `json:"version"`
	Username   string `json:"username"`
}
