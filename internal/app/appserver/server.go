package appserver

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"simplesite/internal/app/controllers/frontend"

	"simplesite/internal/app/services"
	"simplesite/internal/app/store"
)

type Server struct {
	Routing services.Routing
	Store   store.Store
	Session services.SessionService
	View    services.View
	Logger  *logrus.Logger
}

func newServer(
	routing services.Routing,
	store store.Store,
	session services.SessionService,
	view services.View,
	logger *logrus.Logger,
) *Server {
	s := &Server{
		Routing: routing,
		Session: session,
		Store:   store,
		View:    view,
		Logger:  logger,
	}
	s.configureRouter()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Routing.Router.ServeHTTP(w, r)
}

func (s *Server) configureRouter() {

	s.Routing.Router.Use(s.Routing.Middleware.LogRequest)
	//s.Routing.Router.Use(s.setRequestID)
	//s.Routing.Router.Use(s.logRequest)
	//s.Routing.Router.Use(handlers.CORS(handlers.AllowedOrigins([]string{"*"})))
	//
	//private := s.Routing.Router.PathPrefix("/private").Subrouter()
	//private.Use(s.authenticateUser)
	//private.HandleFunc("/whoami", s.handleWhoami())
	c := frontend.BlogController{View: s.View, Logger: s.Logger}
	s.Routing.Router.HandleFunc("/", c.Home)
}

//func (s *Server) handleWhoami() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
//	}
//}

//func (s *Server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
//	s.View.ResponseTemplate()
//	respond(w, r, code, map[string]string{"error": err.Error()})
//}
