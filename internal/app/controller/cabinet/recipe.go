package cabinet

import (
	"net/http"
	"simplesite/internal/app/model"
	"simplesite/internal/app/repository"
	"strconv"
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

	recipesByAuthor, err := rr.FindByAuthor(data["author"].(*model.Author))
	if err != nil {
		c.Error(w, r, err)
		return
	}

	type item struct {
		Id           uint
		Title        string
		Category     string
		Cuisine      string
		FormatedDate string
		Path         string
	}

	var recipes []item

	for _, recipe := range recipesByAuthor {
		path, _ := c.Router.
			Get("CabinetEditRecipe").
			URL("id", strconv.FormatUint(uint64(recipe.GetId()), 10))
		recipes = append(recipes, item{
			Id:           recipe.GetId(),
			Title:        recipe.Title,
			Category:     recipe.Category.Name,
			Cuisine:      recipe.Cuisine.Name,
			FormatedDate: recipe.FormatedDate(),
			Path:         path.Path,
		})
	}

	data["recipes"] = recipes
	err = c.View.ExecuteTemplate(w, data, tmplName)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	return
}
