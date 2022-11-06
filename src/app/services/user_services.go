package services

import (
	"iecare-api/src/app/interfaces"
	"iecare-api/src/app/models"
	pagination2 "iecare-api/src/app/pkg/pagination"
)

type UserServices struct {
	ur interfaces.UserInterface
}

type IUserServices interface {
	interfaces.UserInterface
}

var _ IUserServices = &UserServices{}

func (u *UserServices) List(meta pagination2.Meta) (*pagination2.Pagination, error) {
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

func (u *UserServices) FindManyBy(field []string, value string) (*models.User, error) {
	return u.ur.FindManyBy(field, value)
}
