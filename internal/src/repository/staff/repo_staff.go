package staff

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/plantation-service/internal/src/model"
)

func (r *staff) InsertStaff(ctx context.Context, data model.StaffKebun) (result model.StaffKebun, err error) {
	query := `INSERT INTO staff_kebun 
	(nama_staff, 
	jabatan, 
	kontak) 
	VALUES 
	($1, $2, $3) 
	RETURNING 
	id_staff, nama_staff, jabatan, kontak`

	err = r.plantationDb.QueryRowContext(ctx, query, data.NamaStaff, data.Jabatan, data.Kontak).Scan(
		&result.ID,
		&result.NamaStaff,
		&result.Jabatan,
		&result.Kontak,
	)

	return
}

func (r *staff) GetByID(ctx context.Context, idStaff int) (data model.StaffKebun, err error) {
	query := `SELECT id_staff, 
	nama_staff, 
	jabatan, 
	kontak 
	FROM staff_kebun
	WHERE id_staff = $1`

	err = r.plantationDb.QueryRowContext(ctx, query, idStaff).Scan(
		&data.ID,
		&data.NamaStaff,
		&data.Jabatan,
		&data.Kontak,
	)

	if err == sql.ErrNoRows {
		err = nil
	}

	return data, err
}

func (r *staff) GetAllStaff(ctx context.Context,
	filter model.StaffFilter,
	page, pageSize int) (data model.AllStaffKebun, err error) {

	var (
		filters []string
		args    []interface{}
		argID   = 1
	)

	if filter.NamaStaff != "" {
		filters = append(filters, fmt.Sprintf("LOWER(nama_staff) ILIKE '%%' || $%d || '%%'", argID))
		args = append(args, strings.ToLower(filter.NamaStaff))
		argID++
	}

	if filter.Jabatan != "" {
		filters = append(filters, fmt.Sprintf("jabatan = $%d", argID))
		args = append(args, filter.Jabatan)
		argID++
	}

	whereClause := ""
	if len(filters) > 0 {
		whereClause = "WHERE " + strings.Join(filters, " AND ")
	}

	var totalCount int
	countQuery := fmt.Sprintf(`SELECT COUNT(*) FROM staff_kebun %s`, whereClause)
	err = r.plantationDb.QueryRowContext(ctx, countQuery, args...).Scan(&totalCount)
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
		id_staff, 
		nama_staff, 
		jabatan, 
		kontak
	FROM staff_kebun
	%s
	ORDER BY id_staff DESC
	LIMIT $%d
	OFFSET $%d
	`, whereClause, argID, argID+1)

	rows, err := r.plantationDb.QueryContext(ctx, query, args...)
	if err != nil {
		return data, err
	}

	var result []model.StaffKebun
	defer rows.Close()

	for rows.Next() {
		var staff model.StaffKebun
		err := rows.Scan(
			&staff.ID,
			&staff.NamaStaff,
			&staff.Jabatan,
			&staff.Kontak,
		)
		if err != nil {
			return data, err
		}
		result = append(result, staff)
	}

	data.Data = result
	data.TotalCount = totalCount
	data.Page = page
	data.PageSize = pageSize

	return data, nil
}

func (r *staff) UpdateStaff(ctx context.Context, idStaff int, data model.StaffKebun) (result model.StaffKebun, err error) {
	query := `UPDATE staff_kebun SET 
	nama_staff = $1,
	jabatan = $2,
	kontak = $3
	WHERE id_staff = $4
	RETURNING 
	id_staff, nama_staff, jabatan, kontak`

	err = r.plantationDb.QueryRowContext(ctx, query,
		data.NamaStaff,
		data.Jabatan,
		data.Kontak,
		idStaff,
	).Scan(
		&result.ID,
		&result.NamaStaff,
		&result.Jabatan,
		&result.Kontak,
	)

	return
}

func (r *staff) DeleteStaff(ctx context.Context, idStaff int) (err error) {
	query := `DELETE FROM staff_kebun WHERE id_staff = $1`

	_, err = r.plantationDb.ExecContext(ctx, query, idStaff)
	if err != nil {
		return err
	}

	return nil
}
