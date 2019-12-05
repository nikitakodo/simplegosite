package routing

import (
	"github.com/gorilla/handlers"
	"net/http"
	"simplesite/internal/app/controller/blog"
)

func (r *Routing) setupBlogRoutes() {

	r.Router.Use(
		r.Middleware.SetRequestID,
		r.Middleware.LogRequest,
		handlers.CORS(
			handlers.AllowedOrigins([]string{r.Config.Origin}),
		),
	)
	blogController := blog.Controller{
		View:   r.Di.View,
		Logger: r.Di.Logger,
		Store:  r.Di.Store,
		Router: r.Router,
	}
	routes := r.blogRoutes(blogController)
	for _, route := range routes {
		r.Router.Methods(route.Method).Path(route.Pattern).Handler(route.Function).Name(route.Name)
	}

	r.Router.NotFoundHandler = r.blogNotFound()
	r.Router.MethodNotAllowedHandler = r.blogNotFound()
}

func (r *Routing) blogRoutes(controller blog.Controller) []Route {

	return []Route{
		{
			"/",
			[]string{},
			"GET",
			"Home",
			controller.Home,
		},
		{
			"/recipes",
			[]string{},
			"GET",
			"Recipes",
			controller.Recipes,
		},
		{
			"/recipes/{id}",
			[]string{},
			"GET",
			"Recipe",
			controller.Recipe,
		},
		{
			"/about",
			[]string{},
			"GET",
			"About",
			controller.About,
		},
		{
			"/contact",
			[]string{},
			"GET",
			"Contact",
			controller.Contact,
		},
		{
			"/contact_form",
			[]string{},
			"POST",
			"ContactForm",
			controller.ContactForm,
		},
	}
}

func (r *Routing) blogNotFound() http.Handler {
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
