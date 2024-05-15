package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v4"

	"abitis/internal/model"
	"abitis/internal/services"
)

type InterviewController struct {
	service *services.InterviewService
}

func NewInterviewController(
	service *services.InterviewService,
) *InterviewController {
	return &InterviewController{
		service: service,
	}
}

// CreateInterview godoc
// @Summary Create a new interview
// @Description Add a new interview to the database
// @Tags interview
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param interview body model.InterviewData true "Add interview"
// @Success 201 {object} nil
// @Failure 400 {object} nil "Bad request"
// @Failure 500 {object} nil "Server error"
// @Router /api/v1/interview [post]
func (c *InterviewController) CreateInterview(
	w http.ResponseWriter, r *http.Request,
) {
	var interview model.InterviewData
	if err := json.NewDecoder(r.Body).Decode(&interview); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.service.CreateInterview(r.Context(), &interview); err != nil {
		http.Error(
			w, "Failed to create interview", http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// ListInterviews godoc
// @Summary List all interviews
// @Description get a list of all interviews
// @Tags interview
// @Produce json
// @Success 200 {array} model.Interview
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/interview [get]
func (c *InterviewController) ListInterviews(
	w http.ResponseWriter, r *http.Request,
) {
	interviews, err := c.service.ListInterviews(r.Context())
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to list interviews: %v", err),
			http.StatusInternalServerError,
		)
		return
	}
	if err := json.NewEncoder(w).Encode(interviews); err != nil {
		http.Error(
			w, "Failed to encode interviews", http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetInterview godoc
// @Summary Get details of an interview
// @Description Get details of an interview by id
// @Tags interview
// @Produce json
// @Param id path int true "Interview ID"
// @Success 200 {object} model.Interview
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/interview/{id} [get]
func (c *InterviewController) GetInterview(
	w http.ResponseWriter, r *http.Request,
) {
	idValue := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idValue)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to convert id: %v", err),
			http.StatusBadRequest,
		)
		return
	}
	interview, err := c.service.GetInterview(r.Context(), id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Semester not found", http.StatusNotFound)
			return
		}
		http.Error(
			w,
			fmt.Sprintf("Failed to get semester: %v", err),
			http.StatusInternalServerError,
		)
		return
	}
	if err := json.NewEncoder(w).Encode(interview); err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to get semester: %v", err),
			http.StatusInternalServerError,
		)
		return
	}
}

// UpdateInterview godoc
// @Summary Update an existing interview
// @Description Update an existing interview by id
// @Tags interview
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Interview ID"
// @Param semester body model.InterviewData true "Interview data"
// @Success 200 {object} model.Interview
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/interview/{id} [put]
func (c *InterviewController) UpdateInterview(
	w http.ResponseWriter, r *http.Request,
) {
	idValue := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idValue)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to convert id: %v", err),
			http.StatusBadRequest,
		)
		return
	}
	var interview model.InterviewData
	if err := json.NewDecoder(r.Body).Decode(&interview); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := c.service.UpdateInterview(
		r.Context(), id, &interview,
	); err != nil {
		http.Error(
			w, "Failed to update interview", http.StatusInternalServerError,
		)
		return
	}
	if err := json.NewEncoder(w).Encode(&interview); err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to get semester: %v", err),
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// DeleteInterview godoc
// @Summary Delete a interview
// @Description Delete a interview by id
// @Tags interview
// @Produce json
// @Security ApiKeyAuth
// @Param id path int true "Interview ID"
// @Success 204 "Deleted successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/interview/{id} [delete]
func (c *InterviewController) DeleteInterview(
	w http.ResponseWriter, r *http.Request,
) {
	idValue := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idValue)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to convert num: %v", err),
			http.StatusBadRequest,
		)
		return
	}
	if err := c.service.DeleteInterview(r.Context(), id); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Interview not found", http.StatusNotFound)
			return
		}
		http.Error(
			w, "Failed to delete interview", http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
