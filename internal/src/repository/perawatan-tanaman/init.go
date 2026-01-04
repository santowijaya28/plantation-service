package perawatantanaman

import (
	"database/sql"

	"github.com/plantation-service/internal/src/repository"
)

type perawatanTanaman struct {
	db *sql.DB
}

func Init(db *sql.DB) repository.PerawatanTanamanRepoInterface {
	return &perawatanTanaman{
		db: db,
	}
}
