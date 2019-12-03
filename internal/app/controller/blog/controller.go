package blog

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
	"math"
	"net/http"
	"simplesite/internal/app/repository"
	"simplesite/internal/app/services"
	"simplesite/internal/app/store"
	"strconv"
)

type Controller struct {
	View   *services.View
	Logger *logrus.Logger
	Store  *store.Store
	Router *mux.Router
}

func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	tmplName := "blog_pages_home"
	if res, err := GetCachedView(tmplName, c.Store.Cache); err == nil && len(*res) > 0 {
		_, err := w.Write([]byte(*res))
		if err != nil {
			c.Error(w, r, err)
			return
		}
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

	topRated, err := repository.RecipeRepository{Store: c.Store}.TopRated(5, 0)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	data["top_rated"] = topRated

	mostLiked, err := repository.RecipeRepository{Store: c.Store}.MostLiked(5, 0)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	data["most_liked"] = mostLiked

	content, err := c.View.ProcessTemplate(data, tmplName)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	_ = SetCachedView(*content, tmplName, c.Store.Cache)
	_, err = w.Write([]byte(*content))
	if err != nil {
		c.Error(w, r, err)
		return
	}
}

func (c *Controller) Recipes(w http.ResponseWriter, r *http.Request) {
	tmplName := "blog_pages_recipes"
	limit := 12
	offset := 0
	pageNum := 1
	pageStr := "1"
	keys, ok := r.URL.Query()["page"]
	if ok && len(keys) > 0 && len(keys[0]) > 0 {
		pageStr = keys[0]
		pageNum, _ = strconv.Atoi(pageStr)
	}
	if pageNum <= 1 {
		pageNum = 1
	} else {
		offset = limit * (pageNum - 1)
	}

	if res, err := GetCachedView(tmplName+"_page"+pageStr, c.Store.Cache); err == nil && len(*res) > 0 {
		_, err := w.Write([]byte(*res))
		if err != nil {
			c.Error(w, r, err)
			return
		}
		return
	}

	data := map[string]interface{}{}
	page, err := GetBasicData(c.View, c.Logger, c.Store)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	data["page"] = page

	rr := repository.RecipeRepository{Store: c.Store}
	count, err := rr.GetLatestCount()
	if err != nil {
		c.Error(w, r, err)
		return
	}

	pages := map[string]string{}
	p := int(math.Ceil(float64(*count) / float64(limit)))
	if p > 1 {
		for i := 1; i <= p; i++ {
			ps := strconv.Itoa(i)

			path, err := c.Router.GetRoute("Recipes").URL()
			if err != nil {
				c.Error(w, r, err)
				return
			}
			pages[ps] = path.RequestURI()
		}
	}

	latestRecipes, err := rr.GetLatest(limit, offset)
	if err != nil {
		c.Error(w, r, err)
		return
	}

	data["latest"] = map[string]interface{}{"items": latestRecipes, "pages": pages, "current_page": pageStr}

	content, err := c.View.ProcessTemplate(data, tmplName)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	_ = SetCachedView(*content, tmplName+"_page"+pageStr, c.Store.Cache)
	_, err = w.Write([]byte(*content))
	if err != nil {
		c.Error(w, r, err)
		return
	}
}

func (c *Controller) About(w http.ResponseWriter, r *http.Request) {
	tmplName := "blog_pages_about"
	if res, err := GetCachedView(tmplName, c.Store.Cache); err == nil && len(*res) > 0 {
		_, err := w.Write([]byte(*res))
		if err != nil {
			c.Error(w, r, err)
			return
		}
		return
	}

	data := map[string]interface{}{}
	page, err := GetBasicData(c.View, c.Logger, c.Store)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	data["page"] = page

	aboutRepo := repository.AboutRepository{Store: c.Store}

	data["about"], err = aboutRepo.FindFirst()
	if err != nil {
		c.Error(w, r, err)
		return
	}

	//rr := repository.RecipeRepository{Store: c.Store}

	data["facts"] = map[string]interface{}{
		"Amazing receipies": map[string]interface{}{"icon": "img/icon/1.png", "count": 1850},
		"Wine pairings":     map[string]interface{}{"icon": "img/icon/2.png", "count": 185},
		"Meat receipies":    map[string]interface{}{"icon": "img/icon/3.png", "count": 1500},
		"Desert receipies":  map[string]interface{}{"icon": "img/icon/4.png", "count": 50},
	}

	content, err := c.View.ProcessTemplate(data, tmplName)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	_ = SetCachedView(*content, tmplName, c.Store.Cache)
	_, err = w.Write([]byte(*content))
	if err != nil {
		c.Error(w, r, err)
		return
	}
}

func (c *Controller) Contact(w http.ResponseWriter, r *http.Request) {
	tmplName := "blog_pages_contact"
	if res, err := GetCachedView(tmplName, c.Store.Cache); err == nil && len(*res) > 0 {
		_, err := w.Write([]byte(*res))
		if err != nil {
			c.Error(w, r, err)
			return
		}
		return
	}

	data := map[string]interface{}{}
	page, err := GetBasicData(c.View, c.Logger, c.Store)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	data["page"] = page

	data["info"] = map[string]interface{}{
		"address": "<p>481 Creekside Lane Avila</p><p>Beach, CA 93424</p>",
		"phone":   "<p>+53 345 7953 32453</p><p>+53 345 7557 822112</p>",
		"email":   "<p>yourmail@gmail.com</p>",
	}
	data["form_action"] = "/contact_form"

	content, err := c.View.ProcessTemplate(data, tmplName)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	_ = SetCachedView(*content, tmplName, c.Store.Cache)
	_, err = w.Write([]byte(*content))
	if err != nil {
		c.Error(w, r, err)
		return
	}
}

func (c *Controller) ContactForm(w http.ResponseWriter, r *http.Request) {

	type Form struct {
		Name    string
		Email   string
		Subject string
		Message string
	}
	form := new(Form)
	err := r.ParseForm()
	if err != nil {
		c.Error(w, r, err)
		return
	}
	decoder := schema.NewDecoder()
	err = decoder.Decode(form, r.PostForm)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	a, err := json.Marshal(form)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	_, _ = w.Write(a)
}

func (c *Controller) NotFound(w http.ResponseWriter, r *http.Request) {
	c.Logger.Error(r.RequestURI)
	c.View.Error(w, r, http.StatusNotFound, errors.New(r.RequestURI+" not found"))
	return
}

func (c *Controller) Error(w http.ResponseWriter, r *http.Request, err error) {
	c.Logger.Error(err)
	c.View.Error(w, r, http.StatusInternalServerError, err)
	return
}
