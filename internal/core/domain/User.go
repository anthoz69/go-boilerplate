package domain

type User struct {
	ID       string   `json:"id"`
	Username string   `json:"username"`
	Password string   `json:"password"`
	Role     UserRole `json:"role"`
}

type UserRole int

const (
	Normal UserRole = iota + 1
	Editor
	Admin = 99
)

var userRoleLabels = map[UserRole]string{
	Normal: "Normal",
	Editor: "Editor",
	Admin:  "Admin",
}

func (r UserRole) String() string {
	return userRoleLabels[r]
}
