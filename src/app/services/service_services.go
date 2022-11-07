package services

import (
	"iecare-api/src/app/interfaces"
	"iecare-api/src/app/models"
	"iecare-api/src/app/pkg/paginate"
)

type ServiceServices struct {
	ss interfaces.BaseRepository[models.Service]
}

type IServiceServices interface {
	interfaces.BaseRepository[models.Service]
}

var _ IServiceServices = &ServiceServices{}

func (s ServiceServices) List(meta paginate.Meta) (*paginate.Pagination, error) {
	return s.ss.List(meta)
}

func (s ServiceServices) Get(id string) (*models.Service, error) {
	return s.ss.Get(id)
}

func (s ServiceServices) Store(model *models.Service) (*models.Service, error) {
	return s.ss.Store(model)
}

func (s ServiceServices) Edit(model *models.Service) (*models.Service, error) {
	return s.ss.Edit(model)
}

func (s ServiceServices) Delete(model *models.Service) error {
	return s.ss.Delete(model)
}

func (s ServiceServices) FindBy(field string, value string) (*models.Service, error) {
	return s.ss.FindBy(field, value)
}

func (s ServiceServices) FindByMany(field []string, value string) (*models.Service, error) {
	return s.ss.FindByMany(field, value)
}
