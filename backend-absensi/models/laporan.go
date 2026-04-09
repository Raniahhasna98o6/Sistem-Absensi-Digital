package models

type Laporan struct {
	id_laporan int
	periode    string
}

func (l *Laporan) GenerateLaporan() Laporan {
	return *l
}

func (l *Laporan) TampilkanLaporan() []Absensi {
	return []Absensi{}
}

func (l *Laporan) ExportPDF() string {
	return "file.pdf"
}
