package frontend

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"
	"simplesite/internal/app/model"
	"simplesite/internal/app/repository"
	"simplesite/internal/app/services"
	"simplesite/internal/app/store"
)

type BasicPageData struct {
	Title  string
	Nav    []model.Nav
	Social []model.Interface
}

func GetBasicData(
	View *services.View,
	Logger *logrus.Logger,
	Store *store.Store,
) (*BasicPageData, error) {
	navRepo := repository.NavRepository{
		Repository: repository.Repository{Store: Store},
	}
	navs, err := navRepo.GetOrdered(model.Nav{})
	if err != nil {
		fmt.Println("1", err)
		return nil, err
	}
	socialRepo := repository.SocialRepository{Store: Store}
	socialItems, err := socialRepo.FindAll(model.Social{})
	if err != nil {
		fmt.Println("2", err)
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
