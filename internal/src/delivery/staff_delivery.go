package delivery

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"github.com/plantation-service/internal/src/model"
	"github.com/plantation-service/pkg/log"
	"github.com/plantation-service/pkg/response"
)

func (d *delivery) InsertStaff(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("error when read request body : ", err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	var dataStaff model.StaffKebun

	err = json.Unmarshal(body, &dataStaff)
	if err != nil {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	if dataStaff.NamaStaff == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "nama tidak boleh kosong",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	result, err := d.staffUsecase.InsertStaff(ctx, dataStaff)
	if err != nil {
		log.Error(err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusInternalServerError,
		})
		return
	}

	response.WriteSuccessResponse(w, response.Response{
		Data:           result,
		HttpStatusCode: http.StatusCreated,
	})

}

func (d *delivery) GetStaffByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStaff := r.URL.Query().Get("id")
	if idStaff == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid parameter",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	idStaffInt, _ := strconv.Atoi(idStaff)
	result, err := d.staffUsecase.GetStaffByID(ctx, idStaffInt)
	if err != nil {
		log.Error(err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusInternalServerError,
		})
		return
	}

	response.WriteSuccessResponse(w, response.Response{
		Data: result,
	})
}

func (d *delivery) GetAllStaff(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	page := r.URL.Query().Get("page")
	if page == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid parameter",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	pageSize := r.URL.Query().Get("pagesize")
	if page == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid parameter",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	namaStaff := r.URL.Query().Get("nama_staff")
	jabatan := r.URL.Query().Get("jabatan")

	filter := model.StaffFilter{
		NamaStaff: namaStaff,
		Jabatan:   jabatan,
	}

	result, err := d.staffUsecase.GetAllStaff(ctx, filter, pageInt, pageSizeInt)
	if err != nil {
		log.Error(err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusInternalServerError,
		})
		return
	}

	response.WriteSuccessResponse(w, response.Response{
		Data: result,
	})
}

func (d *delivery) UpdateStaff(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStaff := r.URL.Query().Get("id")
	if idStaff == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid parameter",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	idStaffInt, _ := strconv.Atoi(idStaff)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Error("error when read request body : ", err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	var dataStaff model.StaffKebun

	err = json.Unmarshal(body, &dataStaff)
	if err != nil {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	result, err := d.staffUsecase.UpdateStaff(ctx, idStaffInt, dataStaff)
	if err != nil {
		log.Error(err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusInternalServerError,
		})
		return
	}

	response.WriteSuccessResponse(w, response.Response{
		Data: result,
	})
}

func (d *delivery) DeleteStaff(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idStaff := r.URL.Query().Get("id")
	if idStaff == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid parameter",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	idStaffInt, _ := strconv.Atoi(idStaff)

	err := d.staffUsecase.DeleteStaff(ctx, idStaffInt)
	if err != nil {
		log.Error(err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusInternalServerError,
		})
		return
	}

	response.WriteSuccessResponse(w, response.Response{
		Data: "success delete staff",
	})
}
