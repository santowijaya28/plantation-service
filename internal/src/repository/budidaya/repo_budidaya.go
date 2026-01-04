package budidaya

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/plantation-service/internal/src/model"
)

func (r *BudidayaRepo) InsertBudidaya(ctx context.Context, data model.Budidaya) (result model.Budidaya, err error) {
	query := `
		INSERT INTO 
			budidaya
			(id_kebun, 
			id_komoditas,
			tanggal_tanam, 
			jumlah_tanaman,
			tanggal_estimasi_panen,
			status_tanaman
			)
		VALUES
			($1, 
			$2, 
			$3, 
			$4,
			$5,
			$6
			)
	`

	_, err = r.dbConn.ExecContext(ctx, query,
		data.IDKebun,
		data.IDKomoditas,
		data.TanggalTanam,
		data.JumlahTanaman,
		data.TanggalEstimasiPanen,
		data.StatusTanaman,
	)

	if err != nil {
		return model.Budidaya{}, err
	}

	return data, nil
}

func (r *BudidayaRepo) GetByID(ctx context.Context, idBudidaya int) (data model.Budidaya, err error) {
	query := `
		SELECT 
			id_kebun, 
			id_komoditas,
			tanggal_tanam, 
			jumlah_tanaman,
			tanggal_estimasi_panen,
			status_tanaman
		FROM budidaya
		WHERE id_budidaya = $1
	`

	err = r.dbConn.QueryRowContext(ctx, query, idBudidaya).Scan(
		&data.IDKebun,
		&data.IDKomoditas,
		&data.TanggalTanam,
		&data.JumlahTanaman,
		&data.TanggalEstimasiPanen,
		&data.StatusTanaman,
	)

	if err != nil {
		return model.Budidaya{}, err
	}

	data.ID = idBudidaya
	return data, nil
}

func (r *BudidayaRepo) UpdateBudidaya(ctx context.Context, idBudidaya int, data model.Budidaya) (result model.Budidaya, err error) {
	query := `
		UPDATE 
			budidaya
		SET
			id_kebun = $1, 
			id_komoditas = $2,
			tanggal_tanam = $3, 
			jumlah_tanaman = $4,
			tanggal_estimasi_panen = $5,
			status_tanaman = $6
		WHERE id_budidaya = $7
	`

	_, err = r.dbConn.ExecContext(ctx, query,
		data.IDKebun,
		data.IDKomoditas,
		data.TanggalTanam,
		data.JumlahTanaman,
		data.TanggalEstimasiPanen,
		data.StatusTanaman,
		idBudidaya,
	)

	if err != nil {
		return model.Budidaya{}, err
	}

	return data, nil
}

func (r *BudidayaRepo) GetAllByKebun(
	ctx context.Context,
	filter model.FilterBudidaya,
	page, pageSize int,
) (data []model.Budidaya, err error) {

	var (
		filters []string
		args    []interface{}
		argID   = 1
	)

	// ID Kebun
	if filter.IDKebun > 0 {
		filters = append(filters, "id_kebun = $"+strconv.Itoa(argID))
		args = append(args, filter.IDKebun)
		argID++
	}

	// ID Komoditas
	if filter.IDKomoditas > 0 {
		filters = append(filters, "id_komoditas = $"+strconv.Itoa(argID))
		args = append(args, filter.IDKomoditas)
		argID++
	}

	// Status tanaman (ILIKE)
	if filter.StatusTanaman != "" {
		filters = append(filters, "LOWER(status_tanaman) ILIKE LOWER('%' || $"+strconv.Itoa(argID)+" || '%')")
		args = append(args, filter.StatusTanaman)
		argID++
	}

	if filter.NamaKomoditas != "" {
		filters = append(filters,
			"LOWER(komoditas.nama_komoditas) ILIKE LOWER($"+strconv.Itoa(argID)+")",
		)
		args = append(args, "%"+filter.NamaKomoditas+"%")
		argID++
	}

	// Tanggal Estimasi Panen
	if filter.TanggalEstimasiPanen != "" {
		filters = append(filters, "tanggal_estimasi_panen = $"+strconv.Itoa(argID))
		args = append(args, filter.TanggalEstimasiPanen)
		argID++
	}

	// Tanggal Tanam
	if filter.TanggalTanam != "" {
		filters = append(filters, "tanggal_tanam = $"+strconv.Itoa(argID))
		args = append(args, filter.TanggalTanam)
		argID++
	}

	// Final WHERE
	whereClause := ""
	if len(filters) > 0 {
		whereClause = "WHERE " + strings.Join(filters, " AND ")
	}

	// Hitung total
	var totalCount int
	countQuery := "SELECT COUNT(*) FROM budidaya JOIN komoditas ON budidaya.id_komoditas = komoditas.id_komoditas " + whereClause

	if err := r.dbConn.QueryRowContext(ctx, countQuery, args...).Scan(&totalCount); err != nil {
		return nil, err
	}
	if totalCount == 0 {
		return []model.Budidaya{}, nil
	}

	// Pagination
	offset := (page - 1) * pageSize
	args = append(args, pageSize, offset)

	query := fmt.Sprintf(`
        SELECT 
            id_budidaya,
			komoditas.nama_komoditas,
            id_kebun, 
            komoditas.id_komoditas,
            tanggal_tanam, 
            jumlah_tanaman,
            tanggal_estimasi_panen,
            status_tanaman
        FROM budidaya
		JOIN komoditas ON budidaya.id_komoditas = komoditas.id_komoditas
        %s
        ORDER BY id_budidaya DESC
        LIMIT $%d OFFSET $%d
    `, whereClause, argID, argID+1)

	rows, err := r.dbConn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan rows
	for rows.Next() {
		var b model.Budidaya
		if err := rows.Scan(
			&b.ID,
			&b.NamaKomoditas,
			&b.IDKebun,
			&b.IDKomoditas,
			&b.TanggalTanam,
			&b.JumlahTanaman,
			&b.TanggalEstimasiPanen,
			&b.StatusTanaman,
		); err != nil {
			return nil, err
		}
		data = append(data, b)
	}

	return data, rows.Err()
}
