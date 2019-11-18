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
	tmplName := "blog_pages_home"
	data := map[string]interface{}{}
	page, err := GetBasicData(c.View, c.Logger, c.Store, c.Cache)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	data["page"] = page

	slidesRepo := sqlstore.SlidesRepository{Store: c.Store, Cache: c.Cache}
	slides, err := slidesRepo.FindAll()
	if err != nil {
		c.Error(w, r, err)
		return
	}
	data["slides"] = slides

	addRepo := sqlstore.AddRepository{Store: c.Store, Cache: c.Cache}
	add, err := addRepo.Find(1)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	data["add"] = add

	err = c.View.ResponseTemplate(w, data, tmplName)
	if err != nil {
		c.Error(w, r, err)
		return
	}
}

func (c *BlogController) Error(w http.ResponseWriter, r *http.Request, err error) {
	c.Logger.Error(err)
	c.View.Error(w, r, http.StatusInternalServerError, err)
	return
}
