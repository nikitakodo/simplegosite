package routing

import (
	"github.com/gorilla/handlers"
	"net/http"
	"simplesite/internal/app/controller/cabinet"
)

func (r *Routing) setupCabinetRoutes() {

	r.Router.Use(
		r.Middleware.SetRequestID,
		r.Middleware.LogRequest,
		handlers.CORS(
			handlers.AllowedOrigins([]string{r.Config.Origin}),
		),
	)
	cc := cabinet.Controller{
		View:   r.Di.View,
		Logger: r.Di.Logger,
		Store:  r.Di.Store,
		Router: r.Router,
	}
	routes := r.cabinetRoutes(cc)

	subRoute := r.Router.PathPrefix("/cabinet").Subrouter()
	for _, route := range routes {
		subRoute.Methods(route.Method).Path(route.Pattern).Handler(route.Function).Name(route.Name)
	}

	subRoute.NotFoundHandler = r.cabinetNotFound()
	subRoute.MethodNotAllowedHandler = r.cabinetNotFound()
}

func (r *Routing) cabinetRoutes(controller cabinet.Controller) []Route {

	return []Route{
		{
			"/",
			[]string{},
			"GET",
			"CabinetHome",
			controller.Home,
		},
		{
			"/signin",
			[]string{},
			"GET",
			"SignIn",
			controller.SignInPage,
		},
		{
			"/signin",
			[]string{},
			"POST",
			"SignInRequest",
			controller.SignInRequest,
		},
		{
			"/signup",
			[]string{},
			"GET",
			"SignUp",
			controller.SignUpPage,
		},
		{
			"/signup",
			[]string{},
			"POST",
			"SignUpRequest",
			controller.SignUpRequest,
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
