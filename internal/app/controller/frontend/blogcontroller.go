package frontend

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"simplesite/internal/app/repository"
	"simplesite/internal/app/services"
	"simplesite/internal/app/store"
)

type BlogController struct {
	View   *services.View
	Logger *logrus.Logger
	Store  *store.Store
}

func (c *BlogController) Home(w http.ResponseWriter, r *http.Request) {
	tmplName := "blog_pages_home"
	if res := WriteCachedResponse(w, tmplName, c.Store.Cache); res {
		return
	}
	data := map[string]interface{}{}
	page, err := GetBasicData(c.View, c.Logger, c.Store)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	data["page"] = page

	slides, err := repository.SlidesRepository{Store: c.Store}.GetOrdered("order")
	if err != nil {
		c.Error(w, r, err)
		return
	}
	data["slides"] = slides

	add, err := repository.AddRepository{Store: c.Store}.Find(1)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	data["add"] = add

	latestRecipes, err := repository.RecipeRepository{Store: c.Store}.GetLatest(6, 0)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	data["latest"] = latestRecipes

	topRated, err := repository.RecipeRepository{Store: c.Store}.GetLatest(6, 0)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	data["top_rated"] = topRated

	mostLiked, err := repository.RecipeRepository{Store: c.Store}.GetLatest(6, 0)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	data["most_liked"] = mostLiked

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
