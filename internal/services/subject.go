package services

import (
	"context"
	"fmt"

	"abitis/internal/model"
	"abitis/internal/repository"
)

type SubjectService struct {
	repository *repository.SubjectRepository
}

func NewSubjectService(
	repository *repository.SubjectRepository,
) *SubjectService {
	return &SubjectService{
		repository: repository,
	}
}

func (s *SubjectService) CreateSubject(
	ctx context.Context, subject *model.Subject,
) error {
	if err := s.repository.CreateSubject(ctx, subject); err != nil {
		return fmt.Errorf("create subject: %w", err)
	}
	return nil
}

func (s *SubjectService) ListSubjects(
	ctx context.Context,
) ([]model.Subject, error) {
	subjects, err := s.repository.ListSubjects(ctx)
	if err != nil {
		return nil, fmt.Errorf("list subjects: %w", err)
	}
	return subjects, nil
}

func (s *SubjectService) GetSubject(
	ctx context.Context, name string, semester int,
) (*model.Subject, error) {
	subject, err := s.repository.GetSubject(ctx, name, semester)
	if err != nil {
		return nil, fmt.Errorf("get subject: %w", err)
	}
	return subject, nil
}

func (s *SubjectService) UpdateSubject(
	ctx context.Context, subject *model.Subject,
) (*model.Subject, error) {
	subject, err := s.repository.UpdateSubject(ctx, subject)
	if err != nil {
		return nil, fmt.Errorf("update subject: %w", err)
	}
	return subject, nil
}

func (s *SubjectService) DeleteSubject(
	ctx context.Context, name string, semester int,
) error {
	if err := s.repository.DeleteSubject(ctx, name, semester); err != nil {
		return fmt.Errorf("delete semester: %w", err)
	}
	return nil
}
