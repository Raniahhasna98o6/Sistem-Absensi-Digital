package models

type Mahasiswa struct {
	User
	nim string
}

// getter
func (m *Mahasiswa) GetNIM() string {
	return m.nim
}

// setter
func (m *Mahasiswa) SetNIM(nim string) {
	m.nim = nim
}

func (m *Mahasiswa) LihatJadwal() []MataKuliah {
	return []MataKuliah{}
}

func (m *Mahasiswa) LakukanAbsensi(kode_mk string) Absensi {
	return Absensi{}
}

func (m *Mahasiswa) LihatRiwayatAbsensi() []Absensi {
	return []Absensi{}
}

func (m *Mahasiswa) LihatLaporan() Laporan {
	return Laporan{}
}
