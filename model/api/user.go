package api

// User
// Contains information about a user in our system. Please note that the ID is the 'secret key'.
type User struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name" binding:"required,min=1,max=64"`
	IsPublic bool   `json:"is_public" binding:"omitempty,boolean"`
}
