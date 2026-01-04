package model

import "time"

type BudidayaReq struct {
	ID                   int    `db:"id_budidaya" json:"id,omitempty"`
	IDKebun              int    `db:"id_kebun" json:"id_kebun,omitempty"`
	IDKomoditas          int    `db:"id_komoditas" json:"id_komoditas,omitempty"`
	TanggalTanam         string `db:"tanggal_tanam" json:"tanggal_tanam,omitempty"`
	JumlahTanaman        int    `db:"jumlah_tanaman" json:"jumlah_tanaman,omitempty"`
	TanggalEstimasiPanen string `db:"tanggal_estimasi_panen" json:"tanggal_estimasi_panen,omitempty"`
	StatusTanaman        string `db:"status_tanaman" json:"status_tanaman,omitempty"`
}

type Budidaya struct {
	ID                   int       `db:"id_budidaya" json:"id,omitempty"`
	IDKebun              int       `db:"id_kebun" json:"id_kebun,omitempty"`
	IDKomoditas          int       `db:"id_komoditas" json:"id_komoditas,omitempty"`
	NamaKomoditas        string    `db:"nama_komoditas" json:"nama_komoditas,omitempty"`
	TanggalTanam         time.Time `db:"tanggal_tanam" json:"tanggal_tanam,omitempty"`
	JumlahTanaman        int       `db:"jumlah_tanaman" json:"jumlah_tanaman,omitempty"`
	TanggalEstimasiPanen time.Time `db:"tanggal_estimasi_panen" json:"tanggal_estimasi_panen,omitempty"`
	StatusTanaman        string    `db:"status_tanaman" json:"status_tanaman,omitempty"`
}

type AllBudidaya struct {
	Data       []Budidaya `json:"data,omitempty"`
	TotalCount int        `json:"total_count"`
	Page       int        `json:"page"`
	PageSize   int        `json:"page_size"`
}

type FilterBudidaya struct {
	IDKebun              int    `json:"id_kebun,omitempty"`
	IDKomoditas          int    `json:"id_komoditas,omitempty"`
	StatusTanaman        string `json:"status_tanaman,omitempty"`
	TanggalEstimasiPanen string `json:"tanggal_estimasi_panen,omitempty"`
	TanggalTanam         string `json:"tanggal_tanam,omitempty"`
	NamaKomoditas        string `json:"nama_komoditas,omitempty"`
}

const (
	BudidayaStatusAktif = "aktif"
	BudidayaStatusPanen = "panen"
	BudidataStatusGagal = "gagal"
)
