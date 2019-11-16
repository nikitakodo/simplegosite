package appserver

import (
	"net/http"
	"simplesite/internal/app/di"
	"simplesite/internal/app/routing"
)

type App struct {
	*di.GlobalDi
	routing.Routing
}

func newHandler(di *di.GlobalDi, routing routing.Routing) *App {
	return &App{di, routing}
}

func (app *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.Routing.Router.ServeHTTP(w, r)
}

//func (s *App) handleWhoami() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
//	}
//}

//func (s *App) error(w http.ResponseWriter, r *http.Request, code int, err error) {
//	s.View.ResponseTemplate()
//	respond(w, r, code, map[string]string{"error": err.Error()})
//}
