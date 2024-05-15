package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"firebase.google.com/go"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/pflag"
	"github.com/swaggo/http-swagger"
	"google.golang.org/api/option"

	_ "abitis/docs"
	"abitis/internal/controllers"
	mw "abitis/internal/middleware"
	"abitis/internal/repository"
	"abitis/internal/services"
	"abitis/internal/views"
)

const (
	DefaultAppPort = 8080
)

const (
	DefaultDBPortEnv     = "POSTGRES_PORT"
	DefaultDBNameEnv     = "POSTGRES_DB_NAME"
	DefaultDBUserEnv     = "POSTGRES_USER"
	DefaultDBHostEnv     = "POSTGRES_HOST"
	DefaultDBPasswordEnv = "POSTGRES_PASSWORD"
)

// @title Abit IS
// @description Abit service.
// @version 1.0
// @BasePath /
// @schemes https
// @securityDefinitions.apikey BearerAuth
// @In header
// @Name Authorization
func main() {
	if err := run(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}

func run() error {
	fs := pflag.NewFlagSet("abit-is", pflag.ContinueOnError)

	appPort := fs.Int(
		"port",
		DefaultAppPort,
		"port to run app on",
	)
	dbPortEnv := fs.String(
		"db-port-env",
		DefaultDBPortEnv,
		"environment variable to get database port",
	)
	dbNameEnv := fs.String(
		"db-name-env",
		DefaultDBNameEnv,
		"environment variable to get database name",
	)
	dbUserEnv := fs.String(
		"db-user-env",
		DefaultDBUserEnv,
		"environment variable to get database user",
	)
	dbHostEnv := fs.String(
		"db-host-env",
		DefaultDBHostEnv,
		"environment variable to get database host",
	)
	dbPasswordEnv := fs.String(
		"db-password-env",
		DefaultDBPasswordEnv,
		"environment variable to get database password",
	)
	if err := fs.Parse(os.Args[1:]); err != nil {
		if err == pflag.ErrHelp {
			return nil
		}
		return err
	}

	ctx := context.Background()
	pool, err := pgxpool.Connect(ctx, fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		os.Getenv(*dbUserEnv),
		url.QueryEscape(os.Getenv(*dbPasswordEnv)),
		os.Getenv(*dbHostEnv),
		os.Getenv(*dbPortEnv),
		os.Getenv(*dbNameEnv),
	))
	if err != nil {
		return fmt.Errorf("connect db: %w", err)
	}

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(
		os.Getenv("FIREBASE_CREDENTIALS_FILEPATH"),
	))
	if err != nil {
		return fmt.Errorf("init Firebase app: %w", err)
	}
	authClient, err := app.Auth(ctx)
	if err != nil {
		return fmt.Errorf("get auth client: %w", err)
	}

	semesterController := controllers.NewSemesterController(
		services.NewSemesterService(
			repository.NewSemesterRepository(pool),
		),
	)
	subjectController := controllers.NewSubjectController(
		services.NewSubjectService(
			repository.NewSubjectRepository(pool),
		),
	)
	interviewController := controllers.NewInterviewController(
		services.NewInterviewService(
			repository.NewInterviewRepository(pool),
		),
	)
	studyPlanController := controllers.NewStudyPlanController(
		services.NewStudyPlanService(
			repository.NewStudyPlanRepository(pool),
		),
	)
	authController := controllers.NewAuthController(
		services.NewAuthService(app),
	)
	ws := services.NewWSService(
		websocket.Upgrader{
			ReadBufferSize:  services.ReadBufferSize,
			WriteBufferSize: services.WriteBufferSize,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	)

	handler, err := views.NewHandler("templates/*.html")
	if err != nil {
		return fmt.Errorf("init views: %w", err)
	}

	r.Handle("/static/*", http.StripPrefix(
		"/static/", http.FileServer(http.Dir("./static")),
	))
	r.Get("/", handler.IndexHandler())
	r.Get("/study_plan", handler.StudyPlanHandler())
	r.Get("/chat", handler.ChatHandler())
	r.Get("/interviews", handler.InterviewsHandler())
	r.Get("/interview/{id}", handler.InterviewHandler())
	r.Get("/signup", handler.SignUpHandler())
	r.Get("/login", handler.LogInHandler())

	r.Get("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "robots.txt")
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/swagger/*", httpSwagger.WrapHandler)
			r.Post("/signup", authController.SignUp)
			r.Route("/semester", func(r chi.Router) {
				r.Post("/", semesterController.CreateSemester)
				r.Get("/", semesterController.ListSemesters)
				r.Get("/{num}", semesterController.GetSemester)
				r.Put("/", semesterController.UpdateSemester)
				r.Delete("/{num}", semesterController.DeleteSemester)
			})
			r.Route("/subject", func(r chi.Router) {
				r.Get("/", subjectController.ListSubjects)
				r.Get("/{name}/{semester_num}", subjectController.GetSubject)
				r.With(mw.AuthMiddleware(authClient)).Post("/", subjectController.CreateSubject)
				r.With(mw.AuthMiddleware(authClient)).Put("/", subjectController.UpdateSubject)
				r.With(mw.AuthMiddleware(authClient)).Delete("/{name}/{semester_num}", subjectController.DeleteSubject)
			})
			r.Route("/interview", func(r chi.Router) {
				r.Get("/", interviewController.ListInterviews)
				r.Get("/{id}", interviewController.GetInterview)
				r.With(mw.AuthMiddleware(authClient)).Post("/", interviewController.CreateInterview)
				r.With(mw.AuthMiddleware(authClient)).Put("/{id}", interviewController.UpdateInterview)
				r.With(mw.AuthMiddleware(authClient)).Delete("/{id}", interviewController.DeleteInterview)
			})
			r.Route("/study_plan", func(r chi.Router) {
				r.Get("/", studyPlanController.GetStudyPlan)
				r.Put("/", studyPlanController.UpdateStudyPlan)
			})
		})
	})
	r.Get("/ws", ws.Connect)

	go ws.HandleMessages()

	return http.ListenAndServe(fmt.Sprintf(":%d", *appPort), r)
}
