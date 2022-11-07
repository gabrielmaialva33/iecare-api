package services

import (
	"iecare-api/src/app/interfaces"
	"iecare-api/src/app/models"
	paginate "iecare-api/src/app/pkg/paginate"
)

type ProviderServices struct {
	pr interfaces.ProviderInterface
}

type IProviderServices interface {
	interfaces.ProviderInterface
}

var _ IProviderServices = &ProviderServices{}

func (p *ProviderServices) List(meta paginate.Meta) (*paginate.Pagination, error) {
	return p.pr.List(meta)
}

func (p *ProviderServices) Get(id string) (*models.Provider, error) {
	return p.pr.Get(id)
}

func (p *ProviderServices) Store(provider *models.Provider) (*models.Provider, error) {
	return p.pr.Store(provider)
}

func (p *ProviderServices) Edit(provider *models.Provider) (*models.Provider, error) {
	return p.pr.Edit(provider)
}

func (p *ProviderServices) Delete(provider *models.Provider) error {
	return p.pr.Delete(provider)
}

func (p *ProviderServices) FindBy(field string, value string) (*models.Provider, error) {
	return p.pr.FindBy(field, value)
}

func (p *ProviderServices) FindByMany(field []string, value string) (*models.Provider, error) {
	return p.pr.FindByMany(field, value)
}
