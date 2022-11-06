package services

import (
	"iecare-api/src/app/interfaces"
	"iecare-api/src/app/models"
	pagination2 "iecare-api/src/app/pkg/pagination"
)

type ProviderServices struct {
	pr interfaces.ProviderInterface
}

type IProviderServices interface {
	interfaces.ProviderInterface
}

var _ IProviderServices = &ProviderServices{}

func (p *ProviderServices) List(meta pagination2.Meta) (*pagination2.Pagination, error) {
	return p.pr.List(meta)
}

func (p *ProviderServices) Get(id string) (*models.Provider, error) {
	return p.pr.Get(id)
}

func (p *ProviderServices) Store(model *models.Provider) (*models.Provider, error) {
	return p.pr.Store(model)
}

func (p *ProviderServices) Edit(model *models.Provider) (*models.Provider, error) {
	return p.pr.Edit(model)
}

func (p *ProviderServices) Delete(model *models.Provider) error {
	return p.pr.Delete(model)
}

func (p *ProviderServices) FindBy(field string, value string) (*models.Provider, error) {
	return p.pr.FindBy(field, value)
}

func (p *ProviderServices) FindManyBy(field []string, value string) (*models.Provider, error) {
	return p.pr.FindManyBy(field, value)
}
