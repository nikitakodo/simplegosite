package sqlstore

import "simplesite/internal/app/model"

type AdminRepository struct {
	store *Store
}

func (repo AdminRepository) Create(abstractModel *model.Interface) error {
	panic("implement me")
}

func (repo AdminRepository) Find(int) (*model.Interface, error) {
	panic("implement me")
}

func (repo AdminRepository) Update(abstractModel *model.Interface) error {
	panic("implement me")
}

func (repo AdminRepository) Delete(int) error {
	panic("implement me")
}
