package models
import "backend-absensi/config"

type KRS struct {
	IDUser   string `json:"id_user"`
	KodeMK   string `json:"kode_mk"`
	Semester int    `json:"semester"`
}

func GetKRSByUser(id string) ([]KRS, error) {
	rows, err := config.DB.Query("SELECT id_user, kode_mk, semester FROM KRS WHERE id_user = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []KRS
	for rows.Next() {
		var k KRS
		rows.Scan(&k.IDUser, &k.KodeMK, &k.Semester)
		list = append(list, k)
	}
	return list, nil
}
