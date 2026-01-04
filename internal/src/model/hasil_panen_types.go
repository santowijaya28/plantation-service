package model

import "time"

type HasilPanen struct {
	IDPanen       int       `db:"id_panen" json:"id_panen,omitempty"`
	IDBudidaya    int       `db:"id_budidaya" json:"id_budidaya,omitempty"`
	NamaKomoditas string    `db:"nama_komoditas" json:"nama_komoditas,omitempty"`
	TotalKg       float64   `db:"total_kg" json:"total_kg,omitempty"`
	TanggalPanen  time.Time `db:"tanggal_panen" json:"tanggal_panen,omitempty"`
}

type HasilPanenReq struct {
	IDPanen      int     `db:"id_panen" json:"id_panen,omitempty"`
	IDBudidaya   int     `db:"id_budidaya" json:"id_budidaya,omitempty"`
	TotalKg      float64 `db:"total_kg" json:"total_kg,omitempty"`
	TanggalPanen string  `db:"tanggal_panen" json:"tanggal_panen,omitempty"`
}

type FilterHasilPanen struct {
	IDKebun    int
	IDBudidaya int
	StartDate  string
	EndDate    string
}

type AllHasilPanen struct {
	Data       []HasilPanen `json:"data,omitempty"`
	TotalCount int          `json:"total_count"`
	Page       int          `json:"page"`
	PageSize   int          `json:"page_size"`
}
