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

func (d *delivery) InsertPerawatanTanaman(w http.ResponseWriter, r *http.Request) {
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

	var data model.PerawatanTanaman
	err = json.Unmarshal(body, &data)
	if err != nil {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	log.Info("datanya ", string(body))

	result, err := d.perawatanTanamanUsecase.InsertPerawatanTanaman(ctx, data)
	if err != nil {
		log.Error("error insert perawatan tanaman. Err %v, Request %#v", err, data)
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

func (d *delivery) GetPerawatanTanamanByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := r.URL.Query().Get("id_perawatan")

	if id == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid parameter",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	idInt, _ := strconv.Atoi(id)
	result, err := d.perawatanTanamanUsecase.GetByID(ctx, idInt)
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

func (d *delivery) GetAllPerawatanTanaman(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pagesize")

	page := 1
	pageSize := 10

	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}

	if pageSizeStr != "" {
		pageSize, _ = strconv.Atoi(pageSizeStr)
	}

	var filter model.FilterPerawatanTanaman
	idKebun := r.URL.Query().Get("id_kebun")
	idStaff := r.URL.Query().Get("id_staff")
	idBudidaya := r.URL.Query().Get("id_budidaya")

	if idKebun != "" {
		filter.IDKebun, _ = strconv.Atoi(idKebun)
	}
	if idStaff != "" {
		filter.IDStaff, _ = strconv.Atoi(idStaff)
	}
	if idBudidaya != "" {
		filter.IDBudidaya, _ = strconv.Atoi(idBudidaya)
	}

	result, err := d.perawatanTanamanUsecase.GetAllPerawatanTanaman(ctx, filter, page, pageSize)
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

func (d *delivery) UpdatePerawatanTanaman(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	idStr := r.URL.Query().Get("id_perawatan")

	if idStr == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid parameter id_perawatan",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	id, _ := strconv.Atoi(idStr)

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

	var data model.PerawatanTanaman
	err = json.Unmarshal(body, &data)
	if err != nil {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	result, err := d.perawatanTanamanUsecase.UpdatePerawatanTanaman(ctx, id, data)
	if err != nil {
		log.Error("error update perawatan tanaman. Err %v, Request %#v", err, data)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusInternalServerError,
		})
		return
	}

	response.WriteSuccessResponse(w, response.Response{
		Data:           result,
		HttpStatusCode: http.StatusOK,
	})
}
