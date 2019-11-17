package di

import (
	"database/sql"
	"github.com/go-redis/redis/v7"
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
	Cache   *services.Cache
}

func Factory(config *config.Config) (*GlobalDi, error) {
	db, err := newDB(config.DB.Url)
	if err != nil {
		return nil, err
	}

	Store := sqlstore.New(db)

	client, err := newRedis(config.Cache.Addr, config.Cache.Password)
	if err != nil {
		return nil, err
	}

	cache := &services.Cache{
		Client: client,
		Prefix: config.Cache.Prefix,
	}

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
		Cache:   cache,
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

func newRedis(addr string, pass string) (*redis.Client, error) {
	client := redis.NewClient(
		&redis.Options{
			Addr:     addr,
			Password: pass,
			DB:       0,
		},
	)
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return client, nil
}
