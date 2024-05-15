package model

import "github.com/shopspring/decimal"

type StudyPlan struct {
	ID            int             `json:"id"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	EducationForm EduForm         `json:"education_form"`
	Cost          decimal.Decimal `json:"cost"`
	Years         int             `json:"years"`
}
