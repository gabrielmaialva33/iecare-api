package models

import (
	"gorm.io/gorm"
	"time"
)

const (
	RoleRoot     = "root"
	RoleAdmin    = "admin"
	RoleProvider = "provider"
	RoleUser     = "user"
	RoleGuest    = "guest"
)

type Role struct {
	gorm.Model

	Id          string         `gorm:"primaryKey;type:uuid;default:uuid_generate_v4();" json:"id"`
	Name        string         `gorm:"column:name;size:80;not null;unique;index;" json:"name" validate:"required,min=2,max=80,unique,omitempty"`
	Slug        string         `gorm:"column:slug;size:80;not null;unique;" json:"slug" validate:"required,min=2,max=80,unique,omitempty"`
	Description string         `gorm:"column:description;size:255;not null;" json:"description" validate:"required,min=2,max=255,omitempty"`
	CreatedAt   time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt   time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index,default:null" json:"-"`

	// Relationships
	Users Users `gorm:"many2many:user_roles;" json:"users"`
}
type Roles []Role

type RolePublic struct {
	Id          string `json:"id"`
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

func (r *Role) PublicRole() interface{} {
	return &RolePublic{
		Id:          r.Id,
		Slug:        r.Slug,
		Description: r.Description,
	}
}

func (r Roles) PublicRoles() []interface{} {
	roles := make([]interface{}, len(r))
	for i, role := range r {
		roles[i] = role.PublicRole()
	}
	if len(roles) == 0 {
		return []interface{}{}
	}
	return roles
}
