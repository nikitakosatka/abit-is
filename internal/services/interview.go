package services

import (
	"context"
	"fmt"

	"abitis/internal/model"
	"abitis/internal/repository"
)

type InterviewService struct {
	repository *repository.InterviewRepository
}

func NewInterviewService(
	repository *repository.InterviewRepository,
) *InterviewService {
	return &InterviewService{
		repository: repository,
	}
}

func (s *InterviewService) CreateInterview(
	ctx context.Context, interview *model.InterviewData,
) error {
	if err := s.repository.CreateInterview(ctx, interview); err != nil {
		return fmt.Errorf("create interview: %w", err)
	}
	return nil
}

func (s *InterviewService) ListInterviews(
	ctx context.Context,
) ([]model.Interview, error) {
	interviews, err := s.repository.ListInterviews(ctx)
	if err != nil {
		return nil, fmt.Errorf("list interviews: %w", err)
	}
	return interviews, nil
}

func (s *InterviewService) GetInterview(
	ctx context.Context, id int,
) (*model.Interview, error) {
	interview, err := s.repository.GetInterview(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("get interview: %w", err)
	}
	return interview, nil
}

func (s *InterviewService) UpdateInterview(
	ctx context.Context, id int, interview *model.InterviewData,
) (*model.Interview, error) {
	res, err := s.repository.UpdateInterview(ctx, &model.Interview{
		ID:            id,
		InterviewData: *interview,
	})
	if err != nil {
		return nil, fmt.Errorf("update interview: %w", err)
	}
	return res, nil
}

func (s *InterviewService) DeleteInterview(
	ctx context.Context, id int,
) error {
	if err := s.repository.DeleteInterview(ctx, id); err != nil {
		return fmt.Errorf("delete interview: %w", err)
	}
	return nil
}
