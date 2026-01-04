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

func (d *delivery) InsertBahanPerawatan(w http.ResponseWriter, r *http.Request) {
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

	var dataPerawatan model.BahanPerawatan

	err = json.Unmarshal(body, &dataPerawatan)
	if err != nil {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	if dataPerawatan.NamaBahan == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "nama bahan tidak boleh kosong",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	result, err := d.bahanPerawatanusecase.InsertBahanPerawatan(ctx, dataPerawatan)
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

func (d *delivery) GetBahanPerawatanByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idBahanPerawatan := r.URL.Query().Get("id")
	if idBahanPerawatan == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid parameter",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	idBahanPerawatanInt, _ := strconv.Atoi(idBahanPerawatan)
	result, err := d.bahanPerawatanusecase.GetByID(ctx, idBahanPerawatanInt)
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

func (d *delivery) GetAllBahanPerawatan(w http.ResponseWriter, r *http.Request) {
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

	var filter model.FilterBahanPerawatan

	tipePerawatan := r.URL.Query().Get("tipe_perawatan")
	jenisBahan := r.URL.Query().Get("jenis_bahan")
	namaBahan := r.URL.Query().Get("nama_bahan")

	filter.JenisBahan = jenisBahan
	filter.TipePerawatan = tipePerawatan
	filter.NamaBahan = namaBahan

	result, err := d.bahanPerawatanusecase.GetAllBahanPerawatan(ctx, filter, pageInt, pageSizeInt)
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

func (d *delivery) UpdateBahanPerawatan(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idBahanPerawatan := r.URL.Query().Get("id")
	if idBahanPerawatan == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid parameter",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	idBahanPerawatanInt, _ := strconv.Atoi(idBahanPerawatan)
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

	var dataPerawatan model.BahanPerawatan

	err = json.Unmarshal(body, &dataPerawatan)
	if err != nil {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	result, err := d.bahanPerawatanusecase.UpdateBahanPerawatan(ctx, idBahanPerawatanInt, dataPerawatan)
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
