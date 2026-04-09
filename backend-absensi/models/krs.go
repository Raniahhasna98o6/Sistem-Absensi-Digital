package models

type KRS struct {
	semester    int
	nilai_angka float64
	nilai_index string
}

func (k *KRS) AmbilMK(matkul MataKuliah) bool {
	return true
}

func (k *KRS) DropMK(matkul MataKuliah) bool {
	return true
}

func (k *KRS) UpdateNilai(nilaiBaru float64) bool {
	k.nilai_angka = nilaiBaru
	return true
}

func (k *KRS) GetDetailKRS() KRS {
	return *k
}
