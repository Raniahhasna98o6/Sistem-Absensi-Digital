package models

type MataKuliah struct {
	kode_mk string
	nama_mk string
	jadwal  string
	ruangan string
}

func (matkul *MataKuliah) GetJadwal() string {
	return matkul.jadwal
}

func (matkul *MataKuliah) UpdateJadwal(jadwalBaru string) {
	matkul.jadwal = jadwalBaru
}
