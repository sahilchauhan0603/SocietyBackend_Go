package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/sahilchauhan0603/society/internal/repository"
	"github.com/sahilchauhan0603/society/internal/service"
)

func InitializeRoutes(router *mux.Router) {
	registerSystemRoutes(router)

	api := router.PathPrefix("/api/v1").Subrouter()
	registerAuthRoutes(router, api)
	registerSocietyRoutes(api)
	registerStudentRoutes(api)
	registerContentRoutes(api)
	registerAdminRoutes(router, api)
}

func registerSystemRoutes(router *mux.Router) {
	healthRepo := repository.NewHealthRepository()
	healthSvc := service.NewHealthService(healthRepo)

	healthHandler := func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
		defer cancel()

		status := healthSvc.Check(ctx)
		httpStatus := http.StatusOK
		if status.Status != "ok" {
			httpStatus = http.StatusServiceUnavailable
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(httpStatus)
		_ = json.NewEncoder(w).Encode(status)
	}

	router.HandleFunc("/healthz", healthHandler).Methods(http.MethodGet)
	router.HandleFunc("/readyz", healthHandler).Methods(http.MethodGet)
}
