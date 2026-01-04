package perawatantanaman

import (
	"github.com/plantation-service/internal/src/repository"
	"github.com/plantation-service/internal/src/usecase"
)

type perawatanTanaman struct {
	perawatanTanamanRepo repository.PerawatanTanamanRepoInterface
	staffRepo            repository.StaffRepoInterface
	budidayaRepo         repository.BudidayaRepoInterface
	bahanPerawatanRepo   repository.BahanPerawatanRepoInterface
	kebunRepo            repository.KebunRepoInterface
}

func Init(
	perawatanTanamanRepo repository.PerawatanTanamanRepoInterface,
	staffRepo repository.StaffRepoInterface,
	budidayaRepo repository.BudidayaRepoInterface,
	bahanPerawatanRepo repository.BahanPerawatanRepoInterface,
	kebunRepo repository.KebunRepoInterface,
) usecase.PerawatanTanamanUsecase {
	return &perawatanTanaman{
		perawatanTanamanRepo: perawatanTanamanRepo,
		staffRepo:            staffRepo,
		budidayaRepo:         budidayaRepo,
		bahanPerawatanRepo:   bahanPerawatanRepo,
		kebunRepo:            kebunRepo,
	}
}
