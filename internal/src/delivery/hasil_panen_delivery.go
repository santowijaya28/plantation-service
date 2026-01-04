package delivery

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/plantation-service/internal/src/model"
	"github.com/plantation-service/pkg/log"
	"github.com/plantation-service/pkg/response"
)

func (d *delivery) InsertHasilPanen(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var dataPanen model.HasilPanenReq
	err := json.NewDecoder(r.Body).Decode(&dataPanen)
	if err != nil {
		log.Error(err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid request body",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	tanggalPanen, err := time.Parse("2006-01-02", dataPanen.TanggalPanen)
	if err != nil {
		log.Error(err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid tanggal_panen format, expected YYYY-MM-DD",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	hasilPanen := model.HasilPanen{
		IDBudidaya:   dataPanen.IDBudidaya,
		TotalKg:      dataPanen.TotalKg,
		TanggalPanen: tanggalPanen,
	}

	_, err = d.kebunUsecase.InsertHasilPanen(ctx, hasilPanen)
	if err != nil {
		log.Error("gagal memasukkan hasil panen", err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusInternalServerError,
		})
		return
	}

	response.WriteSuccessResponse(w, response.Response{
		HttpStatusCode: http.StatusCreated,
	})
}

func (d *delivery) UpdateHasilPanen(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	queryParams := r.URL.Query()
	idPanenStr := queryParams.Get("id_panen")
	if idPanenStr == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "id_panen is required",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	idPanen, _ := strconv.Atoi(idPanenStr)

	var dataPanen model.HasilPanenReq
	err := json.NewDecoder(r.Body).Decode(&dataPanen)
	if err != nil {
		log.Error(err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid request body",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	tanggalPanen, err := time.Parse("2006-01-02", dataPanen.TanggalPanen)
	if err != nil {
		log.Error(err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid tanggal_panen format, expected YYYY-MM-DD",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	hasilPanen := model.HasilPanen{
		IDPanen:      idPanen,
		IDBudidaya:   dataPanen.IDBudidaya,
		TotalKg:      dataPanen.TotalKg,
		TanggalPanen: tanggalPanen,
	}

	res, err := d.kebunUsecase.UpdateHasilPanen(ctx, hasilPanen)
	if err != nil {
		log.Error("gagal update hasil panen", err)
		code := http.StatusInternalServerError
		if err.Error() == "hasil panen tidak ditemukan" {
			code = http.StatusNotFound
		}
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: code,
		})
		return
	}

	response.WriteSuccessResponse(w, response.Response{
		Data:           res,
		HttpStatusCode: http.StatusOK,
	})
}

func (d *delivery) GetHasilPanenByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	queryParams := r.URL.Query()
	idPanenStr := queryParams.Get("id_panen")
	if idPanenStr == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "id_panen is required",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	idPanen, _ := strconv.Atoi(idPanenStr)

	data, err := d.kebunUsecase.GetHasilPanenByID(ctx, idPanen)
	if err != nil {
		log.Error("gagal gethasilpanenbyid", err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusInternalServerError,
		})
		return
	}

	response.WriteSuccessResponse(w, response.Response{
		Data: data,
	})
}

func (d *delivery) GetAllHasilPanen(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	queryParams := r.URL.Query()

	// Pagination parameters
	pageStr := queryParams.Get("page")
	pageSizeStr := queryParams.Get("page_size")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		pageSize = 10
	}

	idBudidayaStr := queryParams.Get("id_budidaya")
	startDate := queryParams.Get("start")
	endDate := queryParams.Get("end")
	idKebunStr := queryParams.Get("id_kebun")

	idBudidaya, _ := strconv.Atoi(idBudidayaStr)
	idKebun, _ := strconv.Atoi(idKebunStr)

	filter := model.FilterHasilPanen{
		IDBudidaya: idBudidaya,
		StartDate:  startDate,
		EndDate:    endDate,
		IDKebun:    idKebun,
	}

	data, err := d.kebunUsecase.GetAllHasilPanen(ctx, filter, page, pageSize)
	if err != nil {
		log.Error("error getall hasil panen", err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusInternalServerError,
		})
		return
	}

	response.WriteSuccessResponse(w, response.Response{
		Data: data,
	})
}
