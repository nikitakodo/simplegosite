package di

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"simplesite/internal/app/config"
	"simplesite/internal/app/services"
	"simplesite/internal/app/store/sqlstore"
)

type GlobalDi struct {
	Router  *mux.Router
	Store   *sqlstore.Store
	Session *services.SessionService
	View    *services.View
	Logger  *logrus.Logger
}

func Factory(config *config.Config) (*GlobalDi, error) {
	db, err := newDB(config.DB.Url)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	Store := sqlstore.New(db)

	sessionStore := sessions.NewCookieStore([]byte(config.Session.Key))
	session := services.NewSession(config.Session.Name, sessionStore)

	view, err := services.NewView(config.TemplatesDir)
	if err != nil {
		return nil, err
	}

	logger := logrus.New()
	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logger.Error("setup log error ", err)
	} else {
		logger.SetLevel(logLevel)
	}

	di := &GlobalDi{
		Store:   Store,
		Session: &session,
		View:    &view,
		Logger:  logger,
	}

	return di, nil
}

func newDB(dbURL string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
