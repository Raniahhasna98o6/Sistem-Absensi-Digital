package models

import "time"

type Absensi struct {
	id_absensi int
	tanggal    time.Time
	status     string
	lokasi     string
	foto       string
}

// getter
func (a *Absensi) GetStatus() string {
	return a.status
}

// setter
func (a *Absensi) SetStatus(status string) {
	a.status = status
}

func (a *Absensi) SimpanAbsensi() bool {
	return true
}

func (a *Absensi) UpdateStatus(status string) {
	a.status = status
}

func (a *Absensi) ValidasiLokasi(lokasi string) bool {
	return true
}
