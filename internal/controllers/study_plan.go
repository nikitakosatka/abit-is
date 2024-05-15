package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"abitis/internal/model"
	"abitis/internal/services"
)

type StudyPlanController struct {
	service *services.StudyPlanService
}

func NewStudyPlanController(
	service *services.StudyPlanService,
) *StudyPlanController {
	return &StudyPlanController{
		service: service,
	}
}

// GetStudyPlan godoc
// @Summary Get current study plan
// @Description get current study plan
// @Tags study_plan
// @Produce json
// @Success 200 {object} model.StudyPlan
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/study_plan [get]
func (c *StudyPlanController) GetStudyPlan(
	w http.ResponseWriter, r *http.Request,
) {
	studyPlan, err := c.service.GetStudyPlan(r.Context())
	if err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to get study plan: %v", err),
			http.StatusInternalServerError,
		)
		return
	}
	if err := json.NewEncoder(w).Encode(studyPlan); err != nil {
		http.Error(
			w, "Failed to encode study plan", http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// UpdateStudyPlan godoc
// @Summary Update current study plan
// @Description Update an existing study plan
// @Tags study_plan
// @Accept json
// @Produce json
// @Param study_plan body model.StudyPlan true "Study Plan"
// @Success 200 {object} model.StudyPlan
// @Failure 400 {string} string "Bad request"
// @Failure 404 {string} string "Not Found"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/v1/study_plan [put]
func (c *StudyPlanController) UpdateStudyPlan(
	w http.ResponseWriter, r *http.Request,
) {
	var studyPlan model.StudyPlan
	if err := json.NewDecoder(r.Body).Decode(&studyPlan); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, err := c.service.UpdateStudyPlan(
		r.Context(), &studyPlan,
	); err != nil {
		http.Error(
			w, "Failed to update study plan", http.StatusInternalServerError,
		)
		return
	}
	if err := json.NewEncoder(w).Encode(&studyPlan); err != nil {
		http.Error(
			w,
			fmt.Sprintf("Failed to get semester: %v", err),
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusOK)
}
