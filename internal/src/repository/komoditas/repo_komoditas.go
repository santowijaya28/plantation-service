package komoditas

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/plantation-service/internal/src/model"
)

func (r *komoditasRepo) InsertKomoditas(ctx context.Context, data model.Komoditas) (result model.Komoditas, err error) {
	query := `
		INSERT INTO 
			komoditas
			(nama_komoditas, jenis_tanaman)
		VALUES
			($1, $2)
	`

	_, err = r.dbConn.ExecContext(ctx, query,
		data.NamaKomoditas,
		data.JenisTanaman,
	)

	if err != nil {
		return model.Komoditas{}, err
	}

	return data, nil
}

func (r *komoditasRepo) GetByID(ctx context.Context, idKomoditas int) (data model.Komoditas, err error) {
	query := `
		SELECT 
			id_komoditas,
			nama_komoditas,
			jenis_tanaman
		FROM komoditas
		WHERE id_komoditas = $1
	`

	err = r.dbConn.QueryRowContext(ctx, query, idKomoditas).Scan(
		&data.ID,
		&data.NamaKomoditas,
		&data.JenisTanaman,
	)
	if err != nil {
		return model.Komoditas{}, err
	}

	return data, nil
}

func (r *komoditasRepo) GetAllKomoditas(ctx context.Context, filter model.FilterKomoditas, page, pageSize int) (data model.AllKomoditas, err error) {
	var (
		filters []string
		args    []interface{}
		argID   = 1
	)

	if filter.NamaKomoditas != "" {
		filters = append(filters, fmt.Sprintf("LOWER(nama_komoditas) ILIKE '%%' || $%d || '%%'", argID))
		args = append(args, strings.ToLower(filter.NamaKomoditas))
		argID++
	}

	if filter.JenisTanaman != "" {
		filters = append(filters, fmt.Sprintf("LOWER(jenis_tanaman) ILIKE '%%' || $%d || '%%'", argID))
		args = append(args, filter.JenisTanaman)
		argID++
	}

	whereClause := ""
	if len(filters) > 0 {
		whereClause = "WHERE " + strings.Join(filters, " AND ")
	}

	var totalCount int
	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM komoditas %s`, whereClause)
	err = r.dbConn.QueryRowContext(ctx, countQuery, args...).Scan(&totalCount)
	if err != nil {
		return data, fmt.Errorf("count query failed: %w", err)
	}

	offset := (page - 1) * pageSize
	args = append(args, pageSize, offset)

	query := fmt.Sprintf(`
		SELECT 
			id_komoditas,
			nama_komoditas,
			jenis_tanaman
		FROM komoditas
		%s
		ORDER BY id_komoditas DESC
		LIMIT $%d
		OFFSET $%d
		`, whereClause, argID, argID+1)

	rows, err := r.dbConn.QueryContext(ctx, query, args...)
	if err != nil && err != sql.ErrNoRows {
		return data, err
	}

	if err == sql.ErrNoRows {
		data.Data = []model.Komoditas{}
		data.TotalCount = 0
		data.Page = page
		data.PageSize = pageSize
		return data, nil
	}

	defer rows.Close()
	var result []model.Komoditas
	for rows.Next() {
		var komoditas model.Komoditas

		err := rows.Scan(
			&komoditas.ID,
			&komoditas.NamaKomoditas,
			&komoditas.JenisTanaman,
		)
		if err != nil {
			return model.AllKomoditas{}, err
		}

		result = append(result, komoditas)
	}

	data.Data = result
	data.Page = page
	data.PageSize = pageSize
	data.TotalCount = totalCount

	return data, nil
}
