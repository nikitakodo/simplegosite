package services

import (
	"bytes"
	"fmt"
	"html/template"
	"net/http"
	"simplesite/internal/app/store"
	"strconv"
	"time"
)

type View struct {
	Templates  *template.Template
	Cache      *store.Cache
	AssetsDir  string
	UploadsDir string
}

func NewView(
	templatesDir string,
	assetsDir string,
	adminAssetsDir string,
	cabinetAssetsDir string,
	uploadsDir string,
	cache *store.Cache,
) (view View, err error) {
	view.Cache = cache
	view.AssetsDir = assetsDir
	view.UploadsDir = uploadsDir
	templates, err := template.New("default").
		Funcs(template.FuncMap{
			"asset": func(filePath string) string {
				return fmt.Sprintf("/%s/%s", assetsDir, filePath)
			},
			"admin_asset": func(filePath string) string {
				return fmt.Sprintf("/%s/%s", adminAssetsDir, filePath)
			},
			"cabinet_asset": func(filePath string) string {
				return fmt.Sprintf("/%s/%s", cabinetAssetsDir, filePath)
			},
			"upload": func(filePath string) string {
				return fmt.Sprintf("/%s/%s", uploadsDir, filePath)
			},
			"unescape": func(s string) template.HTML {
				return template.HTML(s)
			},
		}).
		ParseGlob(templatesDir + "/*/*/*.html")
	if err != nil {
		return
	}
	view.Templates = templates
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
	return
}

func (view *View) ProcessTemplate(data map[string]interface{}, templateName string) (res *string, err error) {
	var tpl bytes.Buffer
	err = view.Templates.ExecuteTemplate(&tpl, templateName, data)
	if err != nil {
		return nil, err
	}
	content := tpl.String()
	return &content, nil
}

func (view *View) BlogError(w http.ResponseWriter, r *http.Request, code int, err error) {
	view.Error(w, "blog_error_"+strconv.Itoa(code), map[string]interface{}{"code": code})
	return
}

func (view *View) CabinetError(w http.ResponseWriter, r *http.Request, code int, err error) {
	view.Error(w, "cabinet_error_"+strconv.Itoa(code), map[string]interface{}{"code": code})
	return
}

func (view *View) AdminError(w http.ResponseWriter, r *http.Request, code int, err error) {
	view.Error(w, "admin_error_"+strconv.Itoa(code), map[string]interface{}{"code": code})
	return
}

func (view *View) Error(w http.ResponseWriter, tmpl string, data map[string]interface{}) {
	_ = view.ExecuteTemplate(w, data, tmpl)
	return
}
