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

func (d *delivery) InsertKebun(w http.ResponseWriter, r *http.Request) {

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

	var dataKebun model.Kebun

	err = json.Unmarshal(body, &dataKebun)
	if err != nil {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	if dataKebun.NamaKebun == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "nama kebun tidak boleh kosong",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	result, err := d.kebunUsecase.InsertKebun(ctx, dataKebun)
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

func (d *delivery) GetKebunByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idKebun := r.URL.Query().Get("id_kebun")
	if idKebun == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid parameter",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	idKebunInt, _ := strconv.Atoi(idKebun)
	result, err := d.kebunUsecase.GetKebunByID(ctx, idKebunInt)
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

func (d *delivery) GetAllKebun(w http.ResponseWriter, r *http.Request) {
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

	jenisKebun := r.URL.Query().Get("jenis_kebun")
	namaKebun := r.URL.Query().Get("nama_kebun")

	filter := model.KebunFilter{
		JenisKebun: jenisKebun,
		NamaKebun:  namaKebun,
	}

	result, err := d.kebunUsecase.GetAllKebun(ctx, filter, pageInt, pageSizeInt)
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

func (d *delivery) InsertBahanPerawatankebun(w http.ResponseWriter, r *http.Request) {

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

	var dataBahanPerawatanKebun model.BahanPerawatanKebun

	err = json.Unmarshal(body, &dataBahanPerawatanKebun)
	if err != nil {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	if dataBahanPerawatanKebun.IDBahan == 0 || dataBahanPerawatanKebun.IDKebun == 0 {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "id_bahan dan id_kebun tidak boleh kosong",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	result, err := d.kebunUsecase.InsertBahanPerawatankebun(ctx, dataBahanPerawatanKebun)
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

func (d *delivery) GetBahanPerawatanKebun(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idBahan := r.URL.Query().Get("id_bahan")
	idKebun := r.URL.Query().Get("id_kebun")
	if idBahan == "" || idKebun == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid parameter",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	idBahanInt, _ := strconv.Atoi(idBahan)
	idKebunInt, _ := strconv.Atoi(idKebun)

	result, err := d.kebunUsecase.GetBahanPerawatanKebun(ctx, idBahanInt, idKebunInt)
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

func (d *delivery) UpdateBahanPerawatankebun(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	idBahan := r.URL.Query().Get("id_bahan")
	idKebun := r.URL.Query().Get("id_kebun")
	if idBahan == "" || idKebun == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid parameter",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	idBahanInt, _ := strconv.Atoi(idBahan)
	idKebunInt, _ := strconv.Atoi(idKebun)

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

	var dataBahanPerawatanKebun model.BahanPerawatanKebun

	err = json.Unmarshal(body, &dataBahanPerawatanKebun)
	if err != nil {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	if dataBahanPerawatanKebun.IDBahan == 0 || dataBahanPerawatanKebun.IDKebun == 0 {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "id_bahan dan id_kebun tidak boleh kosong",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	result, err := d.kebunUsecase.UpdateBahanPerawatankebun(ctx, idBahanInt, idKebunInt, dataBahanPerawatanKebun)
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

func (d *delivery) GetAllBahanPerawatanKebun(w http.ResponseWriter, r *http.Request) {
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

	idKebun := r.URL.Query().Get("id_kebun")
	namaBahan := r.URL.Query().Get("nama_bahan")

	filter := model.FilterBahanPerawatanKebun{}
	if idKebun != "" {
		idKebunInt, _ := strconv.Atoi(idKebun)
		filter.IDKebun = idKebunInt
	}

	if namaBahan != "" {
		filter.NamaBahan = namaBahan
	}

	result, err := d.kebunUsecase.GetAllBahanPerawatanKebun(ctx, filter, pageInt, pageSizeInt)
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
