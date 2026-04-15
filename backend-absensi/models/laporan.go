package models
import "backend-absensi/config"

type Laporan struct {
	IDLaporan int    `json:"id_laporan"`
	IDUser    string `json:"id_user"`
	Periode   string `json:"periode"`
}
func GetLaporan() ([]Laporan, error) {
	rows, err := config.DB.Query("SELECT id_laporan, id_user, periode FROM Laporan")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Laporan
	for rows.Next() {
		var l Laporan
		rows.Scan(&l.IDLaporan, &l.IDUser, &l.Periode)
		list = append(list, l)
	}

	return list, nil
}
