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

type SemesterController struct {
	service *services.SemesterService
}

func NewSemesterController(
	service *services.SemesterService,
) *SemesterController {
	return &SemesterController{
		service: service,
	}
}

// CreateSemester godoc
// @Summary Create a new semester
// @Description Add a new semester to the database
// @Tags semester
// @Accept json
// @Produce json
// @Param semester body model.Semester true "Add semester"
// @Success 201 {object} nil
// @Failure 400 {object} nil "Bad request"
// @Failure 500 {object} nil "Server error"
// @Router /api/v1/semester [post]
func (c *SemesterController) CreateSemester(
	w http.ResponseWriter, r *http.Request,
) {
	var semester model.Semester
	if err := json.NewDecoder(r.Body).Decode(&semester); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := c.service.CreateSemester(r.Context(), &semester); err != nil {
		http.Error(
			w, "Failed to create semester", http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

// ListSemesters godoc
// @Summary List all semesters
// @Description get a list of all semesters
// @Tags semester
// @Produce json
// @Success 200 {array} model.Semester
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/semester [get]
func (c *SemesterController) ListSemesters(
	w http.ResponseWriter, r *http.Request,
) {
	semesters, err := c.service.ListSemesters(r.Context())
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to list semesters: %v", err),
			http.StatusInternalServerError,
		)
		return
	}
	if err := json.NewEncoder(w).Encode(semesters); err != nil {
		http.Error(
			w, "Failed to encode semesters", http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// GetSemester godoc
// @Summary Get details of a semester
// @Description Get details of a semester by semester number
// @Tags semester
// @Produce json
// @Param num path int true "Semester Number"
// @Success 200 {object} model.Semester
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/semester/{num} [get]
func (c *SemesterController) GetSemester(
	w http.ResponseWriter, r *http.Request,
) {
	numValue := chi.URLParam(r, "num")
	num, err := strconv.Atoi(numValue)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to convert num: %v", err),
			http.StatusBadRequest,
		)
		return
	}
	semester, err := c.service.GetSemester(r.Context(), num)
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

// UpdateSemester godoc
// @Summary Update an existing semester
// @Description Update an existing semester by semester number
// @Tags semester
// @Accept json
// @Produce json
// @Param semester body model.Semester true "Semester object"
// @Success 200 {object} model.Semester
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/semester [put]
func (c *SemesterController) UpdateSemester(
	w http.ResponseWriter, r *http.Request,
) {
	var semester model.Semester
	if err := json.NewDecoder(r.Body).Decode(&semester); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := c.service.UpdateSemester(r.Context(), &semester); err != nil {
		http.Error(
			w, "Failed to create semester", http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(semester); err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to get semester: %v", err),
			http.StatusInternalServerError,
		)
		return
	}
}

// DeleteSemester godoc
// @Summary Delete a semester
// @Description Delete a semester by semester number
// @Tags semester
// @Produce json
// @Param num path int true "Semester Number"
// @Success 204 "Deleted successfully"
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/semester/{num} [delete]
func (c *SemesterController) DeleteSemester(
	w http.ResponseWriter, r *http.Request,
) {
	numValue := chi.URLParam(r, "num")
	num, err := strconv.Atoi(numValue)
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to convert num: %v", err),
			http.StatusBadRequest,
		)
		return
	}
	if err := c.service.DeleteSemester(r.Context(), num); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, "Semester not found", http.StatusNotFound)
			return
		}
		http.Error(
			w, "Failed to delete semester", http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
