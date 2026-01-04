package budidaya

import "database/sql"

type BudidayaRepo struct {
	dbConn *sql.DB
}

func Init(dbConn *sql.DB) *BudidayaRepo {
	return &BudidayaRepo{
		dbConn: dbConn,
	}
}
