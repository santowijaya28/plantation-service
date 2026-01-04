package postgres

import "database/sql"

type pqConn struct {
	plantationDB *sql.DB
}
