package form

// ImportUser
// Contains the private key needed to import the user.
type ImportUser struct {
	Key string `form:"key" binding:"required,uuid4"`
}
