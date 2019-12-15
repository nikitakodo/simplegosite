package cabinet

import (
	"net/http"
	"simplesite/internal/app/model"
	"simplesite/internal/app/repository"
)

func (c *Controller) CreateRecipe(w http.ResponseWriter, r *http.Request) {
	return
}

func (c *Controller) EditRecipe(w http.ResponseWriter, r *http.Request) {
	return
}

func (c *Controller) RecipeList(w http.ResponseWriter, r *http.Request) {
	tmplName := "cabinet_grids_recipes"
	data := c.getBasicData(r)

	rr := repository.RecipeRepository{Store: c.Store}

	recipes, err := rr.FindByAuthor(data["author"].(*model.Author))
	if err != nil {
		c.Error(w, r, err)
		return
	}
	data["recipes"] = recipes
	err = c.View.ExecuteTemplate(w, data, tmplName)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	return
}
