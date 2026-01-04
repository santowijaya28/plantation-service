package delivery

import "github.com/plantation-service/internal/src/usecase"

type delivery struct {
	staffUsecase            usecase.StaffUsecase
	kebunUsecase            usecase.KebunUsecase
	bahanPerawatanusecase   usecase.BahanPerawatanUsecase
	komoditasUsecase        usecase.KomoditasUsecase
	perawatanTanamanUsecase usecase.PerawatanTanamanUsecase
}
