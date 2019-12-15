package routing

import (
	"github.com/gorilla/mux"
	"net/http"
	"simplesite/internal/app/di"
	"simplesite/internal/app/middleware"
	"strings"
)

type Routing struct {
	Router     *mux.Router
	Middleware *middleware.Service
	Di         *di.GlobalDi
	Routes     []Route
	Config     RouterConfig
}

type Route struct {
	Pattern    string
	Middleware []mux.MiddlewareFunc
	Method     []string
	Name       string
	Function   http.HandlerFunc
}

type RouterConfig struct {
	Origin string
}

func Factory(globalDi *di.GlobalDi, conf RouterConfig) Routing {
	router := mux.NewRouter()
	r := Routing{
		Router:     router,
		Middleware: &middleware.Service{Di: globalDi},
		Di:         globalDi,
		Config:     conf,
	}
	r.setupBlogRoutes()
	r.setupCabinetRoutes()

	return r
}

func (r Routing) GetRoute(subRoute string, path string, method string) (*string, error) {
	route, err := r.Router.Get(routeName(subRoute, path, method)).URL()
	if err != nil {
		return nil, err
	}
	return &route.Path, nil
}

//
//func (r Routing) ConfigureRecourse(subRoute string, resource controller.ResourceInterface) {
//	recourseName := resource.GetName()
//	recourseRouter := r.Router.PathPrefix("/" + recourseName).Subrouter()
//	recourseRouter.HandleFunc("/{id}", resource.Read).Methods("GET").Name(routeName(subRoute, recourseName, "read"))
//	recourseRouter.HandleFunc("/", resource.ReadAll).Methods("GET").Name(routeName(subRoute, recourseName, "read_all"))
//	recourseRouter.HandleFunc("/create", resource.Create).Methods("GET", "POST").Name(routeName(subRoute, recourseName, "create"))
//	recourseRouter.HandleFunc("/{id}/update", resource.Update).Methods("POST").Name(routeName(subRoute, recourseName, "update"))
//	recourseRouter.HandleFunc("/{id}/delete", resource.Delete).Methods("POST").Name(routeName(subRoute, recourseName, "delete"))
//}

func routeName(subRoute string, path string, method string) string {
	return strings.Join([]string{subRoute, path, method}, "_")
}
