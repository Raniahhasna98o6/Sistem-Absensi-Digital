package models
import (
	"backend-absensi/config"
)
type User struct {
	ID       string `json:"id_user"`
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
// constructor
func NewUser(id, nama, email, password string) *User {
	return &User{
		ID:       id,
		Nama:     nama,
		Email:    email,
		Password: password,
	}
}
// cek password
func (u *User) CheckPassword(password string) bool {
	return u.Password == password
}
// getter
func (u *User) GetIDUser() string {
	return u.ID
}
func (u *User) GetNama() string {
	return u.Nama
}
func (u *User) GetEmail() string {
	return u.Email
}

// setter
func (u *User) SetNama(nama string) {
	u.Nama = nama
}
func (u *User) SetEmail(email string) {
	u.Email = email
}
func (u *User) SetPassword(password string) {
	u.Password = password
}

// method
func (u *User) Login(email string, password string) bool {
	return u.Email == email && u.Password == password
}
func (u *User) Logout() {}
func (u *User) UpdateProfil(nama string, email string) {
	u.Nama = nama
	u.Email = email
}
func (u *User) GantiPassword(passwordBaru string) {
	u.Password = passwordBaru
}

// ambil user dari database
func GetUserByEmail(email string) (*User, error) {
	var user User

	query := "SELECT id_user, nama, email, password FROM User WHERE email = ?"
	row := config.DB.QueryRow(query, email)

	err := row.Scan(&user.ID, &user.Nama, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
