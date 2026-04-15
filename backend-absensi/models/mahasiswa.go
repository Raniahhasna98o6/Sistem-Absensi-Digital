package models
import "backend-absensi/config"

type Mahasiswa struct {
	IDUser string `json:"id_user"`
	NIM    string `json:"nim"`
}
func GetMahasiswaByID(id string) (*Mahasiswa, error) {
	var m Mahasiswa
	query := "SELECT id_user, nim FROM Mahasiswa WHERE id_user = ?"

	err := config.DB.QueryRow(query, id).Scan(&m.IDUser, &m.NIM)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
