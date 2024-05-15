package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"abitis/internal/schema"
	"abitis/internal/services"
)

type AuthController struct {
	service *services.AuthService
}

func NewAuthController(
	service *services.AuthService,
) *AuthController {
	return &AuthController{
		service: service,
	}
}

// SignUp godoc
// @Summary Register user
// @Description Register a new user using email and password.
// @Tags auth
// @Accept json
// @Produce json
// @Param user body schema.User true "User info"
// @Success 200 {object} schema.UID
// @Failure 400 {string} string "Invalid input"
// @Failure 500 {string} string "Internal server error"
// @Router /api/v1/signup [post]
func (c *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	var user schema.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(
			w,
			fmt.Sprintf("decode user: %v", err),
			http.StatusBadRequest,
		)
	}
	uid, err := c.service.SignUp(r.Context(), &user)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("sign up: %v", err),
			http.StatusInternalServerError,
		)
		return
	}
	if err := json.NewEncoder(w).Encode(uid); err != nil {
		http.Error(
			w,
			fmt.Sprintf("encode uid: %v", err),
			http.StatusInternalServerError,
		)
		return
	}
}
