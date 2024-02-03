package form

// Album
// Validates the add album form.
type Album struct {
	Artist  string `form:"artist" binding:"required,min=1,max=64"`
	Album   string `form:"album" binding:"required,min=1,max=64"`
	Barcode string `form:"barcode" binding:"required" binding:"required,min=11,max=13,numeric"`
	Source  string `form:"source" binding:"required,oneof=CDROM SHMCD SACD DVDA BLURAY DOWNLOAD"`

	ReleaseYear int `form:"release_year" binding:"required,numeric,min=1000,max=3000"`

	DRLog string
	Art   string

	Label         string `form:"label" binding:"omitempty,min=1,max=64"`
	Country       string `form:"country" binding:"omitempty,len=2"`
	Comment       string `form:"comment" binding:"omitempty,min=1,max=1024"`
	CatalogNumber string `form:"catalog_number" binding:"omitempty,min=1,max=16"`
	LabelCode     int    `form:"label_code" binding:"omitempty,min=4,max=5"`
}
