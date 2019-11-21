package frontend

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"simplesite/internal/app/model"
	"simplesite/internal/app/repository"
	"simplesite/internal/app/services"
	"simplesite/internal/app/store"
)

type BasicPageData struct {
	Title  string
	Nav    []*model.Nav
	Social []*model.Social
}

func GetBasicData(
	View *services.View,
	Logger *logrus.Logger,
	Store *store.Store,
) (*BasicPageData, error) {
	navs, err := repository.NavRepository{Store: Store}.GetOrdered("order")
	if err != nil {
		return nil, err
	}
	socialItems, err := repository.SocialRepository{Store: Store}.GetOrdered("order")
	if err != nil {
		return nil, err
	}
	data := &BasicPageData{
		Nav:    navs,
		Social: socialItems,
	}
	return data, nil
}

func WriteCachedResponse(w http.ResponseWriter, templateName string, cache *store.Cache) bool {
	val, err := cache.Get(templateName)
	if err != nil {
		return false
	}
	if len(*val) > 0 {
		_, err := w.Write([]byte(*val))
		if err != nil {
			return false
		}
		return true
	}
	return false
}
