package kebun

import (
	"context"
	"database/sql"
	"errors"

	"github.com/plantation-service/internal/src/model"
)

func (u *kebunUsecase) InsertHasilPanen(ctx context.Context, data model.HasilPanen) (res model.HasilPanen, err error) {
	err = u.validateInsertHasilPanen(ctx, data)
	if err != nil {
		return
	}

	data, err = u.hasilPanenRepo.InsertHasilPanen(ctx, data)
	if err != nil {
		return
	}

	return
}

func (u *kebunUsecase) UpdateHasilPanen(ctx context.Context, data model.HasilPanen) (res model.HasilPanen, err error) {
	err = u.validateUpdateHasilPanen(ctx, data)
	if err != nil {
		return
	}

	res, err = u.hasilPanenRepo.UpdateHasilPanen(ctx, data)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, errors.New("hasil panen tidak ditemukan")
		}
		return
	}

	return
}

func (u *kebunUsecase) GetHasilPanenByID(ctx context.Context, idPanen int) (res model.HasilPanen, err error) {
	if idPanen <= 0 {
		return res, errors.New("id_panen tidak boleh kosong")
	}

	res, err = u.hasilPanenRepo.GetByID(ctx, idPanen)
	if err != nil {
		return
	}

	return
}

func (u *kebunUsecase) GetAllHasilPanen(ctx context.Context, filter model.FilterHasilPanen, page, pagesize int) (data model.AllHasilPanen, err error) {
	data, err = u.hasilPanenRepo.GetAllHasilPanen(ctx, filter, page, pagesize)
	if err != nil {
		return
	}

	return
}

func (u *kebunUsecase) validateInsertHasilPanen(ctx context.Context, data model.HasilPanen) (err error) {
	if data.TanggalPanen.IsZero() {
		return errors.New("tanggal_panen tidak boleh kosong")
	}

	if data.TotalKg <= 0 {
		return errors.New("total_kg tidak boleh kosong")
	}

	if data.IDBudidaya <= 0 {
		return errors.New("id_budidaya tidak boleh kosong")
	}

	budidaya, err := u.GetBudidayaByID(ctx, data.IDBudidaya)
	if err != nil && err == sql.ErrNoRows || budidaya.ID == 0 {
		return errors.New("id_budidaya tidak ditemukan")
	}

	return nil
}

func (u *kebunUsecase) validateUpdateHasilPanen(ctx context.Context, data model.HasilPanen) (err error) {
	if data.IDPanen == 0 {
		return errors.New("id_panen tidak boleh kosong")
	}

	return u.validateInsertHasilPanen(ctx, data)
}
