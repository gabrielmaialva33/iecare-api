package services

import (
	"iecare-api/src/app/modules/accounts/interfaces"
	"iecare-api/src/app/modules/accounts/models"
	"iecare-api/src/app/shared/pkg"
)

type UserServices struct {
	ur interfaces.UserInterface
}

type UserServicesInterface interface {
	interfaces.UserInterface
}

var _ UserServicesInterface = &UserServices{}

func (u *UserServices) List(meta pkg.Meta) (*pkg.Pagination, error) {
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
