package models

import (
	"backend-absensi/config"
	"fmt"
)

type Mahasiswa struct {
	User
	NIM      string `json:"nim"`
	Nama     string `json:"nama"`
	Prodi    string `json:"prodi"`
	Fakultas string `json:"fakultas"`
	Angkatan string `json:"angkatan"` // Tambahin ini!
	Email    string `json:"email"`
	NoHP     string `json:"nohp"`
}

func (m *Mahasiswa) Login(email, password string) bool {
	query := `
        SELECT m.nim, m.nama, u.password 
        FROM mahasiswa m 
        JOIN User u ON m.id_user = u.id_user 
        WHERE u.email = ? AND u.role = 'mahasiswa'`

	err := config.DB.QueryRow(query, email).Scan(&m.NIM, &m.Nama, &m.Password)
	if err != nil {
		fmt.Println("DB Error:", err) // lihat error di terminal
		return false
	}

	fmt.Println("DB Password:", m.Password, "| Input:", password) // lihat nilai passwordnya
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

func (m *Mahasiswa) GetAttribute(nim string) (Mahasiswa, error) {
	var profil Mahasiswa
	// Aliasnya gue ganti 'm', dan 'angkatan' udah masuk radar!
	query := `
		SELECT m.nim, m.nama, m.prodi, m.fakultas, m.angkatan, u.email, m.nohp 
		FROM mahasiswa m 
		JOIN User u ON m.id_user = u.id_user 
		WHERE m.nim = ?`

	// Scan harus nangkep 7 data sesuai urutan SELECT di atas
	err := config.DB.QueryRow(query, nim).Scan(
		&profil.NIM,
		&profil.Nama,
		&profil.Prodi,
		&profil.Fakultas,
		&profil.Angkatan,
		&profil.Email,
		&profil.NoHP,
	)
	if err != nil {
		return profil, err
	}
	return profil, nil
}
