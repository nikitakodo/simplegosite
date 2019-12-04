package routing

import (
	"github.com/gorilla/handlers"
	"net/http"
	"simplesite/internal/app/controller/blog"
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
		r.Router.Methods(route.Method).Path(route.Pattern).Handler(route.Function).Name(route.Name)
	}

	r.Router.NotFoundHandler = r.notFound()
	r.Router.MethodNotAllowedHandler = r.notFound()
}

func (r *Routing) routes() []Route {
	blogController := blog.Controller{
		View:   r.Di.View,
		Logger: r.Di.Logger,
		Store:  r.Di.Store,
		Router: r.Router,
	}
	return []Route{
		{
			"/",
			[]string{},
			"GET",
			"Home",
			blogController.Home,
		},
		{
			"/recipes",
			[]string{},
			"GET",
			"Recipes",
			blogController.Recipes,
		},
		{
			"/recipes/{id}",
			[]string{},
			"GET",
			"Recipe",
			blogController.Recipe,
		},
		{
			"/about",
			[]string{},
			"GET",
			"About",
			blogController.About,
		},
		{
			"/contact",
			[]string{},
			"GET",
			"Contact",
			blogController.Contact,
		},
		{
			"/contact_form",
			[]string{},
			"POST",
			"ContactForm",
			blogController.ContactForm,
		},
	}
}

func (r *Routing) notFound() http.Handler {
	blogController := blog.Controller{
		View:   r.Di.View,
		Logger: r.Di.Logger,
		Store:  r.Di.Store,
		Router: r.Router,
	}
	return http.HandlerFunc(blogController.NotFound)
}

//s.GlobalDi.Routing.Router.Use(s.Routing.Middleware.LogRequest)
//s.Routing.Router.Use(s.setRequestID)
//s.Routing.Router.Use(s.logRequest)
//s.Routing.Router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
//
//private := s.Routing.Router.PathPrefix("/private").Subrouter()
//private.Use(s.authenticateUser)
//private.HandleFunc("/whoami", s.handleWhoami())
