package models

import (
	"gorm.io/gorm"
	"time"
)

type Category struct {
	gorm.Model

	Id          string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"id"`
	Name        string         `gorm:"column:name;size:100;not null;unique;index;" json:"name" validate:"required,min=2,max=255,unique,omitempty"`
	Description string         `gorm:"column:description;type:text;not null;" json:"description" validate:"required,min=2,max=255,omitempty"`
	Icon        string         `gorm:"column:icon;size:255;" json:"icon" validate:"omitempty"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index,default:null" json:"-"`

	// Relationships
	Services Services `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"services"`
}

type Categories []Category

type CategoryPublic struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

// PublicCategory returns a public representation of a category.
func (c *Category) PublicCategory() interface{} {
	return CategoryPublic{
		Id:          c.Id,
		Name:        c.Name,
		Description: c.Description,
		Icon:        c.Icon,
	}
}

// PublicCategories returns a public representation of a list of categories.
func (c *Categories) PublicCategories() []interface{} {
	var categories []interface{}
	for _, category := range *c {
		categories = append(categories, category.PublicCategory())
	}
	if len(categories) == 0 {
		return []interface{}{}
	}
	return categories
}
