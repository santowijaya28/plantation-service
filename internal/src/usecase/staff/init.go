package staff

import "github.com/plantation-service/internal/src/repository"

type staffUsecase struct {
	staffRepo repository.StaffRepoInterface
}

func Init(staffRepo repository.StaffRepoInterface) *staffUsecase {
	return &staffUsecase{
		staffRepo: staffRepo,
	}
}
