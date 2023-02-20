package form

// User
// Validates the new user form.
type User struct {
	Name     string `form:"name" binding:"required,min=1,max=64"`
	IsPublic string `form:"is_public" binding:"max=2"`
}

func (u *User) IsPublicBool() bool {
	return u.IsPublic == "on"
}

func InvalidUsername(username string) bool {
	if len(username) < 1 || len(username) > 64 {
		return true
	}

	return false
}
