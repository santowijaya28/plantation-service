package komoditas

import "database/sql"

type komoditasRepo struct {
	dbConn *sql.DB
}

func Init(dbConn *sql.DB) *komoditasRepo {
	return &komoditasRepo{
		dbConn: dbConn,
	}
}
