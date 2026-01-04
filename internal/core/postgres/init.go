package postgres

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/plantation-service/internal/src/config"
)

// InitConn will initialize postgres database
func InitConn(cfg *config.Config) (*pqConn, error) {
	if cfg == nil {
		return nil, errors.New("missing config")
	}

	// init  database
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		cfg.DatabaseConfig.Host,
		cfg.DatabaseConfig.Port,
		cfg.DatabaseConfig.DbUserName,
		cfg.DatabaseConfig.DbName,
		cfg.DatabaseConfig.DbPwd,
	)

	plantationDb, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return &pqConn{
		plantationDB: plantationDb,
	}, nil
}

func (p *pqConn) GetDB() *sql.DB {
	return p.plantationDB
}
