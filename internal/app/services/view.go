package services

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"
)

type View struct {
	Templates *template.Template
	Cache     *Cache
}

func NewView(templatesDir string, cache *Cache) (view View, err error) {
	templates, err := template.ParseGlob(templatesDir + "/*/*/*.html")
	if err != nil {
		return
	}
	view.Templates = templates
	view.Cache = cache
	return
}

func (view *View) ExecuteTemplate(w http.ResponseWriter, data map[string]interface{}, templateName string) (err error) {

	//TODO create template caching
	r, err := view.Cache.Get(view.Cache.Prefix + templateName)
	if err != nil {
		return
	}
	if r == nil {
		var tpl bytes.Buffer
		if err := view.Templates.Execute(&tpl, data); err != nil {
			return err
		}
		result := tpl.String()
		err = view.Cache.Set(view.Cache.Prefix+templateName, result)
		if err != nil {
			return
		}
	} else {
		_, err = w.Write([]byte(*r))
	}
	return
}

func (view *View) ResponseTemplate(w http.ResponseWriter, data map[string]interface{}, templateName string) (err error) {
	err = view.ExecuteTemplate(w, data, templateName)
	return
}

func (view *View) Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	//w.WriteHeader(code)
	_ = view.ExecuteTemplate(
		w,
		map[string]interface{}{"code": code},
		"blog_error_"+strconv.Itoa(code),
	)
	return
}
