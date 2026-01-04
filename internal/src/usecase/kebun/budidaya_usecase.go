package kebun

import (
	"context"
	"errors"

	"github.com/plantation-service/internal/src/model"
)

func (u *kebunUsecase) InsertBudidaya(ctx context.Context, data model.Budidaya) (result model.Budidaya, err error) {
	//validasi input
	err = u.validateInsertBudidaya(ctx, data)
	if err != nil {
		return
	}

	// get active budidaya in kebun
	// to prevent overlapping budidaya
	activeBudidaya, err := u.budidayaRepo.GetAllByKebun(ctx, model.FilterBudidaya{
		StatusTanaman: model.BudidayaStatusAktif,
		IDKebun:       data.IDKebun,
	}, 1, 100)
	if err != nil {
		return model.Budidaya{}, errors.New("error GetAllByKebun: " + err.Error())
	}

	for _, v := range activeBudidaya {
		if (data.TanggalTanam.Before(v.TanggalEstimasiPanen) || data.TanggalTanam.Equal(v.TanggalEstimasiPanen)) &&
			(data.TanggalEstimasiPanen.After(v.TanggalTanam) || data.TanggalEstimasiPanen.Equal(v.TanggalTanam)) {
			return model.Budidaya{}, errors.New("terdapat budidaya aktif yang tumpang tindih di kebun ini. Budidaya saat aktif saat ini dari " +
				v.TanggalTanam.Format("2006-01-02") + " sampai " + v.TanggalEstimasiPanen.Format("2006-01-02"))
		}
	}

	// insert budidaya
	result, err = u.budidayaRepo.InsertBudidaya(ctx, data)
	if err != nil {
		return
	}

	return
}

func (u *kebunUsecase) GetBudidayaByID(ctx context.Context, idBudidaya int) (data model.Budidaya, err error) {
	if idBudidaya <= 0 {
		return data, errors.New("id_budidaya tidak valid")
	}

	data, err = u.budidayaRepo.GetByID(ctx, idBudidaya)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (u *kebunUsecase) UpdateBudidaya(ctx context.Context, idBudidaya int, data model.Budidaya) (result model.Budidaya, err error) {
	if idBudidaya <= 0 {
		return result, errors.New("id_budidaya tidak valid")
	}

	//validasi input
	err = u.validateInsertBudidaya(ctx, data)
	if err != nil {
		return
	}

	activeBudidaya, err := u.budidayaRepo.GetAllByKebun(ctx, model.FilterBudidaya{
		StatusTanaman: model.BudidayaStatusAktif,
		IDKebun:       data.IDKebun,
	}, 1, 1000)
	if err != nil {
		return model.Budidaya{}, errors.New("error GetAllByKebun: " + err.Error())
	}

	for _, v := range activeBudidaya {
		if v.ID == idBudidaya {
			continue
		}

		if (data.TanggalTanam.Before(v.TanggalEstimasiPanen) || data.TanggalTanam.Equal(v.TanggalEstimasiPanen)) &&
			(data.TanggalEstimasiPanen.After(v.TanggalTanam) || data.TanggalEstimasiPanen.Equal(v.TanggalTanam)) {
			return model.Budidaya{}, errors.New("terdapat budidaya aktif yang tumpang tindih di kebun ini. Budidaya aktif saat ini dari " +
				v.TanggalTanam.Format("2006-01-02") + " sampai " + v.TanggalEstimasiPanen.Format("2006-01-02"))
		}
	}

	// update budidaya
	result, err = u.budidayaRepo.UpdateBudidaya(ctx, idBudidaya, data)
	if err != nil {
		return
	}

	return
}

func (u *kebunUsecase) GetAllBudidayaByKebun(ctx context.Context, idKebun int, filter model.FilterBudidaya, page, pageSize int) (data []model.Budidaya, err error) {
	if idKebun <= 0 {
		return data, errors.New("id_kebun tidak valid")
	}

	if page < 0 || pageSize <= 0 {
		return data, errors.New("pagination tidak valid")
	}

	data, err = u.budidayaRepo.GetAllByKebun(ctx, filter, page, pageSize)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (u *kebunUsecase) validateInsertBudidaya(ctx context.Context, data model.Budidaya) error {
	// validasi kebun
	if data.IDKebun <= 0 {
		return errors.New("id_kebun tidak valid")
	}

	dbKebun, err := u.kebunRepo.GetByID(ctx, data.IDKebun)
	if err != nil {
		return errors.New("error get kebun by id : %v" + err.Error())
	}

	if dbKebun.ID <= 0 {
		return errors.New("kebun tidak ditemukan")
	}

	// validasi komoditas
	if data.IDKomoditas <= 0 {
		return errors.New("id_komoditas tidak valid")
	}

	dbKomoditas, err := u.komoditasRepo.GetByID(ctx, data.IDKomoditas)
	if err != nil {
		return errors.New("error get komoditas by id " + err.Error())
	}

	if dbKomoditas.ID <= 0 {
		return errors.New("komoditas tidak ditemukan")
	}

	return nil
}
