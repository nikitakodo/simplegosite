package appserver

import (
	_ "github.com/lib/pq"
	"net/http"
	"simplesite/internal/app/config"
	"simplesite/internal/app/di"
	"simplesite/internal/app/routing"
)

func Start(config *config.Config) error {

	diContainer, err := di.Factory(config)
	if err != nil {
		return err
	}
	router := routing.Factory(diContainer, routing.RouterConfig{Origin: config.AllowedOrigins})

	return http.ListenAndServe(config.BindAddr, newHandler(diContainer, router))
}
