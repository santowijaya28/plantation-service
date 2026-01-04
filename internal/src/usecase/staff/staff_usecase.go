package staff

import (
	"context"
	"errors"
	"fmt"

	"github.com/plantation-service/internal/src/model"
)

func (u *staffUsecase) InsertStaff(ctx context.Context, data model.StaffKebun) (result model.StaffKebun, err error) {
	result, err = u.staffRepo.InsertStaff(ctx, data)
	if err != nil {
		return
	}

	return
}

func (u *staffUsecase) GetStaffByID(ctx context.Context, idStaff int) (result model.StaffKebun, err error) {
	result, err = u.staffRepo.GetByID(ctx, idStaff)
	if err != nil {
		return
	}

	return
}

func (u *staffUsecase) GetAllStaff(ctx context.Context, filter model.StaffFilter, page, pageSize int) (data model.AllStaffKebun, err error) {
	if page < 0 || pageSize <= 0 {
		return data, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "pagination"))
	}

	data, err = u.staffRepo.GetAllStaff(ctx, filter, page, pageSize)
	if err != nil {
		return data, err
	}

	return
}

func (u *staffUsecase) UpdateStaff(ctx context.Context, idStaff int, data model.StaffKebun) (result model.StaffKebun, err error) {
	if idStaff <= 0 {
		return result, errors.New(fmt.Sprintf(model.ErrMsgInvalidFields, "id_staff"))
	}

	if data.NamaStaff == "" {
		return result, errors.New(fmt.Sprintf(model.ErrMsgEmptyFields, "nama_staff"))
	}

	if data.Jabatan == "" {
		return result, errors.New(fmt.Sprintf(model.ErrMsgEmptyFields, "jabatan"))
	}

	if data.Kontak == "" {
		return result, errors.New(fmt.Sprintf(model.ErrMsgEmptyFields, "kontak"))
	}

	result, err = u.staffRepo.UpdateStaff(ctx, idStaff, data)
	if err != nil {
		return
	}

	return
}

func (u *staffUsecase) DeleteStaff(ctx context.Context, idStaff int) (err error) {
	return u.staffRepo.DeleteStaff(ctx, idStaff)
}
