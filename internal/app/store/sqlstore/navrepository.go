package sqlstore

import (
	"database/sql"
	"fmt"
	"simplesite/internal/app/model"
	"simplesite/internal/app/store"
)

type NavRepository struct {
	Store *Store
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

	return nav, nil
}

func (repo *NavRepository) FindAll() ([]*model.Nav, error) {

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
	res := []*model.Nav{}

	for rows.Next() {
		nav := &model.Nav{}
		err := rows.Scan(&nav.ID,
			&nav.Order,
			&nav.Title,
			&nav.Uri)
		if err != nil {
			return nil, err
		}
		res = append(res, nav)
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

	return err
}

func (repo *NavRepository) Delete(id int) error {
	_, err := repo.Store.Db.Exec(
		fmt.Sprintf("delete from %s where id = $1", model.Nav{}.TableName()),
		id,
	)

	return err
}
