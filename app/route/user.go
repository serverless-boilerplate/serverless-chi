package route

import (
	"github.com/go-chi/chi/v5"
	"github.com/serverless-boilerplate/serverless-chi/app/controller"
)

func UserRoute() chi.Router {
	route := chi.NewRouter()
	route.Post("/signup", controller.SignUp)
	route.Post("/login", controller.Login)
	route.Post("logout", controller.Logout)
	route.Get("/", controller.GetUsers)
	route.Get("/:id", controller.GetUser)
	route.Put("/:id/activate", controller.ActivateUser)
	route.Put("/:id/deactivate", controller.DeactivateUser)
	return route
}
