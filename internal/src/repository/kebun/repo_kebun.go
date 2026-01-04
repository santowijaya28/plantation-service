package kebun

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/plantation-service/internal/src/model"
)

func (repo *kebunRepo) GetByID(ctx context.Context, idKebun int) (data model.Kebun, err error) {
	query := `SELECT 
	id_kebun, 
	nama_kebun, 
	luas_kebun,
	jenis_kebun,
	lat,
	long 
	FROM kebun 
	WHERE id_kebun = $1`

	err = repo.dbConn.QueryRowContext(ctx, query, idKebun).Scan(
		&data.ID,
		&data.NamaKebun,
		&data.LuasKebun,
		&data.JenisKebun,
		&data.Lat,
		&data.Long,
	)

	if err != nil && err != sql.ErrNoRows {
		return model.Kebun{}, err
	}

	return data, nil
}

func (repo *kebunRepo) GetAllKebun(ctx context.Context,
	filter model.KebunFilter,
	page, pageSize int) (data model.AllKebun, err error) {

	var (
		filters []string
		args    []interface{}
		argID   = 1
	)

	if filter.NamaKebun != "" {
		filters = append(filters, fmt.Sprintf("LOWER(nama_kebun) ILIKE '%%' || $%d || '%%'", argID))
		args = append(args, strings.ToLower(filter.NamaKebun))
		argID++
	}

	if filter.JenisKebun != "" {
		filters = append(filters, fmt.Sprintf("jenis_kebun = $%d", argID))
		args = append(args, filter.JenisKebun)
		argID++
	}

	whereClause := ""
	if len(filters) > 0 {
		whereClause = "WHERE " + strings.Join(filters, " AND ")
	}

	var totalCount int
	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM kebun %s`, whereClause)
	err = repo.dbConn.QueryRowContext(ctx, countQuery, args...).Scan(&totalCount)
	if err != nil && err != sql.ErrNoRows {
		return data, fmt.Errorf("count query failed: %w", err)
	}

	if totalCount == 0 {
		data.Data = []model.Kebun{}
		data.TotalCount = 0
		data.Page = page
		data.PageSize = pageSize
		return data, nil
	}

	offset := (page - 1) * pageSize
	args = append(args, pageSize, offset)

	query := fmt.Sprintf(`
	SELECT 
		id_kebun,
		nama_kebun,
		luas_kebun,
		jenis_kebun,
		lat,
		long
	FROM kebun
	%s
	ORDER BY id_kebun DESC
	LIMIT $%d
	OFFSET $%d
	`, whereClause, argID, argID+1)

	rows, err := repo.dbConn.QueryContext(ctx, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return data, err
	}

	var result []model.Kebun
	defer rows.Close()

	for rows.Next() {
		var kebun model.Kebun
		err := rows.Scan(
			&kebun.ID,
			&kebun.NamaKebun,
			&kebun.LuasKebun,
			&kebun.JenisKebun,
			&kebun.Lat,
			&kebun.Long,
		)

		if err != nil {
			return data, err
		}

		result = append(result, kebun)
	}

	data.Data = result
	data.TotalCount = totalCount
	data.Page = page
	data.PageSize = pageSize
	return
}

func (repo *kebunRepo) InsertKebun(ctx context.Context, data model.Kebun) (result model.Kebun, err error) {
	query := `INSERT INTO 
	kebun (
	nama_kebun,
	luas_kebun,
	jenis_kebun,
	lat, long) VALUES ($1, $2, $3, $4, $5)`

	_, err = repo.dbConn.Exec(query,
		data.NamaKebun,
		data.LuasKebun,
		data.JenisKebun,
		data.Lat,
		data.Long,
	)

	return model.Kebun{}, err
}
