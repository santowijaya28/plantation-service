package perawatantanaman

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/plantation-service/internal/src/model"
)

func (r *perawatanTanaman) InsertPerawatanTanaman(ctx context.Context, data model.PerawatanTanaman) (result model.PerawatanTanaman, err error) {
	query := `INSERT INTO perawatan_tanaman 
	(id_staff, 
	id_budidaya, 
	id_bahan_perawatan, 
	jumlah_bahan, 
	biaya_perawatan, 
	tanggal_perawatan,
	jenis_perawatan) 
	VALUES 
	($1, $2, $3, $4, $5, $6, $7) 
	RETURNING 
	id_riwayat_perawatan, id_staff, id_budidaya, id_bahan_perawatan, jumlah_bahan, biaya_perawatan, tanggal_perawatan, jenis_perawatan`

	err = r.db.QueryRowContext(ctx, query,
		data.IDStaff,
		data.IDBudidaya,
		data.IDBahanPerawatan,
		data.JumlahBahan,
		data.BiayaPerawatan,
		data.TanggalPerawatan,
		data.JenisPerawatan,
	).Scan(
		&result.IDRiwayatPerawatan,
		&result.IDStaff,
		&result.IDBudidaya,
		&result.IDBahanPerawatan,
		&result.JumlahBahan,
		&result.BiayaPerawatan,
		&result.TanggalPerawatan,
		&result.JenisPerawatan,
	)

	return
}

func (r *perawatanTanaman) GetbyId(ctx context.Context, idPerwatanTanaman int) (data model.PerawatanTanaman, err error) {
	query := `SELECT 
	p.id_riwayat_perawatan, 
	p.id_staff, 
	p.id_budidaya, 
	p.id_bahan_perawatan, 
	p.jumlah_bahan, 
	p.biaya_perawatan, 
	p.tanggal_perawatan,
	p.jenis_perawatan,
	k.nama_komoditas,
	COALESCE(bp.nama_bahan, '') as nama_bahan,
	COALESCE(s.nama_staff, '') as nama_staff,
	kb.nama_kebun
	FROM perawatan_tanaman p
	JOIN budidaya b ON p.id_budidaya = b.id_budidaya
	JOIN komoditas k ON b.id_komoditas = k.id_komoditas
	JOIN kebun kb ON b.id_kebun = kb.id_kebun
	LEFT JOIN bahan_perawatan bp ON p.id_bahan_perawatan = bp.id_bahan
	LEFT JOIN staff_kebun s ON p.id_staff = s.id_staff
	WHERE p.id_riwayat_perawatan = $1`

	err = r.db.QueryRowContext(ctx, query, idPerwatanTanaman).Scan(
		&data.IDRiwayatPerawatan,
		&data.IDStaff,
		&data.IDBudidaya,
		&data.IDBahanPerawatan,
		&data.JumlahBahan,
		&data.BiayaPerawatan,
		&data.TanggalPerawatan,
		&data.JenisPerawatan,
		&data.NamaKomoditas,
		&data.NamaBahan,
		&data.NamaStaff,
		&data.NamaKebun,
	)

	if err == sql.ErrNoRows {
		err = nil
	}

	return
}

func (r *perawatanTanaman) GetAllPerawatanTanaman(ctx context.Context, filter model.FilterPerawatanTanaman, page, pageSize int) (data model.AllPerawatanTanaman, err error) {
	var (
		filters []string
		args    []interface{}
		argID   = 1
	)

	// Filter by kebun_id requires joining with Budidaya

	if filter.IDKebun != 0 {
		filters = append(filters, fmt.Sprintf("b.id_kebun = $%d", argID))
		args = append(args, filter.IDKebun)
		argID++
	}

	if filter.IDStaff != 0 {
		filters = append(filters, fmt.Sprintf("p.id_staff = $%d", argID))
		args = append(args, filter.IDStaff)
		argID++
	}

	if filter.IDBudidaya != 0 {
		filters = append(filters, fmt.Sprintf("p.id_budidaya = $%d", argID))
		args = append(args, filter.IDBudidaya)
		argID++
	}

	whereClause := ""
	if len(filters) > 0 {
		whereClause = "WHERE " + strings.Join(filters, " AND ")
	}

	// Always join to get names
	joinClause := `
		JOIN budidaya b ON p.id_budidaya = b.id_budidaya
		JOIN komoditas k ON b.id_komoditas = k.id_komoditas
		JOIN kebun kb ON b.id_kebun = kb.id_kebun
		LEFT JOIN bahan_perawatan bp ON p.id_bahan_perawatan = bp.id_bahan
		LEFT JOIN staff_kebun s ON p.id_staff = s.id_staff
	`

	var totalCount int
	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM perawatan_tanaman p %s %s`, joinClause, whereClause)
	err = r.db.QueryRowContext(ctx, countQuery, args...).Scan(&totalCount)
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
		p.id_riwayat_perawatan, 
		p.id_staff, 
		p.id_budidaya, 
		p.id_bahan_perawatan, 
		p.jumlah_bahan, 
		p.biaya_perawatan, 
		p.tanggal_perawatan,
		p.jenis_perawatan,
		k.nama_komoditas,
		COALESCE(bp.nama_bahan, '') as nama_bahan,
		COALESCE(s.nama_staff, '') as nama_staff,
		kb.nama_kebun
	FROM perawatan_tanaman p
	%s
	%s
	ORDER BY p.id_riwayat_perawatan DESC
	LIMIT $%d
	OFFSET $%d
	`, joinClause, whereClause, argID, argID+1)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		return data, err
	}
	defer rows.Close()

	var result []model.PerawatanTanaman

	for rows.Next() {
		var item model.PerawatanTanaman
		err := rows.Scan(
			&item.IDRiwayatPerawatan,
			&item.IDStaff,
			&item.IDBudidaya,
			&item.IDBahanPerawatan,
			&item.JumlahBahan,
			&item.BiayaPerawatan,
			&item.TanggalPerawatan,
			&item.JenisPerawatan,
			&item.NamaKomoditas,
			&item.NamaBahan,
			&item.NamaStaff,
			&item.NamaKebun,
		)
		if err != nil {
			return data, err
		}
		result = append(result, item)
	}

	data.Data = result
	data.TotalCount = totalCount
	data.Page = page
	data.PageSize = pageSize

	return data, nil
}

func (r *perawatanTanaman) UpdatePerawatanTanaman(ctx context.Context, id int, data model.PerawatanTanaman) (result model.PerawatanTanaman, err error) {
	query := `
	UPDATE perawatan_tanaman
	SET
		id_staff = $1,
		id_budidaya = $2,
		id_bahan_perawatan = $3,
		jumlah_bahan = $4,
		biaya_perawatan = $5,
		tanggal_perawatan = $6,
		jenis_perawatan = $7
	WHERE id_riwayat_perawatan = $8
	RETURNING
		id_riwayat_perawatan, id_staff, id_budidaya, id_bahan_perawatan, jumlah_bahan, biaya_perawatan, tanggal_perawatan, jenis_perawatan
	`

	err = r.db.QueryRowContext(ctx, query,
		data.IDStaff,
		data.IDBudidaya,
		data.IDBahanPerawatan,
		data.JumlahBahan,
		data.BiayaPerawatan,
		data.TanggalPerawatan,
		data.JenisPerawatan,
		id,
	).Scan(
		&result.IDRiwayatPerawatan,
		&result.IDStaff,
		&result.IDBudidaya,
		&result.IDBahanPerawatan,
		&result.JumlahBahan,
		&result.BiayaPerawatan,
		&result.TanggalPerawatan,
		&result.JenisPerawatan,
	)

	return
}
