package models

type User struct {
	id_user  string
	nama     string
	email    string
	password string
}

// constructor
func NewUser(id, nama, email, password string) *User {
	return &User{
		id_user:  id,
		nama:     nama,
		email:    email,
		password: password,
	}
}

// getter
func (u *User) GetIDUser() string {
	return u.id_user
}
func (u *User) GetNama() string {
	return u.nama
}
func (u *User) GetEmail() string {
	return u.email
}

// setter
func (u *User) setNama(nama string) {
	u.nama = nama
}
func (u *User) setEmail(email string) {
	u.email = email
}
func (u *User) setPassword(password string) {
	u.password = password
}

// method
func (u *User) Login(email string, password string) bool {
	return u.email == email && u.password == password
}

func (u *User) Logout() {}

func (u *User) UpdateProfil(nama string, email string) {
	u.nama = nama
	u.email = email
}

func (u *User) GantiPassword(passwordBaru string) {
	u.password = passwordBaru
}
