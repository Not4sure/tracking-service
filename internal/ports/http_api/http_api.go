package http_api

import (
	"net/http"

	"github.com/not4sure/tracking-service/internal/app"
)

type APIServer struct{ app app.Application }

func NewAPIServer(app app.Application) APIServer {
	return APIServer{
		app: app,
	}
}

func (s APIServer) RegisterRoutes(r *http.ServeMux) {
	// TODO: Add routes.

}
