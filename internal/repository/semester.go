package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"abitis/internal/model"
)

type SemesterRepository struct {
	pool *pgxpool.Pool
}

func NewSemesterRepository(dbPool *pgxpool.Pool) *SemesterRepository {
	return &SemesterRepository{
		pool: dbPool,
	}
}

func (r *SemesterRepository) CreateSemester(
	ctx context.Context, s *model.Semester,
) error {
	query := `INSERT INTO Semester (semester_num, season) VALUES ($1, $2)`
	if _, err := r.pool.Exec(
		ctx, query, s.SemesterNum, s.Season,
	); err != nil {
		return fmt.Errorf("execute query: %w", err)
	}
	return nil
}

func (r *SemesterRepository) ListSemesters(
	ctx context.Context,
) ([]model.Semester, error) {
	query := `SELECT semester_num, season FROM semester`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("execute query: %w", err)
	}
	defer rows.Close()

	semesters := make([]model.Semester, 0)
	for rows.Next() {
		var sem model.Semester
		if err := rows.Scan(&sem.SemesterNum, &sem.Season); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}
		semesters = append(semesters, sem)
	}
	return semesters, nil
}

func (r *SemesterRepository) GetSemester(
	ctx context.Context, num int,
) (*model.Semester, error) {
	var semester model.Semester
	query := `SELECT semester_num, season FROM semester WHERE semester_num = $1`
	if err := r.pool.QueryRow(ctx, query, num).Scan(
		&semester.SemesterNum, &semester.Season,
	); err != nil {
		return nil, fmt.Errorf("scan row: %w", err)
	}
	return &semester, nil
}

func (r *SemesterRepository) UpdateSemester(
	ctx context.Context, sem *model.Semester,
) (*model.Semester, error) {
	query := `UPDATE semester SET season = $1 WHERE semester_num = $2 RETURNING semester_num`
	if err := r.pool.QueryRow(
		ctx, query, sem.Season, sem.SemesterNum,
	).Scan(&sem.SemesterNum); err != nil {
		return nil, fmt.Errorf("scan row: %w", err)
	}
	return sem, nil
}

func (r *SemesterRepository) DeleteSemester(
	ctx context.Context, num int,
) error {
	query := `DELETE FROM semester WHERE semester_num = $1`
	if _, err := r.pool.Exec(ctx, query, num); err != nil {
		return fmt.Errorf("execute query: %w", err)
	}
	return nil
}
