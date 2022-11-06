package services

import (
	"iecare-api/src/app/interfaces"
	"iecare-api/src/app/models"
	pagination2 "iecare-api/src/app/pkg/pagination"
)

type RoleServices struct {
	rr interfaces.RoleInterface
}

type IRoleServices interface {
	interfaces.RoleInterface
}

var _ IRoleServices = &RoleServices{}

func (r *RoleServices) List(meta pagination2.Meta) (*pagination2.Pagination, error) {
	return r.rr.List(meta)
}

func (r *RoleServices) Get(id string) (*models.Role, error) {
	return r.rr.Get(id)
}

func (r *RoleServices) Store(role *models.Role) (*models.Role, error) {
	return r.rr.Store(role)
}

func (r *RoleServices) Edit(role *models.Role) (*models.Role, error) {
	return r.rr.Edit(role)
}

func (r *RoleServices) Delete(role *models.Role) error {
	return r.rr.Delete(role)
}

func (r *RoleServices) FindBy(field string, value string) (*models.Role, error) {
	return r.rr.FindBy(field, value)
}

func (r *RoleServices) FindManyBy(field []string, value string) (*models.Role, error) {
	return r.rr.FindManyBy(field, value)
}
