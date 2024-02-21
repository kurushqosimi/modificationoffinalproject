package Kurush

import (
	"github.com/gorilla/mux"
	"net/http"
	"one-day-job/config"
	"one-day-job/internal/domain/Kurush"
	"one-day-job/internal/middlewares"
	"one-day-job/pkg/logger"
)

type Handler struct {
	service Kurush.IService
	cfg     *config.Config
	log     *logger.Logger
}

func NewHandler(service Kurush.IService, config *config.Config, log *logger.Logger) *Handler {
	return &Handler{
		service: service,
		cfg:     config,
		log:     log,
	}
}

//GET /api/jobs: Получить список доступных работ.
//GET /api/jobs/{jobId}: Получить информацию о конкретной работе.
//POST /api/jobs: Создать новую работу.
//PUT /api/jobs/{jobId}: Обновить информацию о работе.
//DELETE /api/jobs/{jobId}: Удалить работу.

//GET /api/users/{userId}: Получить информацию о пользователе.
//POST /api/users: Зарегистрировать нового пользователя.
//PUT /api/users/{userId}: Обновить информацию о пользователе.
//DELETE /api/users/{userId}: Удалить пользователя.

//GET /api/jobs/{jobId}/applications: Получить список откликов на конкретную работу.
//POST /api/jobs/{jobId}/applications: Подать отклик на работу.
//PUT /api/applications/{applicationId}: Обновить статус отклика (принят/отклонен).

//GET /api/jobs/{jobId}/reviews: Получить список отзывов о конкретной работе.
//POST /api/jobs/{jobId}/reviews: Оставить отзыв о работе.

//GET /api/categories: Получить список доступных категорий работ.
//GET /api/categories/{categoryId}: Получить информацию о конкретной категории.
//POST /api/categories: Создать новую категорию.

func (h *Handler) InitRoutes() *mux.Router {
	newRouter := mux.NewRouter()
	mainRoute := newRouter.PathPrefix("/api").Subrouter()

	routeVer := mainRoute.PathPrefix("/v1").Subrouter()

	newRouter.Use(middlewares.SetContentType)
	routeVer.HandleFunc("/health", h.HealthCheck).Methods(http.MethodGet)

	//routeVer.HandleFunc("/api/jobs").Methods(http.MethodGet)
	//routeVer.HandleFunc("/api/jobs/{jobID}").Methods(http.MethodGet)
	//routeVer.HandleFunc("/api/jobs").Methods(http.MethodPost)
	//routeVer.HandleFunc("/api/jobs/{jobID}").Methods(http.MethodPut)
	//routeVer.HandleFunc("/api/jobs/{jobId}").Methods(http.MethodDelete)
	//
	//routeVer.HandleFunc("/api/users/{userId}").Methods(http.MethodGet)
	routeVer.HandleFunc("/api/users").Methods(http.MethodPost)
	//routeVer.HandleFunc("/api/users/{userId}").Methods(http.MethodPut)
	//routeVer.HandleFunc("/api/users/{userId}").Methods(http.MethodPut)
	//
	//routeVer.HandleFunc("/api/jobs/{jobId}/applications").Methods(http.MethodGet)
	//routeVer.HandleFunc("/api/jobs/{jobId}/applications").Methods(http.MethodPost)
	//routeVer.HandleFunc("/api/applications/{applicationId}").Methods(http.MethodPut)
	//
	//routeVer.HandleFunc("/api/jobs/{jobId}/reviews").Methods(http.MethodGet)
	//routeVer.HandleFunc(" /api/jobs/{jobId}/reviews").Methods(http.MethodPost)
	//
	//routeVer.HandleFunc("/api/categories").Methods(http.MethodGet)
	//routeVer.HandleFunc("/api/categories/{categoryId}").Methods(http.MethodGet)
	//routeVer.HandleFunc("/api/categories").Methods(http.MethodPost)
	return routeVer
}
