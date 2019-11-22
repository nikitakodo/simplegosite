package routing

import (
	"github.com/gorilla/handlers"
	"simplesite/internal/app/controller/frontend"
)

func (r *Routing) setupBlogRoutes() {

	routes := r.routes()
	r.Router.Use(
		r.Middleware.SetRequestID,
		r.Middleware.LogRequest,
		handlers.CORS(
			handlers.AllowedOrigins([]string{r.Config.Origin}),
		),
	)
	for _, route := range routes {
		r.Router.Methods(route.Method).Path(route.Pattern).Handler(route.Function)
	}
}

func (r *Routing) routes() []Route {
	blogController := frontend.BlogController{
		View:   r.Di.View,
		Logger: r.Di.Logger,
		Store:  r.Di.Store,
	}
	return []Route{
		{
			"/",
			"GET",
			blogController.Home,
		},
	}
}

//s.GlobalDi.Routing.Router.Use(s.Routing.Middleware.LogRequest)
//s.Routing.Router.Use(s.setRequestID)
//s.Routing.Router.Use(s.logRequest)
//s.Routing.Router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
//
//private := s.Routing.Router.PathPrefix("/private").Subrouter()
//private.Use(s.authenticateUser)
//private.HandleFunc("/whoami", s.handleWhoami())
