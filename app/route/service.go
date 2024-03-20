package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/serverless-boilerplate/serverless-chi/app/controller"
)

func ServiceRoute() chi.Router {
	route := chi.NewRouter()
	route.Get("/health", controller.ServiceHealth)
	route.Get("/time", controller.ServiceTime)
	return route
}
