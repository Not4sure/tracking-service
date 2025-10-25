package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
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

	// TODO: add CORS middleware
}
