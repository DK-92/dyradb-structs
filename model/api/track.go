package api

// Track
// Describes one track. A track is a part of an album.
type Track struct {
	ID           string  `json:"id,omitempty"`
	Name         string  `json:"name" binding:"required,min=1,max=128"`
	Peak         float32 `json:"peakDB" binding:"required,numeric,min=-100,max=0"`
	Average      float32 `json:"averageDB" binding:"required,numeric,min=-100,max=0"`
	Position     int8    `json:"position" binding:"required,numeric,min=1,max=99"`
	DynamicRange int8    `json:"dr" binding:"required,numeric,min=1,max=32"`
}
