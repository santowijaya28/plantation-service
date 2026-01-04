package model

type Kebun struct {
	ID         int     `db:"id_kebun" json:"id,omitempty"`
	NamaKebun  string  `db:"nama_kebun" json:"nama_kebun,omitempty"`
	LuasKebun  float64 `db:"luas_kebun" json:"luas_kebun,omitempty"`
	JenisKebun string  `db:"jenis_kebun" json:"jenis_kebun,omitempty"`
	Lat        float64 `db:"lat" json:"lat,omitempty"`
	Long       float64 `db:"long" json:"long,omitempty"`
}

const (
	JENIS_KEBUN_GREEN_HOUSE     = "green_house"
	JENIS_KEBUN_NON_GREEN_HOUSE = "non_green_house"
)

var (
	MapKebun = map[string]bool{
		JENIS_KEBUN_GREEN_HOUSE:     true,
		JENIS_KEBUN_NON_GREEN_HOUSE: true,
	}
)

type AllKebun struct {
	Data       []Kebun `json:"data,omitempty"`
	TotalCount int     `json:"total_count"`
	Page       int     `json:"page"`
	PageSize   int     `json:"page_size"`
}

type KebunFilter struct {
	JenisKebun string
	NamaKebun  string
}
