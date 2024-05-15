package services

import (
	"context"
	"fmt"

	"abitis/internal/model"
	"abitis/internal/repository"
)

type StudyPlanService struct {
	repository *repository.StudyPlanRepository
}

func NewStudyPlanService(
	repository *repository.StudyPlanRepository,
) *StudyPlanService {
	return &StudyPlanService{
		repository: repository,
	}
}

func (s *StudyPlanService) GetStudyPlan(
	ctx context.Context,
) (*model.StudyPlan, error) {
	res, err := s.repository.GetStudyPlan(ctx)
	if err != nil {
		return nil, fmt.Errorf("get study plan: %w", err)
	}
	return res, nil
}

func (s *StudyPlanService) UpdateStudyPlan(
	ctx context.Context, studyPlan *model.StudyPlan,
) (*model.StudyPlan, error) {
	res, err := s.repository.UpdateStudyPlan(ctx, studyPlan)
	if err != nil {
		return nil, fmt.Errorf("update study plan: %w", err)
	}
	return res, nil
}
