package sqlstore

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"simplesite/internal/app/model"
	"simplesite/internal/app/services"
	"simplesite/internal/app/store"
	"strconv"
)

type SocialRepository struct {
	Store *Store
	Cache *services.Cache
}

func (repo *SocialRepository) Create(social *model.Social) error {

	if err := social.Validate(); err != nil {
		return err
	}

	return repo.Store.Db.QueryRow(
		fmt.Sprintf(
			"INSERT INTO %s (order, icon, url) VALUES ($1, $2, $3) RETURNING id",
			social.TableName(),
		),
		social.Order,
		social.Icon,
		social.Url,
	).Scan(social.ID)
}

func (repo *SocialRepository) Find(id int) (*model.Social, error) {
	nav := &model.Social{}

	val, err := repo.Cache.Get(nav.TableName() + "_" + strconv.Itoa(id))
	if err != nil {
		return nil, err
	}

	if val == nil || *val == "" {
		if err := repo.Store.Db.QueryRow(
			fmt.Sprintf(
				"SELECT id, order, icon, url FROM %s WHERE id = $1",
				nav.TableName(),
			),
			id,
		).Scan(
			nav.ID,
			nav.Order,
			nav.Icon,
			nav.Url,
		); err != nil {
			if err == sql.ErrNoRows {
				return nil, store.ErrRecordNotFound
			}

			return nil, err
		}

		err = repo.Cache.SetStruct(nav.TableName()+"_"+strconv.Itoa(id), nav)
		if err != nil {
			return nil, err
		}
	} else {
		err = json.Unmarshal([]byte(*val), nav)
		if err != nil {
			return nil, err
		}
	}
	return nav, nil
}

func (repo *SocialRepository) FindAll() ([]*model.Social, error) {

	val, err := repo.Cache.Get(model.Social{}.TableName() + "_all")
	if err != nil {
		return nil, err
	}

	res := []*model.Social{}

	if val == nil || *val == "" {
		rows, err := repo.Store.Db.Query(
			fmt.Sprintf(
				"select * from %s order by \"order\"",
				model.Social{}.TableName(),
			),
		)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			nav := &model.Social{}
			err := rows.Scan(&nav.ID,
				&nav.Order,
				&nav.Icon,
				&nav.Url,
				&nav.CreateTime,
				&nav.UpdateTime,
			)
			if err != nil {
				return nil, err
			}
			res = append(res, nav)
		}

		err = repo.Cache.SetStruct(model.Social{}.TableName()+"_all", res)
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

func (repo *SocialRepository) Update(nav *model.Social) error {
	_, err := repo.Store.Db.Exec(
		fmt.Sprintf(
			"update %s set order = $1, icon=$2, url=$3 where id = $4",
			nav.TableName(),
		),
		nav.Order,
		nav.Icon,
		nav.Url,
		nav.ID,
	)
	if err != nil {
		repo.Cache.Client.Del(nav.TableName() + "_" + strconv.Itoa(nav.ID))
	}

	return err
}

func (repo *SocialRepository) Delete(id int) error {
	_, err := repo.Store.Db.Exec(
		fmt.Sprintf("delete from %s where id = $1", model.Social{}.TableName()),
		id,
	)
	if err != nil {
		repo.Cache.Client.Del(model.Social{}.TableName() + "_" + strconv.Itoa(id))
	}

	return err
}
