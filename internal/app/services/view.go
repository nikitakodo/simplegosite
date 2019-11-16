package services

import (
	"encoding/json"
	"html/template"
	"net/http"
)

type View struct {
	Templates *template.Template
}

func NewView(templatesDir string) (view View, err error) {
	templates, err := template.ParseGlob(templatesDir + "/*/*/*.html")
	if err != nil {
		return
	}
	view.Templates = templates
	return
}

func (view *View) ExecuteTemplate(w http.ResponseWriter, data map[string]interface{}, templateName string) (err error) {
	err = view.Templates.ExecuteTemplate(w, templateName, data)
	return
}

func (view *View) ResponseTemplate(w http.ResponseWriter, data map[string]interface{}, templateName string) (err error) {
	err = view.ExecuteTemplate(w, data, templateName)
	return
}

func (view *View) Error(w http.ResponseWriter, r *http.Request, code int, err error) {
	w.WriteHeader(code)
	if err != nil {
		_ = json.NewEncoder(w).Encode(map[string]string{"Error": err.Error()})
	}
}
