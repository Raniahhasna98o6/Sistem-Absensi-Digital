package models

import (
	"backend-absensi/config"
	"errors"
)

type Dosen struct {
	User
	NIDN string `json:"nidn"`
	Nama string `json:"nama"`
}

func (d *Dosen) Login(email, password string) bool {
	query := `
		SELECT d.nidn, d.nama, u.password 
		FROM dosen d 
		JOIN user u ON d.id_user = u.id_user 
		WHERE u.email = ? AND u.role = 'dosen'`

	err := config.DB.QueryRow(query, email).Scan(&d.NIDN, &d.Nama, &d.Password)
	if err != nil {
		return false
	}

	return d.Password != "" && d.Password == password
}

func (d *Dosen) Logout() bool {
	return true
}

func (d *Dosen) MintaLaporan(periode string, idKelas int) ([]Absensi, error) {
	query := `
		SELECT a.id_absensi, a.tanggal_abs, a.status_abs, a.lokasi_abs 
		FROM absensi a
		JOIN mahasiswa m ON a.nim = m.nim
		WHERE m.id_kelas = ? AND a.tanggal_abs LIKE ?
		ORDER BY a.tanggal_abs DESC`

	rows, err := config.DB.Query(query, idKelas, periode+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var laporan []Absensi
	for rows.Next() {
		var item Absensi
		var t string
		if err := rows.Scan(&item.IdAbsensi, &t, &item.StatusAbs, &item.LokasiAbs); err != nil {
			continue
		}
		item.TanggalAbs = t
		laporan = append(laporan, item)
	}

	if len(laporan) == 0 {
		return nil, errors.New("tidak ada data absensi")
	}

	return laporan, nil
}
