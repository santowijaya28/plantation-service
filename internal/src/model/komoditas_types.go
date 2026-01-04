package model

type Komoditas struct {
	ID            int    `db:"id_komoditas" json:"id,omitempty"`
	NamaKomoditas string `db:"nama_komoditas" json:"nama_komoditas,omitempty"`
	JenisTanaman  string `db:"jenis_tanaman" json:"jenis_tanaman,omitempty"`
}

type AllKomoditas struct {
	Data       []Komoditas `json:"data,omitempty"`
	TotalCount int         `json:"total_count"`
	Page       int         `json:"page"`
	PageSize   int         `json:"page_size"`
}

type FilterKomoditas struct {
	NamaKomoditas string
	JenisTanaman  string
}
