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
		JOIN User u ON d.id_user = u.id_user 
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
	// 1. Tambahin a.foto_abs di baris SELECT
	query := `
        SELECT a.id_absensi, a.tanggal_abs, m.nama, a.status_abs, a.lokasi_abs, a.foto_abs 
        FROM absensi a
        JOIN mahasiswa m ON a.nim = m.nim
        JOIN dosen_kelas dk ON m.id_kelas = dk.id_kelas
        WHERE dk.nidn = ? AND m.id_kelas = ? AND a.tanggal_abs LIKE ?
        ORDER BY a.tanggal_abs DESC`

	rows, err := config.DB.Query(query, d.NIDN, idKelas, periode+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var laporan []Absensi
	for rows.Next() {
		var item Absensi
		var t string

		// 2. Tambahin &item.FotoAbs di dalem Scan (Pastiin urutannya sama kayak SELECT)
		if err := rows.Scan(&item.IdAbsensi, &t, &item.NamaMhs, &item.StatusAbs, &item.LokasiAbs, &item.FotoAbs); err != nil {
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
