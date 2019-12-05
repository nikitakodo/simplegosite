package cabinet

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"simplesite/internal/app/services"
	"simplesite/internal/app/store"
)

type Controller struct {
	View   *services.View
	Logger *logrus.Logger
	Store  *store.Store
	Router *mux.Router
}

func (c *Controller) SignInPage(w http.ResponseWriter, r *http.Request) {
	tmplName := "cabinet_forms_signin"
	data := map[string]interface{}{}
	err := c.View.ExecuteTemplate(w, data, tmplName)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	return
}

func (c *Controller) SignInRequest(w http.ResponseWriter, r *http.Request) {
	cabinetHome, err := c.Router.GetRoute("CabinetHome").URL()
	if err != nil {
		c.Error(w, r, err)
		return
	}
	http.Redirect(w, r, cabinetHome.Path, http.StatusSeeOther)
	return
}

func (c *Controller) SignUpPage(w http.ResponseWriter, r *http.Request) {
	tmplName := "cabinet_forms_signup"
	data := map[string]interface{}{}
	err := c.View.ExecuteTemplate(w, data, tmplName)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	return
}

func (c *Controller) SignUpRequest(w http.ResponseWriter, r *http.Request) {
	cabinetHome, err := c.Router.GetRoute("CabinetHome").URL()
	if err != nil {
		c.Error(w, r, err)
		return
	}
	http.Redirect(w, r, cabinetHome.Path, http.StatusSeeOther)
	return
}

func (c *Controller) Home(w http.ResponseWriter, r *http.Request) {
	tmplName := "cabinet_grids_dashboard"
	data := map[string]interface{}{}
	err := c.View.ExecuteTemplate(w, data, tmplName)
	if err != nil {
		c.Error(w, r, err)
		return
	}
	return
}

func (c *Controller) NotFound(w http.ResponseWriter, r *http.Request) {
	c.Logger.Error(r.RequestURI)
	c.View.CabinetError(w, r, http.StatusNotFound, errors.New(r.RequestURI+" not found"))
	return
}

func (c *Controller) Error(w http.ResponseWriter, r *http.Request, err error) {
	c.Logger.Error(err)
	c.View.CabinetError(w, r, http.StatusInternalServerError, err)
	return
}
