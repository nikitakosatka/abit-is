package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"abitis/internal/model"
)

type StudyPlanRepository struct {
	pool *pgxpool.Pool
}

func NewStudyPlanRepository(dbPool *pgxpool.Pool) *StudyPlanRepository {
	return &StudyPlanRepository{
		pool: dbPool,
	}
}

func (r *StudyPlanRepository) GetStudyPlan(
	ctx context.Context,
) (*model.StudyPlan, error) {
	var studyPlan model.StudyPlan
	query := `SELECT id, name, description, education_form, cost, years FROM study_plan WHERE id = 1`
	if err := r.pool.QueryRow(ctx, query).Scan(
		&studyPlan.ID,
		&studyPlan.Name,
		&studyPlan.Description,
		&studyPlan.EducationForm,
		&studyPlan.Cost,
		&studyPlan.Years,
	); err != nil {
		return nil, fmt.Errorf("scan row: %w", err)
	}
	return &studyPlan, nil
}

func (r *StudyPlanRepository) UpdateStudyPlan(
	ctx context.Context, studyPlan *model.StudyPlan,
) (*model.StudyPlan, error) {
	query := `UPDATE study_plan SET name = $1, description = $2, education_form = $3, cost = $4, years = $5 WHERE id = $6 RETURNING id`
	if err := r.pool.QueryRow(
		ctx,
		query,
		studyPlan.Name,
		studyPlan.Description,
		studyPlan.EducationForm,
		studyPlan.Cost,
		studyPlan.Years,
		studyPlan.ID,
	).Scan(&studyPlan.ID); err != nil {
		return nil, fmt.Errorf("scan row: %w", err)
	}
	return studyPlan, nil
}
