package api

import (
	"bytes"
	"encoding/gob"
	"errors"
)

const (
	magicPrefix = "DYRADB"
	magicSuffix = "251225"
)

// AlbumList
// Describes a list of albums, together with pagination information.
type AlbumList struct {
	TotalAlbums int         `json:"total_albums"`
	PageCount   int         `json:"page_count,omitempty"`
	ExportMeta  *ExportMeta `json:"export_meta,omitempty"`

	Albums []*Album `json:"albums"`
}

func (al *AlbumList) ExportAsBinary() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(al); err != nil {
		return nil, err
	}

	prefix := []byte(magicPrefix)
	data := buf.Bytes()
	suffix := []byte(magicSuffix)

	result := join(len(prefix)+len(data)+len(suffix), prefix, data, suffix)

	return result, nil
}

func ImportToAlbumList(data []byte) (*AlbumList, error) {
	prefix := []byte(magicPrefix)
	suffix := []byte(magicSuffix)
	minLen := len(prefix) + len(suffix)

	if len(data) < minLen {
		return nil, errors.New("data too short")
	}

	if !bytes.Equal(data[:len(prefix)], prefix) {
		return nil, errors.New("invalid prefix")
	}

	if !bytes.Equal(data[len(data)-len(suffix):], suffix) {
		return nil, errors.New("invalid suffix")
	}

	encodedData := data[len(prefix) : len(data)-len(suffix)]

	var al AlbumList
	dec := gob.NewDecoder(bytes.NewReader(encodedData))
	if err := dec.Decode(&al); err != nil {
		return nil, err
	}

	return &al, nil
}

func join(s int, ba ...[]byte) []byte {
	b, i := make([]byte, s), 0

	for _, v := range ba {
		i += copy(b[i:], v)
	}

	return b
}
