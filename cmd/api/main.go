package main

import (
	"context"
	"net/http"

	"github.com/not4sure/tracking-service/internal/ports/http_api"
	"github.com/not4sure/tracking-service/internal/server"
	"github.com/not4sure/tracking-service/internal/service"
)

func main() {
	ctx := context.Background()

	app := service.NewApplication(ctx)

	server.RunServer(func(router *http.ServeMux) {
		apiServer := http_api.NewAPIServer(app)
		apiServer.RegisterRoutes(router)
	})
}
