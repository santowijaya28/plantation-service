package repository

import (
	"context"

	"github.com/plantation-service/internal/src/model"
)

type StaffRepoInterface interface {
	InsertStaff(ctx context.Context, data model.StaffKebun) (result model.StaffKebun, err error)
	GetByID(ctx context.Context, idStaff int) (data model.StaffKebun, err error)
	GetAllStaff(ctx context.Context, filter model.StaffFilter, page, pageSize int) (data model.AllStaffKebun, err error)
	UpdateStaff(ctx context.Context, idStaff int, data model.StaffKebun) (result model.StaffKebun, err error)
	DeleteStaff(ctx context.Context, idStaff int) (err error)
}

type KomoditasRepoInterface interface {
	InsertKomoditas(ctx context.Context, data model.Komoditas) (result model.Komoditas, err error)
	GetByID(ctx context.Context, idKomoditas int) (data model.Komoditas, err error)
	GetAllKomoditas(ctx context.Context, filter model.FilterKomoditas, page, pageSize int) (data model.AllKomoditas, err error)
}

type BahanPerawatanRepoInterface interface {
	InsertBahanPerawatan(ctx context.Context, data model.BahanPerawatan) (result model.BahanPerawatan, err error)
	GetByID(ctx context.Context, idBahanPerawatan int) (data model.BahanPerawatan, err error)
	GetAllBahanPerawatan(ctx context.Context, filter model.FilterBahanPerawatan, page, pageSize int) (data model.AllBahanPerawatan, err error)
	UpdateBahanPerawatan(ctx context.Context, idBahanPerawatan int, data model.BahanPerawatan) (result model.BahanPerawatan, err error)
}

type BudidayaRepoInterface interface {
	InsertBudidaya(ctx context.Context, data model.Budidaya) (result model.Budidaya, err error)
	GetByID(ctx context.Context, idBudidaya int) (data model.Budidaya, err error)
	GetAllByKebun(ctx context.Context, filter model.FilterBudidaya, page, pageSize int) (data []model.Budidaya, err error)
	UpdateBudidaya(ctx context.Context, idBudidaya int, data model.Budidaya) (result model.Budidaya, err error)
}

type KebunRepoInterface interface {
	InsertKebun(ctx context.Context, data model.Kebun) (result model.Kebun, err error)
	GetByID(ctx context.Context, idKebun int) (data model.Kebun, err error)
	GetAllKebun(ctx context.Context, filter model.KebunFilter, page, pageSize int) (data model.AllKebun, err error)

	InsertBahanPerawatankebun(ctx context.Context, data model.BahanPerawatanKebun) (result model.BahanPerawatanKebun, err error)
	GetBahanPerawatanKebun(ctx context.Context, idBahanPerawatan int, idKebun int) (data model.BahanPerawatanKebun, err error)
	UpdateBahanPerawatankebun(ctx context.Context, idBahanPerawatan int, idKebun int, data model.BahanPerawatanKebun) (result model.BahanPerawatanKebun, err error)
	GetAllBahanPerawatanKebun(ctx context.Context,
		filter model.FilterBahanPerawatanKebun,
		page, pageSize int) (data model.AllBahanPerawatanKebun, err error)
}

type PerawatanTanamanRepoInterface interface {
	InsertPerawatanTanaman(ctx context.Context, data model.PerawatanTanaman) (result model.PerawatanTanaman, err error)
	GetbyId(ctx context.Context, idPerwatanTanaman int) (data model.PerawatanTanaman, err error)
	GetAllPerawatanTanaman(ctx context.Context, filter model.FilterPerawatanTanaman, page, pageSize int) (data model.AllPerawatanTanaman, err error)
	UpdatePerawatanTanaman(ctx context.Context, id int, data model.PerawatanTanaman) (result model.PerawatanTanaman, err error)
}
type HasilPanenRepoInterface interface {
	InsertHasilPanen(ctx context.Context, data model.HasilPanen) (result model.HasilPanen, err error)
	GetByID(ctx context.Context, idPanen int) (data model.HasilPanen, err error)
	GetAllHasilPanen(ctx context.Context, filter model.FilterHasilPanen, page, pageSize int) (data model.AllHasilPanen, err error)
	UpdateHasilPanen(ctx context.Context, data model.HasilPanen) (result model.HasilPanen, err error)
}
