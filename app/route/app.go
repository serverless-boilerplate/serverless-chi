package route

import (
	"fmt"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func GetRoutePath(routeName string) string {
	basePath := "/api/v1"
	return fmt.Sprintf("%s/%s", basePath, routeName)
}

func App() *chi.Mux {
	app := chi.NewRouter()
	app.Use(middleware.RequestID)
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)
	app.Use(middleware.URLFormat)
	app.Use(render.SetContentType(render.ContentTypeJSON))
	app.Mount(GetRoutePath("service"), ServiceRoute())
	app.Mount(GetRoutePath("user"), UserRoute())
	return app
}
