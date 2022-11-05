package models

import (
	"gorm.io/gorm"
	"time"
)

type UserRole struct {
	gorm.Model

	Id     string `gorm:"primaryKey;default:uuid_generate_v4()" json:"id"`
	UserId string `gorm:"column:user_id;not null;index;" json:"user_id"`
	RoleId string `gorm:"column:role_id;not null;index;" json:"role_id"`

	CreatedAt time.Time      `gorm:"column:created_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updated_at;not null;default:CURRENT_TIMESTAMP" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index,default:null" json:"-"`
}
