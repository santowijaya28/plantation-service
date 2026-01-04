package kebun

import (
	"context"
	"errors"
	"fmt"

	"github.com/plantation-service/internal/src/model"
)

func (u *kebunUsecase) InsertKebun(ctx context.Context, data model.Kebun) (result model.Kebun, err error) {
	if data.NamaKebun == "" {
		return result, errors.New(fmt.Sprintf(model.ErrMsgEmptyFields, "nama_kebun"))
	}

	if data.LuasKebun <= 0 {
		return result, errors.New("luas_kebun tidak valid")
	}

	if _, ok := model.MapKebun[data.JenisKebun]; !ok {
		return result, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "jenis_kebun"))
	}

	result, err = u.kebunRepo.InsertKebun(ctx, data)
	if err != nil {
		return
	}

	return
}

func (u *kebunUsecase) GetKebunByID(ctx context.Context, idKebun int) (data model.Kebun, err error) {
	if idKebun <= 0 {
		return data, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "id_kebun"))
	}

	data, err = u.kebunRepo.GetByID(ctx, idKebun)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (u *kebunUsecase) GetAllKebun(ctx context.Context, filter model.KebunFilter, page, pageSize int) (data model.AllKebun, err error) {
	if page < 0 || pageSize <= 0 {
		return data, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "pagination"))
	}

	if filter.JenisKebun != "" {
		if _, ok := model.MapKebun[filter.JenisKebun]; !ok {
			return data, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "jenis_kebun"))
		}
	}

	data, err = u.kebunRepo.GetAllKebun(ctx, filter, page, pageSize)
	if err != nil {
		return data, err
	}

	return
}
