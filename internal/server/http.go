package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/not4sure/tracking-service/internal/common/logs"
	"github.com/sirupsen/logrus"
)

type RegisterFn func(router chi.Router)

func RunServer(register RegisterFn) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	RunServerOnAddr(addr, register)
}

func RunServerOnAddr(addr string, register RegisterFn) {
	apiRouter := chi.NewRouter()
	setMiddlewares(apiRouter)
	register(apiRouter)

	rootRouter := chi.NewRouter()
	rootRouter.Mount("/api", apiRouter)

	server := &http.Server{
		Addr:    addr,
		Handler: rootRouter,
	}

	logrus.WithField("addr", addr).Info("Starting HTTP server")
	err := server.ListenAndServe()
	if err != nil {
		logrus.WithError(err).Panic("Unable to start HTTP server")
	}
}

func setMiddlewares(router *chi.Mux) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(logs.NewStructuredLogger(logrus.StandardLogger()))
	router.Use(middleware.Recoverer)

	addCorsMiddleware(router)
}

func addCorsMiddleware(router *chi.Mux) {
	allowedOrigins := strings.Split(os.Getenv("CORS_ALLOWED_ORIGINS"), ";")
	if len(allowedOrigins) == 0 {
		logrus.Info("Empty allowed origins")
		return
	}
	logrus.WithField("origins", allowedOrigins).Info("Setting allowed origins")

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   allowedOrigins,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
	router.Use(corsMiddleware.Handler)
}
