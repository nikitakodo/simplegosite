package blog

import (
	"github.com/sirupsen/logrus"
	"simplesite/internal/app/model"
	"simplesite/internal/app/repository"
	"simplesite/internal/app/services"
	"simplesite/internal/app/store"
	"time"
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

func GetCachedView(templateName string, cache *store.Cache) (*string, error) {
	result, err := cache.Get(templateName)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func SetCachedView(content string, key string, cache *store.Cache) error {
	err := cache.Set(key, content, 60*10*time.Second)
	if err != nil {
		return err
	}
	return nil
}
