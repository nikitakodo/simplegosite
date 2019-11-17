package frontend

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"simplesite/internal/app/services"
	"simplesite/internal/app/store/sqlstore"
)

type BlogController struct {
	View   *services.View
	Logger *logrus.Logger
	Store  *sqlstore.Store
	Cache  *services.Cache
}

func (c *BlogController) Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{}
	navRepo := sqlstore.NavRepository{Store: c.Store, Cache: c.Cache}
	navs, err := navRepo.FindAll()
	if err != nil {
		c.Logger.Error(err)
	}
	for i, d := range navs {
		c.Logger.Infoln(i, d)
	}
	data["nav"] = navs
	c.Logger.Infoln(data)
	err = c.View.ResponseTemplate(w, data, "blog_pages_home")
	if err != nil {
		c.Logger.Error(err)
	}
}
