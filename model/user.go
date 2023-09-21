package model

type User struct {
	ID       int64  `json:"-"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Email    string `json:"email"`
	Password string `json:",omitempty"`
}
