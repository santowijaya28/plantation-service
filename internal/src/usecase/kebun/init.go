package kebun

import "github.com/plantation-service/internal/src/repository"

type kebunUsecase struct {
	kebunRepo          repository.KebunRepoInterface
	bahanPerawatanRepo repository.BahanPerawatanRepoInterface
	budidayaRepo       repository.BudidayaRepoInterface
	komoditasRepo      repository.KomoditasRepoInterface
	hasilPanenRepo     repository.HasilPanenRepoInterface
}

func Init(kebunRepo repository.KebunRepoInterface,
	bahanPerawatanRepo repository.BahanPerawatanRepoInterface,
	budidayaRepo repository.BudidayaRepoInterface,
	komoditasRepo repository.KomoditasRepoInterface,
	hasilPanenRepo repository.HasilPanenRepoInterface,
) *kebunUsecase {
	return &kebunUsecase{
		kebunRepo:          kebunRepo,
		bahanPerawatanRepo: bahanPerawatanRepo,
		budidayaRepo:       budidayaRepo,
		komoditasRepo:      komoditasRepo,
		hasilPanenRepo:     hasilPanenRepo,
	}
}
