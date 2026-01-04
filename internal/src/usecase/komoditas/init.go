package komoditas

import "github.com/plantation-service/internal/src/repository"

type komoditasUsecase struct {
	komoditasRepo repository.KomoditasRepoInterface
}

func Init(komoditasRepo repository.KomoditasRepoInterface) *komoditasUsecase {
	return &komoditasUsecase{
		komoditasRepo: komoditasRepo,
	}
}
