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

type SubjectController struct {
	service *services.SubjectService
}

func NewSubjectController(
	service *services.SubjectService,
) *SubjectController {
	return &SubjectController{
		service: service,
	}
}

// CreateSubject godoc
// @Summary Create a new subject
// @Description Add a new subject to the database
// @Tags subject
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param semester body model.Subject true "Add semester"
// @Success 201 {object} nil
// @Failure 400 {object} nil "Bad request"
// @Failure 500 {object} nil "Server error"
// @Router /api/v1/subject [post]
func (c *SubjectController) CreateSubject(
	w http.ResponseWriter, r *http.Request,
) {
	var subject model.Subject
	if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.service.CreateSubject(r.Context(), &subject); err != nil {
		http.Error(
			w, "Failed to create subject", http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// ListSubjects godoc
// @Summary List all subjects
// @Description get a list of all subjects
// @Tags subject
// @Produce json
// @Success 200 {array} model.Subject
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/subject [get]
func (c *SubjectController) ListSubjects(
	w http.ResponseWriter, r *http.Request,
) {
	subjects, err := c.service.ListSubjects(r.Context())
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to list semesters: %v", err),
			http.StatusInternalServerError,
		)
		return
	}
	if err := json.NewEncoder(w).Encode(subjects); err != nil {
		http.Error(
			w, "Failed to encode subjects", http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetSubject godoc
// @Summary Get details of a subject
// @Description Get details of a subject by semester number
// @Tags subject
// @Produce json
// @Param name path string true "Subject Name"
// @Param semester_num path int true "Semester Number"
// @Success 200 {object} model.Subject
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/subject/{name}/{semester_num} [get]
func (c *SubjectController) GetSubject(
	w http.ResponseWriter, r *http.Request,
) {
	name := chi.URLParam(r, "name")
	numValue := chi.URLParam(r, "semester_num")
	num, err := strconv.Atoi(numValue)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to convert num: %v", err),
			http.StatusBadRequest,
		)
		return
	}
	semester, err := c.service.GetSubject(r.Context(), name, num)
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
	if err := json.NewEncoder(w).Encode(semester); err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to get semester: %v", err),
			http.StatusInternalServerError,
		)
		return
	}
}

// UpdateSubject godoc
// @Summary Update an existing subject
// @Description Update an existing subject by name and semester
// @Tags subject
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param subject body model.Subject true "Subject object"
// @Success 200 {object} model.Subject
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/subject [put]
func (c *SubjectController) UpdateSubject(
	w http.ResponseWriter, r *http.Request,
) {
	var subject model.Subject
	if err := json.NewDecoder(r.Body).Decode(&subject); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := c.service.UpdateSubject(r.Context(), &subject); err != nil {
		http.Error(
			w, "Failed to create subject", http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(subject); err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to get semester: %v", err),
			http.StatusInternalServerError,
		)
		return
	}
}

// DeleteSubject godoc
// @Summary Delete a subject
// @Description Delete a subject by name and semester
// @Tags subject
// @Produce json
// @Security ApiKeyAuth
// @Param name path string true "Name"
// @Param semester_num path int true "Semester Number"
// @Success 204 "Deleted successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/subject/{name}/{semester_num} [delete]
func (c *SubjectController) DeleteSubject(
	w http.ResponseWriter, r *http.Request,
) {
	name := chi.URLParam(r, "name")
	numValue := chi.URLParam(r, "semester_num")
	num, err := strconv.Atoi(numValue)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to convert num: %v", err),
			http.StatusBadRequest,
		)
		return
	}
	if err := c.service.DeleteSubject(r.Context(), name, num); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Subject not found", http.StatusNotFound)
			return
		}
		http.Error(
			w, "Failed to delete subject", http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
