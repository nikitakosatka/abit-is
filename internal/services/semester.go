package services

import (
	"context"
	"fmt"

	"abitis/internal/model"
	"abitis/internal/repository"
)

type SemesterService struct {
	repository *repository.SemesterRepository
}

func NewSemesterService(
	repository *repository.SemesterRepository,
) *SemesterService {
	return &SemesterService{
		repository: repository,
	}
}

func (s *SemesterService) CreateSemester(
	ctx context.Context, sem *model.Semester,
) error {
	if err := s.repository.CreateSemester(ctx, sem); err != nil {
		return fmt.Errorf("create semester: %w", err)
	}
	return nil
}

func (s *SemesterService) ListSemesters(
	ctx context.Context,
) ([]model.Semester, error) {
	semesters, err := s.repository.ListSemesters(ctx)
	if err != nil {
		return nil, fmt.Errorf("list semesters: %w", err)
	}
	return semesters, nil
}

func (s *SemesterService) GetSemester(
	ctx context.Context, num int,
) (*model.Semester, error) {
	semester, err := s.repository.GetSemester(ctx, num)
	if err != nil {
		return nil, fmt.Errorf("get semester: %w", err)
	}
	return semester, nil
}

func (s *SemesterService) UpdateSemester(
	ctx context.Context, sem *model.Semester,
) (*model.Semester, error) {
	semester, err := s.repository.UpdateSemester(ctx, sem)
	if err != nil {
		return nil, fmt.Errorf("update semester: %w", err)
	}
	return semester, nil
}

func (s *SemesterService) DeleteSemester(
	ctx context.Context, num int,
) error {
	if err := s.repository.DeleteSemester(ctx, num); err != nil {
		return fmt.Errorf("delete semester: %w", err)
	}
	return nil
}
