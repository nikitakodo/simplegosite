package appserver

import (
	"database/sql"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/gorilla/sessions"
	_ "github.com/lib/pq"

	"simplesite/internal/app/services"
	"simplesite/internal/app/store/sqlstore"
)

func Start(config *Config) error {
	db, err := newDB(config.DB.Url)
	if err != nil {
		return err
	}
	defer db.Close()
	store := sqlstore.New(db)

	sessionStore := sessions.NewCookieStore([]byte(config.Session.Key))
	session := services.NewSession(config.Session.Name, sessionStore)

	view, err := services.NewView(config.TemplatesDir)
	if err != nil {
		return err
	}

	logger := logrus.New()
	logLevel, err := logrus.ParseLevel(config.LogLevel)
	if err != nil {
		logger.Error("setup log error ", err)
	} else {
		logger.SetLevel(logLevel)
	}

	router := services.NewRouting(services.MiddlewareService{
		Store:   store,
		Session: session,
		Logger:  logger,
	})

	srv := newServer(router, store, session, view, logger)

	return http.ListenAndServe(config.BindAddr, srv)
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
