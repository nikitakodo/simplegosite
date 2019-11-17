package sqlstore

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v7"
	"simplesite/internal/app/model"
	"simplesite/internal/app/services"
	"simplesite/internal/app/store"
	"strconv"
)

type NavRepository struct {
	Store *Store
	Cache *services.Cache
}

func (repo *NavRepository) Create(nav *model.Nav) error {

	if err := nav.Validate(); err != nil {
		return err
	}

	return repo.Store.Db.QueryRow(
		fmt.Sprintf(
			"INSERT INTO %s (order, title, uri) VALUES ($1, $2, $3) RETURNING id",
			nav.TableName(),
		),
		nav.Order,
		nav.Title,
		nav.Uri,
	).Scan(nav.ID)
}

func (repo *NavRepository) Find(id int) (*model.Nav, error) {
	nav := &model.Nav{}

	val, err := repo.Cache.GetMarshalStruct(nav.TableName() + "_" + strconv.Itoa(id))
	if err != nil && err != redis.Nil {
		return nil, err
	}
	if val == nil || *val == "" {
		if err := repo.Store.Db.QueryRow(
			fmt.Sprintf(
				"SELECT id, order, title, uri FROM %s WHERE id = $1",
				nav.TableName(),
			),
			id,
		).Scan(
			nav.ID,
			nav.Order,
			nav.Title,
			nav.Uri,
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

func (repo *NavRepository) FindAll() ([]*model.Nav, error) {

	val, err := repo.Cache.GetMarshalStruct(model.Nav{}.TableName() + "_all")
	if err != nil && err != redis.Nil {
		return nil, err
	}

	res := []*model.Nav{}

	if val == nil || *val == "" {
		rows, err := repo.Store.Db.Query(
			fmt.Sprintf(
				"select * from %s order by \"order\"",
				model.Nav{}.TableName(),
			),
		)
		if err != nil {
			return nil, err
		}
		defer rows.Close()

		for rows.Next() {
			nav := &model.Nav{}
			err := rows.Scan(&nav.ID,
				&nav.Order,
				&nav.Title,
				&nav.Uri,
			)
			if err != nil {
				return nil, err
			}
			res = append(res, nav)
		}

		err = repo.Cache.SetStruct(model.Nav{}.TableName()+"_all", res)
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

func (repo *NavRepository) Update(nav *model.Nav) error {
	_, err := repo.Store.Db.Exec(
		fmt.Sprintf(
			"update %s set order = $1, title=$2, uri=$3 where id = $4",
			nav.TableName(),
		),
		nav.Order,
		nav.Title,
		nav.Uri,
		nav.ID,
	)
	if err != nil {
		repo.Cache.Client.Del(nav.TableName() + "_" + strconv.Itoa(nav.ID))
	}

	return err
}

func (repo *NavRepository) Delete(id int) error {
	_, err := repo.Store.Db.Exec(
		fmt.Sprintf("delete from %s where id = $1", model.Nav{}.TableName()),
		id,
	)
	if err != nil {
		repo.Cache.Client.Del(model.Nav{}.TableName() + "_" + strconv.Itoa(id))
	}

	return err
}
