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

type ServicesRepo struct {
	db *gorm.DB
}

// NewServicesRepository returns a new instance of ServicesRepo
func NewServicesRepository(db *gorm.DB) *ServicesRepo {
	return &ServicesRepo{db}
}

// *ServicesRepo implements interfaces.ServicesInterface
var _ interfaces.BaseRepository[models.Service] = &ServicesRepo{}

func (s *ServicesRepo) List(meta paginate.Meta) (*paginate.Pagination, error) {
	var services models.Services
	var fields = []string{"name", "description"}
	var pagination paginate.Pagination

	if err := s.db.Scopes(scopes.Paginate(services, fields, &meta, s.db)).Find(&services).Error; err != nil {
		return nil, err
	}

	pagination.SetMeta(meta)
	pagination.SetData(services.PublicServices())

	return &pagination, nil
}

func (s *ServicesRepo) Get(id string) (*models.Service, error) {
	var service models.Service
	if err := s.db.Where("id = ?", id).First(&service).Error; err != nil {
		return nil, err
	}
	return &service, nil
}

func (s *ServicesRepo) Store(services *models.Service) (*models.Service, error) {
	if err := s.db.Create(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}

func (s *ServicesRepo) Edit(services *models.Service) (*models.Service, error) {
	if err := s.db.Clauses(clause.Returning{}).Where("id = ?", services.Id).Updates(&services).Error; err != nil {
		return nil, err
	}
	return services, nil
}

func (s *ServicesRepo) Delete(services *models.Service) error {
	s.db.Where("id = ?", services.Id).Updates(&services)
	if err := s.db.Where("id = ?", services.Id).Delete(&services).Error; err != nil {
		return err
	}
	return nil
}

func (s *ServicesRepo) FindBy(field string, value string) (*models.Service, error) {
	var service models.Service
	if err := s.db.Where(field+" = ?", value).First(&service).Error; err != nil {
		return nil, err
	}
	return &service, nil
}

func (s *ServicesRepo) FindByMany(field []string, value string) (*models.Service, error) {
	var service models.Service
	for _, f := range field {
		s.db.Where(f+" = ?", value).First(&service)
		if service.Id != "" {
			return &service, nil
		}
	}
	return nil, errors.New("record not found")
}
