package models

import (
	"gorm.io/gorm"
	"time"
)

type Provider struct {
	gorm.Model

	Id          string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"id"`
	Name        string         `gorm:"column:name;size:80;not null;unique;index;" json:"name" validate:"required,min=2,max=80,unique,omitempty"`
	Description string         `gorm:"column:description;type:text;not null;" json:"description" validate:"required,min=2,max=255,omitempty"`
	Email       string         `gorm:"column:email;size:255;unique;index;" json:"email" validate:"email,min=2,max=80,unique,omitempty"`
	Phone       string         `gorm:"column:phone;size:80;unique;index;" json:"phone" validate:"min=2,max=80,unique,omitempty"`
	WebSite     string         `gorm:"column:website;type:text;" json:"website" validate:"omitempty"`
	Logo        string         `gorm:"column:image;size:255;" json:"logo" validate:"omitempty"`
	Banner      string         `gorm:"column:banner;type:text;" json:"banner" validate:"omitempty"`
	UserId      string         `gorm:"column:user_id;not null;index;type:uuid" json:"user_id"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index,default:null" json:"-"`

	// Relationships
	User     User     `gorm:"foreignKey:UserId;references:Id;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"user" validate:"-"`
	Services Services `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"services"`
}

type Providers []Provider

type ProviderPublic struct {
	Id          string      `json:"id"`
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Email       string      `json:"email"`
	Phone       string      `json:"phone"`
	Logo        string      `json:"logo"`
	Banner      string      `json:"banner"`
	WebSite     string      `json:"website"`
	User        interface{} `json:"user"`
}

// PublicProvider returns a public representation of a provider.
func (p *Provider) PublicProvider() interface{} {
	return ProviderPublic{
		Id:          p.Id,
		Name:        p.Name,
		Description: p.Description,
		Email:       p.Email,
		Phone:       p.Phone,
		Logo:        p.Logo,
		Banner:      p.Banner,
		WebSite:     p.WebSite,
		User:        p.User.PublicUser(),
	}
}

// PublicProviders returns a public representation of a list of providers.
func (p *Providers) PublicProviders() []interface{} {
	var providers []interface{}
	for _, provider := range *p {
		providers = append(providers, provider.PublicProvider())
	}
	if len(providers) == 0 {
		return []interface{}{}
	}
	return providers
}
