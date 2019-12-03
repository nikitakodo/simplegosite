package repository

import (
	"encoding/json"
	"simplesite/internal/app/model"
	"simplesite/internal/app/store"
)

type SettingsRepository struct {
	Store *store.Store
}

func (r SettingsRepository) Create(m *model.Settings) error {
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

func (r SettingsRepository) Find(id uint) (*model.Settings, error) {
	var m model.Settings
	val, err := r.Store.Cache.Get(m.GetItemCacheKey())
	if err != nil {
		return nil, err
	}
	if val == nil || *val == "" {
		err := r.Store.Db.Find(&m, id).Error
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

func (r SettingsRepository) FindByKey(key string) (*model.Settings, error) {
	var m model.Settings
	val, err := r.Store.Cache.Get(m.GetTableCacheKey() + "_" + key)
	if err != nil {
		return nil, err
	}
	if val == nil || *val == "" {
		err := r.Store.Db.Where("key = ?", key).Find(&m).Error
		if err != nil {
			return nil, err
		}
		err = r.Store.Cache.SetStruct(m.GetTableCacheKey()+"_"+key, m)
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

func (r SettingsRepository) FindAll() ([]*model.Settings, error) {
	var res []*model.Settings
	var m model.Settings
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

func (r SettingsRepository) Update(m *model.Settings) error {
	//TODO remake on callbacks
	if err := m.Validate(); err != nil {
		return err
	}
	err := r.Store.Db.Save(m).Error
	err = r.Store.Cache.Del(m.GetItemCacheKey())
	err = r.Store.Cache.Del(m.GetTableCacheKey())
	return err
}

func (r SettingsRepository) Delete(m model.Settings) (err error) {
	err = r.Store.Db.Delete(&m).Error
	err = r.Store.Cache.Del(m.GetItemCacheKey())
	err = r.Store.Cache.Del(m.GetTableCacheKey())
	return
}
