package routing

import (
	"github.com/gorilla/mux"
	"net/http"
	"simplesite/internal/app/controller/cabinet"
)

func (r *Routing) setupCabinetRoutes() {
	cc := cabinet.Controller{
		View:    r.Di.View,
		Logger:  r.Di.Logger,
		Store:   r.Di.Store,
		Router:  r.Router,
		Session: r.Di.Session,
	}
	routes := r.cabinetRoutes(cc)

	subRoute := r.Router.PathPrefix("/cabinet").Subrouter()
	for _, route := range routes {
		r := subRoute.PathPrefix(route.Pattern).Subrouter()
		r.Use(route.Middleware...)
		r.Path("").Handler(route.Function).Name(route.Name)

	}
	subRoute.NotFoundHandler = r.cabinetNotFound()
	subRoute.MethodNotAllowedHandler = r.cabinetNotFound()
}

func (r *Routing) cabinetRoutes(controller cabinet.Controller) []Route {

	return []Route{
		{
			"/signin",
			nil,
			[]string{"GET", "POST"},
			"SignIn",
			controller.SignIn,
		},
		{
			"/signup",
			nil,
			[]string{"GET", "POST"},
			"SignUp",
			controller.SignUp,
		},
		{
			"",
			[]mux.MiddlewareFunc{r.Middleware.CabinetAuth},
			[]string{"GET"},
			"CabinetHome",
			controller.Home,
		},
	}
}

func (r *Routing) cabinetNotFound() http.Handler {
	cc := cabinet.Controller{
		View:   r.Di.View,
		Logger: r.Di.Logger,
		Store:  r.Di.Store,
		Router: r.Router,
	}
	return http.HandlerFunc(cc.NotFound)
}

//s.GlobalDi.Routing.Router.Use(s.Routing.Middleware.LogRequest)
//s.Routing.Router.Use(s.setRequestID)
//s.Routing.Router.Use(s.logRequest)
//s.Routing.Router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
//
//private := s.Routing.Router.PathPrefix("/private").Subrouter()
//private.Use(s.authenticateUser)
//private.HandleFunc("/whoami", s.handleWhoami())
