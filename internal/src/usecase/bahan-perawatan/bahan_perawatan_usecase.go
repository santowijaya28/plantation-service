package bahanperawatan

import (
	"context"
	"errors"
	"fmt"

	"github.com/plantation-service/internal/src/model"
)

func (u *bpUsecase) InsertBahanPerawatan(ctx context.Context, data model.BahanPerawatan) (result model.BahanPerawatan, err error) {
	if data.NamaBahan == "" {
		return result, errors.New(fmt.Sprintf(model.ErrMsgEmptyFields, "nama_bahan"))
	}

	if data.JenisBahan == "" {
		return result, errors.New(fmt.Sprintf(model.ErrMsgEmptyFields, "jenis_bahan"))
	}

	if data.TipePerawatan == "" {
		return result, errors.New(fmt.Sprintf(model.ErrMsgEmptyFields, "tipe_perawatan"))
	}

	if data.HargaKg <= 0 {
		return result, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "harga_kg"))
	}

	// validasi jenis Pemupukan
	if _, ok := model.MapJenisBahanPerawatan[data.JenisBahan]; !ok {
		return result, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "jenis_bahan"))
	}

	// validasi tipe perawatan
	if _, ok := model.MapTipePerawatan[data.TipePerawatan]; !ok {
		return result, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "tipe_perawatan"))
	}

	_, err = u.bpRepo.InsertBahanPerawatan(ctx, data)
	if err != nil {
		return
	}

	return data, nil
}

func (u *bpUsecase) GetByID(ctx context.Context, idBahanPerawatan int) (data model.BahanPerawatan, err error) {
	if idBahanPerawatan <= 0 {
		return data, errors.New(fmt.Sprintf(model.ErrMsgEmptyFields, "id_bahan_perawatan"))
	}

	data, err = u.bpRepo.GetByID(ctx, idBahanPerawatan)
	if err != nil {
		return
	}

	return data, nil
}

func (u *bpUsecase) GetAllBahanPerawatan(ctx context.Context, filter model.FilterBahanPerawatan, page, pageSize int) (data model.AllBahanPerawatan, err error) {
	if page < 0 || pageSize <= 0 {
		return data, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "pagination"))
	}

	if filter.JenisBahan != "" {
		if _, ok := model.MapJenisBahanPerawatan[filter.JenisBahan]; !ok {
			return data, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "jenis_bahan"))
		}
	}

	if filter.TipePerawatan != "" {
		if _, ok := model.MapTipePerawatan[filter.TipePerawatan]; !ok {
			return data, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "tipe_perawatan"))
		}
	}

	data, err = u.bpRepo.GetAllBahanPerawatan(ctx, filter, page, pageSize)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (u *bpUsecase) UpdateBahanPerawatan(ctx context.Context, idBahanPerawatan int, data model.BahanPerawatan) (result model.BahanPerawatan, err error) {
	if idBahanPerawatan <= 0 {
		return result, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "id_bahan_perawatan"))
	}

	if data.NamaBahan == "" {
		return result, errors.New(fmt.Sprintf(model.ErrMsgEmptyFields, "nama_bahan"))
	}

	if data.JenisBahan == "" {
		return result, errors.New(fmt.Sprintf(model.ErrMsgEmptyFields, "jenis_bahan"))
	}

	if data.TipePerawatan == "" {
		return result, errors.New(fmt.Sprintf(model.ErrMsgEmptyFields, "tipe_perawatan"))
	}

	if data.HargaKg <= 0 {
		return result, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "harga_kg"))
	}

	// validasi jenis Pemupukan
	if _, ok := model.MapJenisBahanPerawatan[data.JenisBahan]; !ok {
		return result, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "jenis_bahan"))
	}

	// validasi tipe perawatan
	if _, ok := model.MapTipePerawatan[data.TipePerawatan]; !ok {
		return result, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "tipe_perawatan"))
	}

	result, err = u.bpRepo.UpdateBahanPerawatan(ctx, idBahanPerawatan, data)
	if err != nil {
		return
	}

	return
}
