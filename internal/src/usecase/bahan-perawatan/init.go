package bahanperawatan

import "github.com/plantation-service/internal/src/repository"

type bpUsecase struct {
	bpRepo repository.BahanPerawatanRepoInterface
}

func Init(bpRepo repository.BahanPerawatanRepoInterface) *bpUsecase {
	return &bpUsecase{
		bpRepo: bpRepo,
	}
}
