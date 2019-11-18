package frontend

import (
	"github.com/sirupsen/logrus"
	"simplesite/internal/app/model"
	"simplesite/internal/app/services"
	"simplesite/internal/app/store/sqlstore"
)

type BasicPageData struct {
	Title  string
	Nav    []*model.Nav
	Social []*model.Social
}

func GetBasicData(
	View *services.View,
	Logger *logrus.Logger,
	Store *sqlstore.Store,
	Cache *services.Cache,
) (*BasicPageData, error) {
	navRepo := sqlstore.NavRepository{Store: Store, Cache: Cache}
	navs, err := navRepo.FindAll()
	if err != nil {
		return nil, err
	}
	socialRepo := sqlstore.SocialRepository{Store: Store, Cache: Cache}
	socialItems, err := socialRepo.FindAll()
	if err != nil {
		return nil, err
	}
	data := &BasicPageData{
		Nav:    navs,
		Social: socialItems,
	}

	return data, nil
}
