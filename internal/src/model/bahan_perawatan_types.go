package model

type BahanPerawatan struct {
	ID            int     `db:"id_bahan" json:"id,omitempty"`
	NamaBahan     string  `db:"nama_bahan" json:"nama_bahan,omitempty"`
	JenisBahan    string  `db:"jenis_bahan" json:"jenis_bahan,omitempty"`
	TipePerawatan string  `db:"tipe_perawatan" json:"tipe_perawatan,omitempty"`
	HargaKg       float64 `db:"harga_kg" json:"harga_kg,omitempty"`
}

type AllBahanPerawatan struct {
	Data       []BahanPerawatan `json:"data,omitempty"`
	TotalCount int              `json:"total_count"`
	Page       int              `json:"page"`
	PageSize   int              `json:"page_size"`
}

type FilterBahanPerawatan struct {
	JenisBahan    string `json:"jenis_bahan,omitempty"`
	TipePerawatan string `json:"tipe_perawatan,omitempty"`
	NamaBahan     string `json:"nama_bahan,omitempty"`
}

var (
	JenisPerawatanPemupukan    = "pemupukan"
	JenisPerawatanPenanaman    = "penanaman"
	JenisPerawatanPenyemprotan = "penyemprotan"
)

var (
	TipePerawatanBenih     = "benih"
	TipePerawatanPestisida = "pestisida"
	TipePerawatanPupuk     = "pupuk"
)

var (
	MapJenisBahanPerawatan = map[string]bool{
		JenisPerawatanPemupukan:    true,
		JenisPerawatanPenanaman:    true,
		JenisPerawatanPenyemprotan: true,
	}

	MapTipePerawatan = map[string]bool{
		TipePerawatanBenih:     true,
		TipePerawatanPestisida: true,
		TipePerawatanPupuk:     true,
	}
)
