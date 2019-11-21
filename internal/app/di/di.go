package di

import (
	"github.com/go-redis/redis/v7"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"simplesite/internal/app/config"
	"simplesite/internal/app/services"
	"simplesite/internal/app/store"
)

type GlobalDi struct {
	Router  *mux.Router
	Store   *store.Store
	Session *services.SessionService
	View    *services.View
	Logger  *logrus.Logger
}

func Factory(config *config.Config) (*GlobalDi, error) {
	db, err := newDB(config.DB.Url)

	if err != nil {
		return nil, err
	}

	client, err := newRedis(config.Cache.Addr, config.Cache.Password)
	if err != nil {
		return nil, err
	}

	cache := &store.Cache{
		Client: client,
		Prefix: config.Cache.Prefix,
	}

	Store := store.New(db, cache)

	sessionStore := sessions.NewCookieStore([]byte(config.Session.Key))
	session := services.NewSession(config.Session.Name, sessionStore)

	view, err := services.NewView(config.TemplatesDir, cache)
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

func newDB(dbURL string) (*gorm.DB, error) {
	db, err := gorm.Open("postgres", dbURL)
	if err != nil {
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
