package models

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Login(email, password string) bool {
	return false
}

func (u *User) Logout() bool {
	return true
}
