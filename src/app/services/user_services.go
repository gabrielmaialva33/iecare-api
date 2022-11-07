package services

import (
	"iecare-api/src/app/interfaces"
	"iecare-api/src/app/models"
	paginate "iecare-api/src/app/pkg/paginate"
)

type UserServices struct {
	ur interfaces.UserInterface
}

type IUserServices interface {
	interfaces.UserInterface
}

var _ IUserServices = &UserServices{}

func (u *UserServices) List(meta paginate.Meta) (*paginate.Pagination, error) {
	return u.ur.List(meta)
}

func (u *UserServices) Get(id string) (*models.User, error) {
	return u.ur.Get(id)
}

func (u *UserServices) Store(user *models.User) (*models.User, error) {
	return u.ur.Store(user)
}

func (u *UserServices) Edit(user *models.User) (*models.User, error) {
	return u.ur.Edit(user)
}

func (u *UserServices) Delete(user *models.User) error {
	return u.ur.Delete(user)
}

func (u *UserServices) FindBy(field string, value string) (*models.User, error) {
	return u.ur.FindBy(field, value)
}

func (u *UserServices) FindByMany(field []string, value string) (*models.User, error) {
	return u.ur.FindByMany(field, value)
}
