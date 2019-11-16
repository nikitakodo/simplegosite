package frontend

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"simplesite/internal/app/services"
)

type BlogController struct {
	View   *services.View
	Logger *logrus.Logger
}

func (c *BlogController) Home(w http.ResponseWriter, r *http.Request) {
	err := c.View.ResponseTemplate(w, map[string]interface{}{}, "blog_pages_home")
	if err != nil {
		c.Logger.Error(err)
	}
}
