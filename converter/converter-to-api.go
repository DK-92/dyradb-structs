package converter

import (
	"github.com/DK-92/dyradb-structs/model/api"
	"github.com/DK-92/dyradb-structs/model/form"
)

func ConvertFormToAlbumAPI(form *form.Album) *api.Album {
	return &api.Album{
		Barcode:     form.Barcode,
		Artist:      form.Artist,
		Album:       form.Album,
		Label:       form.Label,
		Country:     form.Country,
		Source:      form.Source,
		Comment:     form.Comment,
		ReleaseYear: form.ReleaseYear,
		DRLog:       form.DRLog,
		Art:         form.Art,
	}
}

func ConvertFormToUserAPI(form *form.User) *api.User {
	return &api.User{
		Name:     form.Name,
		IsPublic: form.IsPublicBool(),
	}
}
