package models

import (
	"backend-absensi/config"
	"fmt"
)

type Mahasiswa struct {
	User
	NIM  string `json:"nim"`
	Nama string `json:"nama"`
}

func (m *Mahasiswa) Login(email, password string) bool {
	query := `
        SELECT m.nim, m.nama, u.password 
        FROM mahasiswa m 
        JOIN User u ON m.id_user = u.id_user 
        WHERE u.email = ? AND u.role = 'mahasiswa'`

	err := config.DB.QueryRow(query, email).Scan(&m.NIM, &m.Nama, &m.Password)
	if err != nil {
		// TAMBAHIN LOG INI SEAN! Biar ketauan errornya apa
		fmt.Println("DEBUG LOGIN ERROR:", err)
		return false
	}

	return m.Password != "" && m.Password == password
}

func (m *Mahasiswa) Logout() bool {
	return true
}

func (m *Mahasiswa) InputKehadiran(data Absensi) (bool, string) {
	if data.CekRadiusLokasi() {
		return data.SimpanKeDatabase(m.NIM)
	}
	return false, "Gagal! Anda berada di luar radius kampus!"
}

func (m *Mahasiswa) BukaRiwayat() []Absensi {
	var a Absensi
	return a.AmbilDataAbsensi(m.NIM)
}
