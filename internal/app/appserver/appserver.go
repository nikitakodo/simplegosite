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
	defer diContainer.Store.Db.Close()
	defer diContainer.Store.Cache.Client.Close()
	router := routing.Factory(diContainer, routing.RouterConfig{Origin: config.AllowedOrigins})
	diContainer.Router = router.Router

	return http.ListenAndServe(config.BindAddr, newHandler(diContainer, router))
}
