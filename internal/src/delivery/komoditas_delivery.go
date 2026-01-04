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

func (d *delivery) InsertKomoditas(w http.ResponseWriter, r *http.Request) {

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

	var dataKomoditas model.Komoditas

	err = json.Unmarshal(body, &dataKomoditas)
	if err != nil {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          err.Error(),
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	if dataKomoditas.NamaKomoditas == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "nama komoditas tidak boleh kosong",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	if dataKomoditas.JenisTanaman == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "jenis tanaman tidak boleh kosong",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	result, err := d.komoditasUsecase.InsertKomoditas(ctx, dataKomoditas)
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

func (d *delivery) GetKomoditasByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	idKomoditas := r.URL.Query().Get("id_komoditas")
	if idKomoditas == "" {
		response.WriteErrorResponse(w, response.Response{
			Data:           nil,
			Error:          "invalid id_komoditas parameter",
			HttpStatusCode: http.StatusBadRequest,
		})
		return
	}

	idKomoditasInt, _ := strconv.Atoi(idKomoditas)
	result, err := d.komoditasUsecase.GetByID(ctx, idKomoditasInt)
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

func (d *delivery) GetAllKomoditas(w http.ResponseWriter, r *http.Request) {
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

	jenisTanaman := r.URL.Query().Get("jenis_tanaman")
	namaKomoditas := r.URL.Query().Get("nama_komoditas")

	filter := model.FilterKomoditas{
		JenisTanaman:  jenisTanaman,
		NamaKomoditas: namaKomoditas,
	}

	result, err := d.komoditasUsecase.GetAllKomoditas(ctx, filter, pageInt, pageSizeInt)
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
