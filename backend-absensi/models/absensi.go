package models
import "backend-absensi/config"

type Absensi struct {
	IDAbsensi int    `json:"id_absensi"`
	IDUser    string `json:"id_user"`
	KodeMK    string `json:"kode_mk"`
	Tanggal   string `json:"tanggal"`
	Status    string `json:"status"`
	Lokasi    string `json:"lokasi"`
}

// insert absensi
func InsertAbsensi(idUser, kodeMK, lokasi string) error {
	query := `
	INSERT INTO Absensi (id_user, kode_mk, tanggal_abs, status_abs, lokasi_abs)
	VALUES (?, ?, CURDATE(), 'Hadir', ?)
	`
	_, err := config.DB.Exec(query, idUser, kodeMK, lokasi)
	return err
}
// get absensi user
func GetAbsensiByUser(id string) ([]Absensi, error) {
	rows, err := config.DB.Query(`
		SELECT id_absensi, id_user, kode_mk, tanggal_abs, status_abs, lokasi_abs
		FROM Absensi WHERE id_user = ?
	`, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []Absensi
	for rows.Next() {
		var a Absensi
		rows.Scan(&a.IDAbsensi, &a.IDUser, &a.KodeMK, &a.Tanggal, &a.Status, &a.Lokasi)
		list = append(list, a)
	}
	return list, nil
}
