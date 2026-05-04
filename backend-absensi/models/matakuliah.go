package models
import "backend-absensi/config"

type MataKuliah struct {
	KodeMK    string `json:"kode_mk"`
	IDUser    string `json:"id_user"`
	IDRuangan int    `json:"id_ruangan"`
	NamaMK    string `json:"nama_mk"`
	Jadwal    string `json:"jadwal"`
}

func GetAllMataKuliah() ([]MataKuliah, error) {
	rows, err := config.DB.Query("SELECT kode_mk, id_user, id_ruangan, nama_mk, jadwal FROM Mata_Kuliah")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []MataKuliah
	for rows.Next() {
		var mk MataKuliah
		rows.Scan(&mk.KodeMK, &mk.IDUser, &mk.IDRuangan, &mk.NamaMK, &mk.Jadwal)
		list = append(list, mk)
	}

	return list, nil
}
