package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/plantation-service/internal/core/postgres"
	"github.com/plantation-service/internal/src/config"
	"github.com/plantation-service/internal/src/delivery"
	bahanPerawatanRepository "github.com/plantation-service/internal/src/repository/bahan-perawatan"
	budidayaRepository "github.com/plantation-service/internal/src/repository/budidaya"
	hasilPanenRepository "github.com/plantation-service/internal/src/repository/hasil-panen"

	kebunRepository "github.com/plantation-service/internal/src/repository/kebun"
	komoditasRepository "github.com/plantation-service/internal/src/repository/komoditas"
	perawatanTanamanRepository "github.com/plantation-service/internal/src/repository/perawatan-tanaman"
	staffRepository "github.com/plantation-service/internal/src/repository/staff"
	bahanperawatan "github.com/plantation-service/internal/src/usecase/bahan-perawatan"

	"github.com/plantation-service/internal/src/usecase/kebun"
	"github.com/plantation-service/internal/src/usecase/komoditas"
	perawatantanaman "github.com/plantation-service/internal/src/usecase/perawatan-tanaman"
	"github.com/plantation-service/internal/src/usecase/staff"
	"github.com/plantation-service/pkg/log"
)

func main() {
	// init logger
	log.Init()
	defer log.Sync()

	// init config
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal("can't load config", err)
	}

	// init db connection
	dbConn, err := postgres.InitConn(cfg)
	if err != nil {
		log.Fatal("can't connect to database", err)
	}

	// init repository
	staffRepo := staffRepository.Init(dbConn.GetDB())
	kebunRepo := kebunRepository.Init(dbConn.GetDB())
	bahanPerawatanRepo := bahanPerawatanRepository.Init(dbConn.GetDB())
	komoditasRepo := komoditasRepository.Init(dbConn.GetDB())
	budidayaRepo := budidayaRepository.Init(dbConn.GetDB())
	hasilPanenRepo := hasilPanenRepository.Init(dbConn.GetDB())

	perawatanTanamanRepo := perawatanTanamanRepository.Init(dbConn.GetDB())

	// init usecase
	staffUsecase := staff.Init(staffRepo)
	kebunUsecase := kebun.Init(kebunRepo, bahanPerawatanRepo, budidayaRepo, komoditasRepo, hasilPanenRepo)
	bahanPerawatanUsecase := bahanperawatan.Init(bahanPerawatanRepo)
	komoditasUsecase := komoditas.Init(komoditasRepo)
	perawatanTanamanUsecase := perawatantanaman.Init(perawatanTanamanRepo, staffRepo, budidayaRepo, bahanPerawatanRepo, kebunRepo)

	// init delivery
	dHttp := delivery.Init(staffUsecase, kebunUsecase, bahanPerawatanUsecase, komoditasUsecase, perawatanTanamanUsecase)

	// init router
	r := mux.NewRouter()

	r.Use(corsMiddleware)

	// handler func
	r.Handle("/ping", http.HandlerFunc(dHttp.Ping)).Methods("GET")

	// staff
	r.Handle("/staff", http.HandlerFunc(dHttp.InsertStaff)).Methods("POST", "OPTIONS")
	r.Handle("/staff", http.HandlerFunc(dHttp.GetStaffByID)).Methods("GET", "OPTIONS")
	r.Handle("/staff/all", http.HandlerFunc(dHttp.GetAllStaff)).Methods("GET", "OPTIONS")
	r.Handle("/staff", http.HandlerFunc(dHttp.UpdateStaff)).Methods("PATCH", "OPTIONS")
	r.Handle("/staff", http.HandlerFunc(dHttp.DeleteStaff)).Methods("DELETE", "OPTIONS")

	// kebun
	r.Handle("/kebun", http.HandlerFunc(dHttp.InsertKebun)).Methods("POST", "OPTIONS")
	r.Handle("/kebun", http.HandlerFunc(dHttp.GetKebunByID)).Methods("GET", "OPTIONS")
	r.Handle("/kebun/all", http.HandlerFunc(dHttp.GetAllKebun)).Methods("GET", "OPTIONS")

	// bahan perawatan
	r.Handle("/bahan-perawatan", http.HandlerFunc(dHttp.InsertBahanPerawatan)).Methods("POST", "OPTIONS")
	r.Handle("/bahan-perawatan", http.HandlerFunc(dHttp.GetBahanPerawatanByID)).Methods("GET", "OPTIONS")
	r.Handle("/bahan-perawatan/all", http.HandlerFunc(dHttp.GetAllBahanPerawatan)).Methods("GET", "OPTIONS")
	r.Handle("/bahan-perawatan", http.HandlerFunc(dHttp.UpdateBahanPerawatan)).Methods("PATCH", "OPTIONS")

	// komoditas
	r.Handle("/komoditas", http.HandlerFunc(dHttp.InsertKomoditas)).Methods("POST", "OPTIONS")
	r.Handle("/komoditas", http.HandlerFunc(dHttp.GetKomoditasByID)).Methods("GET", "OPTIONS")
	r.Handle("/komoditas/all", http.HandlerFunc(dHttp.GetAllKomoditas)).Methods("GET", "OPTIONS")

	// relasi
	// bahan perawatan kebun
	r.Handle("/kebun/bahan-perawatan", http.HandlerFunc(dHttp.InsertBahanPerawatankebun)).Methods("POST", "OPTIONS")
	r.Handle("/kebun/bahan-perawatan", http.HandlerFunc(dHttp.GetBahanPerawatanKebun)).Methods("GET", "OPTIONS")
	r.Handle("/kebun/bahan-perawatan", http.HandlerFunc(dHttp.UpdateBahanPerawatankebun)).Methods("PATCH", "OPTIONS")
	r.Handle("/kebun/bahan-perawatan/all", http.HandlerFunc(dHttp.GetAllBahanPerawatanKebun)).Methods("GET", "OPTIONS")

	// budidaya
	r.Handle("/kebun/budidaya", http.HandlerFunc(dHttp.InsertBudidaya)).Methods("POST", "OPTIONS")
	r.Handle("/kebun/budidaya", http.HandlerFunc(dHttp.GetBudidayaByID)).Methods("GET", "OPTIONS")
	r.Handle("/kebun/budidaya", http.HandlerFunc(dHttp.UpdateBudidaya)).Methods("PATCH", "OPTIONS")
	r.Handle("/kebun/budidaya/all", http.HandlerFunc(dHttp.GetAllBudidayaByKebun)).Methods("GET", "OPTIONS")

	// hasil panen
	r.Handle("/hasil-panen", http.HandlerFunc(dHttp.InsertHasilPanen)).Methods("POST", "OPTIONS")
	r.Handle("/hasil-panen", http.HandlerFunc(dHttp.GetHasilPanenByID)).Methods("GET", "OPTIONS")
	r.Handle("/hasil-panen/all", http.HandlerFunc(dHttp.GetAllHasilPanen)).Methods("GET", "OPTIONS")
	r.Handle("/hasil-panen", http.HandlerFunc(dHttp.UpdateHasilPanen)).Methods("PATCH", "OPTIONS")

	// perawatan
	r.Handle("/perawatan", http.HandlerFunc(dHttp.InsertPerawatanTanaman)).Methods("POST", "OPTIONS")
	r.Handle("/perawatan", http.HandlerFunc(dHttp.GetPerawatanTanamanByID)).Methods("GET", "OPTIONS")
	r.Handle("/perawatan", http.HandlerFunc(dHttp.UpdatePerawatanTanaman)).Methods("PATCH", "OPTIONS")
	r.Handle("/perawatan/all", http.HandlerFunc(dHttp.GetAllPerawatanTanaman)).Methods("GET", "OPTIONS")

	// start server
	log.Info("Starting the server...")
	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Fatal("can't start server", err)
	}
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
