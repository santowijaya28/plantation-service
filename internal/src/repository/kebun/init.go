package kebun

import "database/sql"

type kebunRepo struct {
	dbConn *sql.DB
}

func Init(dbConn *sql.DB) *kebunRepo {
	return &kebunRepo{
		dbConn: dbConn,
	}
}
