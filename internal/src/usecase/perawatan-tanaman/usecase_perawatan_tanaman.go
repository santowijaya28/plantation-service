package perawatantanaman

import (
	"context"
	"fmt"

	"github.com/plantation-service/internal/src/model"
)

func (u *perawatanTanaman) InsertPerawatanTanaman(ctx context.Context, data model.PerawatanTanaman) (result model.PerawatanTanaman, err error) {
	// Add business logic validation here if needed, in the future.

	// Validate IDStaff
	if data.IDStaff != nil {
		resStaff, err := u.staffRepo.GetByID(ctx, *data.IDStaff)
		if err != nil {
			return result, fmt.Errorf("staff with id %d not found", *data.IDStaff)
		}
		if resStaff.ID == 0 {
			return result, fmt.Errorf("staff with id %d not found", *data.IDStaff)
		}
	}

	// Validate IDBudidaya
	resBudidaya, err := u.budidayaRepo.GetByID(ctx, data.IDBudidaya)
	if err != nil {
		return result, fmt.Errorf("budidaya with id %d not found", data.IDBudidaya)
	}
	// Budidaya repo seemingly returns error, but safe to check ID
	if resBudidaya.ID == 0 {
		return result, fmt.Errorf("budidaya with id %d not found", data.IDBudidaya)
	}

	// Validate IDBahanPerawatan
	if data.IDBahanPerawatan != nil {
		resBahan, err := u.bahanPerawatanRepo.GetByID(ctx, *data.IDBahanPerawatan)
		if err != nil {
			return result, fmt.Errorf("bahan perawatan with id %d not found", *data.IDBahanPerawatan)
		}
		if resBahan.ID == 0 {
			return result, fmt.Errorf("bahan perawatan with id %d not found", *data.IDBahanPerawatan)
		}

		// Validation stock
		bahanKebun, err := u.kebunRepo.GetBahanPerawatanKebun(ctx, *data.IDBahanPerawatan, resBudidaya.IDKebun)
		if err != nil {
			return result, fmt.Errorf("bahan perawatan kebun not found for bahan id %d and kebun id %d", *data.IDBahanPerawatan, resBudidaya.IDKebun)
		}

		if bahanKebun.StokKg-data.JumlahBahan < 0 {
			return result, fmt.Errorf("stock bahan perawatan tidak cukup. stock bahan saat ini: %.2f Kg, stock bahan yang dibutuhkan: %.2f Kg", bahanKebun.StokKg, data.JumlahBahan)
		}

		// Decrement stock
		bahanKebun.StokKg -= data.JumlahBahan
		_, err = u.kebunRepo.UpdateBahanPerawatankebun(ctx, *data.IDBahanPerawatan, resBudidaya.IDKebun, bahanKebun)
		if err != nil {
			return result, fmt.Errorf("failed to update stock bahan perawatan: %v", err)
		}
	}

	return u.perawatanTanamanRepo.InsertPerawatanTanaman(ctx, data)
}

func (u *perawatanTanaman) GetByID(ctx context.Context, idPerawatanTanaman int) (data model.PerawatanTanaman, err error) {
	return u.perawatanTanamanRepo.GetbyId(ctx, idPerawatanTanaman)
}

func (u *perawatanTanaman) GetAllPerawatanTanaman(ctx context.Context, filter model.FilterPerawatanTanaman, page, pageSize int) (data model.AllPerawatanTanaman, err error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	return u.perawatanTanamanRepo.GetAllPerawatanTanaman(ctx, filter, page, pageSize)
}

func (u *perawatanTanaman) UpdatePerawatanTanaman(ctx context.Context, id int, data model.PerawatanTanaman) (result model.PerawatanTanaman, err error) {
	return u.perawatanTanamanRepo.UpdatePerawatanTanaman(ctx, id, data)
}
