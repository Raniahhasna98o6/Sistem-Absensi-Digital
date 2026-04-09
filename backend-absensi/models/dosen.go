package models

type Dosen struct {
	User
	nidn string
}

func (d *Dosen) TambahMataKuliah(matkul MataKuliah) {}

func (d *Dosen) UbahMataKuliah(matkul MataKuliah) {}

func (d *Dosen) HapusMataKuliah(matkul MataKuliah) {}

func (d *Dosen) LihatAbsensiMahasiswa(matkul MataKuliah) []Absensi {
	return []Absensi{}
}

func (d *Dosen) GenerateLaporan(periode string) Laporan {
	return Laporan{}
}
