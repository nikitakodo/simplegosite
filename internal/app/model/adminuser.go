package model

type Admin struct {
	ID                int
	Email             string
	Password          string
	EncryptedPassword string
}

func (admin *Admin) BeforeCreate(model *Interface) error {
	return nil
}

func (admin *Admin) AfterCreate(model *Interface) error {
	return nil
}

func (admin *Admin) BeforeUpdate(model *Interface) error {
	return nil
}

func (admin *Admin) AfterUpdate(model *Interface) error {
	return nil
}

func (admin *Admin) BeforeDelete(model *Interface) error {
	return nil
}

func (admin *Admin) AfterDelete(model *Interface) error {
	return nil
}

func (admin *Admin) TableName() string {
	return "admins"
}
