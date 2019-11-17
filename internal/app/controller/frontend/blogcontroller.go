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
	pageData, err := GetBasicData(c.View, c.Logger, c.Store, c.Cache)
	if err != nil {
		c.Logger.Error(err)
		c.View.Error(w, r, http.StatusInternalServerError, err)
		return
	}
	data["page"] = pageData
	err = c.View.ResponseTemplate(w, data, "blog_pages_home")
	if err != nil {
		c.Logger.Error(err)
	}
}
