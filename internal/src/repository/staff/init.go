package staff

import "database/sql"

type staff struct {
	plantationDb *sql.DB
}

func Init(db *sql.DB) *staff {
	return &staff{
		plantationDb: db,
	}
}
