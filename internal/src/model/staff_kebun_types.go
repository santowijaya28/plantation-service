package model

type StaffKebun struct {
	ID        int    `db:"id_staff" json:"id,omitempty"`
	NamaStaff string `db:"nama_staff" json:"nama_staff,omitempty"`
	Jabatan   string `db:"jabatan" json:"jabatan,omitempty"`
	Kontak    string `db:"kontak" json:"kontak,omitempty"`
}

type AllStaffKebun struct {
	Data       []StaffKebun `json:"data,omitempty"`
	TotalCount int          `json:"total_count"`
	Page       int          `json:"page"`
	PageSize   int          `json:"page_size"`
}

type StaffFilter struct {
	NamaStaff string
	Jabatan   string
}
