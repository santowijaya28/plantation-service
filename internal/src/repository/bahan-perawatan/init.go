package bahanperawatan

import "database/sql"

type BPRepo struct {
	dbConn *sql.DB
}

func Init(dbConn *sql.DB) *BPRepo {
	return &BPRepo{
		dbConn: dbConn,
	}
}
