package komoditas

import (
	"context"
	"errors"
	"fmt"

	"github.com/plantation-service/internal/src/model"
)

func (u *komoditasUsecase) InsertKomoditas(ctx context.Context, data model.Komoditas) (result model.Komoditas, err error) {
	data, err = u.komoditasRepo.InsertKomoditas(ctx, data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (u *komoditasUsecase) GetByID(ctx context.Context, idKomoditas int) (data model.Komoditas, err error) {
	if idKomoditas <= 0 {
		return data, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "id_komoditas"))
	}

	data, err = u.komoditasRepo.GetByID(ctx, idKomoditas)
	if err != nil {
		return data, err
	}

	return
}

func (u *komoditasUsecase) GetAllKomoditas(ctx context.Context, filter model.FilterKomoditas, page, pageSize int) (data model.AllKomoditas, err error) {
	if page < 0 || pageSize <= 0 {
		return data, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "pagination"))
	}

	data, err = u.komoditasRepo.GetAllKomoditas(ctx, filter, page, pageSize)
	if err != nil {
		return model.AllKomoditas{}, err
	}

	return data, nil
}
