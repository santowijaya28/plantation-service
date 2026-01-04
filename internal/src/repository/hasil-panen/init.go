package hasilpanen

import "database/sql"

type HasilPanenRepo struct {
	dbConn *sql.DB
}

func Init(dbConn *sql.DB) *HasilPanenRepo {
	return &HasilPanenRepo{
		dbConn: dbConn,
	}
}
