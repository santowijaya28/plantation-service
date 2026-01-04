package bahanperawatan

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/plantation-service/internal/src/model"
)

func (r *BPRepo) InsertBahanPerawatan(ctx context.Context, data model.BahanPerawatan) (result model.BahanPerawatan, err error) {
	query := `INSERT INTO 
	bahan_perawatan(
		nama_bahan,
		jenis_bahan,
		tipe_perawatan,
		harga_kg
	) VALUES (
	 	$1, 
		$2,
		$3, 
		$4
	)
	`

	_, err = r.dbConn.Exec(query,
		data.NamaBahan,
		data.JenisBahan,
		data.TipePerawatan,
		data.HargaKg,
	)

	return model.BahanPerawatan{}, err
}

func (r *BPRepo) GetByID(ctx context.Context, idBahanPerawatan int) (data model.BahanPerawatan, err error) {
	query := `
	SELECT 
		id_bahan,
		nama_bahan,
		jenis_bahan,
		tipe_perawatan,
		harga_kg
	FROM bahan_perawatan
	WHERE 
		id_bahan = $1
	`

	err = r.dbConn.QueryRowContext(ctx,
		query,
		idBahanPerawatan).Scan(
		&data.ID,
		&data.NamaBahan,
		&data.JenisBahan,
		&data.TipePerawatan,
		&data.HargaKg,
	)

	if err != nil && err != sql.ErrNoRows {
		return
	}

	return data, nil
}

func (r *BPRepo) GetAllBahanPerawatan(ctx context.Context,
	filter model.FilterBahanPerawatan,
	page, pageSize int) (
	data model.AllBahanPerawatan, err error) {

	var (
		filters []string
		args    []interface{}
		argID   = 1
	)

	if filter.NamaBahan != "" {
		filters = append(filters, fmt.Sprintf("nama_bahan ILIKE '%%' || $%d || '%%'", argID))
		args = append(args, filter.NamaBahan)
		argID++
	}

	if filter.JenisBahan != "" {
		filters = append(filters, fmt.Sprintf("jenis_bahan = $%d", argID))
		args = append(args, filter.JenisBahan)
		argID++
	}

	if filter.TipePerawatan != "" {
		filters = append(filters, fmt.Sprintf("tipe_perawatan = $%d", argID))
		args = append(args, filter.TipePerawatan)
		argID++
	}

	whereClause := ""
	if len(filters) > 0 {
		whereClause = "WHERE " + strings.Join(filters, " AND ")
	}

	var totalCount int
	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM bahan_perawatan %s`, whereClause)
	err = r.dbConn.QueryRowContext(ctx, countQuery, args...).Scan(&totalCount)
	if err != nil {
		return data, fmt.Errorf("count query failed: %w", err)
	}

	if totalCount == 0 {
		return data, nil
	}

	offset := (page - 1) * pageSize
	args = append(args, pageSize, offset)

	query := fmt.Sprintf(`
	SELECT 
		id_bahan,
		nama_bahan,
		jenis_bahan,
		tipe_perawatan,
		harga_kg
	FROM bahan_perawatan
		%s
	ORDER BY id_bahan DESC
	LIMIT $%d
	OFFSET $%d
	`, whereClause, argID, argID+1)

	rows, err := r.dbConn.QueryContext(ctx,
		query,
		args...,
	)

	if err != nil {
		return data, err
	}

	defer rows.Close()

	var result []model.BahanPerawatan
	for rows.Next() {
		var bahanPerawatan model.BahanPerawatan

		err := rows.Scan(
			&bahanPerawatan.ID,
			&bahanPerawatan.NamaBahan,
			&bahanPerawatan.JenisBahan,
			&bahanPerawatan.TipePerawatan,
			&bahanPerawatan.HargaKg,
		)
		if err != nil {
			return data, err
		}

		result = append(result, bahanPerawatan)
	}

	data.Data = result
	data.TotalCount = totalCount
	data.Page = page
	data.PageSize = pageSize
	return data, nil
}

func (r *BPRepo) UpdateBahanPerawatan(ctx context.Context, idBahanPerawatan int, data model.BahanPerawatan) (result model.BahanPerawatan, err error) {
	query := `UPDATE 
	bahan_perawatan SET
		nama_bahan = $1,
		jenis_bahan = $2,
		tipe_perawatan = $3,
		harga_kg = $4
	WHERE 
		id_bahan = $5
	`

	_, err = r.dbConn.Exec(query,
		data.NamaBahan,
		data.JenisBahan,
		data.TipePerawatan,
		data.HargaKg,
		idBahanPerawatan,
	)

	return model.BahanPerawatan{}, err
}
