package services

import (
	"context"
	"fmt"

	"firebase.google.com/go"
	"firebase.google.com/go/auth"

	"abitis/internal/schema"
)

type AuthService struct {
	app *firebase.App
}

func NewAuthService(app *firebase.App) *AuthService {
	return &AuthService{
		app: app,
	}
}

func (s *AuthService) SignUp(ctx context.Context, user *schema.User) (*schema.UID, error) {
	client, err := s.app.Auth(ctx)
	if err != nil {
		return nil, fmt.Errorf("get auth client: %w", err)
	}
	params := (&auth.UserToCreate{}).
		Email(user.Email).
		EmailVerified(false).
		Password(user.Password).
		Disabled(false)

	createdUser, err := client.CreateUser(ctx, params)
	if err != nil {
		return nil, fmt.Errorf("create user: %w", err)
	}
	return &schema.UID{
		ID: createdUser.UID,
	}, nil
}
