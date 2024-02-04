package api

// Album
// Describes an album.
type Album struct {
	ID            string `json:"id,omitempty"`
	Barcode       string `json:"barcode" binding:"required,min=11,max=13,numeric"`
	Artist        string `json:"artist" binding:"required,min=1,max=64"`
	Album         string `json:"album" binding:"required,min=1,max=64"`
	Label         string `json:"label,omitempty" binding:"omitempty,min=1,max=64"`
	Country       string `json:"country,omitempty" binding:"omitempty,len=2"`
	Source        string `json:"source" binding:"required,oneof=CDROM SHMCD SACD DVDA BLURAY DOWNLOAD"`
	CatalogNumber string `json:"catalogNumber" binding:"omitempty,min=1,max=16"`
	Comment       string `json:"comment,omitempty" binding:"omitempty,min=1,max=1024"`

	DRLog string `json:"drLog,omitempty" binding:"required,drlog"`
	Art   string `json:"art,omitempty" binding:"required,encodedimage"`

	LabelCode    string `json:"labelCode" binding:"omitempty,min=4,max=5"`
	ReleaseYear  int    `json:"releaseYear" binding:"required,numeric,min=1000,max=3000"`
	AverageDR    int8   `json:"averageDR"`
	MinimumDR    int8   `json:"minimumDR"`
	MaximumDR    int8   `json:"maximumDR"`
	InCollection bool   `json:"in_collection,omitempty"`

	Tracks []*Track `json:"tracks,omitempty"`
}

func (a *Album) SourceFriendly() string {
	if a.Source == "CDROM" {
		return "CD"
	}

	if a.Source == "SHMCD" {
		return "SHM-CD"
	}

	if a.Source == "SACD" {
		return "SACD"
	}

	if a.Source == "DVDA" {
		return "DVD-Audio"
	}

	if a.Source == "BLURAY" {
		return "Blu-ray"
	}

	if a.Source == "DOWNLOAD" {
		return "Download"
	}

	return "Unknown"
}
