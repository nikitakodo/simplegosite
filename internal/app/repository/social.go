package repository

import (
	"encoding/json"
	"simplesite/internal/app/model"
	"simplesite/internal/app/store"
)

type SocialRepository struct {
	Store *store.Store
}

func (r SocialRepository) Create(m *model.Social) error {
	//TODO remake on callbacks
	if err := m.Validate(); err != nil {
		return err
	}
	err := r.Store.Db.Create(m).Error
	if err != nil {
		return err
	}
	err = r.Store.Cache.Del(m.GetTableCacheKey())
	if err != nil {
		return err
	}
	return nil
}

func (r SocialRepository) Find(id uint) (*model.Social, error) {
	var m model.Social
	val, err := r.Store.Cache.Get(m.GetItemCacheKey())
	if err != nil {
		return nil, err
	}
	if val == nil || *val == "" {
		err := r.Store.Db.Find(&m, m.GetId()).Error
		if err != nil {
			return nil, err
		}
		err = r.Store.Cache.SetStruct(m.GetItemCacheKey(), m)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(*val), &m)
		if err != nil {
			return nil, err
		}
	}
	return &m, nil
}

func (r SocialRepository) FindAll() ([]*model.Social, error) {
	var res []*model.Social
	var m model.Social
	val, err := r.Store.Cache.Get(m.GetTableCacheKey())
	if err != nil {
		return nil, err
	}
	if val == nil || *val == "" {
		err := r.Store.Db.Find(&res).Error
		if err != nil {
			return nil, err
		}
		err = r.Store.Cache.SetStruct(m.GetTableCacheKey(), res)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(*val), &res)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func (r SocialRepository) Update(m *model.Social) error {
	//TODO remake on callbacks
	if err := m.Validate(); err != nil {
		return err
	}
	err := r.Store.Db.Save(m).Error
	err = r.Store.Cache.Del(m.GetItemCacheKey())
	err = r.Store.Cache.Del(m.GetTableCacheKey())
	return err
}

func (r SocialRepository) Delete(m model.Social) (err error) {
	err = r.Store.Db.Delete(&m).Error
	err = r.Store.Cache.Del(m.GetItemCacheKey())
	err = r.Store.Cache.Del(m.GetTableCacheKey())
	return
}

func (r SocialRepository) GetOrdered(column string) ([]*model.Social, error) {
	var m model.Social
	var res []*model.Social
	val, err := r.Store.Cache.Get(m.GetTableCacheKey() + "_ordered")
	if err != nil {
		return nil, err
	}
	if val == nil || *val == "" {
		err := r.Store.Db.Order(column).Find(&res).Error
		if err != nil {
			return nil, err
		}
		err = r.Store.Cache.SetStruct(m.GetTableCacheKey()+"_ordered", res)
		if err != nil {
			return nil, err
		}

	} else {
		err = json.Unmarshal([]byte(*val), &res)
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}
