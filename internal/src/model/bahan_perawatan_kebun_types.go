package model

type BahanPerawatanKebun struct {
	IDBahan int     `db:"id_bahan" json:"id_bahan,omitempty"`
	IDKebun int     `db:"id_kebun" json:"id_kebun,omitempty"`
	StokKg  float64 `db:"stok_kg" json:"stok_kg,omitempty"`
}

type FilterBahanPerawatanKebun struct {
	IDKebun   int
	NamaBahan string
}

type AllBahanPerawatanKebun struct {
	NamaKebun  string                   `json:"nama_kebun"`
	Data       []BahanPerawatanWithMeta `json:"data,omitempty"`
	TotalCount int                      `json:"total_count"`
	Page       int                      `json:"page"`
	PageSize   int                      `json:"page_size"`
}

type BahanPerawatanWithMeta struct {
	IDBahan   int     `db:"id_bahan" json:"id_bahan,omitempty"`
	IDKebun   int     `db:"id_kebun" json:"id_kebun,omitempty"`
	StokKg    float64 `db:"stok_kg" json:"stok_kg,omitempty"`
	NamaBahan string  `db:"nama_bahan" json:"nama_bahan,omitempty"`
}
