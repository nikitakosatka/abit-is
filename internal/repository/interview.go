package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"

	"abitis/internal/model"
)

type InterviewRepository struct {
	pool *pgxpool.Pool
}

func NewInterviewRepository(dbPool *pgxpool.Pool) *InterviewRepository {
	return &InterviewRepository{
		pool: dbPool,
	}
}

func (r *InterviewRepository) CreateInterview(
	ctx context.Context, s *model.InterviewData,
) (*model.Interview, error) {
	var id int
	if err := r.pool.QueryRow(
		ctx, `INSERT INTO interview (title, text) VALUES ($1, $2) RETURNING interview_id`, s.Title, s.Text,
	).Scan(&id); err != nil {
		return nil, fmt.Errorf("scan id: %w", err)
	}

	return &model.Interview{
		ID: id,
		InterviewData: model.InterviewData{
			Title: s.Title,
			Text:  s.Text,
		},
	}, nil
}

func (r *InterviewRepository) ListInterviews(
	ctx context.Context,
) ([]model.Interview, error) {
	query := `SELECT interview_id, title, text FROM interview`
	rows, err := r.pool.Query(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("execute query: %w", err)
	}
	defer rows.Close()

	interviews := make([]model.Interview, 0)
	for rows.Next() {
		var interview model.Interview
		if err := rows.Scan(
			&interview.ID, &interview.Title, &interview.Text,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}
		interviews = append(interviews, interview)
	}
	return interviews, nil
}

func (r *InterviewRepository) GetInterview(
	ctx context.Context, id int,
) (*model.Interview, error) {
	var interview model.Interview
	query := `SELECT interview_id, title, text FROM interview WHERE interview_id = $1`
	if err := r.pool.QueryRow(ctx, query, id).Scan(
		&interview.ID, &interview.Title, &interview.Text,
	); err != nil {
		return nil, fmt.Errorf("scan row: %w", err)
	}
	return &interview, nil
}

func (r *InterviewRepository) UpdateInterview(
	ctx context.Context, interview *model.Interview,
) (*model.Interview, error) {
	query := `UPDATE interview SET title = $1, text = $2 WHERE interview_id = $3 RETURNING interview_id`
	if err := r.pool.QueryRow(
		ctx, query, interview.Title, interview.Text, interview.ID,
	).Scan(&interview.ID); err != nil {
		return nil, fmt.Errorf("scan row: %w", err)
	}
	return interview, nil
}

func (r *InterviewRepository) DeleteInterview(
	ctx context.Context, id int,
) error {
	query := `DELETE FROM interview WHERE interview_id = $1`
	if _, err := r.pool.Exec(ctx, query, id); err != nil {
		return fmt.Errorf("execute query: %w", err)
	}
	return nil
}
