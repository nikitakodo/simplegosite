package cabinet

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"simplesite/internal/app/repository"
	"simplesite/internal/app/services"
	"simplesite/internal/app/store"
)

type Controller struct {
	View    *services.View
	Logger  *logrus.Logger
	Store   *store.Store
	Router  *mux.Router
	Session *services.SessionService
}

func (c *Controller) SignIn(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {
		cabinetHome, err := c.Router.GetRoute("CabinetHome").URL()
		if err != nil {
			c.View.CabinetError(w, r, http.StatusInternalServerError, err)
			return
		}
		session, err := c.Session.SessionStore.Get(r, c.Session.SessionName)
		if err != nil {
			c.Error(w, r, err)
			return
		}

		id, ok := session.Values["user_id"].(uint)
		if ok {
			authorRepo := repository.AuthorRepository{Store: c.Store}
			_, err := authorRepo.Find(id)
			if err != nil {
				c.Error(w, r, err)
				return
			}
			http.Redirect(w, r, cabinetHome.Path, http.StatusSeeOther)
			return
		}
		tmplName := "cabinet_forms_signin"
		data := map[string]interface{}{}
		err = c.View.ExecuteTemplate(w, data, tmplName)
		if err != nil {
			c.Error(w, r, err)
			return
		}
		return
	} else if r.Method == http.MethodPost {
		type Form struct {
			Email    string
			Password string
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

		authorRepo := repository.AuthorRepository{Store: c.Store}
		author, err := authorRepo.FindByLogin(form.Email)
		if err != nil {
			c.Error(w, r, err)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(author.Password), []byte(form.Password))
		if err != nil {
			signin, err := c.Router.GetRoute("SignIn").URL()
			if err != nil {
				c.Error(w, r, err)
				return
			}
			http.Redirect(w, r, signin.Path, http.StatusSeeOther)
		}

		session, err := c.Session.SessionStore.Get(r, c.Session.SessionName)
		if err != nil {
			c.Error(w, r, err)
			return
		}

		session.Values["user_id"] = author.ID
		if err := c.Session.SessionStore.Save(r, w, session); err != nil {
			c.Error(w, r, err)
			return
		}

		cabinetHome, err := c.Router.GetRoute("CabinetHome").URL()
		if err != nil {
			c.Error(w, r, err)
			return
		}
		http.Redirect(w, r, cabinetHome.Path, http.StatusSeeOther)
		return
	}
	c.Error(w, r, errors.New("method not allowed"))
	return
}

func (c *Controller) SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmplName := "cabinet_forms_signup"
		data := map[string]interface{}{}
		err := c.View.ExecuteTemplate(w, data, tmplName)
		if err != nil {
			c.Error(w, r, err)
			return
		}
		return
	} else if r.Method == http.MethodPost {
		cabinetHome, err := c.Router.GetRoute("CabinetHome").URL()
		if err != nil {
			c.Error(w, r, err)
			return
		}
		http.Redirect(w, r, cabinetHome.Path, http.StatusSeeOther)
		return
	}
	c.Error(w, r, errors.New("method not allowed"))
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
