package delivery

import "github.com/plantation-service/internal/src/usecase"

func Init(staffUsecase usecase.StaffUsecase,
	kebunUsecase usecase.KebunUsecase,
	bahanPerawatanUsecase usecase.BahanPerawatanUsecase,
	komoditasUsecase usecase.KomoditasUsecase,
	perawatanTanamanUsecase usecase.PerawatanTanamanUsecase,
) delivery {
	return delivery{
		staffUsecase:            staffUsecase,
		kebunUsecase:            kebunUsecase,
		bahanPerawatanusecase:   bahanPerawatanUsecase,
		komoditasUsecase:        komoditasUsecase,
		perawatanTanamanUsecase: perawatanTanamanUsecase,
	}
}
