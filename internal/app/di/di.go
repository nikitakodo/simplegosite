package di

import (
	"fmt"
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
	db, err := newDB(
		config.DB.DbHost,
		config.DB.DbUser,
		config.DB.DbName,
		config.DB.DbPassword,
		config.DB.DbSsl,
	)

	if err != nil {
		return nil, err
	}

	client, err := newRedis(config.Cache.Host, config.Cache.Port, config.Cache.Password)
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

	view, err := services.NewView(config.TemplatesDir, config.AssetsUrl, config.UploadsUrl, cache)
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

func newDB(dbHost string, username string, dbName string, password string, ssl string) (*gorm.DB, error) {
	dbUri := fmt.Sprintf(
		"host=%s user=%s dbname=%s password=%s sslmode=%s",
		dbHost,
		username,
		dbName,
		password,
		ssl,
	)
	db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func newRedis(host string, port string, pass string) (*redis.Client, error) {
	client := redis.NewClient(
		&redis.Options{
			Addr:     fmt.Sprintf("%s:%s", host, port),
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
