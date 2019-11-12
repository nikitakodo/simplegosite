package services

import (
	"github.com/gorilla/mux"
	"simplesite/internal/app/controllers"
	"strings"
)

func NewRouting(service MiddlewareService) Routing {
	router := mux.NewRouter()
	return Routing{Router: router}
}

type Routing struct {
	Router     *mux.Router
	Middleware MiddlewareService
}

func (routing *Routing) GetRoute(subRoute string, path string, method string) (*string, error) {
	route, err := routing.Router.Get(routeName(subRoute, path, method)).URL()
	if err != nil {
		return nil, err
	}
	return &route.Path, nil
}

func (routing *Routing) ConfigureRecourse(subRoute string, resource controllers.ResourceInterface) {
	recourseName := resource.GetName()
	recourseRouter := routing.Router.PathPrefix("/" + recourseName).Subrouter()
	recourseRouter.HandleFunc("/{id}", resource.Read).Methods("GET").Name(routeName(subRoute, recourseName, "read"))
	recourseRouter.HandleFunc("/", resource.ReadAll).Methods("GET").Name(routeName(subRoute, recourseName, "read_all"))
	recourseRouter.HandleFunc("/create", resource.Create).Methods("GET", "POST").Name(routeName(subRoute, recourseName, "create"))
	recourseRouter.HandleFunc("/{id}/update", resource.Update).Methods("POST").Name(routeName(subRoute, recourseName, "update"))
	recourseRouter.HandleFunc("/{id}/delete", resource.Delete).Methods("POST").Name(routeName(subRoute, recourseName, "delete"))
}

func routeName(subRoute string, path string, method string) string {
	return strings.Join([]string{subRoute, path, method}, "_")
}
