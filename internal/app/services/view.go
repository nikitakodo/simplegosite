package services

import (
	"bytes"
	"html/template"
	"net/http"
	"strconv"
	"time"
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
	return view.Templates.ExecuteTemplate(w, templateName, data)
}

func (view *View) ResponseTemplate(w http.ResponseWriter, data map[string]interface{}, templateName string) (err error) {
	var tpl bytes.Buffer
	err = view.Templates.ExecuteTemplate(&tpl, templateName, data)
	if err != nil {
		return err
	}
	content := tpl.String()
	err = view.Cache.Set(templateName, content, 60*10*time.Second)
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(content))
	//err = view.ExecuteTemplate(w, data, templateName)

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
