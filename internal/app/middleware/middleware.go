package middleware

import (
	"context"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"net/http"
	"simplesite/internal/app/di"
	"time"
)

type ctxKey int8

const (
	ctxKeyUser ctxKey = iota
	ctxKeyRequestID
)

type Service struct {
	Di *di.GlobalDi
}

func (m Service) SetRequestID(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := uuid.New().String()
		w.Header().Set("X-Request-ID", id)
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyRequestID, id)))
	})
}

func (m Service) LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := logrus.New()
		logger := log.WithFields(logrus.Fields{
			"remote_addr": r.RemoteAddr,
			"request_id":  r.Context().Value(ctxKeyRequestID),
		})
		logger.Infof("started %s %s", r.Method, r.RequestURI)

		start := time.Now()
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		var level logrus.Level
		switch {
		case rw.code >= 500:
			level = logrus.ErrorLevel
		case rw.code >= 400:
			level = logrus.WarnLevel
		default:
			level = logrus.InfoLevel
		}
		logger.Logf(
			level,
			"completed with %d %s in %v",
			rw.code,
			http.StatusText(rw.code),
			time.Now().Sub(start),
		)
	})
}

func (m Service) AdminAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			//session, err := m.Di.Session.SessionStore.Get(r, m.Di.Session.SessionName)
			//if err != nil {
			//	m.Di.View.Error(w, r, http.StatusInternalServerError, err)
			//	return
			//}
			//
			//id, ok := session.Values["user_id"]
			//if !ok {
			//	//View{}.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			//	return
			//}
			//
			//u, err := m.Di.Store.User().Find(id.(int))
			//if err != nil {
			//	//s.error(w, r, http.StatusUnauthorized, errNotAuthenticated)
			//	return
			//}

			//next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), ctxKeyUser, u)))
		})
}

//func (s *App) handleWhoami() http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		s.respond(w, r, http.StatusOK, r.Context().Value(ctxKeyUser).(*model.User))
//	}
//}
//
//func (s *App) error(w http.ResponseWriter, r *http.Request, code int, err error) {
//	s.View.ResponseTemplate()
//	respond(w, r, code, map[string]string{"error": err.Error()})
//}
