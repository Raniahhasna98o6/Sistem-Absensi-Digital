package models

type Kelas struct {
	ID_Kelas        int    `json:"id_kelas"`
	Nama_Kelas      string `json:"nama_kelas"`
	Kode_Matakuliah string `json:"kode_mk"`
}
