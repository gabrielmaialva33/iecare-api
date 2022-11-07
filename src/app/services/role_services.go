package services

import (
	"iecare-api/src/app/interfaces"
	"iecare-api/src/app/models"
	"iecare-api/src/app/pkg/paginate"
)

type RoleServices struct {
	rr interfaces.BaseRepository[models.Role]
}

type IRoleServices interface {
	interfaces.BaseRepository[models.Role]
}

var _ IRoleServices = &RoleServices{}

func (r *RoleServices) List(meta paginate.Meta) (*paginate.Pagination, error) {
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

func (r *RoleServices) FindByMany(field []string, value string) (*models.Role, error) {
	return r.rr.FindByMany(field, value)
}
