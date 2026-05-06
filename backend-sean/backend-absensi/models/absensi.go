package models

import (
	"backend-absensi/config"
	"math"
	"time"
)

type Absensi struct {
	IdAbsensi  int     `json:"id_absensi"`
	TanggalAbs string  `json:"tanggal_abs"`
	StatusAbs  string  `json:"status_abs"`
	LokasiAbs  string  `json:"lokasi_abs"`
	FotoAbs    string  `json:"foto_abs"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
}

const (
	TelyuLat  = -6.974001
	TelyuLon  = 107.630339
	MaxRadius = 500000
)

func (a *Absensi) CekRadiusLokasi() bool {
	const R = 6371000

	dLat := (a.Latitude - TelyuLat) * (math.Pi / 180)
	dLon := (a.Longitude - TelyuLon) * (math.Pi / 180)

	phi1 := TelyuLat * (math.Pi / 180)
	phi2 := a.Latitude * (math.Pi / 180)

	x := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(phi1)*math.Cos(phi2)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(x), math.Sqrt(1-x))
	jarak := R * c

	return jarak <= MaxRadius
}

func (a *Absensi) SimpanKeDatabase(nim string) (bool, string) {
	query := `INSERT INTO absensi (nim, tanggal_abs, status_abs, lokasi_abs, foto_abs, latitude, longitude) 
			  VALUES (?, NOW(), 'Hadir', ?, ?, ?, ?)`

	_, err := config.DB.Exec(query, nim, a.LokasiAbs, a.FotoAbs, a.Latitude, a.Longitude)

	if err != nil {
		return false, "data tidak dapat disimpan"
	}

	return true, "absensi status berhasil"
}

func (a *Absensi) AmbilDataAbsensi(nim string) []Absensi {
	query := "SELECT id_absensi, tanggal_abs, status_abs, lokasi_abs FROM absensi WHERE nim LIKE ? ORDER BY tanggal_abs DESC"

	rows, err := config.DB.Query(query, nim)
	if err != nil {
		return nil
	}
	defer rows.Close()

	var riwayat []Absensi
	for rows.Next() {
		var item Absensi
		var t time.Time
		// Sesuaikan urutan scan dengan SELECT (tidak ada nim yang di-scan)
		rows.Scan(&item.IdAbsensi, &t, &item.StatusAbs, &item.LokasiAbs)

		item.TanggalAbs = t.Format("2006-01-02 15:04:05")
		riwayat = append(riwayat, item)
	}
	return riwayat
}
