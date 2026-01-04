package hasilpanen

import (
	"context"
	"strconv"
	"strings"

	"github.com/plantation-service/internal/src/model"
)

func (r *HasilPanenRepo) InsertHasilPanen(ctx context.Context, data model.HasilPanen) (model.HasilPanen, error) {
	query := `
		INSERT INTO hasil_panen (
			id_budidaya,
			tanggal_panen,
			total_kg
		)
		VALUES ($1, $2, $3)
		RETURNING
			id_panen,
			id_budidaya,
			tanggal_panen,
			total_kg
	`

	var result model.HasilPanen
	_, err := r.dbConn.ExecContext(ctx, query,
		data.IDBudidaya,
		data.TanggalPanen,
		data.TotalKg,
	)
	if err != nil {
		return model.HasilPanen{}, err
	}

	return result, nil
}

func (r *HasilPanenRepo) UpdateHasilPanen(ctx context.Context, data model.HasilPanen) (model.HasilPanen, error) {
	query := `
		UPDATE hasil_panen
		SET
			id_budidaya = $1,
			tanggal_panen = $2,
			total_kg = $3
		WHERE id_panen = $4
		RETURNING
			id_panen,
			id_budidaya,
			tanggal_panen,
			total_kg
	`

	var result model.HasilPanen
	err := r.dbConn.QueryRowContext(ctx, query,
		data.IDBudidaya,
		data.TanggalPanen,
		data.TotalKg,
		data.IDPanen,
	).Scan(
		&result.IDPanen,
		&result.IDBudidaya,
		&result.TanggalPanen,
		&result.TotalKg,
	)

	if err != nil {
		return model.HasilPanen{}, err
	}

	return result, nil
}

func (r *HasilPanenRepo) GetByID(ctx context.Context, idPanen int) (data model.HasilPanen, err error) {
	query := `
		SELECT 
			id_budidaya,
			tanggal_panen,
			total_kg
		FROM hasil_panen
		WHERE id_panen = $1
	`

	err = r.dbConn.QueryRowContext(ctx, query, idPanen).Scan(
		&data.IDBudidaya,
		&data.TanggalPanen,
		&data.TotalKg,
	)
	if err != nil {
		return model.HasilPanen{}, err
	}

	return data, nil
}

func (r *HasilPanenRepo) GetAllHasilPanen(
	ctx context.Context,
	filter model.FilterHasilPanen,
	page, pageSize int,
) (data model.AllHasilPanen, err error) {

	var (
		filters []string
		args    []interface{}
		argID   = 1
	)

	// Filter ID Kebun
	if filter.IDKebun != 0 {
		filters = append(filters, "budidaya.id_kebun = $"+strconv.Itoa(argID))
		args = append(args, filter.IDKebun)
		argID++
	}

	// Filter ID Budidaya
	if filter.IDBudidaya != 0 {
		filters = append(filters, "hasil_panen.id_budidaya = $"+strconv.Itoa(argID))
		args = append(args, filter.IDBudidaya)
		argID++
	}

	// Filter tanggal
	if filter.StartDate != "" && filter.EndDate != "" {
		filters = append(filters, "tanggal_panen BETWEEN $"+strconv.Itoa(argID)+" AND $"+strconv.Itoa(argID+1))
		args = append(args, filter.StartDate, filter.EndDate)
		argID += 2
	}

	// WHERE final
	whereClause := ""
	if len(filters) > 0 {
		whereClause = "WHERE " + strings.Join(filters, " AND ")
	}

	// ==== Hitung total ====
	countQuery := `
		SELECT COUNT(*)
		FROM hasil_panen
		JOIN budidaya ON hasil_panen.id_budidaya = budidaya.id_budidaya
		` + whereClause

	if err := r.dbConn.QueryRowContext(ctx, countQuery, args...).Scan(&data.TotalCount); err != nil {
		return data, err
	}

	// ==== Pagination ====
	offset := (page - 1) * pageSize
	args = append(args, pageSize, offset)

	limitArg := strconv.Itoa(argID)
	offsetArg := strconv.Itoa(argID + 1)

	// ==== Query data ====
	query := `
		SELECT 
			hasil_panen.id_panen,
			hasil_panen.id_budidaya,
			komoditas.nama_komoditas,
			hasil_panen.tanggal_panen,
			hasil_panen.total_kg
		FROM hasil_panen
		JOIN budidaya 
			ON hasil_panen.id_budidaya = budidaya.id_budidaya
		JOIN komoditas
			ON budidaya.id_komoditas = komoditas.id_komoditas
		` + whereClause + `
		ORDER BY tanggal_panen DESC
		LIMIT $` + limitArg + ` OFFSET $` + offsetArg

	rows, err := r.dbConn.QueryContext(ctx, query, args...)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	for rows.Next() {
		var record model.HasilPanen
		if err := rows.Scan(
			&record.IDPanen,
			&record.IDBudidaya,
			&record.NamaKomoditas,
			&record.TanggalPanen,
			&record.TotalKg,
		); err != nil {
			return data, err
		}
		data.Data = append(data.Data, record)
	}

	data.Page = page
	data.PageSize = pageSize
	return data, nil
}
