package model

type Semester struct {
	SemesterNum int    `json:"semester_num"`
	Season      Season `json:"season"`
}
