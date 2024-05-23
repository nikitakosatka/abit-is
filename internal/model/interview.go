package model

type Interview struct {
	ID int `json:"interview_id"`
	InterviewData
}

type InterviewData struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}
