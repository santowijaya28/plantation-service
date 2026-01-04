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

func (d *delivery) InsertBudidaya(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var dataBudidaya model.BudidayaReq
	err := json.NewDecoder(r.Body).Decode(&dataBudidaya)
	if err != nil {
		log.Error(err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid request body",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	tanggalTanam, err := time.Parse("2006-01-02", dataBudidaya.TanggalTanam)
	if err != nil {
		log.Error(err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid tanggal_tanam format, expected YYYY-MM-DD",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	tanggalEstimasiPanen, err := time.Parse("2006-01-02", dataBudidaya.TanggalEstimasiPanen)
	if err != nil {
		log.Error(err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid tanggal_estimasi_panen format, expected YYYY-MM-DD",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	budidaya := model.Budidaya{
		IDKebun:              dataBudidaya.IDKebun,
		IDKomoditas:          dataBudidaya.IDKomoditas,
		TanggalTanam:         tanggalTanam,
		JumlahTanaman:        dataBudidaya.JumlahTanaman,
		TanggalEstimasiPanen: tanggalEstimasiPanen,
		StatusTanaman:        dataBudidaya.StatusTanaman,
	}

	result, err := d.kebunUsecase.InsertBudidaya(ctx, budidaya)
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

func (d *delivery) GetBudidayaByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	queryParams := r.URL.Query()
	idBudidayaStr := queryParams.Get("id_budidaya")
	if idBudidayaStr == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "id_budidaya is required",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	idBudidaya, err := strconv.Atoi(idBudidayaStr)
	if err != nil {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid id_budidaya",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	result, err := d.kebunUsecase.GetBudidayaByID(ctx, idBudidaya)
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

func (d *delivery) GetAllBudidayaByKebun(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	queryParams := r.URL.Query()
	idKebunStr := queryParams.Get("id_kebun")
	if idKebunStr == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "id_kebun is required",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	idKebun, err := strconv.Atoi(idKebunStr)
	if err != nil {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid id_kebun",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

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

	statusTanaman := queryParams.Get("status_tanaman")
	idKomoditasStr := queryParams.Get("id_komoditas")

	idKomoditas, _ := strconv.Atoi(idKomoditasStr)
	tanggalTanam := queryParams.Get("tanggal_tanam")
	estTanggalPanen := queryParams.Get("est_tanggal_panen")
	namaKomoditas := queryParams.Get("nama_komoditas")

	filter := model.FilterBudidaya{
		IDKebun:              idKebun,
		StatusTanaman:        statusTanaman,
		IDKomoditas:          idKomoditas,
		TanggalEstimasiPanen: estTanggalPanen,
		TanggalTanam:         tanggalTanam,
		NamaKomoditas:        namaKomoditas,
	}

	result, err := d.kebunUsecase.GetAllBudidayaByKebun(ctx, idKebun, filter, page, pageSize)
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

func (d *delivery) UpdateBudidaya(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	queryParams := r.URL.Query()
	idBudidayaStr := queryParams.Get("id_budidaya")
	if idBudidayaStr == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "id_budidaya is required",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	idBudidaya, err := strconv.Atoi(idBudidayaStr)
	if err != nil {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid id_budidaya",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	var dataBudidaya model.BudidayaReq
	err = json.NewDecoder(r.Body).Decode(&dataBudidaya)
	if err != nil {
		log.Error(err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid request body",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	tanggalTanam, err := time.Parse("2006-01-02", dataBudidaya.TanggalTanam)
	if err != nil {
		log.Error(err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		})
	}

	estTanggalPanen, err := time.Parse("2006-01-02", dataBudidaya.TanggalEstimasiPanen)
	if err != nil {
		log.Error(err)
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		})
	}

	budidaya := model.Budidaya{
		ID:                   dataBudidaya.ID,
		IDKebun:              dataBudidaya.IDKebun,
		IDKomoditas:          dataBudidaya.IDKomoditas,
		TanggalTanam:         tanggalTanam,
		TanggalEstimasiPanen: estTanggalPanen,
		JumlahTanaman:        dataBudidaya.JumlahTanaman,
		StatusTanaman:        dataBudidaya.StatusTanaman,
	}

	result, err := d.kebunUsecase.UpdateBudidaya(ctx, idBudidaya, budidaya)
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
		HttpStatusCode: http.StatusOK,
	})
}
