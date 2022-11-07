package models

import (
	"gorm.io/gorm"
	"time"
)

type Service struct {
	gorm.Model

	Id          string  `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"id"`
	Name        string  `gorm:"column:name;size:80;not null;index;" json:"name" validate:"required,min=2,max=80,omitempty"`
	Description string  `gorm:"column:description;type:text;not null;" json:"description" validate:"required,min=2,max=255,omitempty"`
	Price       float64 `gorm:"column:price;not null;" json:"price" validate:"required,min=0.01"`
	Image       string  `gorm:"column:image;size:255;" json:"image" validate:"omitempty"`

	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index,default:null" json:"-"`

	// Relationships
}

type Services []Service

type ServicePublic struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
}

// PublicService returns a public representation of a service.
func (s *Service) PublicService() interface{} {
	return ServicePublic{
		Id:          s.Id,
		Name:        s.Name,
		Description: s.Description,
		Price:       s.Price,
		Image:       s.Image,
	}
}

// PublicServices returns a public representation of a list of services.
func (s *Services) PublicServices() []interface{} {
	var services []interface{}
	for _, service := range *s {
		services = append(services, service.PublicService())
	}
	if len(services) == 0 {
		return []interface{}{}
	}
	return services
}
