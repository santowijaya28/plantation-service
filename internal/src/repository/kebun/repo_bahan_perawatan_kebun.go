package kebun

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/plantation-service/internal/src/model"
)

func (r *kebunRepo) InsertBahanPerawatankebun(ctx context.Context, data model.BahanPerawatanKebun) (result model.BahanPerawatanKebun, err error) {
	query := `INSERT INTO bahan_perawatan_kebun
	(id_bahan, id_kebun, stok_kg)
	VALUES ($1, $2, $3)
	ON CONFLICT (id_bahan, id_kebun) DO UPDATE
	SET stok_kg = EXCLUDED.stok_kg
	RETURNING stok_kg`

	var stokKg float64
	err = r.dbConn.QueryRowContext(ctx, query,
		data.IDBahan,
		data.IDKebun,
		data.StokKg,
	).Scan(&stokKg)

	if err == sql.ErrNoRows {
		err = nil
	}

	result = model.BahanPerawatanKebun{
		IDBahan: data.IDBahan,
		IDKebun: data.IDKebun,
		StokKg:  stokKg,
	}

	return result, err
}

func (r *kebunRepo) GetBahanPerawatanKebun(ctx context.Context, idBahanPerawatan int, idKebun int) (data model.BahanPerawatanKebun, err error) {
	query := `SELECT 
	stok_kg 
	FROM bahan_perawatan_kebun
	WHERE id_bahan = $1 AND id_kebun = $2`

	err = r.dbConn.QueryRowContext(ctx, query, idBahanPerawatan, idKebun).Scan(
		&data.StokKg,
	)

	if err == sql.ErrNoRows {
		err = nil
		return
	}

	data.IDBahan = idBahanPerawatan
	data.IDKebun = idKebun

	return data, err
}

func (r *kebunRepo) UpdateBahanPerawatankebun(ctx context.Context, idBahanPerawatan int, idKebun int, data model.BahanPerawatanKebun) (result model.BahanPerawatanKebun, err error) {
	query := `UPDATE bahan_perawatan_kebun SET 
	stok_kg = $1
	WHERE id_bahan = $2 AND id_kebun = $3`

	_, err = r.dbConn.ExecContext(ctx, query,
		data.StokKg,
		idBahanPerawatan,
		idKebun,
	)

	return model.BahanPerawatanKebun{
		IDBahan: idBahanPerawatan,
		IDKebun: idKebun,
		StokKg:  data.StokKg,
	}, err
}

func (r *kebunRepo) GetAllBahanPerawatanKebun(
	ctx context.Context,
	filter model.FilterBahanPerawatanKebun,
	page, pageSize int,
) (data model.AllBahanPerawatanKebun, err error) {

	var (
		filters []string
		args    []interface{}
		argID   = 1
	)

	// ===========================
	// FILTER: ID Kebun
	// ===========================
	if filter.IDKebun > 0 {
		filters = append(filters, fmt.Sprintf("bpk.id_kebun = $%d", argID))
		args = append(args, filter.IDKebun)
		argID++
	}

	// ===========================
	// FILTER: Nama Bahan
	// ===========================
	if filter.NamaBahan != "" {
		filters = append(filters,
			fmt.Sprintf("LOWER(bp.nama_bahan) ILIKE $%d", argID),
		)
		args = append(args, "%"+strings.ToLower(filter.NamaBahan)+"%")
		argID++
	}

	// ===========================
	// WHERE CLAUSE
	// ===========================
	whereClause := ""
	if len(filters) > 0 {
		whereClause = "WHERE " + strings.Join(filters, " AND ")
	}

	// ===========================
	// HITUNG TOTAL DATA
	// ===========================
	var totalCount int
	countQuery := fmt.Sprintf(`
		SELECT COUNT(*) 
		FROM bahan_perawatan_kebun bpk
		JOIN bahan_perawatan bp ON bpk.id_bahan = bp.id_bahan
		%s
	`, whereClause)

	err = r.dbConn.QueryRowContext(ctx, countQuery, args...).Scan(&totalCount)
	if err != nil && err != sql.ErrNoRows {
		return
	}

	// totalCount disimpan dalam response
	data.TotalCount = totalCount

	// ===========================
	// PAGINATION
	// ===========================
	offset := (page - 1) * pageSize

	// ===========================
	// QUERY DATA
	// ===========================
	query := fmt.Sprintf(`
		SELECT 
			bpk.id_bahan,
			bpk.id_kebun,
			bp.nama_bahan,
			bpk.stok_kg
		FROM bahan_perawatan_kebun bpk
		JOIN bahan_perawatan bp ON bpk.id_bahan = bp.id_bahan
		%s
		ORDER BY bpk.id_kebun DESC
		LIMIT $%d OFFSET $%d
	`, whereClause, argID, argID+1)

	args = append(args, pageSize, offset)

	rows, err := r.dbConn.QueryContext(ctx, query, args...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var item model.BahanPerawatanWithMeta
		err = rows.Scan(
			&item.IDBahan,
			&item.IDKebun,
			&item.NamaBahan,
			&item.StokKg,
		)
		if err != nil {
			return
		}

		data.Data = append(data.Data, item)
	}

	return
}
