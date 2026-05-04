package models
import "backend-absensi/config"

type Dosen struct {
	IDUser string `json:"id_user"`
	NIDN   string `json:"nidn"`
}

func GetDosenByID(id string) (*Dosen, error) {
	var d Dosen
	query := "SELECT id_user, nidn FROM Dosen WHERE id_user = ?"

	err := config.DB.QueryRow(query, id).Scan(&d.IDUser, &d.NIDN)
	if err != nil {
		return nil, err
	}
	return &d, nil
}
