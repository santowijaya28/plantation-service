package model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type JenisPerawatan int

const (
	Pemupukan        JenisPerawatan = 1
	PengendalianHama JenisPerawatan = 2
	Penyiraman       JenisPerawatan = 3
	PerawatanLain    JenisPerawatan = 4
)

func (j *JenisPerawatan) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		switch strings.ToLower(s) {
		case "pemupukan":
			*j = Pemupukan
		case "pengendalian_hama":
			*j = PengendalianHama
		case "penyiraman":
			*j = Penyiraman
		case "perawatan_lain":
			*j = PerawatanLain
		default:
			return fmt.Errorf("invalid jenis_perawatan: %s", s)
		}
		return nil
	}

	var i int
	if err := json.Unmarshal(b, &i); err == nil {
		*j = JenisPerawatan(i)
		return nil
	}

	return fmt.Errorf("invalid type for jenis_perawatan")
}

func (j JenisPerawatan) String() string {
	switch j {
	case Pemupukan:
		return "pemupukan"
	case PengendalianHama:
		return "pengendalian_hama"
	case Penyiraman:
		return "penyiraman"
	case PerawatanLain:
		return "perawatan_lain"
	default:
		return "unknown"
	}
}

type PerawatanTanaman struct {
	IDRiwayatPerawatan int            `db:"id_riwayat_perawatan" json:"id_riwayat_perawatan"`
	IDStaff            *int           `db:"id_staff" json:"id_staff"` // nullable (SET NULL on delete)
	IDBudidaya         int            `db:"id_budidaya" json:"id_budidaya"`
	IDBahanPerawatan   *int           `db:"id_bahan_perawatan" json:"id_bahan_perawatan"` // nullable (SET NULL on delete)
	JumlahBahan        float64        `db:"jumlah_bahan" json:"jumlah_bahan"`
	BiayaPerawatan     float64        `db:"biaya_perawatan" json:"biaya_perawatan"`
	TanggalPerawatan   string         `db:"tanggal_perawatan" json:"tanggal_perawatan"`
	JenisPerawatan     JenisPerawatan `db:"jenis_perawatan" json:"jenis_perawatan"`
	NamaKomoditas      string         `db:"nama_komoditas" json:"nama_komoditas"`
	NamaBahan          string         `db:"nama_bahan" json:"nama_bahan"`
	NamaStaff          string         `db:"nama_staff" json:"nama_staff"`
	NamaKebun          string         `db:"nama_kebun" json:"nama_kebun"`
}

type FilterPerawatanTanaman struct {
	IDKebun    int `json:"id_kebun"`
	IDStaff    int `json:"id_staff"`
	IDBudidaya int `json:"id_budidaya"`
}

type AllPerawatanTanaman struct {
	Data       []PerawatanTanaman `json:"data"`
	TotalCount int                `json:"total_count"`
	Page       int                `json:"page"`
	PageSize   int                `json:"page_size"`
}
