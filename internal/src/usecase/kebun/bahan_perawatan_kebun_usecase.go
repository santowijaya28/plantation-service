package kebun

import (
	"context"
	"errors"
	"fmt"

	"github.com/plantation-service/internal/src/model"
)

func (u *kebunUsecase) InsertBahanPerawatankebun(ctx context.Context, data model.BahanPerawatanKebun) (result model.BahanPerawatanKebun, err error) {
	//validasi input
	err = u.validasteBahanPerawatanKebun(ctx, data)
	if err != nil {
		return
	}

	// insert bahan perawatan kebun
	result, err = u.kebunRepo.InsertBahanPerawatankebun(ctx, data)
	if err != nil {
		return
	}

	return
}

func (u *kebunUsecase) GetBahanPerawatanKebun(ctx context.Context, idBahanPerawatan int, idKebun int) (data model.BahanPerawatanKebun, err error) {
	data, err = u.kebunRepo.GetBahanPerawatanKebun(ctx, idBahanPerawatan, idKebun)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (u *kebunUsecase) UpdateBahanPerawatankebun(ctx context.Context, idBahanPerawatan int, idKebun int, data model.BahanPerawatanKebun) (result model.BahanPerawatanKebun, err error) {
	//validasi input
	err = u.validasteBahanPerawatanKebun(ctx, data)
	if err != nil {
		return
	}

	// update bahan perawatan kebun
	fmt.Println("data", data)
	result, err = u.kebunRepo.UpdateBahanPerawatankebun(ctx, idBahanPerawatan, idKebun, data)
	if err != nil {
		return
	}

	return
}

func (u *kebunUsecase) validasteBahanPerawatanKebun(ctx context.Context, data model.BahanPerawatanKebun) error {
	// validasi kebun
	if data.IDKebun <= 0 {
		return errors.New("id_kebun tidak valid")
	}

	dbKebun, err := u.kebunRepo.GetByID(ctx, data.IDKebun)
	if err != nil {
		return err
	}

	if dbKebun.ID == 0 {
		return errors.New("kebun tidak ditemukan")
	}

	// validasi bahan perawatan
	if data.IDBahan <= 0 {
		return errors.New("id_bahan tidak valid")
	}

	dbBahanPerawatan, err := u.bahanPerawatanRepo.GetByID(ctx, data.IDBahan)
	if err != nil {
		return err
	}

	if dbBahanPerawatan.ID == 0 {
		return errors.New("bahan perawatan tidak ditemukan")
	}

	return nil
}

func (u *kebunUsecase) GetAllBahanPerawatanKebun(ctx context.Context, filter model.FilterBahanPerawatanKebun, page, pageSize int) (data model.AllBahanPerawatanKebun, err error) {
	if page < 0 || pageSize <= 0 {
		return data, errors.New("pagination tidak valid")
	}

	// get data kebun

	kebun, err := u.kebunRepo.GetByID(ctx, filter.IDKebun)
	if err != nil {
		return data, err
	}

	if kebun.ID == 0 {
		return data, errors.New("kebun tidak ditemukan")
	}

	data, err = u.kebunRepo.GetAllBahanPerawatanKebun(ctx, filter, page, pageSize)
	if err != nil {
		return data, err
	}

	data.NamaKebun = kebun.NamaKebun
	return
}
