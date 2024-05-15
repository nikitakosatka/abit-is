package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"abitis/internal/model"
)

type SubjectRepository struct {
	pool *pgxpool.Pool
}

func NewSubjectRepository(dbPool *pgxpool.Pool) *SubjectRepository {
	return &SubjectRepository{
		pool: dbPool,
	}
}

func (r *SubjectRepository) CreateSubject(
	ctx context.Context, s *model.Subject,
) error {
	query := `INSERT INTO subject (name, description, semester, study_plan) VALUES ($1, $2, $3, 1)`
	if _, err := r.pool.Exec(
		ctx, query, s.Name, s.Description, s.SemesterNum,
	); err != nil {
		return fmt.Errorf("execute query: %w", err)
	}
	return nil
}

func (r *SubjectRepository) ListSubjects(
	ctx context.Context,
) ([]model.Subject, error) {
	query := `SELECT name, description, semester FROM subject`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("execute query: %w", err)
	}
	defer rows.Close()

	var subjects []model.Subject
	for rows.Next() {
		var s model.Subject
		if err := rows.Scan(
			&s.Name, &s.Description, &s.SemesterNum,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}
		subjects = append(subjects, s)
	}
	return subjects, nil
}

func (r *SubjectRepository) GetSubject(
	ctx context.Context, name string, semester int,
) (*model.Subject, error) {
	var subject model.Subject
	query := `SELECT name, description, semester FROM subject
                                   WHERE name = $1 and semester = $2`
	if err := r.pool.QueryRow(ctx, query, name, semester).Scan(
		&subject.Name, &subject.Description, &subject.SemesterNum,
	); err != nil {
		return nil, fmt.Errorf("scan row: %w", err)
	}
	return &subject, nil
}

func (r *SubjectRepository) UpdateSubject(
	ctx context.Context, sem *model.Subject,
) (*model.Subject, error) {
	query := `UPDATE subject SET description = $1
               WHERE name = $2 and semester = $3 RETURNING description`
	if err := r.pool.QueryRow(
		ctx, query, sem.Description, sem.Name, sem.SemesterNum,
	).Scan(&sem.Description); err != nil {
		return nil, fmt.Errorf("scan row: %w", err)
	}
	return sem, nil
}

func (r *SubjectRepository) DeleteSubject(
	ctx context.Context, name string, semester int,
) error {
	query := `DELETE FROM subject WHERE name = $1 and semester = $2`
	if _, err := r.pool.Exec(ctx, query, name, semester); err != nil {
		return fmt.Errorf("execute query: %w", err)
	}
	return nil
}
