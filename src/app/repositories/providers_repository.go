package repositories

import (
	"errors"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"iecare-api/src/app/interfaces"
	"iecare-api/src/app/models"
	"iecare-api/src/app/pkg/paginate"
	"iecare-api/src/app/scopes"
)

type ProviderRepo struct {
	db *gorm.DB
}

func NewProvidersRepository(db *gorm.DB) *ProviderRepo {
	return &ProviderRepo{db}
}

var _ interfaces.BaseRepository[models.Provider] = &ProviderRepo{}

func (p *ProviderRepo) List(meta paginate.Meta) (*paginate.Pagination, error) {
	var providers models.Providers
	var pagination paginate.Pagination
	var fields = []string{"name", "description", "email", "phone"}

	if err := p.db.Preload(clause.Associations).Scopes(scopes.Paginate(providers, fields, &meta, p.db)).Find(&providers).Error; err != nil {
		return nil, err
	}

	pagination.SetMeta(meta)
	pagination.SetData(providers.PublicProviders())

	return &pagination, nil
}

func (p *ProviderRepo) Get(id string) (*models.Provider, error) {
	var provider models.Provider
	if err := p.db.Preload("User").Preload("User.Roles").Where("id = ?", id).First(&provider).Error; err != nil {
		return nil, err
	}
	return &provider, nil
}

func (p *ProviderRepo) Store(provider *models.Provider) (*models.Provider, error) {
	if err := p.db.Create(&provider).Error; err != nil {
		return nil, err
	}
	p.db.Preload("User").Preload("User.Roles").Where("id = ?", provider.Id).First(&provider)
	return provider, nil
}

func (p *ProviderRepo) Edit(model *models.Provider) (*models.Provider, error) {
	if err := p.db.Preload("User").Clauses(clause.Returning{}).Where("id = ?", model.Id).Updates(&model).Error; err != nil {
		return nil, err
	}
	return model, nil
}

func (p *ProviderRepo) Delete(model *models.Provider) error {
	p.db.Where("id = ?", model.Id).Updates(&model)
	if err := p.db.Where("id = ?", model.Id).Delete(&model).Error; err != nil {
		return err
	}
	return nil
}

func (p *ProviderRepo) FindBy(field string, value string) (*models.Provider, error) {
	var provider models.Provider
	if err := p.db.Preload("User").Where(field+" = ?", value).First(&provider).Error; err != nil {
		return nil, err
	}
	return &provider, nil
}

func (p *ProviderRepo) FindByMany(field []string, value string) (*models.Provider, error) {
	var provider models.Provider
	for _, v := range field {
		p.db.Preload("User").Where(v+" = ?", value).First(&provider)
		if provider.Id != "" {
			return &provider, nil
		}
	}
	return nil, errors.New("record not found")
}
