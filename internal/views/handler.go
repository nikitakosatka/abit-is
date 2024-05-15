package views

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"

	"abitis/internal/model"
)

type ViewHandler struct {
	tmpl *template.Template
}

func NewHandler(tmplPattern string) (*ViewHandler, error) {
	tmpl, err := template.ParseGlob(tmplPattern)
	if err != nil {
		return nil, fmt.Errorf("parse template: %w", err)
	}
	return &ViewHandler{
		tmpl: tmpl,
	}, nil
}

func (v *ViewHandler) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		resp, err := http.Get("http://localhost:8080/api/v1/study_plan")
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("Ошибка при формировании страницы: %v", err),
				http.StatusInternalServerError,
			)
			return
		}
		defer resp.Body.Close()
		var plan model.StudyPlan
		if err := json.NewDecoder(resp.Body).Decode(&plan); err != nil {
			http.Error(
				w,
				fmt.Sprintf("Ошибка при формировании страницы: %v", err),
				http.StatusInternalServerError,
			)
			return
		}
		data := map[string]interface{}{
			"Title":          "Главная",
			"About":          plan.Description,
			"Form":           plan.EducationForm,
			"Years":          plan.Years,
			"Cost":           plan.Cost,
			"ServerLoadTime": time.Since(startTime).Milliseconds(),
		}
		if err := v.tmpl.ExecuteTemplate(
			w, "index_page", data,
		); err != nil {
			http.Error(
				w,
				fmt.Sprintf("Ошибка при формировании страницы: %v", err),
				http.StatusInternalServerError,
			)
			return
		}
	}
}

func (v *ViewHandler) StudyPlanHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		data := map[string]interface{}{
			"Title":          "Учебный план",
			"ServerLoadTime": time.Since(startTime).Milliseconds(),
		}
		if err := v.tmpl.ExecuteTemplate(
			w, "study_plan_page", data,
		); err != nil {
			http.Error(
				w,
				fmt.Sprintf("Ошибка при формировании страницы: %v", err),
				http.StatusInternalServerError,
			)
			return
		}
	}
}

func (v *ViewHandler) ChatHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		data := map[string]interface{}{
			"Title":          "Чат",
			"ServerLoadTime": time.Since(startTime).Milliseconds(),
		}
		if err := v.tmpl.ExecuteTemplate(
			w, "chat_page", data,
		); err != nil {
			http.Error(
				w,
				fmt.Sprintf("Ошибка при формировании страницы: %v", err),
				http.StatusInternalServerError,
			)
			return
		}
	}
}

func (v *ViewHandler) InterviewsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		data := map[string]interface{}{
			"Title":          "Интервью",
			"ServerLoadTime": time.Since(startTime).Milliseconds(),
		}
		if err := v.tmpl.ExecuteTemplate(
			w, "albums_page", data,
		); err != nil {
			http.Error(
				w,
				fmt.Sprintf("Ошибка при формировании страницы: %v", err),
				http.StatusInternalServerError,
			)
			return
		}
	}
}

func (v *ViewHandler) InterviewHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		id := chi.URLParam(r, "id")
		resp, err := http.Get(
			fmt.Sprintf("http://localhost:8080/api/v1/interview/%v", id),
		)
		if err != nil {
			http.Error(
				w,
				fmt.Sprintf("Ошибка при формировании страницы: %v", err),
				http.StatusInternalServerError,
			)
			return
		}
		defer resp.Body.Close()
		if resp.StatusCode == http.StatusNotFound {
			http.Error(w, "Интервью не найдено", http.StatusNotFound)
		}
		var interview model.Interview
		if err := json.NewDecoder(resp.Body).Decode(&interview); err != nil {
			http.Error(
				w,
				fmt.Sprintf("Ошибка при формировании страницы: %v", err),
				http.StatusInternalServerError,
			)
			return
		}
		data := map[string]interface{}{
			"Interview":      interview,
			"ServerLoadTime": time.Since(startTime).Milliseconds(),
		}
		if err := v.tmpl.ExecuteTemplate(
			w, "interview_page", data,
		); err != nil {
			http.Error(
				w,
				fmt.Sprintf("Ошибка при формировании страницы: %v", err),
				http.StatusInternalServerError,
			)
			return
		}
	}
}

func (v *ViewHandler) SignUpHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		data := map[string]interface{}{
			"Title":          "Регистрация",
			"ServerLoadTime": time.Since(startTime).Milliseconds(),
		}
		if err := v.tmpl.ExecuteTemplate(
			w, "signup_page", data,
		); err != nil {
			http.Error(
				w,
				fmt.Sprintf("Ошибка при формировании страницы: %v", err),
				http.StatusInternalServerError,
			)
			return
		}
	}
}

func (v *ViewHandler) LogInHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		data := map[string]interface{}{
			"Title":          "Вход",
			"ServerLoadTime": time.Since(startTime).Milliseconds(),
		}
		if err := v.tmpl.ExecuteTemplate(
			w, "login_page", data,
		); err != nil {
			http.Error(
				w,
				fmt.Sprintf("Ошибка при формировании страницы: %v", err),
				http.StatusInternalServerError,
			)
			return
		}
	}
}
